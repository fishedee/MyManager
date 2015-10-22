package user;

import (
	"time"
)

type User struct{
	UserId int `json:"userId" gorm:"column:userId"`
	Name string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
	Type int `json:"type" gorm:"column:type"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	ModifyTime time.Time `json:"modifyTime" gorm:"column:modifyTime"`
}

func (this *User) TableName() string {
    return "t_user"
}