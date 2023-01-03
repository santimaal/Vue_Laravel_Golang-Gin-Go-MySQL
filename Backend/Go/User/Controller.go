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

func UserLogin(c *gin.Context) {
	err, usr := UserLoginService(c)
	if err != nil || len(usr.Name) == 0 {
		c.JSON(http.StatusInternalServerError, "Email or password is not correct")
	} else {
		UpdateContextUserModel(c, usr.Id)
		// c.Set("my_user_id", user.Id)
		// c.Set("my_user_model", user)
		serializer := UserSerializer{c, usr}
		c.JSON(http.StatusOK, serializer.Response())
		// c.JSON(http.StatusOK, usr)
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
