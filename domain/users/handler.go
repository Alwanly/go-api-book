package users

import (
	"go-codebase/infrastructure/authentication"
	"net/http"

	"go-codebase/infrastructure/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UseCase UserUseCase
}

func ConstructUserHandler(router *gin.Engine, usecase UserUseCase, jwtAuth authentication.IJwtAuth) UserHandler {
	handler := &UserHandler{
		UseCase: usecase,
	}

	v1 := router.Group("/api/v1/user")
	v1.POST("/login", authentication.BasicAuthMiddleware(), handler.Login)
	v1.POST("/register", authentication.BasicAuthMiddleware(), handler.Register)
	return *handler
}

// login godoc
// @Summary      Login
// @Description  Authenticate a user to get access token
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        data body LoginModel true "LoginModel data"
// @Router       /v1/user/login [post]
// @Security     BasicAuth
func (u *UserHandler) Login(c *gin.Context) {
	var model LoginModel
	if err := c.ShouldBindJSON(&model); err != nil {
		utils.WriteResponse(c, utils.WrapperReponse(http.StatusBadRequest, "Validate Failed", err.Error()))
		return
	}
	result := u.UseCase.Sign(c, model)
	utils.WriteResponse(c, result)
}

func (u *UserHandler) Register(c *gin.Context) {
	var model RegisterModel

	if err := c.ShouldBindJSON(&model); err != nil {
		utils.WriteResponse(c, utils.WrapperReponse(http.StatusBadRequest, "Validate Failed", err.Error()))
		return
	}

	result := u.UseCase.Register(c, model)
	utils.WriteResponse(c, result)
}
