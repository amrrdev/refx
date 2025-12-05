package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewClient(redisConn string) (*RedisClient, error) {
	opts, err := redis.ParseURL(redisConn)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	return &RedisClient{
		Client: client,
	}, nil
}

func (c *RedisClient) SetLongUrl(ctx context.Context, shortUrl string, longUrl string) error {
	key := fmt.Sprintf("short:%s", shortUrl)
	return c.Client.Set(ctx, key, longUrl, time.Hour).Err()
}

func (c *RedisClient) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	key := fmt.Sprintf("short:%s", shortUrl)
	return c.Client.Get(ctx, key).Result()
}
