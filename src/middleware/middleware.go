package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var supportedCultures = []string{"tr-TR", "en-US"}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t)
		log.Println(ctx.Writer.Status(), " | ", latency, " | "+ctx.Request.URL.Path)
	}
}

func Header() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Accept-Language", strings.Join(supportedCultures, "; "))
		ctx.Next()
	}
}
