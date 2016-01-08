package category

type Category struct {
	CategoryId     int
	UserId       int
	Name   string
	Remark	string
	CreateTime string `xorm:"created"`
	ModifyTime string `xorm:"updated"`
}

type Categorys struct {
	Count int
	Data  []Category
}
