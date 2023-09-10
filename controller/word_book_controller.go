package controller

import (
	"dictService/midllewares"
	"dictService/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type WordBookController struct {
}

// Index
// 获取所有词书
func (WordBookController) Index(c *gin.Context) {
	var wordBooks []models.WordBook
	models.DB.Find(&wordBooks)

	returnResult(c, true, wordBooks)
}

// GetDetailById
// 查询词书详情
func (WordBookController) GetDetailById(c *gin.Context) {
	var searchResult []interface{}
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	var wordBook models.WordBook
	result := models.DB.Find(&wordBook, parameters.QueryParams["id"])

	if result.RowsAffected == 0 {
		returnResult(c, false, searchResult, "未找到记录")
		return
	}
	if result.Error == nil {
		returnResult(c, true, wordBook)
	} else {
		returnResult(c, false, searchResult)
	}
}

// GetAllByUserId
// 查询用户所有的词书
func (WordBookController) GetAllByUserId(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	userId, err := strconv.Atoi(parameters.QueryParams["user_id"])
	if err != nil {
		returnResult(c, false, nil, "user_id格式错误")
		return
	}
	var wordBooks []models.WordBook
	result := models.DB.Where("user_id = ?", userId).Find(&wordBooks)
	if result.Error == nil {
		returnResult(c, true, wordBooks)
	} else {
		returnResult(c, false, nil)
	}
}

// Creat
// 用户创建词书
func (WordBookController) Creat(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	userId, err := strconv.Atoi(parameters.QueryParams["user_id"])
	if err != nil {
		returnResult(c, false, nil, "user_id格式错误")
		return
	}
	wordBook := models.WordBook{}.FromMap(userId, parameters.QueryParams)
	result := models.DB.Create(&wordBook)
	if result.Error == nil && result.RowsAffected > 0 {
		returnResult(c, true, wordBook)
	} else {
		returnResult(c, false, nil)
	}
}

// Delete
// 用户删除词书
func (WordBookController) Delete(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	userId, err := strconv.Atoi(parameters.QueryParams["user_id"])
	if err != nil {
		returnResult(c, false, nil, "user_id格式错误")
		return
	}
	wordBookId, err1 := strconv.Atoi(parameters.QueryParams["word_book_id"])
	if err1 != nil {
		returnResult(c, false, nil, "word_book_id格式错误")
		return
	}
	result := models.DB.Where("user_id = ?", userId).Delete(&models.WordBook{}, wordBookId)
	if result.Error == nil && result.RowsAffected > 0 {
		returnResult(c, true, nil)
	} else {
		returnResult(c, false, nil)
	}
}

// Edit
// 用户编辑词书
func (WordBookController) Edit(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	userId, err := strconv.Atoi(parameters.QueryParams["user_id"])
	if err != nil {
		returnResult(c, false, nil, "user_id格式错误")
		return
	}
	wordBookId, err1 := strconv.Atoi(parameters.QueryParams["word_book_id"])
	if err1 != nil {
		returnResult(c, false, nil, "word_book_id格式错误")
		return
	}
	var book models.WordBook
	models.DB.Where("user_id = ?", userId).Find(&book, wordBookId)
	if book.Id == 0 {
		returnResult(c, false, nil, "未找到数据")
		return
	}
	book = book.UpdateByStringMap(parameters.QueryParams)
	result := models.DB.Model(&book).Updates(book)
	if result.Error == nil && result.RowsAffected > 0 {
		returnResult(c, true, book)
	} else {
		returnResult(c, false, nil)
	}
}

type WordBookData struct {
	Id         uint
	UserId     int
	WordId     int
	WordBookId int
	Word       string
}

// GetWordList
//
//	在词书中查询单词
func (WordBookController) GetWordList(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	wordBookId, err := strconv.Atoi(parameters.QueryParams["word_book_id"])
	if err != nil {
		returnResult(c, false, nil, "word_book_id格式错误")
		return
	}

	// 词书数据
	bookResult := models.DB.Table("word_book_data").Where("word_book_id = ?", wordBookId)

	var list []models.Word
	bookResult.Raw(`
		SELECT
			Book.word_id,
			Dict.link_word_id,
			Dict.word,
			Dict.pron_uk,
			Dict.pron_us,
			Dict.captions
		FROM
			word_book_data AS Book
		INNER JOIN
   			colins_cn AS Dict ON Book.word_id = Dict.link_word_id;
		`).Scan(&list)

	returnResult(c, true, list)
}
