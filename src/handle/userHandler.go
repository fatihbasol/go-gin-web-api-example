package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"user": "getUser " + id,
	})
}
func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"user": "getUsers",
	})
}
func PostUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"user": "postUser",
	})
}
func PutUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"user": "putUser",
	})
}
