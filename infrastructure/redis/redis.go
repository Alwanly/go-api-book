package redis

import (
	"books-api/infrastructure/config"
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
)

type Icache interface {
	Set(ctx context.Context, key string, value interface{}, expiration int) error
	Get(ctx context.Context, key string) ([]byte, error)
	Exist(ctx context.Context, key string) bool
	Delete(ctx context.Context, key string) error
	RemainingTime(ctx context.Context, key string) int
	Close() error
}

type Cache struct {
	client *redis.Client
	ns     string
}

func Initialize() (Icache, error) {
	url, _ := url.Parse(config.GlobalConfig.Redis.URI)
	p, _ := url.User.Password()
	rClient := redis.NewClient(&redis.Options{
		Addr:     url.Host,
		Password: p,
		DB:       0,
	})

	cache := &Cache{
		client: rClient,
		ns:     strings.TrimPrefix(url.Path, "/"),
	}

	_, err := cache.client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return cache, nil
}

// Set set value
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration int) error {
	switch value.(type) {
	case string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, []byte:
		return c.client.Set(ctx, c.ns+key, value, time.Duration(expiration)*time.Second).Err()
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return c.client.Set(ctx, c.ns+key, b, time.Duration(expiration)*time.Second).Err()
	}
}

// Get get value
func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	return c.client.Get(ctx, c.ns+key).Bytes()
}

// Exist check if key exist
func (c *Cache) Exist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, c.ns+key).Val() > 0
}

// Delete delete record
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, c.ns+key).Err()
}

// RemainingTime get remaining time
func (c *Cache) RemainingTime(ctx context.Context, key string) int {
	return int(c.client.TTL(ctx, c.ns+key).Val().Seconds())
}

// Close close connection
func (c *Cache) Close() error {
	return c.client.Close()
}
