package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		ArticlesByUserId(ctx context.Context, userId, likeNum int64, pubTime, sortField string, limit int) ([]*Article, error)
		UpdateArticleStatus(ctx context.Context, id int64, status int) error
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn, c, opts...),
	}
}

func (m *customArticleModel) ArticlesByUserId(ctx context.Context, userId, likeNum int64, pubTime, sortField string, limit int) ([]*Article, error) {
	var (
		err      error
		sql      string
		anyField any
		articles []*Article
	)
	if sortField == "like_num" {
		anyField = likeNum
		sql = fmt.Sprintf("select "+articleRows+" from "+m.table+" where author_id=? and like_num < ? order by %s desc limit ?", sortField)
	} else {
		anyField = pubTime
		sql = fmt.Sprintf("select "+articleRows+" from "+m.table+" where author_id=? and publish_time < ? order by %s desc limit ?", sortField)
	}
	err = m.QueryRowsNoCacheCtx(ctx, &articles, sql, userId, anyField, limit)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (m *customArticleModel) UpdateArticleStatus(ctx context.Context, id int64, status int) error {
	beyondArticleArticleIdKey := fmt.Sprintf("%s%v", cacheBeyondArticleArticleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set status = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, beyondArticleArticleIdKey)

	return err
}
