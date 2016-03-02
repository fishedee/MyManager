package blog

import (
	"fmt"
	. "mymanager/models/common"
	"testing"
)

type BlogGitAoTest struct {
	BaseTest
	BlogGitAo BlogGitAoModel
}

func (this *BlogGitAoTest) TestMarkdown() {
	markdown := `
# 你好
这是啥
为什么会这样子
	`
	data := this.BlogGitAo.markdownToHtml(markdown)
	fmt.Println(data)
}

func TestBlogGit(t *testing.T) {
	InitTest(t, &BlogGitAoTest{})
}
