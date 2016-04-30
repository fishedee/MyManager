package account

import (
	"fmt"
	"github.com/fishedee/language"
	"github.com/fishedee/web"
	"mymanager/models/card"
	"mymanager/models/category"
	"mymanager/models/common"
)

type AccountAoModel struct {
	common.BaseModel
	AccountDb  AccountDbModel
	CategoryAo category.CategoryAoModel
	CardAo     card.CardAoModel
}

func (this *AccountAoModel) checkAccountData(userId int, account Account) {
	//校验分类ID
	this.CategoryAo.Get(userId, account.CategoryId)

	//校验银卡ID
	this.CardAo.Get(userId, account.CardId)

	//校验类型
	for language.ArrayIn(AccountTypeEnum.Keys(), account.Type) == -1 {
		language.Throw(1, fmt.Sprintf("类型ID不合法 %v", account.Type))
	}

	//校验金额
	if account.Money < 0 {
		language.Throw(1, "金额必2须大于或等于0")
	}
}

func (this *AccountAoModel) Search(userId int, where Account, limit common.CommonPage) Accounts {
	where.UserId = userId
	return this.AccountDb.Search(where, limit)
}

func (this *AccountAoModel) Get(userId int, accountId int) Account {
	account := this.AccountDb.Get(accountId)

	if account.UserId != userId {
		language.Throw(1, "权限不足")
	}
	return account
}

func (this *AccountAoModel) Del(userId int, accountId int) {
	this.Get(userId, accountId)

	this.AccountDb.Del(accountId)
}

func (this *AccountAoModel) Add(userId int, account Account) {
	this.checkAccountData(userId, account)

	account.UserId = userId
	this.AccountDb.Add(account)
}

func (this *AccountAoModel) Mod(userId int, accountId int, account Account) {
	this.checkAccountData(userId, account)

	this.Get(userId, accountId)

	this.AccountDb.Mod(accountId, account)
}

func (this *AccountAoModel) whenCategoryDelete(categoryId int) {
	this.AccountDb.ResetCategory(categoryId)
}

func (this *AccountAoModel) whenCardDelete(cardId int) {
	this.AccountDb.ResetCard(cardId)
}

func init() {
	web.InitDaemon(func(this *AccountAoModel) {
		this.Queue.Subscribe("category_del", this.whenCategoryDelete)
		this.Queue.Subscribe("card_del", this.whenCardDelete)
	})
}
