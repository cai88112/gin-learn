package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code string         `json:"code"`
	Result interface{} `json:"result"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = "4444"
	SUCCESS = "0000"
)

func Result(code string, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code string, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
