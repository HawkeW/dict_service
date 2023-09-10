package controller

import (
	"dictService/midllewares"
	"github.com/gin-gonic/gin"
)

type WordController struct {
}

// GetWordDetail
//
//	在词典中查询单词
func (WordController) GetWordDetail(c *gin.Context) {
	success := false
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	if parameters.QueryParams["word"] == "" {
		returnResult(c, false, nil)
		return
	}

	result := map[string]interface{}{
		"word": parameters.QueryParams["word"],
		"dict": parameters.QueryParams["dict"],
	}
	if result != nil {
		success = true
	}

	returnResult(c, success, result)
}
