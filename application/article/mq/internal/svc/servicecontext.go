package svc

import (
	"beyond/application/article/mq/internal/config"
	"beyond/application/article/mq/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Datasource)
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(conn),
	}
}
