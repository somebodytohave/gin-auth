package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mecm/gin-auth/pkg/e"
	"net/http"
	"strconv"
)

// Gin 实体
type Gin struct {
	C *gin.Context
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"msg" example:"ok" `
	Data    string `json:"data" example:"null"`
}

// GetGin 获取Gin
func GetGin(c *gin.Context) Gin {
	return Gin{c}
}

// Response 返回的数据
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}

// ResponseSuc 返回成功
func (g *Gin) ResponseSuc(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
	return
}

// ResponseFail 返回失败
func (g *Gin) ResponseFail() {
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ERROR,
		"msg":  e.GetMsg(e.ERROR),
		"data": nil,
	})
	return
}

// ResponseFailErrCode 返回失败
func (g *Gin) ResponseFailErrCode(errCode int) {
	errMsg := "code : " + strconv.Itoa(errCode) + "msg : " + e.GetMsg(errCode)
	MarkError(errMsg)

	g.C.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": nil,
	})
	return
}

// ResponseFailMsg 返回失败
func (g *Gin) ResponseFailMsg(msg string) {
	MarkError(msg)
	g.C.JSON(http.StatusOK, gin.H{
		"code": http.StatusBadRequest,
		"msg":  msg,
		"data": nil,
	})
	return
}
