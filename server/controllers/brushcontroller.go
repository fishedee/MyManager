package controllers

import (
	. "mymanager/models/brush"
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type BrushController struct {
	BaseController
	UserLoginAo UserLoginAoModel
	BrushAo     BrushAoModel
}

func (this *BrushController) GetTaskType_Json() interface{} {
	return BrushTaskTypeEnum.Names()
}

func (this *BrushController) GetTaskState_Json() interface{} {
	return BrushTaskStateEnum.Names()
}

func (this *BrushController) GetCrawlState_Json() interface{} {
	return BrushCrawlStateEnum.Names()
}

func (this *BrushController) SearchTask_Json() interface{} {
	//检查输入参数
	var where BrushTask
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	user := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BrushAo.SearchTask(user.UserId, where, limit)
}

func (this *BrushController) GetTask_Json() interface{} {
	//检查输入参数
	var task BrushTask
	this.CheckGet(&task)

	//检查权限
	user := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BrushAo.GetTask(user.UserId, task.BrushTaskId)
}

func (this *BrushController) AddTask_Json() {
	//检查输入参数
	var task BrushTask
	this.CheckPost(&task)

	//检查权限
	user := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	this.BrushAo.AddTask(user.UserId, task)
}

func (this *BrushController) SearchCrawl_Json() interface{} {
	//检查输入参数
	var where BrushCrawl
	this.CheckGet(&where)

	var limit CommonPage
	this.CheckGet(&limit)

	//检查权限
	user := this.UserLoginAo.CheckMustLogin()

	//执行业务逻辑
	return this.BrushAo.SearchCrawl(user.UserId, where, limit)
}
