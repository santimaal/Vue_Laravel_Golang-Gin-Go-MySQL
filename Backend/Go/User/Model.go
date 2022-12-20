package User

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Id        uint   `json:"id"`
	Is_active bool   `json:"is_active"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}
