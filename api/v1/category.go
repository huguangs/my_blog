package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddCategory
// @Tags 分类模块
// @Summary 添加分类
// @Param name formData string true "name"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/category-add [post]
func AddCategory(c *gin.Context) {
	name := c.PostForm("name")
	data := &model.Category{
		Name: name,
	}
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetCategories
// @Tags 分类模块
// @Summary 查询分类列表
// @Param pagesize query int false "pagesize"
// @Param page_num query int false "page_num"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/categories [get]
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategories(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// EditCategory
// @Tags 分类模块
// @Summary 编辑分类信息
// @Param id query int false "id"
// @Param name formData string false "name"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/category-edit [put]
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var data model.Category
	data.Name = c.PostForm("name")
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.UpdateCategory(id, &data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}

// DeleteCategory
// @Tags 分类模块
// @Summary 删除分类
// @Param id  query  string false "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/category/:id [delete]
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
