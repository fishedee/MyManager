package brush

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type BrushAoModel struct {
	Model
	BrushProxyAo BrushProxyAoModel
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
	if data.RetryNum < 0 {
		Throw(1, "任务重试次数需要大于或等于0")
	}

	data.UserId = userId
	data.State = BrushTaskStateEnum.STATE_BEGIN
	data.StateMessage = ""
	brushTaskId := this.BrushTaskDb.Add(data)
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
			State:        BrushTaskStateEnum.STATE_SUCCESS,
			StateMessage: "成功",
		})
	}
}

func (this *BrushAoModel) handleAddTask(taskId int) {
	defer CatchCrash(func(e Exception) {
		this.BrushTaskDb.Mod(taskId, BrushTask{
			State:        BrushTaskStateEnum.STATE_FAIL,
			StateMessage: "失败：" + e.GetMessage(),
		})
		panic(e)
	})
	this.BrushTaskDb.Mod(taskId, BrushTask{
		State:        BrushTaskStateEnum.STATE_PROGRESS,
		StateMessage: "进行中",
	})
	task := this.BrushTaskDb.Get(taskId)
	addIds := []int{}
	for i := 0; i != task.TotalNum; i++ {
		addId := this.BrushCrawlDb.Add(BrushCrawl{
			BrushTaskId:  taskId,
			UserId:       task.UserId,
			Proxy:        "",
			RetryNum:     0,
			State:        BrushCrawlStateEnum.STATE_BEGIN,
			StateMessage: "",
		})
		addIds = append(addIds, addId)
	}
	for _, singleCrawl := range addIds {
		this.Queue.Produce(BrushQueueEnum.TASK_CRAWL, singleCrawl, task)
	}
}

func (this *BrushAoModel) handleCrawlTask(brushCrawlId int, task BrushTask) {
	crawl := this.BrushCrawlDb.Get(brushCrawlId)
	defer CatchCrash(func(e Exception) {
		if crawl.RetryNum >= task.RetryNum {
			this.BrushCrawlDb.Mod(brushCrawlId, BrushCrawl{
				State:        BrushCrawlStateEnum.STATE_FAIL,
				StateMessage: "失败：" + e.GetMessage(),
			})
			this.refreshTask(task.BrushTaskId, false)
		} else {
			this.BrushCrawlDb.Mod(brushCrawlId, BrushCrawl{
				State:        BrushCrawlStateEnum.STATE_RETRY,
				StateMessage: "失败重试中：" + e.GetMessage(),
			})
			this.BrushCrawlDb.IncrRetryNum(brushCrawlId)
		}
		panic(e)
	})
	this.BrushCrawlDb.Mod(brushCrawlId, BrushCrawl{
		State:        BrushCrawlStateEnum.STATE_PROGRESS,
		StateMessage: "进行中",
	})
	var ajaxPool *AjaxPool
	var proxy string
	if task.Type == BrushTaskTypeEnum.DIRECT {
		proxy = ""
		ajaxPool = DefaultAjaxPool
	} else if task.Type == BrushTaskTypeEnum.PROXY_XICI {
		proxy = this.BrushProxyAo.GetXiciProxy()
		ajaxPool = NewAjaxPool(&AjaxPoolOption{
			Proxy: proxy,
		})
	} else if task.Type == BrushTaskTypeEnum.PROXY_MIMVP {
		proxy = this.BrushProxyAo.GetMimvpProxy()
		ajaxPool = NewAjaxPool(&AjaxPoolOption{
			Proxy: proxy,
		})
	} else {
		panic("不合法的task type")
	}
	err := ajaxPool.Get(&Ajax{
		Url: task.Url,
	})
	if err != nil {
		panic(err)
	}
	this.BrushCrawlDb.Mod(brushCrawlId, BrushCrawl{
		Proxy:        proxy,
		State:        BrushCrawlStateEnum.STATE_SUCCESS,
		StateMessage: "成功",
	})
	this.refreshTask(task.BrushTaskId, true)
}

func (this *BrushAoModel) handleCrawlRetry() {
	crawls := this.BrushCrawlDb.GetByState(BrushCrawlStateEnum.STATE_RETRY)
	if len(crawls) == 0 {
		return
	}
	crawlIds := QueryColumn(crawls, "BrushCrawlId").([]int)
	this.BrushCrawlDb.ModByIds(crawlIds, BrushCrawl{
		State:        BrushCrawlStateEnum.STATE_BEGIN,
		StateMessage: "",
	})

	brushTaskIds := ArrayUnique(QueryColumn(crawls, "BrushTaskId")).([]int)
	tasks := this.BrushTaskDb.GetByIds(brushTaskIds)
	tasksMap := ArrayColumnMap(tasks, "BrushTaskId").(map[int]BrushTask)

	for _, singleCrawl := range crawls {
		brushCrawlId := singleCrawl.BrushCrawlId
		brushTaskId := singleCrawl.BrushTaskId
		task := tasksMap[brushTaskId]
		this.Queue.Produce(BrushQueueEnum.TASK_CRAWL, brushCrawlId, task)
	}
}
