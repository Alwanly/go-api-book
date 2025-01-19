package utils

import (
	"math"
	"net/http"
)

type JSONResult struct {
	Code       int             `json:"-"`
	StatusCode StatusCode      `json:"statusCode"`
	Message    string          `json:"message"`
	Meta       *PaginationMeta `json:"meta,omitempty"`
	Data       interface{}     `json:"data"`
}

type PaginationMeta struct {
	Page            int         `json:"page"`
	TotalData       int         `json:"totalData"`
	TotalPage       int         `json:"totalPage"`
	TotalDataOnPage int         `json:"totalDataOnPage"`
	MetaData        interface{} `json:"metadata,omitempty"`
}

func CreateStatusCode(code string) StatusCode {
	return StatusCode(code)
}

func ResponseSuccess(code int, data interface{}) JSONResult {
	return JSONResult{
		Code:       code,
		StatusCode: StatusCodeSuccess,
		Message:    "Success",
		Data:       data,
	}
}

func ResponseFailed(httpCode int, statusCode StatusCode, message string, data interface{}) JSONResult {
	return JSONResult{
		Code:       httpCode,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func ResponsePagination(page int, limit int, count int, total int, data interface{}, metaData interface{}) JSONResult {
	return JSONResult{
		Code:       http.StatusOK,
		StatusCode: StatusCodeSuccess,
		Message:    "Success",
		Data:       data,
		Meta: &PaginationMeta{
			Page:            page,
			TotalData:       total,
			TotalDataOnPage: count,
			TotalPage:       int(math.Ceil(float64(total) / float64(limit))),
			MetaData:        metaData,
		},
	}
}
