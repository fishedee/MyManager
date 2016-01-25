package account

import (
	. "github.com/fishedee/language"
)

var AccountTypeEnum struct {
	EnumStruct
	TYPE_IN           int `enum:"1,收入"`
	TYPE_OUT          int `enum:"2,支出"`
	TYPE_TRANSFER_IN  int `enum:"3,转账收入"`
	TYPE_TRANSFER_OUT int `enum:"4,转账支出"`
	TYPE_BORROW_IN    int `enum:"5,借还款收入"`
	TYPE_BORROW_OUT   int `enum:"6,借还款支出"`
}

func init() {
	InitEnumStruct(&AccountTypeEnum)
}
