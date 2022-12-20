package User

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ALL Users
func GetAllUsers(c *gin.Context) {
	var users []UserModel = GetAllUsersService(c)
	println(users)
	serializer := UsersSerializer{c, users}
	c.JSON(http.StatusOK, serializer.Response())
}

// REGISTER User
func UserRegister(c *gin.Context) {
	err, bool := UserRegisterService(c)
	if err != nil || bool {
		c.JSON(http.StatusInternalServerError, "Email is registered")
	} else {
		c.JSON(http.StatusOK, "User created correctly")
	}
}

// GET ONE User
func GetUserByID(c *gin.Context) {
	user, err := GetOneUserService(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "User doesn't exist")
	} else {
		serializer := UserSerializer{c, user}
		c.JSON(http.StatusOK, serializer.Response())
	}
}
