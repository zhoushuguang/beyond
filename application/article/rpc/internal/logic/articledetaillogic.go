package logic

import (
	"context"

	"beyond/application/article/rpc/internal/svc"
	"beyond/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ArticleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDetailLogic {
	return &ArticleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleDetailLogic) ArticleDetail(in *pb.ArticleDetailRequest) (*pb.ArticleDetailResponse, error) {
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.ArticleId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return &pb.ArticleDetailResponse{}, nil
		}
		return nil, err
	}

	return &pb.ArticleDetailResponse{
		Article: &pb.ArticleItem{
			Id:    article.Id,
			Title: article.Title,
		},
	}, nil
}
