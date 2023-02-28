package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"
)

var code int

// SelectUser
// @Tags 用户模块
// @Summary 查询用户
// @Param id query int false "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/user-id [get]
func SelectUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, err := model.SelectUser(id)
	if err != nil {
		code = errmsg.ERROR_USER_NOT_EXIST
	} else {
		code = errmsg.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// AddUser
// @Tags 用户模块
// @Summary 添加用户
// @Param user_name formData string true "user_name"
// @Param pass_word formData string true "pass_word"
// @Param role formData string false "role"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/user-add [post]
func AddUser(c *gin.Context) {
	username := c.PostForm("user_name")
	password := c.PostForm("pass_word")
	role, _ := strconv.Atoi(c.PostForm("role"))
	fmt.Println(username, "sd")
	data := &model.User{
		UserName: username,
		PassWord: password,
		Role:     role,
	}
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		model.CreateUser(data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetUsers
// @Tags 用户模块
// @Summary 查询用户列表
// @Param pagesize query int false "pagesize"
// @Param page_num query int false "page_num"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// EditUser
// @Tags 用户模块
// @Summary 编辑用户信息
// @Param id query int false "id"
// @Param user_name formData string false "user_name"
// @Param role formData string false "role"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/user-edit [put]
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var data model.User
	data.UserName = c.PostForm("user_name")
	data.Role, _ = strconv.Atoi(c.PostForm("role"))
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		model.UpdateUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}

// DeleteUser
// @Tags 用户模块
// @Summary 删除用户
// @Param id  query  string false "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/user-delete [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Println(id)
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
