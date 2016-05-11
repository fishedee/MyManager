package blog

import (
	"fmt"
	. "github.com/fishedee/web"
)

type blogGitAoTest struct {
	Test
	BlogGitAo blogGitAoModel
}

func (this *blogGitAoTest) TestMarkdown() {
	markdown := `
# 你好
这是啥
为什么会这样子
	`
	data := this.BlogGitAo.markdownToHtml(markdown)
	fmt.Println(data)
}

func (this *blogGitAoTest) TestGit() {
	fmt.Println("download git...")
	data := this.BlogGitAo.Get("https://github.com/fishedee/poj.git", func(message string) {
		fmt.Println(message)
	})
	fmt.Println(data)
}
