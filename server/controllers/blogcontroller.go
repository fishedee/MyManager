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

func (this *BlogController) SearchAuto_Json() interface{} {
	//检查输入参数
	var where BlogSyncAuto
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.SearchAuto(loginUser.UserId, where, limit)
}

func (this *BlogController) GetAuto_Json() interface{} {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckGet(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.GetAuto(loginUser.UserId, blogSync.BlogSyncAutoId)
}

func (this *BlogController) AddAuto_Json() {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.AddAuto(loginUser.UserId, blogSync)
}

func (this *BlogController) ModAuto_Json() {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.ModAuto(loginUser.UserId, blogSync.BlogSyncAutoId, blogSync)
}

func (this *BlogController) DelAuto_Json() {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.DelAuto(loginUser.UserId, blogSync.BlogSyncAutoId)
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
