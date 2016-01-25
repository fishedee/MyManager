package category

import (
	"time"
)

type Category struct {
	CategoryId int `xorm:"autoincr"`
	UserId     int
	Name       string
	Remark     string
	CreateTime time.Time `xorm:"created"`
	ModifyTime time.Time `xorm:"updated"`
}

type Categorys struct {
	Count int
	Data  []Category
}
