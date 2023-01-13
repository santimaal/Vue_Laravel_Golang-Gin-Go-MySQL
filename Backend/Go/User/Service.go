package User

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsersService(c *gin.Context) []UserModel {
	return GetAllUsersRepo(c)
}

func GetOneUserService(c *gin.Context) (UserModel, error) {
	idUser := c.Param("id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		println("error")
	}

	return GetOneUserRepo(id, c)
}

func GetUserServiceByID(c *gin.Context, id uint) (UserModel, error) {
	return GetOneUserByID(id, c)
}

func UserRegisterService(c *gin.Context) (error, bool) {
	var usrModel UserModel
	c.BindJSON(&usrModel)
	usrModel.Type = "client"
	usrModel.setPassword(usrModel.Password)
	exists, err := CheckUserEmail(&usrModel, c)
	if exists.Id != 0 {
		println("exist")
		return err, true
	} else {
		return UserRegisterRepo(&usrModel, c)
	}
}

func UserLoginService(c *gin.Context) (error, UserModel) {
	var usrModel UserModel
	c.BindJSON(&usrModel)
	exists, err := CheckUserEmail(&usrModel, c)
	if exists.Id != 0 {
		fmt.Println(exists)
		if exists.checkPassword(usrModel.Password) != nil {
			usrModel.clean()
			return err, usrModel
		} else {
			return err, exists
		}
	}
	return err, exists
	// usrModel.setPassword(usrModel.Password)
}

func ComproveCodeService(c *gin.Context, code string, usrID int64) bool {
	usr, err := ComproveCodeRepo(c, code)
	if err != nil {
		panic(err)
	}
	if usr.Id == 0 {
		return false
	} else {
		usr.Chat_id = strconv.FormatInt(usrID, 10)
		UpdateUserRepo(c, usr)
		return true
	}
}
