package controller

import (
	"dictService/midllewares"
	"dictService/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type WordController struct {
}

// GetWordDetail
//
//	在词典中查询单词
func (WordController) GetWordDetail(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	wordStr := parameters.QueryParams["word"]
	if wordStr == "" {
		returnResult(c, false, nil, "单词不能为空！")
		return
	}

	dictId := DefaultDictId
	// 若传入词典, 则验证词典有效性
	searchedDict := parameters.QueryParams["dict"]
	if (DictController{}.ExistDictId(searchedDict)) {
		dictId = searchedDict
	}

	word := models.Word{
		DictId: dictId,
	}
	result := models.DB.Table(dictId).Where("word = ?", wordStr).First(&word)

	if result.Error == nil {
		returnResult(c, true, word, "查询成功！")
	} else {
		returnResult(c, false, nil, "查询成功！")
	}
}

// SearchWordList
//
//	在词典中搜索单词
func (WordController) SearchWordList(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	wordStr := parameters.QueryParams["word"]
	if wordStr == "" {
		returnResult(c, false, nil, "单词不能为空！")
		return
	}

	dictId := DefaultDictId
	// 若传入词典, 则验证词典有效性
	searchedDict := parameters.QueryParams["dict"]
	if (DictController{}.ExistDictId(searchedDict)) {
		dictId = searchedDict
	}

	var word []models.Word
	search := fmt.Sprintf("%%%s%%", wordStr)
	result := models.DB.Table(dictId).Limit(10).Where("word LIKE ?", search).Find(&word)

	for i := 0; i < len(word); i++ {
		word[i].DictId = dictId
	}
	if result.Error == nil {
		returnResult(c, true, word, "查询成功！")
	} else {
		returnResult(c, false, nil, "查询成功！")
	}
}
