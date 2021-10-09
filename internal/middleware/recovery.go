package middleware

import (
	"fmt"
	"time"

	"github.com/blog-service/pkg/email"

	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/errcode"

	"github.com/blog-service/global"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
