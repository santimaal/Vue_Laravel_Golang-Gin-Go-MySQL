package User

import (
	"github.com/gin-gonic/gin"
)

func GetAllUsersService(c *gin.Context) []UserModel {
	return GetAllUsersRepo(c)
}

func GetOneUserService(id int, c *gin.Context) (UserModel, error) {
	return GetOneUserRepo(id, c)
}
