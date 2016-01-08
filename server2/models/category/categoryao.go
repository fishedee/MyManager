package category

import (
	. "mymanager/models/common"
	. "github.com/fishedee/language"
)

type CategoryAoModel struct {
	BaseModel
	CategoryDb CategoryDbModel
}

func (this *CategoryAoModel) Search(userId int, where Category, limit CommonPage) Categorys {
	where.UserId = userId
	return this.CategoryDb.Search(where, limit)
}

func (this *CategoryAoModel) Get(userId int, categoryId int) Category {
	categoryInfo := this.CategoryDb.Get(categoryId)
	if categoryInfo.UserId != userId {
		Throw(1, "你没有该权限")
	}
	return categoryInfo
}

func (this *CategoryAoModel) Del(userId int, categoryId int) {
	this.Get(userId, categoryId)

	this.CategoryDb.Del(categoryId)

	this.Queue.Publish("category_del", categoryId)
}

func (this *CategoryAoModel) Add(userId int, categoryInfo Category) {
	categoryInfo.UserId = userId
	this.CategoryDb.Add(categoryInfo)
}

func (this *CategoryAoModel) Mod(userId int, categoryId int, categoryInfo Category) {
	this.Get(userId, categoryId)

	categoryInfo.UserId = userId
	this.CategoryDb.Mod(categoryId, categoryInfo)
}
