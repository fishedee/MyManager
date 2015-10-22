package user;

import (
	"time"
)

type User struct{
	UserId int `sql:"AUTO_INCREMENT"`
	Name string `sql:"size:32"`
	Password string `sql:"size:48"`
	Type int
	CreateTime time.Time
	ModifyTime time.Time 
}