package blog

import (
	. "github.com/fishedee/web"
)

type blogCsdnCrawlAoTest struct {
	Test
	BlogCsdnCrawlAo BlogCsdnCrawlAoModel
}

func (this *blogCsdnCrawlAoTest) testCategoryClear() {
	categoryList := this.BlogCsdnCrawlAo.GetCategoryList()

	for i := 0; i != len(categoryList); i++ {
		this.BlogCsdnCrawlAo.DelCategory(categoryList[i].Id)
	}

	categoryList = this.BlogCsdnCrawlAo.GetCategoryList()
	this.AssertEqual(len(categoryList), 0)
}

func (this *blogCsdnCrawlAoTest) testCategoryAdd() []BlogCategory {
	count := 3

	data := []string{}
	for i := 0; i != count; i++ {
		singleData := this.RandomString(16)
		this.BlogCsdnCrawlAo.AddCategory(BlogCategory{
			Name: singleData,
		})
		data = append(data, singleData)
	}

	categoryList := this.BlogCsdnCrawlAo.GetCategoryList()
	this.AssertEqual(len(categoryList), len(data))
	for i := 0; i != count; i++ {
		this.AssertEqual(categoryList[i].Name, data[i])
	}
	return categoryList
}

func (this *blogCsdnCrawlAoTest) testCategoryMod(categoryList []BlogCategory) {
	for i := 0; i != len(categoryList); i++ {
		categoryList[i].Name = this.RandomString(16)
		this.BlogCsdnCrawlAo.ModCategory(categoryList[i].Id, categoryList[i])
	}

	categoryList2 := this.BlogCsdnCrawlAo.GetCategoryList()
	this.AssertEqual(categoryList, categoryList2)
}

func (this *blogCsdnCrawlAoTest) testCategory() {
	this.testCategoryClear()

	data := this.testCategoryAdd()

	this.testCategoryMod(data)

	this.testCategoryClear()
}

func (this *blogCsdnCrawlAoTest) getAllArticle(name string) []BlogArticle {
	var result []BlogArticle
	var articleList []BlogArticle

	totalCount := 10
	pageIndex := 1
	for len(articleList) < totalCount {
		var singleArticleList []BlogArticle
		singleArticleList, totalCount = this.BlogCsdnCrawlAo.GetArticleList(pageIndex)
		pageIndex++
		if len(singleArticleList) == 0 {
			break
		}
		articleList = append(articleList, singleArticleList...)
	}

	for _, singleData := range articleList {
		result = append(result, this.BlogCsdnCrawlAo.GetArticle(
			singleData.Id,
			name,
		))
	}
	return result
}

func (this *blogCsdnCrawlAoTest) testArticleClear(name string) {
	articleList := this.getAllArticle(name)

	for i := 0; i != len(articleList); i++ {
		this.BlogCsdnCrawlAo.DelArticle(articleList[i].Id)
	}

	articleList = this.getAllArticle(name)
	this.AssertEqual(len(articleList), 0)
}

func (this *blogCsdnCrawlAoTest) testArticleAdd(name string) []BlogArticle {
	count := 3

	data := []BlogArticle{}
	for i := 0; i != count; i++ {
		singleData := BlogArticle{
			Title:       this.RandomString(16),
			Content:     "#" + this.RandomString(16),
			HtmlContent: "<h1>" + this.RandomString(16) + "</h1>",
			Category:    this.RandomString(8),
		}
		this.BlogCsdnCrawlAo.AddArticle(singleData)
		data = append(data, singleData)
	}

	articleList := this.getAllArticle(name)
	this.AssertEqual(len(articleList), len(data))
	for i := 0; i != count; i++ {
		this.AssertEqual(articleList[i].Title, data[i].Title)
		this.AssertEqual(articleList[i].Content, data[i].Content)
		this.AssertEqual(articleList[i].HtmlContent, data[i].HtmlContent)
	}
	return articleList
}

func (this *blogCsdnCrawlAoTest) testArticleMod(name string, articleList []BlogArticle) {
	for i := 0; i != len(articleList); i++ {
		articleList[i].Content = "#" + this.RandomString(16)
		articleList[i].HtmlContent = "<h1>" + this.RandomString(16) + "</h1>"
		this.BlogCsdnCrawlAo.ModArticle(articleList[i].Id, articleList[i])
	}

	articleList2 := this.getAllArticle(name)
	this.AssertEqual(len(articleList), len(articleList2))
	for i := 0; i != len(articleList); i++ {
		this.AssertEqual(articleList[i].Title, articleList2[i].Title)
		this.AssertEqual(articleList[i].Content, articleList2[i].Content)
		this.AssertEqual(articleList[i].HtmlContent, articleList2[i].HtmlContent)
	}
}

func (this *blogCsdnCrawlAoTest) testArticle(name string) {
	this.testArticleClear(name)

	data := this.testArticleAdd(name)

	this.testArticleMod(name, data)

	this.testArticleClear(name)
}

func (this *blogCsdnCrawlAoTest) TestBasic() {
	username := "fishmei2"
	password := "woaini520"

	this.BlogCsdnCrawlAo.Login(username, password)

	//测试文章
	//this.testArticle(username)

	//测试分类
	//this.testCategory()
}
