package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "go_blog/api/v1"
	_ "go_blog/docs"
	"go_blog/middleware"
	"go_blog/utils"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口

		auth.PUT("/user-edit", v1.EditUser)
		auth.DELETE("/user-delete", v1.DeleteUser)
		// 分类模块的接口
		auth.POST("/category-add", v1.AddCategory)
		auth.PUT("/category-edit", v1.EditCategory)
		auth.DELETE("/category/:id", v1.DeleteCategory)
		// 文章模块的接口
		auth.POST("/article-add", v1.AddArticle)
		auth.PUT("/article-edit", v1.EditArticle)
		auth.DELETE("/article/:id", v1.DeleteArticle)
		// 上次文件
		auth.POST("/upload", v1.UpLoad)
	}
	router := r.Group("api/v1")
	{
		router.POST("/user-add", v1.AddUser)
		router.GET("/user-id", v1.SelectUser)
		router.GET("/users", v1.GetUsers)
		router.GET("/categories", v1.GetCategories)
		router.GET("/articles", v1.GetArticles)
		router.GET("/article", v1.GetArtInfo)
		router.GET("/cate-art", v1.GetCateArt)
		router.POST("login", v1.Login)
	}
	return r

}
