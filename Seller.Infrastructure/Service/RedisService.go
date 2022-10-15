package Service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisService struct {
	RedisClinet *redis.Client
}

func (service RedisService) Init(data any, key string) {
	service.RedisClinet = redis.NewClient(&redis.Options{
		Addr:     "185.231.115.166",
		Password: "10002", // no password set
		DB:       0,       // use default DB
	})

}

func (service RedisService) Store(data string, key string, expire time.Duration) {
	ctx := context.Background()
	err := service.RedisClinet.Set(ctx, key, data, expire).Err()
	if err != nil {
		panic(err)
	}
}
func (service RedisService) Read(key string) string {
	ctx := context.Background()
	val, err := service.RedisClinet.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func (service RedisService) Remove(key string) {
	ctx := context.Background()
	err := service.RedisClinet.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
}
func (service RedisService) Any(key string) bool {
	ctx := context.Background()
	val, err := service.RedisClinet.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val == 1
}
