package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	KqConsumerConf        kq.KqConf
	ArticleKqConsumerConf kq.KqConf
	Datasource            string
	BizRedis              redis.RedisConf
}
