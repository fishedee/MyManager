package account

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestAccount(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
