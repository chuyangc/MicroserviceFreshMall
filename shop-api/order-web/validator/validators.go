package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateMobile(f1 validator.FieldLevel) bool {
	mobile := f1.Field().String()
	//使用正则表达式判断手机号码合法性
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
