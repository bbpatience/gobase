package common

import (
	"github.com/gin-gonic/gin"
	"gobase/global/consts"
	"net/http"
)

func returnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	if data == nil {
		Context.JSON(httpCode, gin.H{
			"code": dataCode,
			"msg":  msg,
		})
	} else {
		Context.JSON(httpCode, gin.H{
			"code": dataCode,
			"msg":  msg,
			"data": data,
		})
	}
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// 直接返回成功
func Success(c *gin.Context, data interface{}) {
	returnJson(c, http.StatusOK, consts.RspStatusOkCode, consts.RspStatusOkMsg, data)
}

// 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string) {
	returnJson(c, http.StatusOK, dataCode, msg, nil)
	c.Abort()
}

//权限校验失败
func ErrorAuthFail(c *gin.Context) {
	returnJson(c, http.StatusUnauthorized, consts.CommonAuthFailCode, consts.CommonAuthFailMsg, "")
	//暂停执行
	c.Abort()
}

//参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	returnJson(c, http.StatusBadRequest, consts.CommonParamsCheckFailCode, consts.CommonParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	returnJson(c, http.StatusInternalServerError, consts.CommonServerFailCode, consts.CommonServerFailMsg+msg, data)
	c.Abort()
}
