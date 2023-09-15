package routers

import (
	"dictService/controller"
	"github.com/gin-gonic/gin"
)

func WordBookRoutersInit(r *gin.Engine) {
	wordRouters := r.Group("/word_book")
	{
		wordRouters.GET("/", controller.WordBookController{}.Index)
		wordRouters.GET("/detail", controller.WordBookController{}.GetDetailById)
		wordRouters.GET("/by_user", controller.WordBookController{}.GetAllByUserId)
		wordRouters.GET("/words", controller.WordBookController{}.GetWordList)
		wordRouters.GET("/create", controller.WordBookController{}.Creat)
		wordRouters.GET("/update", controller.WordBookController{}.Edit)
		wordRouters.GET("/delete", controller.WordBookController{}.Delete)
		wordRouters.POST("/delete/list", controller.WordBookController{}.DeleteList)
		wordRouters.POST("/add_words", controller.WordBookController{}.AddWordList)

	}
}
