package routers

import (
	"dictService/controller"
	"github.com/gin-gonic/gin"
)

func WordRoutersInit(r *gin.Engine) {
	wordRouters := r.Group("/word")
	{
		wordRouters.GET("/detail", controller.WordController{}.GetWordDetail)
		wordRouters.GET("/search", controller.WordController{}.SearchWordList)
	}
}
