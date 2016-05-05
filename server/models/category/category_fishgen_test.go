package category

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestCategory(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
