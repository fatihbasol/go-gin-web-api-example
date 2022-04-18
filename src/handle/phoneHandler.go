package handle

import (
	"net/http"
	"strconv"

	"github.com/fatihbasol/GoGinExample/src/data"
	"github.com/fatihbasol/GoGinExample/src/model"
	"github.com/gin-gonic/gin"
)

func GetPhone(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	phone := &model.Phone{}
	result := data.DB.First(&phone, id)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	context.JSON(http.StatusOK, phone)

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
