package users

import (
	"context"
	"errors"
	"go-codebase/infrastructure/authentication"
	"go-codebase/infrastructure/redis"
	"go-codebase/infrastructure/utils"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type UserUseCase interface {
	Register(c context.Context, model RegisterModel) utils.BaseResponse
	Sign(c context.Context, model LoginModel) utils.BaseResponse
}

type UserUseCaseImpl struct {
	Database *gorm.DB
	JwtAuth  authentication.IJwtAuth
	Cache    redis.Icache
}

func ConstructUserUseCase(db *gorm.DB, jwtAuth authentication.IJwtAuth, cache redis.Icache) UserUseCase {
	return &UserUseCaseImpl{
		Database: db,
		JwtAuth:  jwtAuth,
		Cache:    cache,
	}
}

func (u *UserUseCaseImpl) Sign(c context.Context, model LoginModel) utils.BaseResponse {

	db := u.Database.WithContext(c)
	var user Users
	dbResult := db.First(&user, "email = ? ", model.Email)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		return utils.WrapperReponse(http.StatusBadRequest, "User not found", nil)
	}
	// verify password
	if !authentication.VerifyPassword(model.Password, *user.Password) {
		return utils.WrapperReponse(http.StatusBadRequest, "Invalid password", nil)
	}

	accessToken, _ := u.JwtAuth.Sign(strconv.Itoa(int(user.ID)), user.Email)
	response := UserLoginDto{
		Token: accessToken,
	}
	defer u.Cache.Set(c, strconv.Itoa(int(user.ID)), response, 1000)

	return utils.WrapperReponse(http.StatusOK, "User Registered", response)

}

func (u *UserUseCaseImpl) Register(c context.Context, model RegisterModel) utils.BaseResponse {
	db := u.Database.WithContext(c)

	var user Users

	dbResult := db.First(&user, "email=?", model.Email)

	if !errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		return utils.WrapperReponse(http.StatusBadRequest, "User already registered", nil)
	}

	hashedPassword, _ := authentication.HashPassword(model.Password)

	user = Users{
		Email:    model.Email,
		Password: &hashedPassword,
		Name:     model.Name,
	}

	db.Create(&user)

	accessToken, _ := u.JwtAuth.Sign(strconv.Itoa(int(user.ID)), model.Email)
	response := UserLoginDto{
		Token: accessToken,
	}
	return utils.WrapperReponse(http.StatusOK, "User Registered", response)
}
