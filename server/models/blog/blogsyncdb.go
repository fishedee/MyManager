package blog

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type BlogSyncDbModel struct {
	Model
}

func (this *BlogSyncDbModel) Search(where BlogSync, limit CommonPage) BlogSyncs {
	db := this.DB.NewSession()

	if limit.PageSize == 0 && limit.PageIndex == 0 {
		return BlogSyncs{
			Count: 0,
			Data:  []BlogSync{},
		}
	}

	if where.GitUrl != "" {
		db = db.And("gitUrl like ?", where.GitUrl)
	}
	if where.State != 0 {
		db = db.And("state = ?", where.State)
	}
	if where.SyncType != 0 {
		db = db.And("syncType = ?", where.SyncType)
	}
	if where.UserId != 0 {
		db = db.And("userId = ?", where.UserId)
	}

	data := []BlogSync{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return BlogSyncs{
		Count: int(count),
		Data:  data,
	}
}

func (this *BlogSyncDbModel) Get(blogSyncId int) BlogSync {
	var blogSyncs []BlogSync
	err := this.DB.Where("blogSyncId = ?", blogSyncId).Find(&blogSyncs)
	if err != nil {
		panic(err)
	}
	if len(blogSyncs) == 0 {
		Throw(1, "不存在该记录"+strconv.Itoa(blogSyncId))
	}
	return blogSyncs[0]
}

func (this *BlogSyncDbModel) Add(blogSync BlogSync) int {
	_, err := this.DB.Insert(&blogSync)
	if err != nil {
		panic(err)
	}
	return blogSync.BlogSyncId
}

func (this *BlogSyncDbModel) Mod(blogSyncId int, blogSync BlogSync) {
	_, err := this.DB.Where("blogSyncId = ?", blogSyncId).Update(&blogSync)
	if err != nil {
		panic(err)
	}
}
