package request

//用来存放所有用户相关的请求结构体，并实现 Validator 接口

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
	Signature string `form:"signature" json:"signature"`
}

// GetMessages 自定义错误信息
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "用户名称不能为空",
		"Mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"Password.required": "用户密码不能为空",
	}
}


//Login 验证器结构体
type Login struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required": "手机号码不能为空",
		"mobile.mobile": "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}