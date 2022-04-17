package main

import (
	"github.com/fatihbasol/GoGinExample/src/data"
	"github.com/fatihbasol/GoGinExample/src/handle"
	"github.com/gin-gonic/gin"
)

func main() {
	data.ConnectToDatabase()
	r := gin.Default()

	handle.InitRoutes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
