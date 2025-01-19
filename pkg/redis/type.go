package redis

import (
	"time"

	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
)

const ContextName = "Components.Database"

const (
	PingTimeout = 10 * time.Second
)

// DBServiceOpts represents the options for configuring the database service.
type DBServiceOpts struct {
	// Debug enables debug mode.
	Debug bool
	// Logger is the logger.
	Logger *zap.Logger

	// Redis database connection string (DSN)
	RedisUri *string

	// Application Name (for tracing)
	ApplicationName *string
}

// DBService represents the database service.
type DBService struct {
	Redis redis.UniversalClient
}

// IDBService represents the interface for the database service.
type IDBService interface {

	// ---- Redis

	// PingRedis pings the Redis database to check if it's available.
	//
	// Returns:
	//   - bool: true if the database is available, false otherwise.
	PingRedis() bool

	// PingRedis pings the Redis database to check if it's available.
	//
	// Returns:
	//   - bool: true if the database is available, false otherwise.
	//   - error: error stack trace.
	PingRedisWithError() (bool, error)
}
