package card

import (
	. "github.com/fishedee/language"
)

var CardQueueEnum struct {
	EnumStructString
	EVENT_DEL string `enum:"/card/_del,银行卡被删除"`
}

func init() {
	InitEnumStructString(&CardQueueEnum)
}
