package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikeRecordModel = (*customLikeRecordModel)(nil)

type (
	// LikeRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikeRecordModel.
	LikeRecordModel interface {
		likeRecordModel
	}

	customLikeRecordModel struct {
		*defaultLikeRecordModel
	}
)

// NewLikeRecordModel returns a model for the database table.
func NewLikeRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LikeRecordModel {
	return &customLikeRecordModel{
		defaultLikeRecordModel: newLikeRecordModel(conn, c, opts...),
	}
}
