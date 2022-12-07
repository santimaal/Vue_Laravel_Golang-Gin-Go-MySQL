package User

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET ALL Users
func GetAllUsers(c *gin.Context) {
	var users []UserModel = GetAllUsersService(c)
	println(users)
	serializer := UsersSerializer{c, users}
	c.JSON(http.StatusOK, serializer.Response())
}

//GET ONE User
func GetUserByID(c *gin.Context) {
	idUser := c.Param("id")
	var user UserModel
	var id int
	id, err := strconv.Atoi(idUser)
	if err != nil {
		println("error")
	}

	user, err = GetOneUserService(id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, "User doesn't exist")
	} else {
		serializer := UserSerializer{c, user}
		c.JSON(http.StatusOK, serializer.Response())
	}

}
