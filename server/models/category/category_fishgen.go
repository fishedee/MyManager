package category

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type CategoryAoModel interface {
	Search(userId int, where Category, limit CommonPage) (_fishgen1 Categorys)
	Search_WithError(userId int, where Category, limit CommonPage) (_fishgen1 Categorys, _fishgenErr Exception)
	Get(userId int, categoryId int) (_fishgen1 Category)
	Get_WithError(userId int, categoryId int) (_fishgen1 Category, _fishgenErr Exception)
	Del(userId int, categoryId int)
	Del_WithError(userId int, categoryId int) (_fishgenErr Exception)
	Add(userId int, categoryInfo Category)
	Add_WithError(userId int, categoryInfo Category) (_fishgenErr Exception)
	Mod(userId int, categoryId int, categoryInfo Category)
	Mod_WithError(userId int, categoryId int, categoryInfo Category) (_fishgenErr Exception)
	TestQueue(id int, str string)
	TestQueue_WithError(id int, str string) (_fishgenErr Exception)
}

type CategoryAoTest interface {
	InitEmpty()
	InitSample()
	TestBasic()
}

type CategoryDbModel interface {
	Search(where Category, limit CommonPage) (_fishgen1 Categorys)
	Search_WithError(where Category, limit CommonPage) (_fishgen1 Categorys, _fishgenErr Exception)
	Get(categoryId int) (_fishgen1 Category)
	Get_WithError(categoryId int) (_fishgen1 Category, _fishgenErr Exception)
	Del(categoryId int)
	Del_WithError(categoryId int) (_fishgenErr Exception)
	Add(category Category)
	Add_WithError(category Category) (_fishgenErr Exception)
	Mod(categoryId int, category Category)
	Mod_WithError(categoryId int, category Category) (_fishgenErr Exception)
}

func (this *categoryAoModel) Search_WithError(userId int, where Category, limit CommonPage) (_fishgen1 Categorys, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(userId, where, limit)
	return
}

func (this *categoryAoModel) Get_WithError(userId int, categoryId int) (_fishgen1 Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId, categoryId)
	return
}

func (this *categoryAoModel) Del_WithError(userId int, categoryId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId, categoryId)
	return
}

func (this *categoryAoModel) Add_WithError(userId int, categoryInfo Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(userId, categoryInfo)
	return
}

func (this *categoryAoModel) Mod_WithError(userId int, categoryId int, categoryInfo Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(userId, categoryId, categoryInfo)
	return
}

func (this *categoryAoModel) TestQueue_WithError(id int, str string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.TestQueue(id, str)
	return
}

func (this *categoryDbModel) Search_WithError(where Category, limit CommonPage) (_fishgen1 Categorys, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *categoryDbModel) Get_WithError(categoryId int) (_fishgen1 Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(categoryId)
	return
}

func (this *categoryDbModel) Del_WithError(categoryId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(categoryId)
	return
}

func (this *categoryDbModel) Add_WithError(category Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(category)
	return
}

func (this *categoryDbModel) Mod_WithError(categoryId int, category Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(categoryId, category)
	return
}
func init() {
	v0 := CategoryAoModel(&categoryAoModel{})
	InitModel(&v0)
	v1 := CategoryAoTest(&categoryAoTest{})
	InitTest(&v1)
	v2 := CategoryDbModel(&categoryDbModel{})
	InitModel(&v2)
}
