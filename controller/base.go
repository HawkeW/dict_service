package controller

import "github.com/gin-gonic/gin"

func returnResult(c *gin.Context, success bool, data interface{}, params ...string) {
	msg := ""
	code := 1
	if success {
		msg = "success"
		code = 0
	} else {
		msg = "request failed"
		code = 1
	}
	if len(params) > 0 {
		msg = params[0]
	}
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
