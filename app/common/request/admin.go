package request

type AdminLogin struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (admin AdminLogin) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"account.required": "账号格式不正确",
		"password.required": "用户密码不能为空",
	}
}