package card

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestCard(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
