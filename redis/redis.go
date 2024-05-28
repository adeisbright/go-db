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
	ctx    context.Context
}

func ConnectToRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.GetConfig().RedisURL,
	})

	return redisClient, nil
}

func NewRedisFactory() (*RedisFactory, error) {
	client, err := ConnectToRedis()
	if err != nil {
		return nil, err
	}
	return &RedisFactory{
		client,
		context.Background(),
	}, nil
}

// Set a new value in the redis store
func (rf *RedisFactory) SetRedisValue(key string, value interface{}) error {
	return rf.client.Set(rf.ctx, key, value, 0).Err()
}
func (rf *RedisFactory) GetRedisValue(key string) (*redis.StringCmd, error) {
	return rf.client.Get(rf.ctx, key), nil
}

func (rf *RedisFactory) Ping() error {
	return rf.client.Ping(rf.ctx).Err()
}

func Init() {
	rf, err := NewRedisFactory()
	if err != nil {
		log.Fatal("Cannot Connect to Redis")
	}
	if err := rf.Ping(); err != nil {
		log.Fatal("Cannot Connect to Redis")
	}
	fmt.Println("Connected to Redis")
}
