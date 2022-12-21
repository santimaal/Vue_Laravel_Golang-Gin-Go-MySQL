package User

type UserModelValidator struct {
	User struct {
		Name      string `form:"username" json:"username" binding:"exists,alphanum,min='4',max='255'"`
		Email     string `form:"email" json:"email" binding:"exists,email"`
		Password  string `form:"password" json:"password" binding:"exists,min='8',max='255'"`
		Is_active bool   `form:"is_active" json:"is_active" valid:"required"`
		Img       string `form:"img" json:"img" binding:"omitempty,url"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

// func (self *UserModelValidator) Bind(c *gin.Context) error {
// err := common.Bind(c, self)
// if err != nil {
// 	return err
// }
// self.userModel.Name = self.User.Name
// self.userModel.Email = self.User.Email

// if self.User.Password != common.NBRandomPassword {
// 	self.userModel.setPassword(self.User.Password)
// }
// if self.User.Img != "" {
// 	self.userModel.Img = self.User.Img
// }
// return nil
// }

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	//userModelValidator.User.Email ="w@g.cn"
	return userModelValidator
}
