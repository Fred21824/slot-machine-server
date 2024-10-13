package cache

import (
	"context"
	"encoding/json"
	"time"

	"slot-machine-server/internal/logger"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("Failed to connect to Redis", zap.Error(err))
	} else {
		logger.Info("Connected to Redis")
	}
}

func Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, json, expiration).Err()
}

func Get(key string, dest interface{}) error {
	ctx := context.Background()
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}
