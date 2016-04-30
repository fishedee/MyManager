package blog

import (
	"github.com/fishedee/language"
	"mymanager/models/common"
	"strconv"
)

type BlogSyncAutoDbModel struct {
	common.BaseModel
}

func (this *BlogSyncAutoDbModel) Search(where BlogSyncAuto, limit common.CommonPage) BlogSyncAutos {
	db := this.DB.NewSession()

	if limit.PageSize == 0 && limit.PageIndex == 0 {
		return BlogSyncAutos{
			Count: 0,
			Data:  []BlogSyncAuto{},
		}
	}

	if where.GitUrl != "" {
		db = db.And("gitUrl like ?", where.GitUrl)
	}
	if where.UserId != 0 {
		db = db.And("userId = ?", where.UserId)
	}

	data := []BlogSyncAuto{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return BlogSyncAutos{
		Count: int(count),
		Data:  data,
	}
}

func (this *BlogSyncAutoDbModel) GetAll() []BlogSyncAuto {
	var data []BlogSyncAuto
	err := this.DB.Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func (this *BlogSyncAutoDbModel) Get(blogSyncAutoId int) BlogSyncAuto {
	var blogSyncs []BlogSyncAuto
	err := this.DB.Where("blogSyncAutoId = ?", blogSyncAutoId).Find(&blogSyncs)
	if err != nil {
		panic(err)
	}
	if len(blogSyncs) == 0 {
		language.Throw(1, "不存在该记录"+strconv.Itoa(blogSyncAutoId))
	}
	return blogSyncs[0]
}

func (this *BlogSyncAutoDbModel) Add(blogSync BlogSyncAuto) int {
	_, err := this.DB.Insert(&blogSync)
	if err != nil {
		panic(err)
	}
	return blogSync.BlogSyncAutoId
}

func (this *BlogSyncAutoDbModel) Mod(blogSyncAutoId int, blogSync BlogSyncAuto) {
	_, err := this.DB.Where("blogSyncAutoId = ?", blogSyncAutoId).Update(&blogSync)
	if err != nil {
		panic(err)
	}
}

func (this *BlogSyncAutoDbModel) Del(blogSyncAutoId int) {
	_, err := this.DB.Where("blogSyncAutoId = ?", blogSyncAutoId).Delete(&BlogSyncAuto{})
	if err != nil {
		panic(err)
	}
}
