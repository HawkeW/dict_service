package routers

import (
	"dictService/controller"
	"github.com/gin-gonic/gin"
)

func WordBookRoutersInit(r *gin.Engine) {
	wordRouters := r.Group("/word_book")
	{
		wordRouters.GET("/", controller.WordBookController{}.Index)
	}
}
