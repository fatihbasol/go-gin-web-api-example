package main

import (
	"github.com/fatihbasol/GoGinExample/src/data"
	"github.com/fatihbasol/GoGinExample/src/handle"
	"github.com/fatihbasol/GoGinExample/src/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	data.ConnectToDatabase()
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Header())

	handle.InitRoutes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
