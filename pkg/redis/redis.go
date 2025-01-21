package redis

import (
	"context"

	"github.com/Alwanly/go-codebase/pkg/logger"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
)

func NewRedis(opts *Opts) (*Service, error) {
	l := logger.WithID(opts.Logger, ContextName, "NewRedis")

	if opts.RedisURI == nil {
		l.Debug("Redis URI is not set, skipping")
		return nil, nil
	}

	// create redis client
	opt, _ := redis.ParseURL(*opts.RedisURI)
	cl := redis.NewClient(opt)

	// setup cancellation
	ctxTimeout, cancel := context.WithTimeout(context.Background(), PingTimeout)
	defer cancel()

	// ping redis
	if res := cl.Ping(ctxTimeout); res.Err() != nil {
		l.Error("Cannot ping redis")
		return nil, res.Err()
	}

	return &Service{
		Redis: cl,
	}, nil
}

func (db *Service) PingRedis() bool {
	l := logger.NewLogger(ContextName, "PingRedis")
	ctx, cancel := context.WithTimeout(context.Background(), PingTimeout)
	defer cancel()

	res := db.Redis.Ping(ctx)
	if res.Err() != nil {
		l.Error("Cannot check redis", zap.Error(res.Err()))
	}
	return res.Err() == nil
}

func (db *Service) PingRedisWithError() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PingTimeout)
	defer cancel()

	res := db.Redis.Ping(ctx)
	return res.Err() == nil, res.Err()
}

func (db *Service) CloseRedis() error {
	return db.Redis.Close()
}

func (db *Service) GetTransaction() (redis.Pipeliner, error) {
	return db.Redis.TxPipeline(), nil
}
