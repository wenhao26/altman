package response

import (
	"altman/common/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Success(c *gin.Context, data interface{}) {
	code := types.RespCode.OK
	msg := types.RespCode.GetMsg(code)
	Result(c, code, msg, data)
}

func Error(c *gin.Context, code int, data interface{}) {
	msg := types.RespCode.GetMsg(code)
	Result(c, code, msg, data)
}
