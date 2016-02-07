package account

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
	"strconv"
)

type AccountDbModel struct {
	BaseModel
}

func (this *AccountDbModel) Search(where Account, limit CommonPage) Accounts {
	db := this.DB.NewSession()

	if limit.PageSize == 0 && limit.PageIndex == 0 {
		return Accounts{
			Count: 0,
			Data:  []Account{},
		}
	}

	if where.Name != "" {
		db = db.And("name like ?", "%"+where.Name+"%")
	}
	if where.Remark != "" {
		db = db.And("remark like ?", "%"+where.Remark+"%")
	}
	if where.CategoryId != 0 {
		db = db.And("categoryId = ?", where.CategoryId)
	}
	if where.CardId != 0 {
		db = db.And("cardId = ?", where.CardId)
	}
	if where.Type != 0 {
		db = db.And("type = ?", where.Type)
	}
	if where.UserId != 0 {
		db = db.And("userId = ?", where.UserId)
	}

	data := []Account{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return Accounts{
		Count: int(count),
		Data:  data,
	}
}

func (this *AccountDbModel) Get(accountId int) Account {
	var accounts []Account
	err := this.DB.Where("accountId = ?", accountId).Find(&accounts)
	if err != nil {
		panic(err)
	}
	if len(accounts) == 0 {
		Throw(1, "不存在该记录"+strconv.Itoa(accountId))
	}
	return accounts[0]
}

func (this *AccountDbModel) Del(accountId int) {
	_, err := this.DB.Where("accountId = ?", accountId).Delete(&Account{})
	if err != nil {
		panic(err)
	}
}

func (this *AccountDbModel) Add(account Account) {
	_, err := this.DB.Insert(account)
	if err != nil {
		panic(err)
	}
}

func (this *AccountDbModel) Mod(accountId int, account Account) {
	_, err := this.DB.Where("accountId = ?", accountId).Update(&account)
	if err != nil {
		panic(err)
	}
}

func (this *AccountDbModel) ResetCategory(categoryId int) {
	_, err := this.DB.Where("categoryId = ?", categoryId).Cols("categoryId").Update(&Account{
		CategoryId: 0,
	})
	if err != nil {
		panic(err)
	}
}

func (this *AccountDbModel) ResetCard(cardId int) {
	_, err := this.DB.Where("cardId = ?", cardId).Cols("cardId").Update(&Account{
		CardId: 0,
	})
	if err != nil {
		panic(err)
	}
}
