package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/adeisbright/go-db/config"
	"github.com/redis/go-redis/v9"
)

type RedisFactory struct {
	client *redis.Client
}

func ConnectToRedis() *redis.Client {
	fmt.Println(config.GetConfig().RedisURL)
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.GetConfig().RedisURL,
	})

	return redisClient
}

func NewRedisFactory() *RedisFactory {
	return &RedisFactory{
		client: ConnectToRedis(),
	}
}

// Set a new value in the redis store
func (rf *RedisFactory) SetRedisValue(key string, value interface{}) error {
	ctx := context.Background()
	return rf.client.Set(ctx, key, value, 0).Err()
}
func (rf *RedisFactory) GetRedisValue(key string) (*redis.StringCmd, error) {
	ctx := context.Background()
	return rf.client.Get(ctx, key), nil
}

func (rf *RedisFactory) Ping() error {
	ctx := context.Background()
	return rf.client.Ping(ctx).Err()
}

func Init() {
	rf := NewRedisFactory()
	if err := rf.Ping(); err != nil {
		log.Fatal("Cannot Connect to Redis")
	}
	fmt.Println("Connected to Redis")
}
