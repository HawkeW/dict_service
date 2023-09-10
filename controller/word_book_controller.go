package controller

import (
	"dictService/models"
	"github.com/gin-gonic/gin"
)

type WordBookController struct {
}

func (WordBookController) Index(c *gin.Context) {
	var wordBooks []models.WordBook
	models.DB.Find(&wordBooks)

	returnResult(c, true, wordBooks)
}
