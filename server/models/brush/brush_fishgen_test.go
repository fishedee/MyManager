package brush

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestBrush(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
