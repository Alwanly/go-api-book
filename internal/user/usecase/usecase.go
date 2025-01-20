package usecase

import (
	"go-codebase/internal/user/repository"
	"go-codebase/pkg/authentication"
	"go-codebase/pkg/config"

	"go.uber.org/zap"
)

const ContextName = "Internal.User.Usecase"

type (
	UseCase struct {
		Config     *config.GlobalConfig
		Logger     *zap.Logger
		Jwt        authentication.IJwtService
		Repository repository.IRepository
	}

	IUseCase interface {
		// Define methods for the use case
		ExampleMethod() error
	}
)

func NewUseCase(uc UseCase) IUseCase {
	return &UseCase{
		Config: uc.Config,
		Logger: uc.Logger,
	}
}

func (u *UseCase) ExampleMethod() error {
	u.Jwt.GenerateToken(authentication.JWTClaims{
		UserID: "123",
	})

	u.Repository.ExampleMethod()
	return nil
}
