package logic

import (
	"context"

	"beyond/application/follow/rpc/internal/svc"
	"beyond/application/follow/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FansListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansListLogic {
	return &FansListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FansList 粉丝列表.
func (l *FansListLogic) FansList(in *pb.FansListRequest) (*pb.FansListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.FansListResponse{}, nil
}
