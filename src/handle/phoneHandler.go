package handle

import (
	"net/http"
	"strconv"

	"github.com/fatihbasol/GoGinExample/src/data"
	"github.com/fatihbasol/GoGinExample/src/handle/model/response"
	"github.com/fatihbasol/GoGinExample/src/mapper"
	"github.com/fatihbasol/GoGinExample/src/model"
	"github.com/gin-gonic/gin"
)

func GetPhone(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	phone := &model.Phone{}
	result := data.DB.First(&phone, id)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	responseModel := mapper.MarshalMap[response.PhoneResponseModel](phone)

	context.JSON(http.StatusOK, responseModel)
}
func GetPhones(context *gin.Context) {

	phones := []model.Phone{}
	result := data.DB.Find(&phones)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": "Fields cannot be empty"})
		return
	}

	responseModel := mapper.MarshalMap[[]response.PhoneResponseModel](phones)

	context.JSON(http.StatusOK, responseModel)

}
func PostPhone(context *gin.Context) {
	//TODO:
	context.JSON(http.StatusOK, gin.H{
		"phone": "postPhone",
	})
}
func PutPhone(context *gin.Context) {
	//TODO:
	context.JSON(http.StatusOK, gin.H{
		"phone": "putPhone",
	})
}
