package validator

import (
	"log"
	"reflect"
	"regexp"
	"strings"
	"zhu/myrest/utils/errmsg"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	// To translate err msg into Chinese
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		// 这个地方应该是翻译失败之类的错误
		log.Println("err:  ", err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	_ = validate.RegisterValidation("customPhoneNumber", customPhoneNumber)
	_ = validate.RegisterValidation("customPassword", customPassword)
	err = validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// phone number is not valid
			if err.Tag() == "customPhoneNumber" {
				return errmsg.PHONE_NUMBER_NOT_VALID, errmsg.ERROR
			} else if err.Tag() == "customPassword" {
				return errmsg.PASSWORD_NOT_VALID, errmsg.ERROR
			} else {
				return err.Translate(trans), errmsg.ERROR
			}
		}
	}
	return "", errmsg.SUCCESS
}

// 自定义约束条件
// Add our own Tag to validate
func customPhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	// phonr number must be 11 numbers starting with 84
	matched, err := regexp.MatchString(`^84\d{9}$`, phoneNumber)
	return err == nil && matched
}
func customPassword(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	// password must contain 1 special character
	if !strings.ContainsAny(password, "!@#$%^&*") {
		return false
	}
	return true
}
