package admin

import (
	"dictService/controller"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	userRouters := r.Group("admin/")
	{
		userRouters.GET("/user/", controller.UserController{}.Index)
		userRouters.GET("/user/get", controller.UserController{}.GetUserById)
		userRouters.GET("/user/add", controller.UserController{}.Add)
		userRouters.GET("/user/edit", controller.UserController{}.EditById)
		userRouters.GET("/user/delete", controller.UserController{}.DeleteById)
	}
}
