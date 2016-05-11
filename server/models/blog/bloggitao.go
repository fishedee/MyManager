package blog

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
	. "github.com/fishedee/web"
	"github.com/russross/blackfriday"
	"io/ioutil"
	. "mymanager/models/file"
	"os"
	"os/exec"
	"path"
	"strings"
)

type blogGitAoModel struct {
	Model
	UploadAo UploadAoModel
}

func (this *blogGitAoModel) getFileContent(fileAddress string) string {
	data, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (this *blogGitAoModel) markdownToHtml(data string) string {
	htmlOption := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES |
		blackfriday.HTML_SMARTYPANTS_ANGLED_QUOTES
	extensionOption := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS
	renderer := blackfriday.HtmlRenderer(htmlOption, "", "")
	result := blackfriday.MarkdownOptions([]byte(data), renderer, blackfriday.Options{
		Extensions: extensionOption})
	return string(result)
}

func (this *blogGitAoModel) convertImage(dir string, content string) string {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(content)))
	if err != nil {
		panic(err)
	}
	doc.Find("img").Each(func(index int, s *goquery.Selection) {
		src := strings.Trim(s.AttrOr("src", ""), " ")
		if src == "" ||
			strings.HasPrefix(src, "http") ||
			strings.HasPrefix(src, "https") ||
			strings.HasPrefix(src, "data") {
			return
		}
		src, err := DecodeUrl(src)
		if err != nil {
			panic(err)
		}
		localPath := dir + "/" + src
		newSrc := this.UploadAo.UploadFileFromLocal(localPath)
		s.SetAttr("src", newSrc)
	})
	docHtml, err := doc.Html()
	if err != nil {
		panic(err)
	}
	return docHtml
}

func (this *blogGitAoModel) analyseSingleDir(fileAddress string) []BlogArticle {
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
			htmlContent = this.convertImage(fileAddress, htmlContent)
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

func (this *blogGitAoModel) analyse(dir string) Blog {
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

func (this *blogGitAoModel) download(gitUrl string, tempDir string) string {
	if strings.HasPrefix(gitUrl, "https://github.com") == false {
		Throw(1, "请输入https://github.com开头的git仓库地址")
	}
	gitFile := path.Base(tempDir)
	gitDir := path.Dir(tempDir)

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

func (this *blogGitAoModel) Get(gitUrl string, progressUpdater BlogSyncProgress) Blog {
	tempDir, err := CreateTempFile("mymanager", "-git")
	defer os.RemoveAll(tempDir)
	if err != nil {
		panic(err)
	}

	progressUpdater("正在从git中下载博客数据")
	localDir := this.download(gitUrl, tempDir)

	progressUpdater("从git中分析博客数据")
	return this.analyse(localDir)
}
