package res

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-blog-server/utils"
)

type Response struct {
	Code int    `json:"code" default:"200"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type FailResponse struct {
	Code int    `json:"code" default:"-1"`
	Msg  string `json:"msg"`
}

type ListResponse struct {
	Code      int    `json:"code" default:"200"`
	Data      any    `json:"data"`
	Msg       string `json:"msg" default:"success"`
	TotalPage int64  `json:"total_page" default:"2"`
	Total     int64  `json:"total" default:"20"`
}

const (
	SUCCESS = 200
	ERROR   = -1
)

func ResultWithNoData(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, FailResponse{
		Code: code,
		Msg:  msg,
	})
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
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

func OkWithNoData(msg string, c *gin.Context) {
	ResultWithNoData(SUCCESS, msg, c)
}

func Ok(c *gin.Context) {
	ResultWithNoData(SUCCESS, "success", c)
}

func Fail(data any, msg string, c *gin.Context) {
	ResultWithNoData(ERROR, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	ResultWithNoData(ERROR, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		ResultWithNoData(int(code), msg, c)
		return
	} else {
		ResultWithNoData(ERROR, "未知错误", c)
	}
}

// 提示 validate 校验错误
func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
