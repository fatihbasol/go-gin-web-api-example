package handle

import (
	"github.com/fatihbasol/GoGinExample/src/data"
	"github.com/fatihbasol/GoGinExample/src/handle/model/request"
	"github.com/fatihbasol/GoGinExample/src/handle/model/response"
	"github.com/fatihbasol/GoGinExample/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func GetUser(context *gin.Context) {
	idParam := context.Param("id")
	id, _ := strconv.Atoi(idParam)

	var user model.User
	result := data.DB.Preload("Phone").First(&user, id)

	userResponseModel := response.UserResponseModel{Id: user.Id, Name: user.Name, Email: user.Email,
		Phone: &response.PhoneResponseModel{Id: user.Phone.Id, CountryCode: user.Phone.CountryCode, Number: user.Phone.Number}}

	if result.Error == gorm.ErrRecordNotFound {
		log.Default().Printf("%s %s: %s", gorm.ErrRecordNotFound.Error(), reflect.TypeOf(user).Name(), idParam)
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}
	context.JSON(http.StatusOK, userResponseModel)
}
func GetUsers(context *gin.Context) {
	// get from db with  relational data(include phones)
	var users []model.User
	data.DB.Preload("Phone").Find(&users)

	var usersResponseModel []response.UserResponseModel

	for _, user := range users {
		usersResponseModel = append(usersResponseModel,
			response.UserResponseModel{Id: user.Id, Name: user.Name, Email: user.Email,
				Phone: &response.PhoneResponseModel{Id: user.Phone.Id, CountryCode: user.Phone.CountryCode, Number: user.Phone.Number}})
	}

	context.JSON(http.StatusOK, usersResponseModel)
}
func PostUser(context *gin.Context) {
	var userModel request.UserRequestModel
	userModel = request.UserRequestModel{Name: context.PostForm("name"), Email: context.PostForm("email")}

	if userModel.Name == "" || userModel.Email == "" {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	var dbUser = model.User{Name: userModel.Name, Email: userModel.Email}
	result := data.DB.Create(&dbUser)

	if result.Error != nil {
		log.Default().Printf("%s %s: %s", gorm.ErrInvalidData.Error(), reflect.TypeOf(dbUser).Name(), userModel)
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	userResponseModel := response.UserResponseModel{
		Id:    dbUser.Id,
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}
	context.JSON(http.StatusOK, userResponseModel)

}
func PutUser(context *gin.Context) {
	idToPut, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		log.Default().Println("Bad request : invalid id")
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	var dbUser *model.User
	result := data.DB.First(&dbUser, idToPut)

	if result.Error != nil {
		log.Default().Printf("An error occurred while fetching user: %d", idToPut)
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if result.RowsAffected == 0 {
		log.Default().Printf("Cannot find user: %d", idToPut)
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	if dbUser != nil {
		name := context.PostForm("name")
		email := context.PostForm("email")

		if name == "" || email == "" {
			log.Default().Printf("Fields cannot be empty: %d", idToPut)
			context.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		updatedUser := model.User{Id: idToPut}
		result := data.DB.Model(&updatedUser).Select("name", "email").Updates(model.User{Name: name, Email: email})

		if result.Error != nil {
			log.Default().Printf("An error occurred while updating user: %d", idToPut)
			context.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		context.JSON(http.StatusOK, updatedUser)
		return
	}

	context.JSON(http.StatusBadRequest, gin.H{})
}
