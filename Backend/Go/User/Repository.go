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

func GetOneUserByID(id uint, c *gin.Context) (UserModel, error) {
	var user UserModel
	err := Config.DB.Where("id = ?", id).Find(&user).Error
	return user, err
}

func UserRegisterRepo(user *UserModel, c *gin.Context) (err error, exist bool) {
	err = Config.DB.Create(user).Error
	return err, false
}

func UserLoginRepo(usrinput *UserModel, c *gin.Context) (err error, user UserModel) {
	err = Config.DB.Raw("SELECT * FROM users u WHERE u.email = ? AND u.password = ?", usrinput.Email, usrinput.Password).Scan(&user).Error
	fmt.Println(user)
	return err, user
}
func CheckUserEmail(user *UserModel, c *gin.Context) (exists UserModel, err error) {
	err = Config.DB.Where("email = ?", user.Email).Find(&exists).Error
	return exists, err
}

func ComproveCodeRepo(c *gin.Context, code string) (usr UserModel, err error) {
	err = Config.DB.Where("chat_id = ?", code).Find(&usr).Error
	return usr, err
}

func UpdateUserRepo(c *gin.Context, usr UserModel) bool {
	err := Config.DB.Save(&usr).Error
	return err != nil
}
