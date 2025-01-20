package usecase

import (
	"go-codebase/pkg/authentication"
	"go-codebase/pkg/config"

	"go.uber.org/zap"
)

const ContextName = "Internal.User.Usecase"

type (
	UseCase struct {
		Config *config.GlobalConfig
		Logger *zap.Logger
		Jwt    authentication.IJwtService
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

func (uc *UseCase) ExampleMethod() error {
	uc.Jwt.GenerateToken(authentication.JWTClaims{
		UserID: "123",
	})
	return nil
}
