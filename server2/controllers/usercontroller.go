package controllers

import (
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type UserController struct {
	BaseController
	UserLoginAo UserLoginAoModel
	UserAo      UserAoModel
}

func (this *UserController) GetType_Json() interface{} {
	return UserTypeEnum.Names()
}

func (this *UserController) Search_Json() interface{} {
	//检查输入参数
	var where User
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	this.UserLoginAo.CheckMustAdmin()

	//执行业务逻辑
	return this.UserAo.Search(where, limit)
}

func (this *UserController) Get_Json() interface{} {
	//检查输入参数
	var user User
	this.CheckGet(&user)

	//检查权限
	this.UserLoginAo.CheckMustAdmin()

	//执行业务逻辑
	return this.UserAo.Get(user.UserId)
}

func (this *UserController) Add_Json() {
	//检查输入参数
	var user User
	this.CheckPost(&user)

	//检查权限
	this.UserLoginAo.CheckMustAdmin()

	//执行业务逻辑
	this.UserAo.Add(user)
}

func (this *UserController) Del_Json() {
	//检查输入参数
	var user User
	this.CheckPost(&user)

	//检查权限
	this.UserLoginAo.CheckMustAdmin()

	//执行业务逻辑
	this.UserAo.Del(user.UserId)
}

func (this *UserController) ModType_Json() {
	//检查输入参数
	var user User
	this.CheckPost(&user)

	//检查权限
	this.UserLoginAo.CheckMustAdmin()

	//执行业务逻辑
	this.UserAo.ModType(user.UserId, user.Type)
}

func (this *UserController) ModPassword_Json() {
	//检查输入参数
	var user User
	this.CheckPost(&user)

	//检查权限
	this.UserLoginAo.CheckMustAdmin()

	//执行业务逻辑
	this.UserAo.ModPassword(user.UserId, user.Password)
}

func (this *UserController) ModMyPassword_Json() {
	//检查输入参数
	var input struct {
		OldPassword string
		NewPassword string
	}
	this.CheckPost(&input)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.UserAo.ModPasswordByOld(loginUser.UserId, input.OldPassword, input.NewPassword)
}
