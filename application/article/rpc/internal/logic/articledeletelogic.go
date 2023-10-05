package logic

import (
	"context"

	"beyond/application/article/rpc/internal/code"
	"beyond/application/article/rpc/internal/svc"
	"beyond/application/article/rpc/internal/types"
	"beyond/application/article/rpc/pb"
	"beyond/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDeleteLogic {
	return &ArticleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleDeleteLogic) ArticleDelete(in *pb.ArticleDeleteRequest) (*pb.ArticleDeleteResponse, error) {
	if in.UserId <= 0 {
		return nil, code.UserIdInvalid
	}
	if in.ArticleId <= 0 {
		return nil, code.ArticleIdInvalid
	}
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.ArticleId)
	if err != nil {
		l.Logger.Errorf("ArticleDelete FindOne req: %v error: %v", in, err)
		return nil, err
	}
	if article.AuthorId != in.UserId {
		return nil, xcode.AccessDenied
	}
	err = l.svcCtx.ArticleModel.UpdateArticleStatus(l.ctx, in.ArticleId, types.ArticleStatusUserDelete)
	if err != nil {
		l.Logger.Errorf("UpdateArticleStatus req: %v error: %v", in, err)
		return nil, err
	}
	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, articlesKey(in.UserId, types.SortPublishTime), in.ArticleId)
	if err != nil {
		l.Logger.Errorf("ZremCtx req: %v error: %v", in, err)
	}
	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, articlesKey(in.UserId, types.SortLikeCount), in.ArticleId)
	if err != nil {
		l.Logger.Errorf("ZremCtx req: %v error: %v", in, err)
	}

	return &pb.ArticleDeleteResponse{}, nil
}
