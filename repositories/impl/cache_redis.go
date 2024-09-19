package impl

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go-redis-url-shortener/repositories"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) repositories.ICacheRepository {
	return &redisRepository{client: client}
}

func (rs *redisRepository) Set(key string, value interface{}, ttl time.Duration) error {
	return rs.client.Set(context.Background(), key, value, ttl*time.Second).Err()
}

func (rs *redisRepository) SetStruct(key string, value interface{}, ttl time.Duration) error {
	serializedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return rs.client.Set(context.Background(), key, string(serializedValue), ttl*time.Second).Err()
}

func (rs *redisRepository) Get(key string) (string, error) {
	return rs.client.Get(context.Background(), key).Result()
}

func (rs *redisRepository) GetInt(key string) (int, error) {
	str, err := rs.client.Get(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str)
}

func (rs *redisRepository) GetTTL(key string) (time.Duration, error) {
	ttl, err := rs.client.TTL(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

func (rs *redisRepository) GetStruct(key string, outputStruct interface{}) error {
	serializedValue, err := rs.client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(serializedValue), &outputStruct); err != nil {
		return err
	}
	return nil
}

func (rs *redisRepository) GetIteratorForPattern(pattern string) (*redis.ScanIterator, error) {
	iterator := rs.client.Scan(context.Background(), 0, pattern, 0).Iterator()
	if err := iterator.Err(); err != nil {
		return nil, err
	}
	return iterator, nil
}

func (rs *redisRepository) Del(keys ...string) error {
	return rs.client.Del(context.Background(), keys...).Err()
}

func (rs *redisRepository) DelPattern(pattern string) error {
	iter := rs.client.Scan(context.Background(), 0, pattern, 0).Iterator()

	for iter.Next(context.Background()) {
		err := rs.client.Del(context.Background(), iter.Val()).Err()
		if err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
