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
	usrModel.Img = "https://i.postimg.cc/W41QygPj/descarga.png"
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
}

func UserUpdateService(c *gin.Context) bool {
	var usrModel UserModel
	c.BindJSON(&usrModel)

	usr, _ := c.Get("my_user_model")
	u, ok := usr.(UserModel)
	if !ok {
		fmt.Println("No se ha podido convertir")
	}

	if len(usrModel.Chat_id) != 0 {
		u.Chat_id = usrModel.Chat_id
	}
	if len(usrModel.Name) != 0 {
		u.Name = usrModel.Name
	}
	if len(usrModel.Password) != 0 {
		usrModel.setPassword(usrModel.Password)
		u.Password = usrModel.Password
	}
	if len(usrModel.Email) != 0 {
		exists, err := CheckUserEmail(&usrModel, c)
		if exists.Id == 0 {
			u.Email = usrModel.Email
		} else {
			fmt.Println(err)
			return false
		}

	}
	if len(usrModel.Img) != 0 {
		u.Img = usrModel.Img
	}
	return UpdateUserRepo(c, u)

}

func ComproveCodeService(c *gin.Context, code string, usrID int64) bool {
	usr, _ := ComproveCodeRepo(c, code)
	if usr.Id == 0 {
		return false
	} else {
		usr.Chat_id = strconv.FormatInt(usrID, 10)
		UpdateUserRepo(c, usr)
		return true
	}
}
