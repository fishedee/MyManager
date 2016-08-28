package brush

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type BrushAoModel struct {
	Model
	BrushTaskDb  BrushTaskDbModel
	BrushCrawlDb BrushCrawlDbModel
}

func (this *BrushAoModel) SearchTask(userId int, where BrushTask, limit CommonPage) BrushTasks {
	where.UserId = userId
	return this.BrushTaskDb.Search(where, limit)
}

func (this *BrushAoModel) GetTask(userId int, brushTaskId int) BrushTask {
	data := this.BrushTaskDb.Get(brushTaskId)
	if data.UserId != userId {
		Throw(1, "你没有权限获取该信息")
	}
	return data
}

func (this *BrushAoModel) AddTask(userId int, data BrushTask) {
	if data.TotalNum <= 0 {
		Throw(1, "任务数量必须为正数")
	}
	if data.Url == "" {
		Throw(1, "任务链接为空")
	}

	data.UserId = userId
	data.State = BushTaskStateEnum.STATE_BEGIN
	data.StateMessage = ""
	brushTaskId := this.BrushTaskDb.Add(data)
	this.Log.Debug("addTask %v", brushTaskId)
	this.Queue.Produce(BrushQueueEnum.TASK_ADD, brushTaskId)
}

func (this *BrushAoModel) SearchCrawl(userId int, where BrushCrawl, limit CommonPage) BrushCrawls {
	where.UserId = userId
	return this.BrushCrawlDb.Search(where, limit)
}

func (this *BrushAoModel) refreshTask(taskId int, isSuccess bool) {
	if isSuccess {
		this.BrushTaskDb.AddSuccessNum(taskId)
	} else {
		this.BrushTaskDb.AddFailNum(taskId)
	}
	task := this.BrushTaskDb.Get(taskId)
	if task.SuccessNum+task.FailNum == task.TotalNum {
		this.BrushTaskDb.Mod(taskId, BrushTask{
			State:        BushTaskStateEnum.STATE_SUCCESS,
			StateMessage: "成功",
		})
	}
}

func (this *BrushAoModel) handleAddTask(taskId int) {
	defer CatchCrash(func(e Exception) {
		this.BrushTaskDb.Mod(taskId, BrushTask{
			State:        BushTaskStateEnum.STATE_FAIL,
			StateMessage: "失败：" + e.GetMessage(),
		})
		panic(e)
	})
	this.BrushTaskDb.Mod(taskId, BrushTask{
		State:        BushTaskStateEnum.STATE_PROGRESS,
		StateMessage: "进行中",
	})
	task := this.BrushTaskDb.Get(taskId)
	if task.Type == BushTaskTypeEnum.DIRECT {
		for i := 0; i != task.TotalNum; i++ {
			brushCrawlId := this.BrushCrawlDb.Add(BrushCrawl{
				BrushTaskId:  taskId,
				UserId:       task.UserId,
				Proxy:        "",
				RetryNum:     0,
				State:        BushCrawlStateEnum.STATE_BEGIN,
				StateMessage: "",
			})
			this.Queue.Produce(BrushQueueEnum.TASK_CRAWL, brushCrawlId, task)
		}
	} else {
		Throw(1, "不合法或仍未实现的task type ["+strconv.Itoa(task.Type)+"]")
	}
}

func (this *BrushAoModel) handleCrawlTask(brushCrawlId int, task BrushTask) {
	defer CatchCrash(func(e Exception) {
		this.BrushCrawlDb.Mod(brushCrawlId, BrushCrawl{
			State:        BushCrawlStateEnum.STATE_FAIL,
			StateMessage: "失败：" + e.GetMessage(),
		})
		this.refreshTask(task.BrushTaskId, false)
		panic(e)
	})
	this.BrushCrawlDb.Mod(brushCrawlId, BrushCrawl{
		State:        BushCrawlStateEnum.STATE_PROGRESS,
		StateMessage: "进行中",
	})
	crawl := this.BrushCrawlDb.Get(brushCrawlId)
	if crawl.Proxy == "" {
		err := DefaultAjaxPool.Get(&Ajax{
			Url: task.Url,
		})
		if err != nil {
			panic(err)
		}
	} else {
		Throw(1, "不合法或仍未实现的crawl的代理")
	}
	this.BrushCrawlDb.Mod(brushCrawlId, BrushCrawl{
		State:        BushCrawlStateEnum.STATE_SUCCESS,
		StateMessage: "成功",
	})
	this.refreshTask(task.BrushTaskId, true)
}
