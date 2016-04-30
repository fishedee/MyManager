package controllers

import (
	"mymanager/models/user"
)

type LoginController struct {
	BaseController
	UserLoginAo user.UserLoginAoModel
}

func (this *LoginController) Islogin_Json() interface{} {
	return this.UserLoginAo.CheckMustLogin()
}

func (this *LoginController) Checkin_Json() {
	//检查输入参数
	var user user.User
	this.CheckPost(&user)

	this.UserLoginAo.Login(user.Name, user.Password)
}

func (this *LoginController) Checkout_Json() {
	this.UserLoginAo.Logout()
}
