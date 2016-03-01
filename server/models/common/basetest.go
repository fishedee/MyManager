package common

import (
	. "github.com/fishedee/web"
	"testing"
)

type BaseTest struct {
	BeegoValidateTest
}

func InitTest(t *testing.T, test BeegoValidateTestInterface) {
	InitBeegoVaildateTest(t, test)
}
