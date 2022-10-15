package Infrastructure

import "time"

type IRedisService interface {
	Init(data any, key string)
	Store(data string, key string, expire time.Duration)
	Read(key string) string
	Remove(key string)
	Any(key string) bool
}
