package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"zhu/myrest/utils/errmsg"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:  ", err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	_ = validate.RegisterValidation("customPhoneNumber", customPhoneNumber)

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
func customPhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	matched, err := regexp.MatchString(`^84\d{9}$`, phoneNumber)
	return err == nil && matched
}

// func customPhoneNumber(v validator.FieldLevel) bool {
// 	phoneNumber := v.Field().String()
// 	matched, err := regexp.MatchString(`^84\d{9}$`, phoneNumber)
// 	return err == nil && matched
// }
