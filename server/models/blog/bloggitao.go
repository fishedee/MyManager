package blog

import (
	"fmt"
	. "github.com/fishedee/language"
	"github.com/russross/blackfriday"
	"io/ioutil"
	. "mymanager/models/common"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

type BlogGitAoModel struct {
	BaseModel
}

func (this *BlogGitAoModel) getFileContent(fileAddress string) string {
	data, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (this *BlogGitAoModel) markdownToHtml(data string) string {
	return string(blackfriday.MarkdownCommon([]byte(data)))
}

func (this *BlogGitAoModel) analyseSingleDir(fileAddress string) []BlogArticle {
	category := path.Base(fileAddress)

	dirFileInfo, err := ioutil.ReadDir(fileAddress)
	if err != nil {
		panic(err)
	}
	result := []BlogArticle{}
	for _, singleFile := range dirFileInfo {
		singleFileName := singleFile.Name()
		var content string
		var htmlContent string
		var article BlogArticle
		if strings.HasSuffix(singleFileName, ".md") {
			content = this.getFileContent(fileAddress + "/" + singleFileName)
			content = strings.Trim(content, " ")
			if content == "" {
				continue
			}
			htmlContent = this.markdownToHtml(content)
		} else {
			continue
		}
		article.Title = singleFileName[0 : len(singleFileName)-len(path.Ext(singleFileName))]
		article.Content = content
		article.HtmlContent = htmlContent
		article.Category = category
		result = append(result, article)
	}
	return result
}

func (this *BlogGitAoModel) analyse(dir string) Blog {
	dirFileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	result := Blog{}
	for _, singleFile := range dirFileInfo {
		if singleFile.Name() == ".git" {
			continue
		}
		if singleFile.Mode().IsDir() == false {
			continue
		}
		articles := this.analyseSingleDir(dir + "/" + singleFile.Name())
		result.Categorys = append(
			result.Categorys,
			BlogCategory{
				Name: singleFile.Name(),
			},
		)
		for _, singleArticle := range articles {
			result.Articles = append(
				result.Articles,
				singleArticle,
			)
		}
	}

	return result
}

func (this *BlogGitAoModel) download(gitUrl string) string {
	if strings.HasPrefix(gitUrl, "https://github.com") == false {
		Throw(1, "请输入https://github.com开头的git仓库地址")
	}
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	now := time.Now()
	gitDir := workingDir + "/../data/git"
	gitFile := fmt.Sprintf("%s_%v-%v-%v_%v:%v:%v", path.Base(gitUrl), now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	cmd := exec.Cmd{
		Path: "/usr/bin/git",
		Args: []string{
			"git",
			"clone",
			gitUrl,
			gitFile,
		},
		Dir: gitDir,
	}
	result, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("%v\n%v", err.Error(), string(result)))
	}
	return gitDir + "/" + gitFile
}

func (this *BlogGitAoModel) Get(gitUrl string, progressUpdater BlogSyncProgress) Blog {
	progressUpdater("正在从git中下载博客数据")
	localDir := this.download(gitUrl)

	progressUpdater("从git中分析博客数据")
	return this.analyse(localDir)
}
