package account

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
)

func (this *AccountAoModel) Search_WithError(userId int, where Account, limit CommonPage) (_fishgen1 Accounts, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(userId, where, limit)
	return
}

func (this *AccountAoModel) Get_WithError(userId int, accountId int) (_fishgen1 Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId, accountId)
	return
}

func (this *AccountAoModel) Del_WithError(userId int, accountId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId, accountId)
	return
}

func (this *AccountAoModel) Add_WithError(userId int, account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(userId, account)
	return
}

func (this *AccountAoModel) Mod_WithError(userId int, accountId int, account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(userId, accountId, account)
	return
}

func (this *AccountDbModel) Search_WithError(where Account, limit CommonPage) (_fishgen1 Accounts, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *AccountDbModel) Get_WithError(accountId int) (_fishgen1 Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(accountId)
	return
}

func (this *AccountDbModel) Del_WithError(accountId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(accountId)
	return
}

func (this *AccountDbModel) Add_WithError(account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(account)
	return
}

func (this *AccountDbModel) Mod_WithError(accountId int, account Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(accountId, account)
	return
}

func (this *AccountDbModel) ResetCategory_WithError(categoryId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ResetCategory(categoryId)
	return
}

func (this *AccountDbModel) ResetCard_WithError(cardId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ResetCard(cardId)
	return
}

func (this *AccountStatisticAoModel) GetWeekTypeStatistic_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypeStatistic(userId)
	return
}

func (this *AccountStatisticAoModel) GetWeekTypeStatisticDetail_WithError(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypeStatisticDetail(userId, year, week, accountType)
	return
}

func (this *AccountStatisticAoModel) GetWeekCardStatistic_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatistic(userId)
	return
}

func (this *AccountStatisticAoModel) GetWeekCardStatisticDetail_WithError(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatisticDetail(userId, year, week, cardId)
	return
}

func (this *AccountStatisticDbModel) GetWeekTypStatisticByUser_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypStatisticByUser(userId)
	return
}

func (this *AccountStatisticDbModel) GetWeekTypeStatisticDetailByUser_WithError(userId int, year int, week int, accountType int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypeStatisticDetailByUser(userId, year, week, accountType)
	return
}

func (this *AccountStatisticDbModel) GetWeekCardStatisticByUser_WithError(userId int) (_fishgen1 []AccountStatistic, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatisticByUser(userId)
	return
}

func (this *AccountStatisticDbModel) GetWeekCardStatisticDetailByUser_WithError(userId int, year int, week int, cardId int) (_fishgen1 []AccountStatisticDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatisticDetailByUser(userId, year, week, cardId)
	return
}
