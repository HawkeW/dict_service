package controller

import (
	"dictService/midllewares"
	"dictService/models"
	"encoding/json"
	"fmt"
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

// DeleteList
// 用户删除词书数组
func (WordBookController) DeleteList(c *gin.Context) {
	type reqParam struct {
		Ids    []int `json:"ids"`
		UserId int   `json:"user_id"`
	}
	var req reqParam
	err := json.NewDecoder(c.Request.Body).Decode(&req)

	if err != nil {
		returnResult(c, false, nil, "参数错误")
		return
	}
	result := models.DB.Where("user_id = ?", req.UserId).Delete(&models.WordBook{}, req.Ids)
	if result.Error == nil && result.RowsAffected > 0 {
		returnResult(c, true, nil)
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
	var list []models.Word
	searchSql := fmt.Sprintf(`
		SELECT
			Book.word_id,
			Dict.link_word_id,
			Dict.word,
			Dict.pron_uk,
			Dict.pron_us,
			Dict.captions
		FROM
			word_book_data AS Book
		JOIN 
			words AS WordBase ON Book.word_id = WordBase.id
		JOIN 
			colins_cn AS Dict ON WordBase.word = Dict.word
		WHERE 
			Book.word_book_id = %d;
		`, wordBookId)
	models.DB.Table("word_book_data").Raw(searchSql).Scan(&list)

	returnResult(c, true, list)
}

type ReqParam struct {
	UserId int      `json:"user_id"`
	BookId int      `json:"word_book_id"`
	Words  []string `json:"words"`
}

// AddWordList
// 添加单词到词书
func (con WordBookController) AddWordList(c *gin.Context) {
	var req ReqParam
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		returnResult(c, false, nil, err.Error())
		return
	}

	err1 := con.addWordsToWordBooks(req.UserId, req.BookId, req.Words)
	if err1 == nil {
		returnResult(c, true, len(req.Words))
	} else {
		returnResult(c, false, nil)
	}
}

// 查询单词对应的 word_id
func (WordBookController) findWordIDs(words []string) ([]uint, error) {
	var results []models.Word
	if err := models.DB.Model(&models.Word{}).Select("id").Where("word IN (?)", words).Find(&results).Error; err != nil {
		return nil, err
	}

	var wordIDs []uint
	for _, result := range results {
		wordIDs = append(wordIDs, result.Id)
	}

	return wordIDs, nil
}

// 向 word_books 表中添加单词
func (con WordBookController) addWordsToWordBooks(userId int, wordBookId int, words []string) error {
	wordIDs, err := con.findWordIDs(words)
	if err != nil {
		return err
	}

	var wordBookData []models.WordBookData
	for _, wordID := range wordIDs {
		wordBook := models.WordBookData{
			WordId:     wordID,
			UserId:     userId,
			WordBookId: wordBookId,
		}
		wordBookData = append(wordBookData, wordBook)
	}

	if err := models.DB.Model(&models.WordBookData{}).Create(&wordBookData).Error; err != nil {
		return err
	}

	return nil
}
