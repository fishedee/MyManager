package blog

import (
	. "github.com/fishedee/language"
)

var BlogSyncTypeEnum struct {
	EnumStruct
	TYPE_INCREMENTAL_UPDATE int `enum:"1,增量更新"`
	TYPE_ALL_UPDATE         int `enum:"2,全量更新"`
}

var BlogStateEnum struct {
	EnumStruct
	STATE_BEGIN    int `enum:"1,未开始"`
	STATE_PROGRESS int `enum:"2,进行中"`
	STATE_FAIL     int `enum:"3,失败"`
	STATE_SUCCESS  int `enum:"4,成功"`
}

func init() {
	InitEnumStruct(&BlogSyncTypeEnum)
	InitEnumStruct(&BlogStateEnum)
}
