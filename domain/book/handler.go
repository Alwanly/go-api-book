package book

import (
	"books-api/infrastructure/authentication"
	"books-api/infrastructure/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHander struct {
	UseCase BookUseCase
}

func ConstructBookHanlder(router *gin.Engine, usecase BookUseCase, jwtAuth authentication.IJwtAuth) BookHander {
	handler := &BookHander{
		UseCase: usecase,
	}
	v1 := router.Group("/api/v1/books/")
	v1.POST("", jwtAuth.JwtAuthMidldleware(), handler.Create)
	v1.GET("", jwtAuth.JwtAuthMidldleware(), handler.GetAll)
	v1.GET(":id", jwtAuth.JwtAuthMidldleware(), handler.Get)
	return *handler
}

func (u *BookHander) Create(c *gin.Context) {
	var model CreateModel

	if err := c.Copy().ShouldBindJSON((&model)); err != nil {
		utils.WriteAbortResponse(c, utils.WrapperReponse(http.StatusBadRequest, "Validate fail", err))
		return
	}
	result := u.UseCase.Create(c, model)

	utils.WriteResponse(c, result)
}

func (u *BookHander) GetAll(c *gin.Context) {
	var model GetAllModel
	if err := c.Copy().ShouldBindQuery((&model)); err != nil {
		utils.WriteAbortResponse(c, utils.WrapperReponse(http.StatusBadRequest, "Validate fail", err))
		return
	}
	result := u.UseCase.GetAll(c, model)

	utils.WritePaginateResponse(c, result)
}

func (u *BookHander) Get(c *gin.Context) {
	var model GetModel

	if err := c.Copy().ShouldBindUri((&model)); err != nil {
		utils.WriteAbortResponse(c, utils.WrapperReponse(http.StatusBadRequest, "Validate fail", err))
		return
	}

	result := u.UseCase.Get(c, model)

	utils.WriteResponse(c, result)
}
