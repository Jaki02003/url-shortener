package connections

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-redis-url-shortener/config"
	"go-redis-url-shortener/utils/log"
)

var redisClient *redis.Client

func NewRedisClient(redisConfig *config.RedisConfig) {
	log.Info("connecting to redis at ", redisConfig.Host, ":", redisConfig.Port, "...")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Pass,
		DB:       redisConfig.Db,
	})

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Error("failed to connect redis: ", err)
		panic(err.Error())
	}

	log.Info("redis connections successful...")
}

func RedisClient() *redis.Client {
	return redisClient
}
