package controllers

import (
	. "mymanager/models/blog"
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type BlogController struct {
	BaseController
	BlogSyncAo  BlogSyncAoModel
	UserLoginAo UserLoginAoModel
}

func (this *BlogController) SearchTask_Json() interface{} {
	//检查输入参数
	var where BlogSync
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.SearchTask(loginUser.UserId, where, limit)
}

func (this *BlogController) GetTask_Json() interface{} {
	//检查输入参数
	var blogSync BlogSync
	this.CheckGet(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.GetTask(loginUser.UserId, blogSync.BlogSyncId)
}

func (this *BlogController) AddTask_Json() {
	//检查输入参数
	var blogSync BlogSync
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.AddTask(loginUser.UserId, blogSync.AccessToken, blogSync.GitUrl, blogSync.SyncType)
}

func (this *BlogController) RestartTask_Json() {
	//检查输入参数
	var blogSync BlogSync
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.RestartTask(loginUser.UserId, blogSync.BlogSyncId)
}

func (this *BlogController) GetAuthUrl_Json() interface{} {
	//检查输入参数
	var data struct {
		RedirectUrl string
	}
	this.CheckGet(&data)

	//检查权限
	this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.GetAuthUrl(data.RedirectUrl)
}

func (this *BlogController) GetAccessToken_Json() interface{} {
	//检查输入参数
	var data struct {
		RedirectUrl string
		Code        string
	}
	this.CheckGet(&data)

	//检查权限
	this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.GetAccessToken(data.RedirectUrl, data.Code)
}
