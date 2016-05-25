package category

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
)

func (this *CategoryAoModel) Search_WithError(userId int, where Category, limit CommonPage) (_fishgen1 Categorys, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(userId, where, limit)
	return
}

func (this *CategoryAoModel) Get_WithError(userId int, categoryId int) (_fishgen1 Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId, categoryId)
	return
}

func (this *CategoryAoModel) Del_WithError(userId int, categoryId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId, categoryId)
	return
}

func (this *CategoryAoModel) Add_WithError(userId int, categoryInfo Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(userId, categoryInfo)
	return
}

func (this *CategoryAoModel) Mod_WithError(userId int, categoryId int, categoryInfo Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(userId, categoryId, categoryInfo)
	return
}

func (this *CategoryDbModel) Search_WithError(where Category, limit CommonPage) (_fishgen1 Categorys, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *CategoryDbModel) Get_WithError(categoryId int) (_fishgen1 Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(categoryId)
	return
}

func (this *CategoryDbModel) Del_WithError(categoryId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(categoryId)
	return
}

func (this *CategoryDbModel) Add_WithError(category Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(category)
	return
}

func (this *CategoryDbModel) Mod_WithError(categoryId int, category Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(categoryId, category)
	return
}
