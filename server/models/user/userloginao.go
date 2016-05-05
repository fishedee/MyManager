package user

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
)

type UserLoginAoModel struct {
	Model
	UserAo UserAoModel
}

func (this *UserLoginAoModel) IsLogin() User {
	sess, err := this.Session.SessionStart()
	if err != nil {
		panic("session启动失败")
	}
	defer sess.SessionRelease()

	userId := sess.Get("userId")
	userIdInt, ok := userId.(int)
	if ok && userIdInt >= 1000 {
		return this.UserAo.Get(userIdInt)
	} else {
		return User{}
	}
}

func (this *UserLoginAoModel) CheckMustLogin() User {
	user := this.IsLogin()
	if user.UserId == 0 {
		Throw(1, "用户未登录")
	}
	return user
}

func (this *UserLoginAoModel) CheckMustAdmin() User {
	user := this.CheckMustLogin()
	if user.Type != UserTypeEnum.ADMIN {
		Throw(1, "非管理员没有权限执行此操作")
	}
	return user
}

func (this *UserLoginAoModel) Logout() {
	sess, err := this.Session.SessionStart()
	if err != nil {
		panic("session启动失败")
	}
	defer sess.SessionRelease()

	sess.Set("userId", 0)
}

func (this *UserLoginAoModel) Login(name string, password string) {
	sess, err := this.Session.SessionStart()
	if err != nil {
		panic("session启动失败")
	}
	defer sess.SessionRelease()
	users := this.UserAo.GetByName(name)
	if len(users) == 0 {
		Throw(1, "不存在此帐号")
	}
	this.UserAo.CheckMustVaildPassword(password, users[0].Password)
	sess.Set("userId", users[0].UserId)
}
