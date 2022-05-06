package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 创建成功的响应 响应码 200
func SuccessResponse(data any) *Response {
	return &Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

// 创建失败的响应 400
func FailureResponse(msg string) *Response {
	return &Response{
		Code: 400,
		Msg:  msg,
	}
}

// 创建错误的响应 500
func ErrorResponse(err error) *Response {
	return &Response{
		Code: 500,
		Msg:  err.Error(),
	}
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r)
}
