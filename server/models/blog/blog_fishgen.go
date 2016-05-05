package blog

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
)

func (this *BlogCsdnAoModel) Sync_WithError(accessToken string, syncType int, src Blog, progressUpdater BlogSyncProgress) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Sync(accessToken, syncType, src, progressUpdater)
	return
}

func (this *BlogCsdnCrawlAoModel) Login_WithError(name string, password string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Login(name, password)
	return
}

func (this *BlogCsdnCrawlAoModel) GetCategoryList_WithError() (_fishgen1 []BlogCategory, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetCategoryList()
	return
}

func (this *BlogCsdnCrawlAoModel) AddCategory_WithError(category BlogCategory) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddCategory(category)
	return
}

func (this *BlogCsdnCrawlAoModel) DelCategory_WithError(id int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DelCategory(id)
	return
}

func (this *BlogCsdnCrawlAoModel) ModCategory_WithError(id int, data BlogCategory) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModCategory(id, data)
	return
}

func (this *BlogCsdnCrawlAoModel) ModArticle_WithError(id int, data BlogArticle) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModArticle(id, data)
	return
}

func (this *BlogCsdnCrawlAoModel) AddArticle_WithError(data BlogArticle) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.AddArticle(data)
	return
}

func (this *BlogCsdnCrawlAoModel) DelArticle_WithError(id int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DelArticle(id)
	return
}

func (this *BlogCsdnCrawlAoModel) GetArticle_WithError(id int, name string) (_fishgen1 BlogArticle, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetArticle(id, name)
	return
}

func (this *BlogCsdnCrawlAoModel) GetArticleList_WithError(page int) (_fishgen1 []BlogArticle, _fishgen2 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1, _fishgen2 = this.GetArticleList(page)
	return
}

func (this *BlogGitAoModel) Get_WithError(gitUrl string, progressUpdater BlogSyncProgress) (_fishgen1 Blog, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(gitUrl, progressUpdater)
	return
}

func (this *BlogSyncAoModel) SearchAuto_WithError(userId int, where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.SearchAuto(userId, where, limit)
	return
}

func (this *BlogSyncAoModel) GetAuto_WithError(userId int, blogSyncAutoId int) (_fishgen1 BlogSyncAuto, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetAuto(userId, blogSyncAutoId)
	return
}

func (this *BlogSyncAoModel) DelAuto_WithError(userId int, blogSyncAutoId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DelAuto(userId, blogSyncAutoId)
	return
}

func (this *BlogSyncAoModel) AddAuto_WithError(userId int, blogSyncAuto BlogSyncAuto) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddAuto(userId, blogSyncAuto)
	return
}

func (this *BlogSyncAoModel) ModAuto_WithError(userId int, blogSyncAutoId int, blogSyncAuto BlogSyncAuto) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModAuto(userId, blogSyncAutoId, blogSyncAuto)
	return
}

func (this *BlogSyncAoModel) SearchTask_WithError(userId int, where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.SearchTask(userId, where, limit)
	return
}

func (this *BlogSyncAoModel) AddTask_WithError(userId int, accessToken string, gitUrl string, syncType int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddTask(userId, accessToken, gitUrl, syncType)
	return
}

func (this *BlogSyncAoModel) GetTask_WithError(userId int, blogSyncId int) (_fishgen1 BlogSync, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetTask(userId, blogSyncId)
	return
}

func (this *BlogSyncAoModel) RestartTask_WithError(userId int, blogSyncId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.RestartTask(userId, blogSyncId)
	return
}

func (this *BlogSyncAutoDbModel) Search_WithError(where BlogSyncAuto, limit CommonPage) (_fishgen1 BlogSyncAutos, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *BlogSyncAutoDbModel) GetAll_WithError() (_fishgen1 []BlogSyncAuto, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetAll()
	return
}

func (this *BlogSyncAutoDbModel) Get_WithError(blogSyncAutoId int) (_fishgen1 BlogSyncAuto, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(blogSyncAutoId)
	return
}

func (this *BlogSyncAutoDbModel) Add_WithError(blogSync BlogSyncAuto) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Add(blogSync)
	return
}

func (this *BlogSyncAutoDbModel) Mod_WithError(blogSyncAutoId int, blogSync BlogSyncAuto) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(blogSyncAutoId, blogSync)
	return
}

func (this *BlogSyncAutoDbModel) Del_WithError(blogSyncAutoId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(blogSyncAutoId)
	return
}

func (this *BlogSyncDbModel) Search_WithError(where BlogSync, limit CommonPage) (_fishgen1 BlogSyncs, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *BlogSyncDbModel) Get_WithError(blogSyncId int) (_fishgen1 BlogSync, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(blogSyncId)
	return
}

func (this *BlogSyncDbModel) Add_WithError(blogSync BlogSync) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Add(blogSync)
	return
}

func (this *BlogSyncDbModel) Mod_WithError(blogSyncId int, blogSync BlogSync) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(blogSyncId, blogSync)
	return
}
