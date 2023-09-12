package routers

import (
	"dictService/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	wordRouters := r.Group("/user")
	{
		wordRouters.POST("/login", controller.UserController{}.Login)
	}
}
