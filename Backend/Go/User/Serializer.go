package User

import (
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	// Id        uint   `json:"id"`
	Is_active bool   `json:"is_active"`
	Name      string `json:"name"`
	// Password  string `json:"password"`
	Email string `json:"email"`
	Img   string `json:"img"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

type UserSerializer struct {
	C    *gin.Context
	user UserModel
}

type UsersSerializer struct {
	C     *gin.Context
	users []UserModel
}

func (ss *UserSerializer) Response() UserResponse {
	response := UserResponse{
		// Id:        ss.user.Id,
		Is_active: ss.user.Is_active,
		Name:      ss.user.Name,
		// Password:  ss.user.Password,
		Email: ss.user.Email,
		Img:   ss.user.Img,
		Type:  ss.user.Type,
		Token: GenToken(ss.user.Id),
	}

	return response
}

func (ss *UsersSerializer) Response() []UserResponse {
	response := []UserResponse{}

	for _, user := range ss.users {
		serializer := UserSerializer{ss.C, user}
		response = append(response, serializer.Response())
	}

	return response
}
