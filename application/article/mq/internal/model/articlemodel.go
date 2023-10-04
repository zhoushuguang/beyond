package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		UpdateLikeNum(ctx context.Context, id, likeNum int64) error
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn),
	}
}

func (m *customArticleModel) UpdateLikeNum(ctx context.Context, id, likeNum int64) error {
	query := fmt.Sprintf("update %s set like_num = ? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, likeNum, id)
	return err
}
