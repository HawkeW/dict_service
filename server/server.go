package server

import (
	"dictService/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Gin路由实例
var router *gin.Engine

func StartServer() {
	// 创建Gin路由实例
	router = gin.Default()

	// 参数读取中间件
	router.Use(logParamsMiddleware())

	registerRouter()

	// 启动HTTP服务器并监听8080端口
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("福娃五启动失败")
		return
	}
}

// 定义接口映射
var routers = map[string]gin.HandlerFunc{
	"/":                HandleHome,
	"/get_word_detail": GetWordDetail,
}

func registerRouter() {
	for path, handlerFunc := range routers {
		// 注册处理器函数，将根路径与处理函数关联起来
		router.Any(path, handlerFunc)
	}
}

func HandleHome(context *gin.Context) {
	_, err := context.Writer.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println("Success")
	}
}

func GetWordDetail(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(Params)

	result := controller.GetWordDetail(map[string]interface{}{
		"id":   parameters.QueryParams["id"],
		"dict": parameters.QueryParams["dict"],
	})

	c.JSON(200, result)
}
