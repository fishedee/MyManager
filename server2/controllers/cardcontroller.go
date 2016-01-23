package controllers

import (
	. "mymanager/models/card"
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type CardController struct {
	BaseController
	CardAo      CardAoModel
	UserLoginAo UserLoginAoModel
}

func (this *CardController) Search_Json() interface{} {
	//检查输入参数
	var where Card
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.CardAo.Search(loginUser.UserId, where, limit)
}

func (this *CardController) Get_Json() interface{} {
	//检查输入参数
	var card Card
	this.CheckGet(&card)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.CardAo.Get(loginUser.UserId, card.CardId)
}

func (this *CardController) Add_Json() {
	//检查输入参数
	var card Card
	this.CheckPost(&card)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.CardAo.Add(loginUser.UserId, card)
}

func (this *CardController) Del_Json() {
	//检查输入参数
	var card Card
	this.CheckPost(&card)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.CardAo.Del(loginUser.UserId, card.CardId)
}

func (this *CardController) Mod_Json() {
	//检查输入参数
	var card Card
	this.CheckPost(&card)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.CardAo.Mod(loginUser.UserId, card.CardId, card)
}
