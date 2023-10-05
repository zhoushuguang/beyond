package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"beyond/application/article/mq/internal/svc"
	"beyond/application/article/mq/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleLogic) Consume(_, val string) error {
	logx.Infof("Consume msg val: %s", val)
	var msg *types.CanalArticleMsg
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}

	return l.articleOperate(l.ctx, msg)
}

func (l *ArticleLogic) articleOperate(ctx context.Context, msg *types.CanalArticleMsg) error {
	if len(msg.Data) == 0 {
		return nil
	}
	for _, d := range msg.Data {
		status, _ := strconv.Atoi(d.Status)
		likNum, _ := strconv.ParseInt(d.LikeNum, 10, 64)

		t, err := time.ParseInLocation("2006-01-02 15:04:05", d.PublishTime, time.Local)
		publishTimeKey := articlesKey(d.AuthorId, 0)
		likeNumKey := articlesKey(d.AuthorId, 1)

		switch status {
		case types.ArticleStatusVisible:
			b, _ := l.svcCtx.BizRedis.ExistsCtx(ctx, publishTimeKey)
			if b {
				_, err = l.svcCtx.BizRedis.ZaddCtx(ctx, publishTimeKey, t.Unix(), d.ID)
				if err != nil {
					l.Logger.Errorf("ZaddCtx key: %s req: %v error: %v", publishTimeKey, d, err)
				}
			}
			b, _ = l.svcCtx.BizRedis.ExistsCtx(ctx, likeNumKey)
			if b {
				_, err = l.svcCtx.BizRedis.ZaddCtx(ctx, likeNumKey, likNum, d.ID)
				if err != nil {
					l.Logger.Errorf("ZaddCtx key: %s req: %v error: %v", likeNumKey, d, err)
				}
			}
		case types.ArticleStatusUserDelete:
			_, err = l.svcCtx.BizRedis.ZremCtx(ctx, publishTimeKey, d.ID)
			if err != nil {
				logx.Errorf("ZremCtx key: %s req: %v error: %v", publishTimeKey, d, err)
			}
			_, err = l.svcCtx.BizRedis.ZremCtx(ctx, likeNumKey, d.ID)
			if err != nil {
				logx.Errorf("ZremCtx key: %s req: %v error: %v", likeNumKey, d, err)
			}
		}
	}

	return nil
}

func articlesKey(uid string, sortType int32) string {
	return fmt.Sprintf("biz#articles#%s#%d", uid, sortType)
}
