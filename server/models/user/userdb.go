package user

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type UserDbModel struct {
	Model
}

func (this *UserDbModel) Search(where User, limit CommonPage) Users {
	db := this.DB.NewSession()

	if limit.PageSize == 0 && limit.PageIndex == 0 {
		return Users{
			Count: 0,
			Data:  []User{},
		}
	}

	if where.Name != "" {
		db = db.Where("name like ?", "%"+where.Name+"%")
	}
	if where.Type != 0 {
		db = db.Where("type = ? ", where.Type)
	}

	data := []User{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return Users{
		Count: int(count),
		Data:  data,
	}
}

func (this *UserDbModel) Get(id int) User {
	var users []User
	err := this.DB.Where("userId=?", id).Find(&users)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		Throw(1, "不存在该用户"+strconv.Itoa(id))
	}
	return users[0]
}

func (this *UserDbModel) GetByName(name string) []User {
	var users []User
	err := this.DB.Where("name = ?", name).Find(&users)
	if err != nil {
		panic(err)
	}
	return users
}

func (this *UserDbModel) Del(id int) {
	_, err := this.DB.Where("userId = ?", id).Delete(&User{})
	if err != nil {
		panic(err)
	}
}

func (this *UserDbModel) Add(user User) {
	_, err := this.DB.Insert(user)
	if err != nil {
		panic(err)
	}
}

func (this *UserDbModel) Mod(id int, user User) {
	_, err := this.DB.Where("userId = ?", id).Update(&user)
	if err != nil {
		panic(err)
	}
}
