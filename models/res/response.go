package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type FailResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ListResponse struct {
	Code      int    `json:"code"`
	Data      any    `json:"data"`
	Msg       string `json:"msg"`
	TotalPage int64  `json:"total_page"`
	Total     int64  `json:"total"`
}

const (
	SUCCESS = 200
	ERROR   = -1
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailResult(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, FailResponse{
		Code: code,
		Msg:  msg,
	})
}

func ListResult(data any, total int64, total_page int64, c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{
		Code:      SUCCESS,
		Data:      data,
		TotalPage: total_page,
		Total:     total,
		Msg:       "OK",
	})
}

func OkWith(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "success", c)

}

func Fail(data any, msg string, c *gin.Context) {
	FailResult(ERROR, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	FailResult(ERROR, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		FailResult(int(code), msg, c)
		return
	} else {
		FailResult(ERROR, "未知错误", c)
	}
}
