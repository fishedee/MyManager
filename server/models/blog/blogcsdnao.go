package blog

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/sdk"
	. "mymanager/models/common"
)

type BlogCsdnAoModel struct {
	BaseModel
}

func (this *BlogCsdnAoModel) getSdk() *CsdnSdk {
	return &CsdnSdk{
		AppKey:    "1100258",
		AppSecert: "47720345d6024a3eb65ee0620d6b7748",
	}
}

func (this *BlogCsdnAoModel) getData(accessToken string) Blog {
	data := Blog{}
	categoryList, err := this.getSdk().GetCategoryList(CsdnSdkGetCategoryListRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		panic(err)
	}
	for _, singleCategory := range categoryList {
		data.Categorys = append(data.Categorys, BlogCategory{
			Name: singleCategory.Name,
		})
	}
	articleList, err := this.getSdk().GetArticleList(CsdnSdkGetArticleListRequest{
		AccessToken: accessToken,
		Status:      "enabled",
		Page:        1,
		Size:        1000,
	})
	if err != nil {
		panic(err)
	}
	for _, singleArticle := range articleList.List {
		detailArticle, err := this.getSdk().GetArticle(CsdnSdkGetArticleRequest{
			AccessToken: accessToken,
			Id:          singleArticle.Id,
		})
		if err != nil {
			panic(err)
		}
		data.Articles = append(data.Articles, BlogArticle{
			Id:       detailArticle.Id,
			Title:    detailArticle.Title,
			Content:  detailArticle.Content,
			Category: detailArticle.Categories,
		})
	}
	return data
}

func (this *BlogCsdnAoModel) diffData(src Blog, dist Blog, syncType int) BlogDiff {
	result := BlogDiff{}

	//比较Category
	mapSrcCategory := ArrayColumnMap(src.Categorys, "Name").(map[string]BlogCategory)
	mapDistCategory := ArrayColumnMap(dist.Categorys, "Name").(map[string]BlogCategory)

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
	mapSrcArticle := ArrayColumnMap(src.Articles, "Name").(map[string]BlogArticle)
	mapDistArticle := ArrayColumnMap(dist.Articles, "Name").(map[string]BlogArticle)
	for name, singleArticle := range mapSrcArticle {
		_, ok := mapDistArticle[name]
		if !ok {
			result.AddArticles = append(result.AddArticles, singleArticle)
		}
	}

	if syncType == BlogSyncTypeEnum.TYPE_ALL_UPDATE {
		for name, singleArticle := range mapDistArticle {
			singleSrcArticle, ok := mapSrcArticle[name]
			if !ok {
				result.DelArticles = append(result.DelArticles, singleArticle)
			} else {
				if singleArticle.Content != singleSrcArticle.Content ||
					singleArticle.Category != singleSrcArticle.Category {
					result.ModArticles = append(result.ModArticles, singleArticle)
				}
			}
		}
	}
	return result
}

func (this *BlogCsdnAoModel) setData(accessToken string, diff BlogDiff) {
	//FIXME csdn开放接口不支持增加和删除category

	//FIXME csdn开放接口不支持删除article

	//添加文章
	for _, singleArticle := range diff.AddArticles {
		_, err := this.getSdk().SaveArticle(CsdnSdkSaveArticleRequest{
			AccessToken: accessToken,
			Title:       singleArticle.Title,
			Content:     singleArticle.Content,
			Categories:  singleArticle.Category,
		})
		if err != nil {
			panic(err)
		}
	}

	//修改文章
	for _, singleArticle := range diff.ModArticles {
		_, err := this.getSdk().SaveArticle(CsdnSdkSaveArticleRequest{
			AccessToken: accessToken,
			Id:          singleArticle.Id,
			Title:       singleArticle.Title,
			Content:     singleArticle.Content,
			Categories:  singleArticle.Category,
		})
		if err != nil {
			panic(err)
		}
	}
}

func (this *BlogCsdnAoModel) Sync(accessToken string, syncType int, src Blog, progressUpdater BlogSyncProgress) {
	//获取数据
	progressUpdater("正在获取csdn博客数据中")
	dist := this.getData(accessToken)

	//比对数据
	progressUpdater("正在比较数据中")
	diff := this.diffData(src, dist, syncType)

	//更新数据
	progressUpdater("正在设置csdn数据中")
	this.setData(accessToken, diff)
}

func (this *BlogCsdnAoModel) GetAuthUrl(redirectUrl string) string {
	data, err := this.getSdk().GetAuthUrl(redirectUrl)
	if err != nil {
		Throw(1, err.Error())
	}
	return data
}

func (this *BlogCsdnAoModel) GetAccessToken(redirectUrl string, code string) string {
	data, err := this.getSdk().GetAccessToken(redirectUrl, code)
	if err != nil {
		panic(err)
	}
	return data.AccessToken
}
