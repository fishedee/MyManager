package mail

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestMail(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
