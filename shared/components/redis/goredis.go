package redis

import (
	"context"
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
	onceGoRedis   sync.Once
	goredisClient *GoRedisClient
)

func goredis_init() {
	onceGoRedis.Do(func() {
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

		log.Infof("redis_init() %+v", conf)

		client := redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:      conf.Addrs,
			DB:         conf.Database,
			MaxRetries: conf.MaxRetries,
		})

		goredisClient = NewRedisClient(context.Background(), client)
	})

}

// DB 获取数据库
func GoRedisRD() *GoRedisClient {
	goredis_init()
	return goredisClient
}

type GoRedisClient struct {
	ctx    context.Context
	client redis.UniversalClient // Replace with the actual Redis client type
}

func NewRedisClient(ctx context.Context, client redis.UniversalClient) *GoRedisClient {
	return &GoRedisClient{
		ctx:    ctx,
		client: client,
	}
}

func (rc *GoRedisClient) Get(key string) string {
	val, err := rc.client.Get(rc.ctx, key).Result()
	if err != nil {
		log.Infof("redis.Get(%s) failed: %v", key, err)
		return ""
	}
	return val
}

func (rc *GoRedisClient) MGet(keys ...string) []string {
	val, err := rc.client.MGet(rc.ctx, keys...).Result()
	if err != nil {
		return []string{}
	}
	return convertToSliceOfStrings(val)
}

func (rc *GoRedisClient) Set(key, value string) bool {
	err := rc.client.Set(rc.ctx, key, value, 0).Err()
	if err != nil {
		log.Infof("redis.Set(%s) failed: %v", key, err)
		return false
	}
	return true
}

func (rc *GoRedisClient) Incr(key string) int64 {
	val, err := rc.client.Incr(rc.ctx, key).Result()
	if err != nil {
		return 0
	}
	return val
}

func (rc *GoRedisClient) HGet(key, field string) string {
	val, err := rc.client.HGet(rc.ctx, key, field).Result()
	if err != nil {
		return ""
	}
	return val
}

func (rc *GoRedisClient) HMGet(key string, field ...string) []string {
	val, err := rc.client.HMGet(rc.ctx, key, field...).Result()
	if err != nil {
		return []string{}
	}
	return convertToSliceOfStrings(val)
}

func (rc *GoRedisClient) HGetAll(key string) map[string]string {
	val, err := rc.client.HGetAll(rc.ctx, key).Result()
	if err != nil {
		return map[string]string{}
	}
	return val
}

func (rc *GoRedisClient) ZRange(key string, min, max int64) []string {
	val, err := rc.client.ZRange(rc.ctx, key, min, max).Result()
	if err != nil {
		return []string{}
	}
	return val
}

func (rc *GoRedisClient) ZRank(key, member string) int64 {
	val, err := rc.client.ZRank(rc.ctx, key, member).Result()
	if err != nil {
		return 0
	}
	return val
}

func (rc *GoRedisClient) ZScore(key, member string) float64 {
	val, err := rc.client.ZScore(rc.ctx, key, member).Result()
	if err != nil {
		return 0
	}
	return val
}

func convertToSliceOfStrings(input []interface{}) []string {
	result := make([]string, len(input))

	for i, v := range input {
		// 使用类型断言将interface{}转换为string
		str, ok := v.(string)
		if !ok {
			// 如果类型断言失败，可以根据实际需求进行处理，这里简单地打印错误信息
			log.Errorf("Error: Unable to convert element at index %d to string\n", i)
			continue
		}

		result[i] = str
	}

	return result
}
