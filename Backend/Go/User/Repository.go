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
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return user, err
}
