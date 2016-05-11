package controllers

import (
	. "mymanager/models/blog"
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type blogController struct {
	baseController
	BlogSyncAo  BlogSyncAoModel
	UserLoginAo UserLoginAoModel
}

func (this *blogController) SearchAuto_Json() interface{} {
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

func (this *blogController) GetAuto_Json() interface{} {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckGet(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.GetAuto(loginUser.UserId, blogSync.BlogSyncAutoId)
}

func (this *blogController) AddAuto_Json() {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.AddAuto(loginUser.UserId, blogSync)
}

func (this *blogController) ModAuto_Json() {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.ModAuto(loginUser.UserId, blogSync.BlogSyncAutoId, blogSync)
}

func (this *blogController) DelAuto_Json() {
	//检查输入参数
	var blogSync BlogSyncAuto
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.DelAuto(loginUser.UserId, blogSync.BlogSyncAutoId)
}

func (this *blogController) SearchTask_Json() interface{} {
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

func (this *blogController) GetTask_Json() interface{} {
	//检查输入参数
	var blogSync BlogSync
	this.CheckGet(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BlogSyncAo.GetTask(loginUser.UserId, blogSync.BlogSyncId)
}

func (this *blogController) AddTask_Json() {
	//检查输入参数
	var blogSync BlogSync
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.AddTask(loginUser.UserId, blogSync.AccessToken, blogSync.GitUrl, blogSync.SyncType)
}

func (this *blogController) RestartTask_Json() {
	//检查输入参数
	var blogSync BlogSync
	this.CheckPost(&blogSync)

	//检查权限
	loginUser := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BlogSyncAo.RestartTask(loginUser.UserId, blogSync.BlogSyncId)
}
