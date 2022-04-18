package handle

import (
	"net/http"
	"strconv"

	"github.com/fatihbasol/GoGinExample/src/data"
	"github.com/fatihbasol/GoGinExample/src/handle/model/response"
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
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	context.JSON(http.StatusOK,
		response.PhoneResponseModel{
			Id:          phone.Id,
			UserId:      phone.UserId,
			Number:      phone.Number,
			CountryCode: phone.CountryCode,
		})

}
func GetPhones(context *gin.Context) {

	phones := []model.Phone{}
	result := data.DB.Find(&phones)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	responseModel := []response.PhoneResponseModel{}

	for _, v := range phones {

		responseModel = append(responseModel, response.PhoneResponseModel{
			Id:          v.Id,
			UserId:      v.UserId,
			Number:      v.Number,
			CountryCode: v.CountryCode,
		})
	}

	context.JSON(http.StatusOK, responseModel)

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
