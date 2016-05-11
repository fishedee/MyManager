package account

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type AccountAoModel interface {
	Search(userId int, where Account, limit CommonPage) (_fishgen1 Accounts)
	Search_WithError(userId int, where Account, limit CommonPage) (_fishgen1 Accounts, _fishgenErr Exception)
	Get(userId int, accountId int) (_fishgen1 Account)
	Get_WithError(userId int, accountId int) (_fishgen1 Account, _fishgenErr Exception)
	Del(userId int, accountId int)
	Del_WithError(userId int, accountId int) (_fishgenErr Exception)
	Add(userId int, account Account)
	Add_WithError(userId int, account Account) (_fishgenErr Exception)
	Mod(userId int, accountId int, account Account)
	Mod_WithError(userId int, accountId int, account Account) (_fishgenErr Exception)
}

type AccountAoTest interface {
	InitEmpty()
	TestBasic()
	TestStatistic()
}

type AccountDbModel interface {
	Search(where Account, limit CommonPage) (_fishgen1 Accounts)
	Search_WithError(where Account, limit CommonPage) (_fishgen1 Accounts, _fishgenErr Exception)
	Get(accountId int) (_fishgen1 Account)
	Get_WithError(accountId int) (_fishgen1 Account, _fishgenErr Exception)
	Del(accountId int)
	Del_WithError(accountId int) (_fishgenErr Exception)
	Add(account Account)
	Add_WithError(account Account) (_fishgenErr Exception)
	Mod(accountId int, account Account)
	Mod_WithError(accountId int, account Account) (_fishgenErr Exception)
	ResetCategory(categoryId int)
	ResetCategory_WithError(categoryId int) (_fishgenErr Exception)
	ResetCard(cardId int)
	ResetCard_WithError(cardId int) (_fishgenErr Exception)
}

type AccountStatisticAoModel interface {
	GetWeekTypeStatistic(userId int) (_fishgen1 []AccountStatistic)
	GetWeekTypeStatistic_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception)
	GetWeekTypeStatisticDetail(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail)
	GetWeekTypeStatisticDetail_WithError(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception)
	GetWeekCardStatistic(userId int) (_fishgen1 []AccountStatistic)
	GetWeekCardStatistic_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception)
	GetWeekCardStatisticDetail(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail)
	GetWeekCardStatisticDetail_WithError(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception)
}

type AccountStatisticDbModel interface {
	GetWeekTypStatisticByUser(userId int) (_fishgen1 []AccountStatistic)
	GetWeekTypStatisticByUser_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception)
	GetWeekTypeStatisticDetailByUser(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail)
	GetWeekTypeStatisticDetailByUser_WithError(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception)
	GetWeekCardStatisticByUser(userId int) (_fishgen1 []AccountStatistic)
	GetWeekCardStatisticByUser_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception)
	GetWeekCardStatisticDetailByUser(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail)
	GetWeekCardStatisticDetailByUser_WithError(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception)
}

func (this *accountAoModel) Search_WithError(userId int, where Account, limit CommonPage) (_fishgen1 Accounts, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(userId, where, limit)
	return
}

func (this *accountAoModel) Get_WithError(userId int, accountId int) (_fishgen1 Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId, accountId)
	return
}

func (this *accountAoModel) Del_WithError(userId int, accountId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId, accountId)
	return
}

func (this *accountAoModel) Add_WithError(userId int, account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(userId, account)
	return
}

func (this *accountAoModel) Mod_WithError(userId int, accountId int, account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(userId, accountId, account)
	return
}

func (this *accountDbModel) Search_WithError(where Account, limit CommonPage) (_fishgen1 Accounts, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *accountDbModel) Get_WithError(accountId int) (_fishgen1 Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(accountId)
	return
}

func (this *accountDbModel) Del_WithError(accountId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(accountId)
	return
}

func (this *accountDbModel) Add_WithError(account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(account)
	return
}

func (this *accountDbModel) Mod_WithError(accountId int, account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(accountId, account)
	return
}

func (this *accountDbModel) ResetCategory_WithError(categoryId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ResetCategory(categoryId)
	return
}

func (this *accountDbModel) ResetCard_WithError(cardId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ResetCard(cardId)
	return
}

func (this *accountStatisticAoModel) GetWeekTypeStatistic_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypeStatistic(userId)
	return
}

func (this *accountStatisticAoModel) GetWeekTypeStatisticDetail_WithError(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypeStatisticDetail(userId, year, week, accountType)
	return
}

func (this *accountStatisticAoModel) GetWeekCardStatistic_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatistic(userId)
	return
}

func (this *accountStatisticAoModel) GetWeekCardStatisticDetail_WithError(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatisticDetail(userId, year, week, cardId)
	return
}

func (this *accountStatisticDbModel) GetWeekTypStatisticByUser_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypStatisticByUser(userId)
	return
}

func (this *accountStatisticDbModel) GetWeekTypeStatisticDetailByUser_WithError(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypeStatisticDetailByUser(userId, year, week, accountType)
	return
}

func (this *accountStatisticDbModel) GetWeekCardStatisticByUser_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatisticByUser(userId)
	return
}

func (this *accountStatisticDbModel) GetWeekCardStatisticDetailByUser_WithError(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatisticDetailByUser(userId, year, week, cardId)
	return
}
func init() {
	v0 := AccountAoModel(&accountAoModel{})
	InitModel(&v0)
	v1 := AccountAoTest(&accountAoTest{})
	InitTest(&v1)
	v2 := AccountDbModel(&accountDbModel{})
	InitModel(&v2)
	v3 := AccountStatisticAoModel(&accountStatisticAoModel{})
	InitModel(&v3)
	v4 := AccountStatisticDbModel(&accountStatisticDbModel{})
	InitModel(&v4)
}
