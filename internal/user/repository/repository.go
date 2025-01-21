package repository

import (
	"context"

	"github.com/Alwanly/go-codebase/model"
	"github.com/Alwanly/go-codebase/pkg/database"
	"github.com/Alwanly/go-codebase/pkg/redis"
)

const ContextName = "Internal.User.Repository"

type (
	Repository struct {
		DB    database.IDBService
		Redis redis.IRedisService
	}

	IRepository interface {
		Login(ctx context.Context, username string) (*model.User, error)
		Register(ctx context.Context, model *model.User) (*model.User, error)
	}
)

func NewRepository(r Repository) IRepository {
	return &Repository{
		DB:    r.DB,
		Redis: r.Redis,
	}
}

func (r *Repository) Login(ctx context.Context, username string) (*model.User, error) {
	tx := r.DB.GetTransaction(ctx)

	var model model.User
	cmd := tx.Where("username = ?", username).First(&model).Error
	if cmd != nil {
		return nil, cmd
	}

	tx.Commit()

	return &model, nil
}

func (r *Repository) Register(ctx context.Context, model *model.User) (*model.User, error) {
	tx := r.DB.GetTransaction(ctx)

	cmd := tx.Create(model).Error
	if cmd != nil {
		return nil, cmd
	}

	return model, nil
}
