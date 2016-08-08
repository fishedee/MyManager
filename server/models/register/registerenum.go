package register

import (
	. "github.com/fishedee/language"
)

var RegisterNeedDealType struct {
	EnumStruct
	NO  int `enum:"1,不需要"`
	YES int `enum:"2,需要"`
}

var RegisterHaveDealType struct {
	EnumStruct
	NO  int `enum:"1,未挂号"`
	YES int `enum:"2,已挂号"`
}

func init() {
	InitEnumStruct(&RegisterNeedDealType)
	InitEnumStruct(&RegisterHaveDealType)
}
