package user

import (
	"time"
)

type User struct {
	UserId     int
	Name       string
	Password   string
	Type       int
	CreateTime time.Time `xorm:"created"`
	ModifyTime time.Time `xorm:"updated"`
}

type Users struct {
	Count int
	Data  []User
}
