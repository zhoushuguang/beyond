package logic

import (
	"context"
	"strconv"
	"time"

	"beyond/application/follow/code"
	"beyond/application/follow/rpc/internal/model"
	"beyond/application/follow/rpc/internal/svc"
	"beyond/application/follow/rpc/internal/types"
	"beyond/application/follow/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

const userFollowExpireTime = 3600 * 24 * 2

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FollowList 关注列表 .
func (l *FollowListLogic) FollowList(in *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	if in.UserId == 0 {
		return nil, code.UserIdEmpty
	}
	if in.PageSize == 0 {
		in.PageSize = types.DefaultPageSize
	}
	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}

	var (
		err             error
		isCache, isEnd  bool
		lastId, cursor  int64
		followedUserIds []int64
		follows         []*model.Follow
		curPage         []*pb.FollowItem
	)

	followUserIds, _ := l.cacheFollowUserIds(l.ctx, in.UserId, in.Cursor, in.PageSize)
	if len(followUserIds) > 0 {
		isCache = true
		if followUserIds[len(followUserIds)-1] == -1 {
			followUserIds = followUserIds[:len(followUserIds)-1]
			isEnd = true
		}
		if len(followUserIds) == 0 {
			return &pb.FollowListResponse{}, nil
		}
		follows, err = l.svcCtx.FollowModel.FindByFollowedUserIds(l.ctx, in.UserId, followUserIds)
		if err != nil {
			l.Logger.Errorf("[FollowList] FollowModel.FindByFollowedUserIds error: %v req: %v", err, in)
			return nil, err
		}
		for _, follow := range follows {
			followedUserIds = append(followedUserIds, follow.FollowedUserID)
			curPage = append(curPage, &pb.FollowItem{
				Id:             follow.ID,
				FollowedUserId: follow.FollowedUserID,
				CreateTime:     follow.CreateTime.Unix(),
			})
		}
	} else {
		follows, err = l.svcCtx.FollowModel.FindByUserId(l.ctx, in.UserId, types.CacheMaxFollowCount)
		if err != nil {
			l.Logger.Errorf("[FollowList] FollowModel.FindByUserId error: %v req: %v", err, in)
			return nil, err
		}
		if len(follows) == 0 {
			return &pb.FollowListResponse{}, nil
		}
		var firstPageFollows []*model.Follow
		if len(follows) > int(in.PageSize) {
			firstPageFollows = follows[:in.PageSize]
		} else {
			firstPageFollows = follows
			isEnd = true
		}
		for _, follow := range firstPageFollows {
			followedUserIds = append(followedUserIds, follow.FollowedUserID)
			curPage = append(curPage, &pb.FollowItem{
				Id:             follow.ID,
				FollowedUserId: follow.FollowedUserID,
				CreateTime:     follow.CreateTime.Unix(),
			})
		}
	}
	if len(curPage) > 0 {
		pageLast := curPage[len(curPage)-1]
		lastId = pageLast.Id
		cursor = pageLast.CreateTime
		if cursor < 0 {
			cursor = 0
		}
		for k, follow := range curPage {
			if follow.CreateTime == in.Cursor && follow.Id == in.Id {
				curPage = curPage[k:]
				break
			}
		}
	}
	fc, err := l.svcCtx.FollowCountModel.FindByUserIds(l.ctx, followedUserIds)
	if err != nil {
		l.Logger.Errorf("[FollowList] FollowCountModel.FindByUserIds error: %v followedUserIds: %v", err, followedUserIds)
	}
	uidFansCount := make(map[int64]int)
	for _, f := range fc {
		uidFansCount[f.UserID] = f.FansCount
	}
	for _, cur := range curPage {
		cur.FansCount = int64(uidFansCount[cur.FollowedUserId])
	}
	ret := &pb.FollowListResponse{
		IsEnd:  isEnd,
		Cursor: cursor,
		Id:     lastId,
		Items:  curPage,
	}

	if !isCache {
		threading.GoSafe(func() {
			if len(follows) < types.CacheMaxFollowCount && len(follows) > 0 {
				follows = append(follows, &model.Follow{FollowedUserID: -1})
			}
			err = l.addCacheFollow(context.Background(), in.UserId, follows)
			if err != nil {
				logx.Errorf("addCacheFollow error: %v", err)
			}
		})
	}

	return ret, nil
}

func (l *FollowListLogic) cacheFollowUserIds(ctx context.Context, userId, cursor, pageSize int64) ([]int64, error) {
	key := userFollowKey(userId)
	b, err := l.svcCtx.BizRedis.ExistsCtx(ctx, key)
	if err != nil {
		logx.Errorf("[cacheFollowUserIds] BizRedis.ExistsCtx error: %v", err)
	}
	if b {
		err = l.svcCtx.BizRedis.ExpireCtx(ctx, key, userFollowExpireTime)
		if err != nil {
			logx.Errorf("[cacheFollowUserIds] BizRedis.ExpireCtx error: %v", err)
		}
	}
	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, key, 0, cursor, 0, int(pageSize))
	if err != nil {
		logx.Errorf("[cacheFollowUserIds] BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx error: %v", err)
		return nil, err
	}
	var uids []int64
	for _, pair := range pairs {
		uid, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			logx.Errorf("[cacheFollowUserIds] strconv.ParseInt error: %v", err)
			continue
		}
		uids = append(uids, uid)
	}

	return uids, nil
}

func (l *FollowListLogic) addCacheFollow(ctx context.Context, userId int64, follows []*model.Follow) error {
	if len(follows) == 0 {
		return nil
	}
	key := userFollowKey(userId)
	for _, follow := range follows {
		var score int64
		if follow.FollowedUserID == -1 {
			score = 0
		} else {
			score = follow.CreateTime.Unix()
		}
		_, err := l.svcCtx.BizRedis.ZaddCtx(ctx, key, score, strconv.FormatInt(follow.FollowedUserID, 10))
		if err != nil {
			logx.Errorf("[addCacheFollow] BizRedis.ZaddCtx error: %v", err)
			return err
		}
	}

	return l.svcCtx.BizRedis.ExpireCtx(ctx, key, userFollowExpireTime)
}
