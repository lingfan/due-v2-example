package redis

import (
	"context"
	"github.com/dobyte/due/v2/etc"
	"github.com/dobyte/due/v2/log"
	"github.com/redis/rueidis"
	"sync"
)

var (
	onceRueidis   sync.Once
	rueidisClient *RueidisClient
)

func rueidis_init() {
	onceRueidis.Do(func() {
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

		log.Infof("rueidis_init() %+v", conf)

		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: conf.Addrs, SelectDB: conf.Database})
		if err != nil {
			panic(err)
		}
		rueidisClient = NewRueidisClient(context.Background(), client)
	})

}

// DB 获取数据库
func RueidisRD() *RueidisClient {
	rueidis_init()
	return rueidisClient
}

type RueidisClient struct {
	ctx    context.Context
	client rueidis.Client // Replace with the actual Redis client type
}

func NewRueidisClient(ctx context.Context, client rueidis.Client) *RueidisClient {
	return &RueidisClient{
		ctx:    ctx,
		client: client,
	}
}

func (rc *RueidisClient) Get(key string) string {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		return ""
	}
	return val
}

func (rc *RueidisClient) MGet(keys ...string) []string {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Mget().Key(keys...).Build()).AsStrSlice()
	if err != nil {
		return []string{}
	}
	return val
}

func (rc *RueidisClient) Set(key, value string) error {
	return rc.client.Do(rc.ctx, rc.client.B().Set().Key(key).Value(value).Build()).Error()
}

func (rc *RueidisClient) Incr(key string) int64 {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Incr().Key(key).Build()).AsInt64()
	if err != nil {
		return 0
	}
	return val
}

func (rc *RueidisClient) HGet(key, field string) string {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Hget().Key(key).Field(field).Build()).ToString()
	if err != nil {
		return ""
	}
	return val
}

func (rc *RueidisClient) HMGet(key string, field ...string) []string {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Hmget().Key(key).Field(field...).Build()).AsStrSlice()
	if err != nil {
		return []string{}
	}
	return val
}

func (rc *RueidisClient) HGetAll(key string) map[string]string {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Hgetall().Key(key).Build()).AsStrMap()
	if err != nil {
		return map[string]string{}
	}
	return val
}

func (rc *RueidisClient) ZRange(key, min, max string) []string {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Zrange().Key(key).Min(min).Max(max).Build()).AsStrSlice()
	if err != nil {
		return []string{}
	}
	return val
}

func (rc *RueidisClient) ZRank(key, member string) int64 {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Zrank().Key(key).Member(member).Build()).AsInt64()
	if err != nil {
		return 0
	}
	return val
}

func (rc *RueidisClient) ZScore(key, member string) float64 {
	val, err := rc.client.Do(rc.ctx, rc.client.B().Zscore().Key(key).Member(member).Build()).AsFloat64()
	if err != nil {
		return 0
	}
	return val
}
