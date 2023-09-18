package controller

import (
	"dictService/midllewares"
	"dictService/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"strconv"
)

type UserController struct {
}

// Index
// 所有用户
func (con UserController) Index(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	returnResult(c, true, users)
}

// GetUserById
// 查找用户
func (con UserController) GetUserById(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	id, err := strconv.Atoi(parameters.QueryParams["id"])
	if err != nil {
		id = 0
	}
	var users []models.User
	models.DB.Where("id = ?", id).Find(&users)

	if len(users) > 0 {
		returnResult(c, true, users[0])
	} else {
		returnResult(c, false, nil, " 未找到用户")
	}
}

// Add
// 添加用户
func (con UserController) Add(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	sex, err := strconv.Atoi(parameters.QueryParams["sex"])
	if err != nil {
		sex = 0
	}
	user := models.User{
		NickName: parameters.QueryParams["nick_name"],
		Phone:    parameters.QueryParams["phone"],
		Password: parameters.QueryParams["password"],
		Email:    parameters.QueryParams["email"],
		Sex:      sex,
		Name:     parameters.QueryParams["name"],
	}
	// 校验手机号是否注册
	var users []models.User
	models.DB.Where("phone = ?", user.Phone).Find(&users)
	if len(users) > 0 {
		returnResult(c, false, nil, "手机号已注册")
		return
	}

	result := models.DB.Create(&user)

	if result.Error == nil {
		returnResult(c, true, user)
	} else {
		returnResult(c, false, nil)
	}
}

// EditById
// 更新用户
func (con UserController) EditById(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	id, err := strconv.Atoi(parameters.QueryParams["id"])
	if err != nil {
		returnResult(c, false, nil, "id格式错误")
		return
	}
	var user models.User
	models.DB.Where("id = ?", id).First(&user)
	delete(parameters.QueryParams, "id")
	user.UpdateByStringMap(parameters.QueryParams)
	result := models.DB.Model(&user).Updates(user)

	if result.Error == nil {
		returnResult(c, true, user, "更新成功！")
	} else {
		returnResult(c, true, user, result.Error.Error())
	}
}

// DeleteById
// 删除用户
func (con UserController) DeleteById(c *gin.Context) {
	params, _ := c.Get("params")
	parameters := params.(midllewares.Params)
	id, err := strconv.Atoi(parameters.QueryParams["id"])
	if err != nil {
		returnResult(c, false, nil, "id 错误")
		return
	}
	var users []models.User
	models.DB.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&users)
	if len(users) > 0 {
		returnResult(c, true, users[0], "删除成功！")
	} else {
		returnResult(c, false, nil, "未找到用户！")
	}
}

// Login
// 用户登录
func (con UserController) Login(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		returnResult(c, false, nil, err.Error())
		return
	}
	models.DB.Clauses(clause.Returning{}).Where("phone = ?", user.Phone).Where("password = ?", user.Password).First(&user)
	if user.Id > 0 {
		returnResult(c, true, user, "成功！")
	} else {
		returnResult(c, false, nil, "失败！")
	}
}
