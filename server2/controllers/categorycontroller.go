package controllers

import ()

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Test_Json() interface{} {
	//检查输入参数
	return "Hello World"
}
