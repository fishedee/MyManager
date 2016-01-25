package card

import (
	"time"
)

type Card struct {
	CardId     int `xorm:"autoincr"`
	UserId     int
	Name       string
	Bank       string
	Card       string
	Money      int
	Remark     string
	CreateTime time.Time `xorm:"created"`
	ModifyTime time.Time `xorm:"updated"`
}

type Cards struct {
	Count int
	Data  []Card
}
