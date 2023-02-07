package forms

type SendSmsForm struct {
	//动态验证码登录发送验证码
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`
}
