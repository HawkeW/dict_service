package routers

import (
	"dictService/controller"
	"github.com/gin-gonic/gin"
)

func DictRoutersInit(r *gin.Engine) {
	dictRouters := r.Group("/dict")
	{
		dictRouters.GET("/all", controller.DictController{}.GetDictList)
	}
}
