package brush

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type BrushTaskDbModel struct {
	Model
}

func (this *BrushTaskDbModel) Search(where BrushTask, limit CommonPage) BrushTasks {
	db := this.DB.NewSession()

	if limit.PageSize == 0 && limit.PageIndex == 0 {
		return BrushTasks{
			Count: 0,
			Data:  []BrushTask{},
		}
	}

	if where.State != 0 {
		db = db.Where("state = ? ", where.State)
	}
	if where.Type != 0 {
		db = db.Where("type = ? ", where.Type)
	}
	if where.UserId != 0 {
		db = db.Where("userId = ? ", where.UserId)
	}

	data := []BrushTask{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return BrushTasks{
		Count: int(count),
		Data:  data,
	}
}

func (this *BrushTaskDbModel) Get(brushTaskId int) BrushTask {
	var tasks []BrushTask
	err := this.DB.Where("brushTaskId=?", brushTaskId).Find(&tasks)
	if err != nil {
		panic(err)
	}
	if len(tasks) == 0 {
		Throw(1, "该"+strconv.Itoa(brushTaskId)+"任务不存在")
	}
	return tasks[0]
}

func (this *BrushTaskDbModel) Add(task BrushTask) int {
	_, err := this.DB.Insert(&task)
	if err != nil {
		panic(err)
	}
	return task.BrushTaskId
}

func (this *BrushTaskDbModel) Mod(brushTaskId int, task BrushTask) {
	_, err := this.DB.Where("brushTaskId = ?", brushTaskId).Update(&task)
	if err != nil {
		panic(err)
	}
}

func (this *BrushTaskDbModel) AddSuccessNum(brushTaskId int) {
	_, err := this.DB.Where("brushTaskId = ?", brushTaskId).Incr("successNum").Update(&BrushTask{})
	if err != nil {
		panic(err)
	}
}

func (this *BrushTaskDbModel) AddFailNum(brushTaskId int) {
	_, err := this.DB.Where("brushTaskId = ?", brushTaskId).Incr("failNum").Update(&BrushTask{})
	if err != nil {
		panic(err)
	}
}
