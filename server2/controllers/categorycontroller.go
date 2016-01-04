package controllers

import (
	. "mymanager/models/category"
)

type CategoryController struct {
	BaseController
	CategoryAo CategoryAoModel
}

func (this *CategoryController) Test_Json() interface{} {
	this.CategoryAo.Produce()
	return ""
}
