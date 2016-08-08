package register

import (
	"time"
)

type Register struct {
	RegisterId     int `xorm:"autoincr"`
	UserId         int
	Name           string
	BeginTime      time.Time
	EndTime        time.Time
	Mail           string
	NeedDealType   int
	HaveDealType   int
	HaveDealResult string
	CreateTime     time.Time `xorm:"created"`
	ModifyTime     time.Time `xorm:"updated"`
}

type Registers struct {
	Count int
	Data  []Register
}

type RegisterResult struct {
	DeptCode   string
	DeptName   string
	DoctorCode string
	DoctorName string
	LeftCount  int
}
