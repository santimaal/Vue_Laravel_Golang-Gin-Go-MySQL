package User

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Id         uint      `json:"id"`
	Is_active  bool      `json:"is_active"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Img        string    `json:"img"`
	Type       string    `json:"type"`
	Noti       int       `json:"noti"`
	Chat_id    string    `json:"chat_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (b *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) clean() error {
	u.Id = 0
	u.Email = ""
	u.Name = ""
	u.Img = ""
	u.Is_active = true
	u.Password = ""
	u.Noti = 0
	u.Type = ""
	u.Chat_id = ""
	return nil
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	fmt.Println(bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword))
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func GenToken(id uint) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(os.Getenv("SecretPassword")))
	return token
}
