package controllers

import (
	. "mymanager/models/user"
)

type loginController struct {
	baseController
	UserLoginAo UserLoginAoModel
}

func (this *loginController) Islogin_Json() interface{} {
	return this.UserLoginAo.CheckMustLogin()
}

func (this *loginController) Checkin_Json() {
	//检查输入参数
	var user User
	this.CheckPost(&user)

	this.UserLoginAo.Login(user.Name, user.Password)
}

func (this *loginController) Checkout_Json() {
	this.UserLoginAo.Logout()
}
