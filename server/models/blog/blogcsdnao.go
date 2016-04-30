package blog

import (
	"github.com/fishedee/language"
	"mymanager/models/common"
	"strings"
)

type BlogCsdnAoModel struct {
	common.BaseModel
	BlogCsdnCrawlAo BlogCsdnCrawlAoModel
}

func (this *BlogCsdnAoModel) login(accessToken string) string {
	accessTokenArray := strings.Split(accessToken, ",")
	if len(accessTokenArray) != 2 {
		language.Throw(1, "请输入用逗号分割的accessToken")
	}
	username := accessTokenArray[0]
	password := accessTokenArray[1]
	if username == "" || password == "" {
		language.Throw(1, "用户名与密码都不能为空")
	}
	this.BlogCsdnCrawlAo.Login(username, password)
	return username
}

func (this *BlogCsdnAoModel) getData(username string) Blog {
	//获取category数据
	categoryList := this.BlogCsdnCrawlAo.GetCategoryList()

	//获取文章数据
	var articleList []BlogArticle
	pageIndex := 0
	totalCount := 10
	for len(articleList) < totalCount {
		var singleArticleList []BlogArticle
		singleArticleList, totalCount = this.BlogCsdnCrawlAo.GetArticleList(pageIndex)
		if len(singleArticleList) == 0 {
			break
		}
		for _, singleArticle := range singleArticleList {
			singleArticleInfo := this.BlogCsdnCrawlAo.GetArticle(singleArticle.Id, username)
			articleList = append(articleList, singleArticleInfo)
		}
	}
	return Blog{
		Articles:  articleList,
		Categorys: categoryList,
	}
}

func (this *BlogCsdnAoModel) diffData(src Blog, dist Blog, syncType int) BlogDiff {
	result := BlogDiff{}

	//比较Category
	mapSrcCategory := language.ArrayColumnMap(src.Categorys, "Name").(map[string]BlogCategory)
	mapDistCategory := language.ArrayColumnMap(dist.Categorys, "Name").(map[string]BlogCategory)

	for name, singleCategory := range mapSrcCategory {
		_, ok := mapDistCategory[name]
		if !ok {
			result.AddCategorys = append(result.AddCategorys, singleCategory)
		}
	}

	if syncType == BlogSyncTypeEnum.TYPE_ALL_UPDATE {
		for name, singleCategory := range mapDistCategory {
			_, ok := mapSrcCategory[name]
			if !ok {
				result.DelCategorys = append(result.DelCategorys, singleCategory)
			}
		}
	}

	//比较Article
	mapSrcArticle := language.ArrayColumnMap(src.Articles, "Title").(map[string]BlogArticle)
	mapDistArticle := language.ArrayColumnMap(dist.Articles, "Title").(map[string]BlogArticle)
	for name, singleArticle := range mapSrcArticle {
		_, ok := mapDistArticle[name]
		if !ok {
			result.AddArticles = append(result.AddArticles, singleArticle)
		}
	}

	for name, singleArticle := range mapDistArticle {
		singleSrcArticle, ok := mapSrcArticle[name]
		if !ok {
			if syncType == BlogSyncTypeEnum.TYPE_ALL_UPDATE {
				result.DelArticles = append(result.DelArticles, singleArticle)
			}
		} else {
			if strings.Trim(singleArticle.Content, " ") != strings.Trim(singleSrcArticle.Content, " ") ||
				singleArticle.Category != singleSrcArticle.Category {
				singleArticle.HtmlContent = singleSrcArticle.HtmlContent
				singleArticle.Content = singleSrcArticle.Content
				singleArticle.Category = singleSrcArticle.Category
				result.ModArticles = append(result.ModArticles, singleArticle)
			}
		}
	}
	return result
}

func (this *BlogCsdnAoModel) setData(username string, diff BlogDiff) {
	//添加category
	for _, singleCategory := range diff.AddCategorys {
		this.BlogCsdnCrawlAo.AddCategory(BlogCategory{
			Name: singleCategory.Name,
		})
	}

	//删除category
	for _, singleCategory := range diff.DelCategorys {
		this.BlogCsdnCrawlAo.DelCategory(singleCategory.Id)
	}

	//添加article
	for _, singleArticle := range diff.AddArticles {
		this.BlogCsdnCrawlAo.AddArticle(singleArticle)
	}

	//删除article
	for _, singleArticle := range diff.DelArticles {
		this.BlogCsdnCrawlAo.DelArticle(singleArticle.Id)
	}

	//修改article
	for _, singleArticle := range diff.ModArticles {
		this.BlogCsdnCrawlAo.ModArticle(singleArticle.Id, singleArticle)
	}
}

func (this *BlogCsdnAoModel) Sync(accessToken string, syncType int, src Blog, progressUpdater BlogSyncProgress) {
	//登录csdn博客中
	progressUpdater("正在登录博客账号中")
	user := this.login(accessToken)

	//获取数据
	progressUpdater("正在获取csdn博客数据中")
	dist := this.getData(user)

	//比对数据
	progressUpdater("正在比较数据中")
	diff := this.diffData(src, dist, syncType)

	//更新数据
	progressUpdater("正在设置csdn数据中")
	this.setData(user, diff)
}
