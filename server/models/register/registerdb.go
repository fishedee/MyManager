package register

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type RegisterDbModel struct {
	Model
}

func (this *RegisterDbModel) Search(where Register, limit CommonPage) Registers {
	db := this.DB.NewSession()

	if where.Name != "" {
		db = db.Where("name like ?", "%"+where.Name+"%")
	}
	if where.UserId != 0 {
		db = db.Where("userId = ? ", where.UserId)
	}

	data := []Register{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return Registers{
		Count: int(count),
		Data:  data,
	}
}

func (this *RegisterDbModel) GetAll() []Register {
	var registers []Register
	err := this.DB.Find(&registers)
	if err != nil {
		panic(err)
	}
	return registers
}

func (this *RegisterDbModel) Get(registerId int) Register {
	var registers []Register
	err := this.DB.Where("registerId = ?", registerId).Find(&registers)
	if err != nil {
		panic(err)
	}
	if len(registers) == 0 {
		Throw(1, "该"+strconv.Itoa(registerId)+"挂号不存在")
	}
	return registers[0]
}

func (this *RegisterDbModel) Del(registerId int) {
	_, err := this.DB.Where("registerId = ?", registerId).Delete(&Register{})
	if err != nil {
		panic(err)
	}
}

func (this *RegisterDbModel) Add(register Register) {
	_, err := this.DB.Insert(register)
	if err != nil {
		panic(err)
	}
}

func (this *RegisterDbModel) Mod(registerId int, register Register) {
	_, err := this.DB.Where("registerId = ?", registerId).Update(&register)
	if err != nil {
		panic(err)
	}
}
