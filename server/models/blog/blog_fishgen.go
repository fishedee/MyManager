package blog

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type BlogCsdnAoModel interface {
	Sync(accessToken string, syncType int, src Blog, progressUpdater BlogSyncProgress)
	Sync_WithError(accessToken string, syncType int, src Blog, progressUpdater BlogSyncProgress) (_fishgenErr Exception)
}

type BlogCsdnCrawlAoModel interface {
	Login(name string, password string)
	Login_WithError(name string, password string) (_fishgenErr Exception)
	GetCategoryList() (_fishgen1 []BlogCategory)
	GetCategoryList_WithError() (_fishgen1 []BlogCategory, _fishgenErr Exception)
	AddCategory(category BlogCategory)
	AddCategory_WithError(category BlogCategory) (_fishgenErr Exception)
	DelCategory(id int)
	DelCategory_WithError(id int) (_fishgenErr Exception)
	ModCategory(id int, data BlogCategory)
	ModCategory_WithError(id int, data BlogCategory) (_fishgenErr Exception)
	ModArticle(id int, data BlogArticle)
	ModArticle_WithError(id int, data BlogArticle) (_fishgenErr Exception)
	AddArticle(data BlogArticle) (_fishgen1 int)
	AddArticle_WithError(data BlogArticle) (_fishgen1 int, _fishgenErr Exception)
	DelArticle(id int)
	DelArticle_WithError(id int) (_fishgenErr Exception)
	GetArticle(id int, name string) (_fishgen1 BlogArticle)
	GetArticle_WithError(id int, name string) (_fishgen1 BlogArticle, _fishgenErr Exception)
	GetArticleList(page int) (_fishgen1 []BlogArticle, _fishgen2 int)
	GetArticleList_WithError(page int) (_fishgen1 []BlogArticle, _fishgen2 int, _fishgenErr Exception)
}

type BlogCsdnCrawlAoTest interface {
	TestBasic()
}

type BlogGitAoModel interface {
	Get(gitUrl string, progressUpdater BlogSyncProgress) (_fishgen1 Blog)
	Get_WithError(gitUrl string, progressUpdater BlogSyncProgress) (_fishgen1 Blog, _fishgenErr Exception)
}

type BlogGitAoTest interface {
	TestMarkdown()
	TestGit()
}

type BlogSyncAoModel interface {
	SearchAuto(userId int, where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos)
	SearchAuto_WithError(userId int, where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos, _fishgenErr Exception)
	GetAuto(userId int, blogSyncAutoId int) (_fishgen1 BlogSyncAuto)
	GetAuto_WithError(userId int, blogSyncAutoId int) (_fishgen1 BlogSyncAuto, _fishgenErr Exception)
	DelAuto(userId int, blogSyncAutoId int)
	DelAuto_WithError(userId int, blogSyncAutoId int) (_fishgenErr Exception)
	AddAuto(userId int, blogSyncAuto BlogSyncAuto)
	AddAuto_WithError(userId int, blogSyncAuto BlogSyncAuto) (_fishgenErr Exception)
	ModAuto(userId int, blogSyncAutoId int, blogSyncAuto BlogSyncAuto)
	ModAuto_WithError(userId int, blogSyncAutoId int, blogSyncAuto BlogSyncAuto) (_fishgenErr Exception)
	SearchTask(userId int, where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs)
	SearchTask_WithError(userId int, where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs, _fishgenErr Exception)
	AddTask(userId int, accessToken string, gitUrl string, syncType int)
	AddTask_WithError(userId int, accessToken string, gitUrl string, syncType int) (_fishgenErr Exception)
	GetTask(userId int, blogSyncId int) (_fishgen1 BlogSync)
	GetTask_WithError(userId int, blogSyncId int) (_fishgen1 BlogSync, _fishgenErr Exception)
	RestartTask(userId int, blogSyncId int)
	RestartTask_WithError(userId int, blogSyncId int) (_fishgenErr Exception)
}

type BlogSyncAutoDbModel interface {
	Search(where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos)
	Search_WithError(where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos, _fishgenErr Exception)
	GetAll() (_fishgen1 []BlogSyncAuto)
	GetAll_WithError() (_fishgen1 []BlogSyncAuto, _fishgenErr Exception)
	Get(blogSyncAutoId int) (_fishgen1 BlogSyncAuto)
	Get_WithError(blogSyncAutoId int) (_fishgen1 BlogSyncAuto, _fishgenErr Exception)
	Add(blogSync BlogSyncAuto) (_fishgen1 int)
	Add_WithError(blogSync BlogSyncAuto) (_fishgen1 int, _fishgenErr Exception)
	Mod(blogSyncAutoId int, blogSync BlogSyncAuto)
	Mod_WithError(blogSyncAutoId int, blogSync BlogSyncAuto) (_fishgenErr Exception)
	Del(blogSyncAutoId int)
	Del_WithError(blogSyncAutoId int) (_fishgenErr Exception)
}

type BlogSyncDbModel interface {
	Search(where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs)
	Search_WithError(where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs, _fishgenErr Exception)
	Get(blogSyncId int) (_fishgen1 BlogSync)
	Get_WithError(blogSyncId int) (_fishgen1 BlogSync, _fishgenErr Exception)
	Add(blogSync BlogSync) (_fishgen1 int)
	Add_WithError(blogSync BlogSync) (_fishgen1 int, _fishgenErr Exception)
	Mod(blogSyncId int, blogSync BlogSync)
	Mod_WithError(blogSyncId int, blogSync BlogSync) (_fishgenErr Exception)
}

