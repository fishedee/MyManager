package card

type Card struct {
	CardId     int
	UserId       int
	Name   string
	Bank   string
	Card string
	Money int
	Remark	string
	CreateTime string `xorm:"created"`
	ModifyTime string `xorm:"updated"`
}

type Cards struct {
	Count int
	Data  []Card
}
