package brush

import (
	"time"
)

type BrushTask struct {
	BrushTaskId  int `xorm:"autoincr"`
	UserId       int
	Url          string
	Type         int
	RetryNum     int
	TotalNum     int
	SuccessNum   int
	FailNum      int
	State        int
	StateMessage string
	CreateTime   time.Time `xorm:"created"`
	ModifyTime   time.Time `xorm:"updated"`
}

type BrushTasks struct {
	Count int
	Data  []BrushTask
}

type BrushCrawl struct {
	BrushCrawlId int `xorm:"autoincr"`
	UserId       int
	BrushTaskId  int
	Proxy        string
	RetryNum     int
	State        int
	StateMessage string
	CreateTime   time.Time `xorm:"created"`
	ModifyTime   time.Time `xorm:"updated"`
}

type BrushCrawls struct {
	Count int
	Data  []BrushCrawl
}
