package brush

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
)

var BushTaskTypeEnum struct {
	EnumStruct
	DIRECT int `enum:"1,直接连接"`
	PROXY  int `enum:"2,代理连接"`
}

var BushTaskStateEnum struct {
	EnumStruct
	STATE_BEGIN    int `enum:"1,未开始"`
	STATE_PROGRESS int `enum:"2,进行中"`
	STATE_FAIL     int `enum:"3,失败"`
	STATE_SUCCESS  int `enum:"4,成功"`
}

var BushCrawlStateEnum struct {
	EnumStruct
	STATE_BEGIN    int `enum:"1,未开始"`
	STATE_PROGRESS int `enum:"2,进行中"`
	STATE_RETRY    int `enum:"3,失败重试中"`
	STATE_FAIL     int `enum:"4,失败"`
	STATE_SUCCESS  int `enum:"5,成功"`
}

var BrushQueueEnum struct {
	EnumStructString
	TASK_ADD   string `enum:"/brush/_add,添加任务"`
	TASK_CRAWL string `enum:"/brush/_crawl,抓取任务"`
}

func init() {
	InitEnumStruct(&BushTaskTypeEnum)
	InitEnumStruct(&BushTaskStateEnum)
	InitEnumStruct(&BushCrawlStateEnum)
	InitEnumStructString(&BrushQueueEnum)
	InitDaemon(func(this *BrushAoModel) {
		this.Queue.Consume(BrushQueueEnum.TASK_ADD, (*BrushAoModel).handleAddTask)
		this.Queue.Consume(BrushQueueEnum.TASK_CRAWL, (*BrushAoModel).handleCrawlTask)
	})
}