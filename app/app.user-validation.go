package app

type UserValidation struct {
	Username string `form:"username" valid:"required"`
	Email    string `form:"email" valid:"email,required"`
	Password string `form:"password" valid:"minstringlength(6),required"`
}

type LoginValidation struct {
	Email    string `form:"email" valid:"email,required"`
	Password string `form:"password" valid:"required"`
}
