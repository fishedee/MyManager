package controllers

import (
	. "mymanager/models/account"
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type AccountController struct {
	BaseController
	AccountAo          AccountAoModel
	AccountStatisticAo AccountStatisticAoModel
	UserLoginAo        UserLoginAoModel
}

func (this *AccountController) Search_Json() interface{} {
	//检查输入参数
	var where Account
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountAo.Search(loginUser.UserId, where, limit)
}

func (this *AccountController) Get_Json() interface{} {
	//检查输入参数
	var account Account
	this.CheckGet(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountAo.Get(loginUser.UserId, account.AccountId)
}

func (this *AccountController) Add_Json() {
	//检查输入参数
	var account Account
	this.CheckPost(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.AccountAo.Add(loginUser.UserId, account)
}

func (this *AccountController) Del_Json() {
	//检查输入参数
	var account Account
	this.CheckPost(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.AccountAo.Del(loginUser.UserId, account.AccountId)
}

func (this *AccountController) Mod_Json() {
	//检查输入参数
	var account Account
	this.CheckPost(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.AccountAo.Mod(loginUser.UserId, account.AccountId, account)
}

func (this *AccountController) GetType_Json() interface{} {
	return AccountTypeEnum.Names()
}

func (this *AccountController) GetWeekTypeStatistic_Json() interface{} {
	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountStatisticAo.GetWeekTypeStatistic(
		loginUser.UserId,
	)
}

func (this *AccountController) GetWeekDetailTypeStatistic_Json() interface{} {
	//检查输入参数
	var account Account
	this.CheckGet(&account)
	var accountStatistic AccountStatistic
	this.CheckGet(&accountStatistic)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountStatisticAo.GetWeekTypeStatisticDetail(
		loginUser.UserId,
		accountStatistic.Year,
		accountStatistic.Week,
		account.Type,
	)
}

func (this *AccountController) GetWeekCardStatistic_Json() interface{} {
	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountStatisticAo.GetWeekCardStatistic(
		loginUser.UserId,
	)
}

func (this *AccountController) GetWeekDetailCardStatistic_Json() interface{} {
	//检查输入参数
	var account Account
	this.CheckGet(&account)
	var accountStatistic AccountStatistic
	this.CheckGet(&accountStatistic)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountStatisticAo.GetWeekCardStatisticDetail(
		loginUser.UserId,
		accountStatistic.Year,
		accountStatistic.Week,
		account.CardId,
	)
}
