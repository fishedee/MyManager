package common

import (
	"github.com/fishedee/web"
)

type BaseTest struct {
	web.BeegoValidateTest
}

func InitTest(test web.BeegoValidateTestInterface) {
	web.InitBeegoVaildateTest(test)
}
