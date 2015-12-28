package user

type User struct {
	UserId     int
	Name       string
	Password   string
	Type       int
	CreateTime string `xorm:"created"`
	ModifyTime string `xorm:"updated"`
}

type Users struct {
	Count int
	Data  []User
}
