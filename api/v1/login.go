package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/middleware"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
)

// Login
// @Tags 用户模块
// @Summary 用户登录
// @Param user_name formData string true "user_name"
// @Param pass_word formData string true "pass_word"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	var data model.User
	data.UserName = c.PostForm("user_name")
	data.PassWord = c.PostForm("pass_word")
	var token string
	var code int
	code = model.CheckLogin(data.UserName, data.PassWord)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.UserName)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
