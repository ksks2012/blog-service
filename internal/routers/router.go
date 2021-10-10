package routers

import (
	"net/http"
	"time"

	"github.com/blog-service/global"

	v1 "github.com/blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/blog-service/docs"
	"github.com/blog-service/internal/middleware"
	"github.com/blog-service/internal/routers/api"
	"github.com/blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", api.GetAuth)

	article_api := v1.NewArticle()
	tag_api := v1.NewTag()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT()) //middleware.JWT()
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
