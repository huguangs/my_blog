package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddArticle
// @Tags 文章模块
// @Summary 添加文章
// @Param title formData string false "title"
// @Param cid formData int false "cid"
// @Param desc formData string false "desc"
// @Param content formData string false "content"
// @Param img formData string false "img"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/article-add [post]
func AddArticle(c *gin.Context) {
	var data model.Article
	data.Title = c.PostForm("title")
	data.Cid, _ = strconv.Atoi(c.PostForm("cid"))
	data.Desc = c.PostForm("desc")
	data.Content = c.PostForm("content")
	data.Img = c.PostForm("img")
	code = model.CreateArt(&data)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetArtInfo
// @Tags 文章模块
// @Summary 查询单个文章
// @Param id query int true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/article [get]
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// GetCateArt
// @Tags 文章模块
// @Summary 查询分类下的文章
// @Param cid query int true "cid"
// @Param pagesize query int false "pagesize"
// @Param page_num query int false "page_num"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/cate-art [get]
func GetCateArt(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Query("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetCateArt(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// GetArticles
// @Tags 文章模块
// @Summary 查询文章列表
// @Param pagesize query int false "pagesize"
// @Param page_num query int false "page_num"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArts(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// EditArticle
// @Tags 文章模块
// @Summary 编辑文章信息
// @Param id query int false "id"
// @Param title formData string false "title"
// @Param cid formData int false "cid"
// @Param desc formData string false "desc"
// @Param content formData string false "content"
// @Param img formData string false "img"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/article-edit [put]
func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var data model.Article
	data.Title = c.PostForm("title")
	data.Cid, _ = strconv.Atoi(c.PostForm("cid"))
	data.Desc = c.PostForm("desc")
	data.Content = c.PostForm("content")
	data.Img = c.PostForm("img")
	code := model.UpdateArt(id, &data)
	//if code == errmsg.ERROR_CATEGORY_USED {
	//	c.Abort()
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}

// DeleteArticle
// @Tags 文章模块
// @Summary 删除文章
// @Param id  query  string false "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/article/:id [delete]
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
