package routers

import (
	"github.com/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article_api := v1.NewArticle()
	tag_api := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	apiv1.Use() //middleware.JWT()
	{
		// 創建標籤
		apiv1.POST("/tags", tag_api.Create)
		// 刪除指定標籤
		apiv1.DELETE("/tags/:id", tag_api.Delete)
		// 更新指定標籤
		apiv1.PUT("/tags/:id", tag_api.Update)
		// 獲取標籤列表
		apiv1.GET("/tags", tag_api.List)

		// 創建文章
		apiv1.POST("/articles", article_api.Create)
		// 刪除指定文章
		apiv1.DELETE("/articles/:id", article_api.Delete)
		// 更新指定文章
		apiv1.PUT("/articles/:id", article_api.Update)
		// 獲取指定文章
		apiv1.GET("/articles/:id", article_api.Get)
		// 獲取文章列表
		apiv1.GET("/articles", article_api.List)
	}

	return r
}
