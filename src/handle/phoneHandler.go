package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPhone(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"phone": "getPhone " + id,
	})
}
func GetPhones(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"phone": "getPhones",
	})
}
func PostPhone(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"phone": "postPhone",
	})
}
func PutPhone(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"phone": "putPhone",
	})
}
