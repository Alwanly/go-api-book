package utils

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	HttpStatus int
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type BasePagination struct {
	HttpStatus int
	Message    string      `json:"message"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

func WrapperReponse(httpStatus int, message string, data interface{}) BaseResponse {
	return BaseResponse{
		HttpStatus: httpStatus,
		Message:    message,
		Data:       data,
	}
}

func WrapperPaginate(httpStatus int, message string, meta interface{}, data interface{}) BasePagination {
	return BasePagination{
		HttpStatus: httpStatus,
		Message:    message,
		Meta:       meta,
		Data:       data,
	}
}

func WriteResponse(c *gin.Context, response BaseResponse) {
	c.JSON(response.HttpStatus, response)
}

func WritePaginateResponse(c *gin.Context, response BasePagination) {
	c.JSON(response.HttpStatus, response)
}

func WriteAbortResponse(c *gin.Context, response BaseResponse) {
	c.AbortWithStatusJSON(response.HttpStatus, response)
}