func (this *blogCsdnAoModel) Sync_WithError(accessToken string, syncType int, src Blog, progressUpdater BlogSyncProgress) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Sync(accessToken, syncType, src, progressUpdater)
	return
}

func (this *blogCsdnCrawlAoModel) Login_WithError(name string, password string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Login(name, password)
	return
}

func (this *blogCsdnCrawlAoModel) GetCategoryList_WithError() (_fishgen1 []BlogCategory, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetCategoryList()
	return
}

func (this *blogCsdnCrawlAoModel) AddCategory_WithError(category BlogCategory) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddCategory(category)
	return
}

func (this *blogCsdnCrawlAoModel) DelCategory_WithError(id int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DelCategory(id)
	return
}

func (this *blogCsdnCrawlAoModel) ModCategory_WithError(id int, data BlogCategory) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModCategory(id, data)
	return
}

func (this *blogCsdnCrawlAoModel) ModArticle_WithError(id int, data BlogArticle) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModArticle(id, data)
	return
}

func (this *blogCsdnCrawlAoModel) AddArticle_WithError(data BlogArticle) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.AddArticle(data)
	return
}

func (this *blogCsdnCrawlAoModel) DelArticle_WithError(id int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DelArticle(id)
	return
}

func (this *blogCsdnCrawlAoModel) GetArticle_WithError(id int, name string) (_fishgen1 BlogArticle, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetArticle(id, name)
	return
}

func (this *blogCsdnCrawlAoModel) GetArticleList_WithError(page int) (_fishgen1 []BlogArticle, _fishgen2 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1, _fishgen2 = this.GetArticleList(page)
	return
}

func (this *blogGitAoModel) Get_WithError(gitUrl string, progressUpdater BlogSyncProgress) (_fishgen1 Blog, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(gitUrl, progressUpdater)
	return
}

func (this *blogSyncAoModel) SearchAuto_WithError(userId int, where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.SearchAuto(userId, where, limit)
	return
}

func (this *blogSyncAoModel) GetAuto_WithError(userId int, blogSyncAutoId int) (_fishgen1 BlogSyncAuto, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetAuto(userId, blogSyncAutoId)
	return
}

func (this *blogSyncAoModel) DelAuto_WithError(userId int, blogSyncAutoId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DelAuto(userId, blogSyncAutoId)
	return
}

func (this *blogSyncAoModel) AddAuto_WithError(userId int, blogSyncAuto BlogSyncAuto) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddAuto(userId, blogSyncAuto)
	return
}

func (this *blogSyncAoModel) ModAuto_WithError(userId int, blogSyncAutoId int, blogSyncAuto BlogSyncAuto) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModAuto(userId, blogSyncAutoId, blogSyncAuto)
	return
}

func (this *blogSyncAoModel) SearchTask_WithError(userId int, where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.SearchTask(userId, where, limit)
	return
}

func (this *blogSyncAoModel) AddTask_WithError(userId int, accessToken string, gitUrl string, syncType int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddTask(userId, accessToken, gitUrl, syncType)
	return
}

func (this *blogSyncAoModel) GetTask_WithError(userId int, blogSyncId int) (_fishgen1 BlogSync, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetTask(userId, blogSyncId)
	return
}

func (this *blogSyncAoModel) RestartTask_WithError(userId int, blogSyncId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.RestartTask(userId, blogSyncId)
	return
}

func (this *blogSyncAutoDbModel) Search_WithError(where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *blogSyncAutoDbModel) GetAll_WithError() (_fishgen1 []BlogSyncAuto, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetAll()
	return
}

func (this *blogSyncAutoDbModel) Get_WithError(blogSyncAutoId int) (_fishgen1 BlogSyncAuto, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(blogSyncAutoId)
	return
}

func (this *blogSyncAutoDbModel) Add_WithError(blogSync BlogSyncAuto) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Add(blogSync)
	return
}

func (this *blogSyncAutoDbModel) Mod_WithError(blogSyncAutoId int, blogSync BlogSyncAuto) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(blogSyncAutoId, blogSync)
	return
}

func (this *blogSyncAutoDbModel) Del_WithError(blogSyncAutoId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(blogSyncAutoId)
	return
}

func (this *blogSyncDbModel) Search_WithError(where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *blogSyncDbModel) Get_WithError(blogSyncId int) (_fishgen1 BlogSync, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(blogSyncId)
	return
}

func (this *blogSyncDbModel) Add_WithError(blogSync BlogSync) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Add(blogSync)
	return
}

func (this *blogSyncDbModel) Mod_WithError(blogSyncId int, blogSync BlogSync) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(blogSyncId, blogSync)
	return
}
func init() {
	v0 := BlogCsdnAoModel(&blogCsdnAoModel{})
	InitModel(&v0)
	v1 := BlogCsdnCrawlAoModel(&blogCsdnCrawlAoModel{})
	InitModel(&v1)
	v2 := BlogCsdnCrawlAoTest(&blogCsdnCrawlAoTest{})
	InitTest(&v2)
	v3 := BlogGitAoModel(&blogGitAoModel{})
	InitModel(&v3)
	v4 := BlogGitAoTest(&blogGitAoTest{})
	InitTest(&v4)
	v5 := BlogSyncAoModel(&blogSyncAoModel{})
	InitModel(&v5)
	v6 := BlogSyncAutoDbModel(&blogSyncAutoDbModel{})
	InitModel(&v6)
	v7 := BlogSyncDbModel(&blogSyncDbModel{})
	InitModel(&v7)
}
