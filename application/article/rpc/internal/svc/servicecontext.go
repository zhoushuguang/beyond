package svc

import (
	"beyond/application/article/rpc/internal/config"
	"beyond/application/article/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
)

type ServiceContext struct {
	Config            config.Config
	ArticleModel      model.ArticleModel
	BizRedis          *redis.Redis
	SingleFlightGroup singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		BizRedis:     rds,
	}
}
