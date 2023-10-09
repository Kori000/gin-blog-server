package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetValidMsg(err error, obj any) string {
	// 需要传入obj指针
	getObj := reflect.TypeOf(obj)
	// 将 err 接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每个错误信息
			// 根据报错字段名, 获取结构提的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}
