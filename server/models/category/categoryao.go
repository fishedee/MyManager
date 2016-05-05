package category

import (
	"fmt"
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type CategoryAoModel struct {
	Model
	CategoryDb CategoryDbModel
}

func (this *CategoryAoModel) Search(userId int, where Category, limit CommonPage) Categorys {
	where.UserId = userId
	return this.CategoryDb.Search(where, limit)
}

func (this *CategoryAoModel) Get(userId int, categoryId int) Category {
	categoryInfo := this.CategoryDb.Get(categoryId)
	if categoryInfo.UserId != userId {
		Throw(1, "你没有权利查看或编辑等操作")
	}
	return categoryInfo
}

func (this *CategoryAoModel) Del(userId int, categoryId int) {
	this.Get(userId, categoryId)

	this.CategoryDb.Del(categoryId)

	this.Queue.Publish(CategoryQueueEnum.EVENT_DEL, categoryId)
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

func (this *CategoryAoModel) TestQueue(id int, str string) {
	fmt.Println(id, str)
}

func init() {
	InitDaemon(func(this *CategoryAoModel) {
		this.Queue.Consume("uu", this.TestQueue)
	})
}
