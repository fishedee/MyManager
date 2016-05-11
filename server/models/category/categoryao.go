package category

import (
	"fmt"
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type categoryAoModel struct {
	Model
	CategoryDb CategoryDbModel
}

func (this *categoryAoModel) Search(userId int, where Category, limit CommonPage) Categorys {
	where.UserId = userId
	return this.CategoryDb.Search(where, limit)
}

func (this *categoryAoModel) Get(userId int, categoryId int) Category {
	categoryInfo := this.CategoryDb.Get(categoryId)
	if categoryInfo.UserId != userId {
		Throw(1, "你没有权利查看或编辑等操作")
	}
	return categoryInfo
}

func (this *categoryAoModel) Del(userId int, categoryId int) {
	this.Get(userId, categoryId)

	this.CategoryDb.Del(categoryId)

	this.Queue.Publish(CategoryQueueEnum.EVENT_DEL, categoryId)
}

func (this *categoryAoModel) Add(userId int, categoryInfo Category) {
	categoryInfo.UserId = userId
	this.CategoryDb.Add(categoryInfo)
}

func (this *categoryAoModel) Mod(userId int, categoryId int, categoryInfo Category) {
	this.Get(userId, categoryId)

	categoryInfo.UserId = userId
	this.CategoryDb.Mod(categoryId, categoryInfo)
}

func (this *categoryAoModel) TestQueue(id int, str string) {
	fmt.Println(id, str)
}

func init() {
	InitDaemon(func(this *categoryAoModel) {
		this.Queue.Consume("uu", this.TestQueue)
	})
}
