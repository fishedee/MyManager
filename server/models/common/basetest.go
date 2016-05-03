package common

import (
	. "github.com/fishedee/web"
)

type BaseTest struct {
	BeegoValidateTest
}

func InitTest(test BeegoValidateTestInterface) {
	InitBeegoVaildateTest(test)
}
