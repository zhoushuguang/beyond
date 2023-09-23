package logic

import (
	"beyond/application/like/rpc/internal/types"
	"context"
	"encoding/json"

	"beyond/application/like/rpc/internal/svc"
	"beyond/application/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbupLogic) Thumbup(in *service.ThumbupRequest) (*service.ThumbupResponse, error) {
	// TODO 逻辑暂时忽略
	// 1. 查询是否点过赞
	// 2. 计算当前内容的总点赞数和点踩数

	msg := &types.ThumbupMsg{
		BizId:    in.BizId,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
		LikeType: in.LikeType,
	}
	// 发送kafka消息，异步
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("[Thumbup] marshal msg: %+v error: %v", msg, err)
			return
		}
		err = l.svcCtx.KqPusherClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[Thumbup] kq push data: %s error: %v", data, err)
		}

	})

	return &service.ThumbupResponse{}, nil
}
