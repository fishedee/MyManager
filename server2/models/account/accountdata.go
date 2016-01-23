package account

import (
	"time"
)

type Account struct {
	AccountId  int
	UserId     int
	Name       string
	Money      int
	Remark     string
	CategoryId int
	CardId     int
	Type       int
	CreateTime time.Time
	ModifyTime time.Time
}

type Accounts struct {
	Count int
	Data  []Account
}

type AccountStatistic struct {
	Name     string
	Year     int
	Week     int
	Type     int
	TypeName string
	CardId   int
	CardName string
	Money    int
}

func (this *AccountStatistic) TableName() string {
	return "t_account"
}

type AccountStatisticDetail struct {
	CategoryId   int
	CategoryName string
	Type         int
	TypeName     string
	Money        int
	Precent      string
}

func (this *AccountStatisticDetail) TableName() string {
	return "t_account"
}
