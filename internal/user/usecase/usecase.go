package usecase

import (
	"context"

	"net/http"
	"time"

	"github.com/Alwanly/go-codebase/config"
	"github.com/Alwanly/go-codebase/internal/user/repository"
	"github.com/Alwanly/go-codebase/internal/user/schema"
	"github.com/Alwanly/go-codebase/model"
	"github.com/Alwanly/go-codebase/pkg/authentication"
	"github.com/Alwanly/go-codebase/pkg/contract"
	"github.com/Alwanly/go-codebase/pkg/logger"
	"github.com/Alwanly/go-codebase/pkg/wrapper"
	"github.com/google/uuid"
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
		Auth(ctx context.Context, req *schema.AuthLoginRequest) wrapper.JSONResult
		Register(ctx context.Context, req *schema.AuthRegisterRequest) wrapper.JSONResult
		Profile(context.Context, *schema.ProfileRequest) wrapper.JSONResult
	}
)

func NewUseCase(uc UseCase) IUseCase {
	return &UseCase{
		Config:     uc.Config,
		Logger:     uc.Logger,
		Jwt:        uc.Jwt,
		Repository: uc.Repository,
	}
}

func (u *UseCase) Auth(ctx context.Context, req *schema.AuthLoginRequest) wrapper.JSONResult {
	l := logger.WithID(u.Logger, ContextName, "Auth")

	user, err := u.Repository.Login(ctx, req.Username)
	if err != nil {
		l.Error("username not found", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusBadRequest, contract.StatusCodeUserOrPasswordInvalid, "username or password invalid", nil)
	}

	if !authentication.VerifyPassword(req.Password, user.Password) {
		l.Error("password invalid", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusBadRequest, contract.StatusCodeUserOrPasswordInvalid, "username or password invalid", nil)
	}

	dataClaims := make(authentication.JWTClaims)

	dataClaims["userId"] = user.ID
	token, err := u.Jwt.GenerateToken(dataClaims)

	if err != nil {
		l.Error("failed to generate token", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusInternalServerError, contract.StatusCode("00000"), "failed to generate token", nil)
	}

	refreshToken, err := u.Jwt.RefreshToken(token)
	if err != nil {
		l.Error("failed to refresh token", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusInternalServerError, contract.StatusCode("00000"), "failed to refresh token", nil)
	}

	return wrapper.ResponseSuccess(http.StatusOK, schema.AuthLoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func (u *UseCase) Register(ctx context.Context, req *schema.AuthRegisterRequest) wrapper.JSONResult {
	l := logger.WithID(u.Logger, ContextName, "Register")

	hash, err := authentication.HashPassword(req.Password)
	if err != nil {
		l.Error("failed to hash password", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusInternalServerError, contract.StatusCode("00000"), "failed to hash password", nil)
	}

	id, err := uuid.NewV7()
	if err != nil {
		l.Error("failed to generate uuid", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusInternalServerError, contract.StatusCode("00000"), "failed to generate uuid", nil)
	}

	now := time.Now()
	model := &model.User{
		ID:        id.String(),
		Username:  req.Username,
		Password:  hash,
		CreatedAt: now,
	}
	user, err := u.Repository.Register(ctx, model)
	if err != nil {
		l.Error("failed to register", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusBadRequest, contract.CreateStatusCode("00001"), "failed to register", nil)
	}

	dataClaims := make(authentication.JWTClaims)

	dataClaims["userId"] = user.ID
	token, err := u.Jwt.GenerateToken(dataClaims)

	if err != nil {
		l.Error("failed to generate token", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusInternalServerError, contract.StatusCode("00000"), "failed to generate token", nil)
	}

	refreshToken, err := u.Jwt.RefreshToken(token)
	if err != nil {
		l.Error("failed to refresh token", zap.Error(err))
		return wrapper.ResponseFailed(http.StatusInternalServerError, contract.StatusCode("00000"), "failed to refresh token", nil)
	}

	return wrapper.ResponseSuccess(http.StatusOK, schema.AuthRegisterResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func (u *UseCase) Profile(_ context.Context, req *schema.ProfileRequest) wrapper.JSONResult {
	l := logger.WithID(u.Logger, ContextName, "Profile")
	l.Info("payload request", zap.Any("request", req))
	return wrapper.ResponseSuccess(http.StatusOK, req)
}
