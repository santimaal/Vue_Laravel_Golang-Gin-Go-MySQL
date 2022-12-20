package User

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsersService(c *gin.Context) []UserModel {
	return GetAllUsersRepo(c)
}

func GetOneUserService(c *gin.Context) (UserModel, error) {
	idUser := c.Param("id")
	var id int
	id, err := strconv.Atoi(idUser)
	if err != nil {
		println("error")
	}

	return GetOneUserRepo(id, c)
}

func UserRegisterService(c *gin.Context) (error, bool) {
	var usrModel UserModel
	c.BindJSON(&usrModel)
	usrModel.setPassword(usrModel.Password)
	exists, err := CheckUserEmail(&usrModel, c)
	if exists.Id != 0 {
		println("exist")
		return err, true
	} else {
		return UserRegisterRepo(&usrModel, c)
	}
}
