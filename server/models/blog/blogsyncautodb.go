package blog

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type blogSyncAutoDbModel struct {
	Model
}

func (this *blogSyncAutoDbModel) Search(where BlogSyncAuto, limit CommonPage) BlogSyncAutos {
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

func (this *blogSyncAutoDbModel) GetAll() []BlogSyncAuto {
	var data []BlogSyncAuto
	err := this.DB.Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func (this *blogSyncAutoDbModel) Get(blogSyncAutoId int) BlogSyncAuto {
	var blogSyncs []BlogSyncAuto
	err := this.DB.Where("blogSyncAutoId = ?", blogSyncAutoId).Find(&blogSyncs)
	if err != nil {
		panic(err)
	}
	if len(blogSyncs) == 0 {
		Throw(1, "不存在该记录"+strconv.Itoa(blogSyncAutoId))
	}
	return blogSyncs[0]
}

func (this *blogSyncAutoDbModel) Add(blogSync BlogSyncAuto) int {
	_, err := this.DB.Insert(&blogSync)
	if err != nil {
		panic(err)
	}
	return blogSync.BlogSyncAutoId
}

func (this *blogSyncAutoDbModel) Mod(blogSyncAutoId int, blogSync BlogSyncAuto) {
	_, err := this.DB.Where("blogSyncAutoId = ?", blogSyncAutoId).Update(&blogSync)
	if err != nil {
		panic(err)
	}
}

func (this *blogSyncAutoDbModel) Del(blogSyncAutoId int) {
	_, err := this.DB.Where("blogSyncAutoId = ?", blogSyncAutoId).Delete(&BlogSyncAuto{})
	if err != nil {
		panic(err)
	}
}
