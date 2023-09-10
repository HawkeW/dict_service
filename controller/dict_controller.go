package controller

import (
	"dictService/models"
	"github.com/gin-gonic/gin"
)

type DictController struct {
}

var DefaultDictId = "colins_cn" // 默认词典

// GetDictList
//
//	获取所有词典列表
func (dict DictController) GetDictList(c *gin.Context) {
	var dictList []models.Dict
	models.DB.Find(&dictList)

	returnResult(c, true, dictList)
}

// ExistDictId
// 存在词典Id
func (DictController) ExistDictId(dictId string) bool {
	exist := false
	var dict models.Dict
	dictResult := models.DB.Where("dict_id = ?", dictId).First(&dict)
	if dictResult.Error == nil {
		exist = true
	}
	return exist
}
