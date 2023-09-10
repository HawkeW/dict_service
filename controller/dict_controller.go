package controller

import (
	"github.com/gin-gonic/gin"
)

type DictController struct {
}

// GetDictList
//
//	获取所有词典列表
func (dict DictController) GetDictList(c *gin.Context) {
	result := []string{"colins_cn"}
	returnResult(c, true, result)
}
