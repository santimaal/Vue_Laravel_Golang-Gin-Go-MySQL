package User

import (
	"fmt"
	"net/http"
	"sanvic/Config"

	"github.com/gin-gonic/gin"
)

func GetAllUsersRepo(c *gin.Context) []UserModel {

	var users []UserModel

	if err := Config.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return users

}

func GetOneUserRepo(id int, c *gin.Context) (UserModel, error) {

	var user UserModel
	err := Config.DB.Where("id = ?", id).Find(&user).Error
	return user, err
}

func UserRegisterRepo(user *UserModel, c *gin.Context) (err error, exist bool) {
	err = Config.DB.Create(user).Error
	return err, false
}

func CheckUserEmail(user *UserModel, c *gin.Context) (exists UserModel, err error) {
	err = Config.DB.Where("name = ?", user.Name).Find(&exists).Error
	return exists, err
}
