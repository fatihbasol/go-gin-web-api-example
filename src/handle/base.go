package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	welcomeMessage = "Welcome to GoGinExample"
	author         = "fatihbasol"
	authorUrl      = "https://www.linkedin.com/in/fatihbasol/"
)

func InitRoutes(r *gin.Engine) {

	router := r.Group("/api")
	{
		router.GET("/user/:id", GetUser)
		router.GET("/user", GetUsers)
		router.POST("/user", PostUser)
		router.PUT("/user/:id", PutUser)

		router.GET("/phone/:id", GetPhone)
		router.GET("/phone", GetPhones)
		router.POST("/phone", PostPhone)
		router.PUT("/phone/:id", PutPhone)
	}

	r.LoadHTMLGlob("templates/*")
	var endpoints []string
	for _, info := range r.Routes() {
		endpoints = append(endpoints, "["+info.Method+"] "+info.Path)
	}

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "welcome.html", gin.H{
			"title":     welcomeMessage,
			"author":    author,
			"authorUrl": authorUrl,
			"endpoints": endpoints,
		})
	})
}
