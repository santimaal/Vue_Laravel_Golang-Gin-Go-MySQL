package User

type UserModel struct {
	Id uint `json:"id"`
	Is_active bool `json:"is_active"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"Email"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *UserModel) TableName() string {
	return "users"
}
