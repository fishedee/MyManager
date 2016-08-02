package register

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestRegister(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
