package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
)

// UpLoad
// @Tags 文件上传模块
// @Summary 图片上传
// @Param file formData file true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/upload [post]
func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"url":    url,
	})

}
