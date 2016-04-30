package user

import (
	"github.com/fishedee/language"
	"mymanager/models/common"
)

type UserLoginAoModel struct {
	common.BaseModel
	UserAo UserAoModel
}

func (this *UserLoginAoModel) IsLogin() User {
	sess, err := this.Session.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	if err != nil {
		panic("session启动失败")
	}
	defer sess.SessionRelease(this.Ctx.ResponseWriter)

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
		language.Throw(1, "用户未登录")
	}
	return user
}

func (this *UserLoginAoModel) CheckMustAdmin() User {
	user := this.CheckMustLogin()
	if user.Type != UserTypeEnum.ADMIN {
		language.Throw(1, "非管理员没有权限执行此操作")
	}
	return user
}

func (this *UserLoginAoModel) Logout() {
	sess, err := this.Session.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	if err != nil {
		panic("session启动失败")
	}
	defer sess.SessionRelease(this.Ctx.ResponseWriter)

	sess.Set("userId", 0)
}

func (this *UserLoginAoModel) Login(name string, password string) {
	sess, err := this.Session.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	if err != nil {
		panic("session启动失败")
	}
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	users := this.UserAo.GetByName(name)
	if len(users) == 0 {
		language.Throw(1, "不存在此帐号")
	}
	this.UserAo.CheckMustVaildPassword(password, users[0].Password)
	sess.Set("userId", users[0].UserId)
}
