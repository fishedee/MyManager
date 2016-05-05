package blog

import (
	. "github.com/fishedee/web"
	"testing"
)

type testFishGenStruct struct{}

func TestBlog(t *testing.T) {
	RunTest(t, &testFishGenStruct{})
}
