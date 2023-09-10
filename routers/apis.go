package routers

import (
	"dictService/midllewares"
	"dictService/routers/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Gin路由实例
var router *gin.Engine

func StartServer() {
	// 创建Gin路由实例
	router = gin.Default()

	// 参数读取中间件
	router.Use(midllewares.LogParamsMiddleware())

	// 分组路由注册
	admin.RouterInit(router)
	DictRoutersInit(router)
	WordRoutersInit(router)
	WordBookRoutersInit(router)

	registerRouter()

	// 启动HTTP服务器并监听8080端口
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("服务启动失败")
		return
	}
}

// 定义接口映射
var routers = map[string]gin.HandlerFunc{
	"/": HandleHome,
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
