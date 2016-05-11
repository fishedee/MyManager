package controllers

import (
	. "mymanager/models/account"
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type accountController struct {
	baseController
	AccountAo          AccountAoModel
	AccountStatisticAo AccountStatisticAoModel
	UserLoginAo        UserLoginAoModel
}

func (this *accountController) Search_Json() interface{} {
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

func (this *accountController) Get_Json() interface{} {
	//检查输入参数
	var account Account
	this.CheckGet(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountAo.Get(loginUser.UserId, account.AccountId)
}

func (this *accountController) Add_Json() {
	//检查输入参数
	var account Account
	this.CheckPost(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.AccountAo.Add(loginUser.UserId, account)
}

func (this *accountController) Del_Json() {
	//检查输入参数
	var account Account
	this.CheckPost(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.AccountAo.Del(loginUser.UserId, account.AccountId)
}

func (this *accountController) Mod_Json() {
	//检查输入参数
	var account Account
	this.CheckPost(&account)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.AccountAo.Mod(loginUser.UserId, account.AccountId, account)
}

func (this *accountController) GetType_Json() interface{} {
	return AccountTypeEnum.Names()
}

func (this *accountController) GetWeekTypeStatistic_Json() interface{} {
	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountStatisticAo.GetWeekTypeStatistic(
		loginUser.UserId,
	)
}

func (this *accountController) GetWeekDetailTypeStatistic_Json() interface{} {
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

func (this *accountController) GetWeekCardStatistic_Json() interface{} {
	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.AccountStatisticAo.GetWeekCardStatistic(
		loginUser.UserId,
	)
}

func (this *accountController) GetWeekDetailCardStatistic_Json() interface{} {
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
