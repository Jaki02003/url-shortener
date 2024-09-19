package repositories

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type ICacheRepository interface {
	Set(key string, value interface{}, ttl time.Duration) error
	SetStruct(key string, value interface{}, ttl time.Duration) error
	Get(key string) (string, error)
	GetInt(key string) (int, error)
	GetTTL(key string) (time.Duration, error)
	GetStruct(key string, outputStruct interface{}) error
	GetIteratorForPattern(pattern string) (*redis.ScanIterator, error)
	Del(keys ...string) error
	DelPattern(pattern string) error
}
