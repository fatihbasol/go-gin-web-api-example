package handle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatihbasol/GoGinExample/src/data"
	"github.com/fatihbasol/GoGinExample/src/handle/model/request"
	"github.com/fatihbasol/GoGinExample/src/handle/model/response"
	"github.com/fatihbasol/GoGinExample/src/mapper"
	"github.com/fatihbasol/GoGinExample/src/model"
	"github.com/gin-gonic/gin"
)

func GetUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var user model.User
	result := data.DB.Preload("Phone").First(&user, id)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"err": result.Error.Error()})
		return
	}

	responseModel := mapper.MarshalMap[response.UserResponseModel](user)

	context.JSON(http.StatusOK, responseModel)
}
func GetUsers(context *gin.Context) {

	var users []model.User
	result := data.DB.Preload("Phone").Find(&users)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"err": result.Error.Error()})
		return
	}

	usersResponseModel := mapper.MarshalMap[[]response.UserResponseModel](users)

	context.JSON(http.StatusOK, usersResponseModel)
}
func PostUser(context *gin.Context) {
	requestModel := request.UserRequestModel{Name: context.PostForm("name"), Email: context.PostForm("email")}

	if requestModel.Name == "" || requestModel.Email == "" {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var dbUser = mapper.MarshalMap[model.User](requestModel)
	result := data.DB.Create(&dbUser)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"err": result.Error.Error()})
		return
	}

	responseModel := mapper.MarshalMap[response.UserResponseModel](dbUser)
	context.JSON(http.StatusOK, responseModel)

}
func PutUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": "invalid id"})
		return
	}

	requestModel := request.UserRequestModel{Name: context.PostForm("name"), Email: context.PostForm("email")}

	if requestModel.Name == "" || requestModel.Email == "" {
		context.JSON(http.StatusBadRequest, gin.H{"err": "Fields cannot be empty"})
		return
	}

	var dbUser *model.User
	result := data.DB.First(&dbUser, id)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": "user not found"})
		return
	}

	userToUpdate := mapper.MarshalMap[model.User](requestModel)
	fmt.Println(userToUpdate)

	updateResult := data.DB.Model(&dbUser).Select("name", "email").Updates(&userToUpdate)
	if updateResult.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": "An error occurred while updating user"})
		return
	}
	context.JSON(http.StatusOK, dbUser)
}
