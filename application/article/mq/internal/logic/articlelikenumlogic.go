package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"beyond/application/article/mq/internal/svc"
	"beyond/application/article/mq/internal/types"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type ArticleLikeNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleLikeNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLikeNumLogic {
	return &ArticleLikeNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleLikeNumLogic) Consume(_, val string) error {
	var msg *types.CanalLikeMsg
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}

	return l.updateArticleLikeNum(l.ctx, msg)
}

func (l *ArticleLikeNumLogic) updateArticleLikeNum(ctx context.Context, msg *types.CanalLikeMsg) error {
	if len(msg.Data) == 0 {
		return nil
	}

	for _, d := range msg.Data {
		if d.BizID != types.ArticleBizID {
			continue
		}
		id, err := strconv.ParseInt(d.ObjID, 10, 64)
		if err != nil {
			logx.Errorf("strconv.ParseInt id: %s error: %v", d.ID, err)
			continue
		}
		likeNum, err := strconv.ParseInt(d.LikeNum, 10, 64)
		if err != nil {
			logx.Errorf("strconv.ParseInt likeNum: %s error: %v", d.LikeNum, err)
			continue
		}
		err = l.svcCtx.ArticleModel.UpdateLikeNum(ctx, id, likeNum)
		if err != nil {
			logx.Errorf("UpdateLikeNum id: %d like: %d", id, likeNum)
		}
	}

	return nil
}

func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewArticleLikeNumLogic(ctx, svcCtx)),
		kq.MustNewQueue(svcCtx.Config.ArticleKqConsumerConf, NewArticleLogic(ctx, svcCtx)),
	}
}
