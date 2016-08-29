package brush

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type BrushCrawlDbModel struct {
	Model
}

func (this *BrushCrawlDbModel) Search(where BrushCrawl, limit CommonPage) BrushCrawls {
	db := this.DB.NewSession()

	if limit.PageSize == 0 && limit.PageIndex == 0 {
		return BrushCrawls{
			Count: 0,
			Data:  []BrushCrawl{},
		}
	}

	if where.UserId != 0 {
		db = db.Where("userId = ? ", where.UserId)
	}
	if where.State != 0 {
		db = db.Where("state = ? ", where.State)
	}
	if where.BrushTaskId != 0 {
		db = db.Where("brushTaskId = ? ", where.BrushTaskId)
	}
	data := []BrushCrawl{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return BrushCrawls{
		Count: int(count),
		Data:  data,
	}
}

func (this *BrushCrawlDbModel) GetByState(state int) []BrushCrawl {
	var tasks []BrushCrawl
	err := this.DB.Where("state=?", state).Find(&tasks)
	if err != nil {
		panic(err)
	}
	return tasks
}

func (this *BrushCrawlDbModel) Get(brushCrawlId int) BrushCrawl {
	var tasks []BrushCrawl
	err := this.DB.Where("brushCrawlId=?", brushCrawlId).Find(&tasks)
	if err != nil {
		panic(err)
	}
	if len(tasks) == 0 {
		Throw(1, "该"+strconv.Itoa(brushCrawlId)+"任务不存在")
	}
	return tasks[0]
}

func (this *BrushCrawlDbModel) Add(task BrushCrawl) int {
	_, err := this.DB.Insert(&task)
	if err != nil {
		panic(err)
	}
	return task.BrushCrawlId
}

func (this *BrushCrawlDbModel) Mod(brushCrawlId int, task BrushCrawl) {
	_, err := this.DB.Where("brushCrawlId = ?", brushCrawlId).Update(&task)
	if err != nil {
		panic(err)
	}
}

func (this *BrushCrawlDbModel) IncrRetryNum(brushCrawlId int) {
	_, err := this.DB.Where("brushCrawlId = ?", brushCrawlId).Incr("retryNum").Update(&BrushCrawl{})
	if err != nil {
		panic(err)
	}
}
