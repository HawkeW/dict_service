package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Params struct {
	QueryParams map[string]string
	PostParams  map[string]string
}

func logParamsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := Params{
			QueryParams: make(map[string]string),
			PostParams:  make(map[string]string),
		}

		// 读取GET参数
		for key, values := range c.Request.URL.Query() {
			params.QueryParams[key] = values[0]
		}

		// 读取POST参数
		err := c.Request.ParseForm()
		if err != nil {
			fmt.Println("解析POST参数失败:", err)
		}
		for key, values := range c.Request.PostForm {
			params.PostParams[key] = values[0]
		}

		// 将参数存储在上下文中
		c.Set("params", params)

		// 继续处理下一个中间件或路由处理函数
		c.Next()
	}
}
