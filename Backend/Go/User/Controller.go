package User

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		serializer := UserSerializer{c, usr}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

func GetProfile(c *gin.Context) {
	usr, _ := c.Get("my_user_model")
	u, ok := usr.(UserModel)
	if ok {
		serializer := UserSerializer{c, u}
		c.JSON(http.StatusOK, serializer.Response())
	} else {
		fmt.Println("No se ha podido convertir")
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

func ComproveCode(code string, usrID int64) bool {
	var c *gin.Context
	return ComproveCodeService(c, code, usrID)
}

func UserUpdate(c *gin.Context) {
	bool := UserUpdateService(c)
	if !bool {
		c.JSON(http.StatusInternalServerError, "Error Updated User")
	} else {
		id, _ := c.Get("my_user_id")
		idUser, _ := id.(uint)
		user, err := GetUserServiceByID(c, uint(idUser))
		if err != nil {
			c.JSON(http.StatusNotFound, "User doesn't exist")
		} else {
			serializer := UserSerializer{c, user}
			c.JSON(http.StatusOK, serializer.Response())
		}
	}
}

func SendTel(c *gin.Context) {
	type Message struct {
		Id_user uint   `json:"id_user"`
		Message string `json:"message"`
	}

	var message Message
	c.BindJSON(&message)
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	u, _ := GetUserServiceByID(c, message.Id_user)

	if len(u.Chat_id) != 0 {
		var parsed int64
		var err error
		parsed, err = strconv.ParseInt(u.Chat_id, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, "can't send it")
		} else {
			bot.Debug = true
			msg := tgbotapi.NewMessage(parsed, message.Message)
			_, err = bot.Send(msg)
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, "sended correctly")
		}
	}
}
