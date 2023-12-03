package redis

import (
	"github.com/dobyte/due/v2/etc"
	"github.com/dobyte/due/v2/log"
	"github.com/redis/go-redis/v9"
	"sync"
)

const (
	defaultAddr       = "127.0.0.1:6379"
	defaultDB         = 0
	defaultMaxRetries = 3
	defaultPrefix     = "game_data"
	defaultPassword   = 0
)

var (
	once   sync.Once
	client redis.UniversalClient
)

func Client() redis.UniversalClient {
	once.Do(func() {
		conf := &struct {
			Addrs      []string
			Database   int
			Prefix     string
			Password   string
			MaxRetries int
		}{}

		conf.Addrs = etc.Get("etc.redis.addr", []string{defaultAddr}).Strings()
		conf.Database = etc.Get("etc.redis.database", defaultDB).Int()
		conf.Prefix = etc.Get("etc.redis.prefix", defaultPrefix).String()
		conf.Password = etc.Get("etc.redis.password", defaultPassword).String()
		conf.MaxRetries = etc.Get("etc.redis.password", defaultMaxRetries).Int()

		log.Infof("redis Client() %+v", conf)

		client = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:      conf.Addrs,
			DB:         conf.Database,
			MaxRetries: conf.MaxRetries,
		})
	})

	return client
}

// DB 获取数据库
func DB() redis.UniversalClient {
	Client()
	return client
}
