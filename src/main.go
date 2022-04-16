package main

import (
	"github.com/fatihbasol/GoGinExample/src/handle"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handle.InitRoutes(r)

	r.Run()
}
