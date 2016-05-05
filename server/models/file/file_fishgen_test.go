package file

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestFile(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
