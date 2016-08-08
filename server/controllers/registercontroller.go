package controllers

import (
	. "mymanager/models/common"
	. "mymanager/models/register"
	. "mymanager/models/user"
)

type RegisterController struct {
	BaseController
	RegisterAo  RegisterAoModel
	UserLoginAo UserLoginAoModel
}

func (this *RegisterController) Search_Json() interface{} {
	//检查输入参数
	var where Register
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.RegisterAo.Search(loginUser.UserId, where, limit)
}

func (this *RegisterController) Get_Json() interface{} {
	//检查输入参数
	var register Register
	this.CheckGet(&register)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.RegisterAo.Get(loginUser.UserId, register.RegisterId)
}

func (this *RegisterController) Add_Json() {
	//检查输入参数
	var register Register
	this.CheckPost(&register)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.RegisterAo.Add(loginUser.UserId, register)
}

func (this *RegisterController) Del_Json() {
	//检查输入参数
	var register Register
	this.CheckPost(&register)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.RegisterAo.Del(loginUser.UserId, register.RegisterId)
}

func (this *RegisterController) Mod_Json() {
	//检查输入参数
	var register Register
	this.CheckPost(&register)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.RegisterAo.Mod(loginUser.UserId, register.RegisterId, register)
}
