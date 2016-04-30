package category

import (
	"github.com/fishedee/language"
	"mymanager/models/common"
	"strconv"
)

type CategoryDbModel struct {
	common.BaseModel
}

func (this *CategoryDbModel) Search(where Category, limit common.CommonPage) Categorys {
	db := this.DB.NewSession()

	if where.Name != "" {
		db = db.Where("name like ?", "%"+where.Name+"%")
	}
	if where.Remark != "" {
		db = db.Where("remark like ?", "%"+where.Remark+"%")
	}
	if where.UserId != 0 {
		db = db.Where("userId = ? ", where.UserId)
	}

	data := []Category{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return Categorys{
		Count: int(count),
		Data:  data,
	}
}

func (this *CategoryDbModel) Get(categoryId int) Category {
	var categorys []Category
	err := this.DB.Where("categoryId = ?", categoryId).Find(&categorys)
	if err != nil {
		panic(err)
	}
	if len(categorys) == 0 {
		language.Throw(1, "不存在该分类"+strconv.Itoa(categoryId))
	}
	return categorys[0]
}

func (this *CategoryDbModel) Del(categoryId int) {
	_, err := this.DB.Where("categoryId = ?", categoryId).Delete(&Category{})
	if err != nil {
		panic(err)
	}
}

func (this *CategoryDbModel) Add(category Category) {
	_, err := this.DB.Insert(category)
	if err != nil {
		panic(err)
	}
}

func (this *CategoryDbModel) Mod(categoryId int, category Category) {
	_, err := this.DB.Where("categoryId = ?", categoryId).Update(&category)
	if err != nil {
		panic(err)
	}
}
