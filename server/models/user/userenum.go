package user

import (
	. "github.com/fishedee/language"
)

var UserTypeEnum struct {
	EnumStruct
	ADMIN int `enum:"1,超级管理员"`
	USER  int `enum:"2,普通管理员"`
}

func init() {
	InitEnumStruct(&UserTypeEnum)
}
