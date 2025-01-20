package repository

import (
	"go-codebase/pkg/database"
	"go-codebase/pkg/redis"
)

const ContextName = "Internal.User.Repository"

type (
	Repository struct {
		DB    database.IDBService
		Redis redis.IRedisService
	}

	IRepository interface {
		// Define methods for the repository
		ExampleMethod() error
	}
)

func NewRepository(r Repository) IRepository {
	return &Repository{
		DB:    r.DB,
		Redis: r.Redis,
	}
}

func (r *Repository) ExampleMethod() error {
	return nil
}
