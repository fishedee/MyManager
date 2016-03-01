package blog

import (
	"bytes"
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/util"
	. "mymanager/models/common"
	"strconv"
	"strings"
	"time"
)

type BlogCsdnCrawlAoModel struct {
	BaseModel
	AjaxPool *AjaxPool
}

func (this *BlogCsdnCrawlAoModel) apiForHtml(method string, url string, data interface{}, referer string) *goquery.Document {
	result := this.api(method, url, data, referer)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		panic(err)
	}
	return doc
}

func (this *BlogCsdnCrawlAoModel) apiForJson(method string, url string, data interface{}, responseData interface{}, referer string) {
	result := this.api(method, url, data, referer)

	var commonData struct {
		Status bool   `json:status`
		Error  string `json:error`
		Data   interface{}
	}
	commonData.Data = responseData
	err := DecodeJson(result, &commonData)
	if commonData.Status == false {
		panic(url + " error " + commonData.Error)
	}
	if err != nil {
		panic("decode error " + string(result))
	}
}

func (this *BlogCsdnCrawlAoModel) api(method string, url string, data interface{}, referer string) []byte {
	if this.AjaxPool == nil {
		this.AjaxPool = NewAjaxPool(&AjaxPoolOption{
			Timeout:      time.Second * 30,
			HasCookieJar: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
	}

	var result []byte
	header := map[string]string{
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",
		"Upgrade-Insecure-Requests": "1",
	}
	if len(referer) != 0 {
		header["REFERER"] = referer
	}
	option := &Ajax{
		Url:          url,
		Data:         data,
		ResponseData: &result,
		Header:       header,
	}
	var err error
	if method == "get" {
		err = this.AjaxPool.Get(option)
	} else {
		err = this.AjaxPool.Post(option)
	}
	if err != nil {
		panic(err)
	}

	return result
}

func (this *BlogCsdnCrawlAoModel) Login(name string, password string) {
	//获取登录页面
	doc := this.apiForHtml("get", "https://passport.csdn.net/account/login", "", "")
	var argv struct {
		UserName  string `json:"username"`
		Password  string `json:"password"`
		Lt        string `json:"lt"`
		Execution string `json:"execution"`
		EventId   string `json:"_eventId"`
	}
	argv.Lt = doc.Find("#fm1 input[name=lt]").AttrOr("value", "")
	if argv.Lt == "" {
		panic("invalid lt")
	}
	argv.Execution = doc.Find("#fm1 input[name=execution]").AttrOr("value", "")
	if argv.Execution == "" {
		panic("invalid exection")
	}
	argv.EventId = doc.Find("#fm1 input[name=_eventId]").AttrOr("value", "")
	if argv.EventId == "" {
		panic("invalid eventId")
	}
	argv.UserName = name
	argv.Password = password

	//登录
	doc2 := this.apiForHtml("post", "https://passport.csdn.net/account/login", argv, "")
	bodyData := doc2.Find("body").AttrOr("onload", "")
	if strings.Index(bodyData, "redirect_back") == -1 {
		html, _ := doc2.Html()
		panic("invalid bodyData " + html)
	}
}

func (this *BlogCsdnCrawlAoModel) analyseCategory(s *goquery.Selection) BlogCategory {
	singleName := s.Find("td").Eq(0).Text()
	singleName = strings.Trim(singleName, " ")
	if len(singleName) == 0 {
		panic("invalid name " + singleName)
	}
	singleId := s.Find("td").Eq(2).Find("a").Eq(0).AttrOr("href", "")
	singleId = strings.Trim(singleId, "#")
	singleIdInt, err := strconv.Atoi(singleId)
	if err != nil {
		panic("invalid integer " + singleId)
	}
	return BlogCategory{
		Name: singleName,
		Id:   singleIdInt,
	}
}

func (this *BlogCsdnCrawlAoModel) GetCategoryList() []BlogCategory {
	doc := this.apiForHtml("get", "http://write.blog.csdn.net/category", "", "")

	data := []BlogCategory{}
	doc.Find("#lstBody tr").Each(func(index int, s *goquery.Selection) {
		data = append(data, this.analyseCategory(s))
	})
	return data
}

func (this *BlogCsdnCrawlAoModel) AddCategory(category BlogCategory) {
	var argv struct {
		T    string `json:"t"`
		Name string `json:"name`
	}
	argv.T = "add"
	argv.Name = category.Name
	this.apiForHtml("get", "http://write.blog.csdn.net/category", argv, "")
}

func (this *BlogCsdnCrawlAoModel) DelCategory(id int) {
	var argv struct {
		T  string `json:"t"`
		Id int    `json:"id`
	}
	argv.T = "del"
	argv.Id = id
	this.apiForHtml("get", "http://write.blog.csdn.net/category", argv, "")
}

func (this *BlogCsdnCrawlAoModel) ModCategory(id int, data BlogCategory) {
	var argv struct {
		T    string `json:"t"`
		Id   int    `json:"id`
		Name string `json:name`
	}
	argv.T = "edit"
	argv.Id = id
	argv.Name = data.Name
	this.apiForHtml("get", "http://write.blog.csdn.net/category", argv, "")
}

type blogArticleInfo struct {
	ArticleEditType int    `json:"articleedittype"`
	Categories      string `json:"categories"`
	Channel         int    `json:"channel"`
	Content         string `json:"content"`
	Description     string `json:"description"`
	Id              int    `json:"id,omitempty"`
	Level           int    `json:"level"`
	MarkdownContent string `json:"markdowncontent"`
	Status          int    `json:"status"`
	Tags            string `json:"tags"`
	Title           string `json:"title"`
	Type            string `json:"type"`
}

func (this *BlogCsdnCrawlAoModel) analyseArticle(s *goquery.Selection) BlogArticle {
	if s.Find("td").Length() == 0 {
		return BlogArticle{}
	}
	singleName := s.Find("td").Eq(0).Find("a").Text()
	singleName = strings.Trim(singleName, " ")
	if len(singleName) == 0 {
		panic("invalid name " + singleName)
	}
	singleId := s.Find("td").Eq(0).Find("a").AttrOr("href", "")
	singleId = strings.Trim(singleId, "/")
	singleIdIndex := strings.LastIndex(singleId, "/")
	singleId = singleId[singleIdIndex+1:]
	singleIdInt, err := strconv.Atoi(singleId)
	if err != nil {
		panic("invalid integer " + singleId)
	}
	return BlogArticle{
		Title: singleName,
		Id:    singleIdInt,
	}
}

func (this *BlogCsdnCrawlAoModel) getArticleInfo(data BlogArticle) interface{} {
	var argv blogArticleInfo
	argv.ArticleEditType = 1
	argv.Categories = data.Category
	argv.Channel = 7
	argv.Description = ""
	argv.Id = data.Id
	argv.Level = 0
	argv.Content = data.HtmlContent
	argv.MarkdownContent = data.Content
	argv.Status = 0
	argv.Tags = ""
	argv.Title = data.Title
	argv.Type = "original"
	return argv
}

func (this *BlogCsdnCrawlAoModel) ModArticle(id int, data BlogArticle) {
	var responseData interface{}
	data.Id = id
	this.apiForJson("post", "http://write.blog.csdn.net/mdeditor/setArticle", this.getArticleInfo(data), &responseData, "")
}

func (this *BlogCsdnCrawlAoModel) AddArticle(data BlogArticle) int {
	var responseData struct {
		Id int `json:id`
	}
	this.apiForJson("post", "http://write.blog.csdn.net/mdeditor/setArticle", this.getArticleInfo(data), &responseData, "")
	return responseData.Id
}

func (this *BlogCsdnCrawlAoModel) DelArticle(id int) {
	var argv struct {
		T  string `json:"t"`
		Id int    `json:"id"`
	}
	argv.T = "del"
	argv.Id = id
	this.apiForHtml("get", "http://write.blog.csdn.net/postlist", argv, "http://write.blog.csdn.net/postlist")
}

func (this *BlogCsdnCrawlAoModel) GetArticle(id int, name string) BlogArticle {
	var argv struct {
		Id       int    `json:"id"`
		UserName string `json:"username"`
	}
	var responseData blogArticleInfo
	argv.Id = id
	argv.UserName = name
	this.apiForJson("get", "http://write.blog.csdn.net/mdeditor/getArticle", argv, &responseData, "")
	return BlogArticle{
		Id:          id,
		Title:       responseData.Title,
		Content:     responseData.MarkdownContent,
		HtmlContent: responseData.Content,
		Category:    responseData.Categories,
	}
}

func (this *BlogCsdnCrawlAoModel) GetArticleList(page int) ([]BlogArticle, int) {
	doc := this.apiForHtml("get", "http://write.blog.csdn.net/postlist/0/0/enabled/"+strconv.Itoa(page), "", "")

	data := []BlogArticle{}
	doc.Find("#lstBox tr").Each(func(index int, s *goquery.Selection) {
		singleArticle := this.analyseArticle(s)
		if singleArticle.Id == 0 {
			return
		}
		data = append(data, singleArticle)
	})

	countData := doc.Find(".page_nav span").Text()
	countData = strings.Trim(countData, " ")
	countDataResult := []byte{}
	for i := 0; i != len(countData); i++ {
		if countData[i] >= '0' && countData[i] <= '9' {
			countDataResult = append(countDataResult, countData[i])
		} else {
			break
		}
	}
	var countDataInt int
	if len(countDataResult) == 0 {
		countDataInt = 0
	} else {
		var err error
		countDataInt, err = strconv.Atoi(string(countDataResult))
		if err != nil {
			panic(err)
		}
	}
	return data, countDataInt
}
