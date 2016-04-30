package controllers

import (
	"github.com/astaxie/beego"
	"github.com/fishedee/encoding"
	"github.com/fishedee/language"
	"github.com/fishedee/web"
)

type BaseController struct {
	web.BeegoValidateController
}

func InitRoute(namespace string, target beego.ControllerInterface) {
	web.InitBeegoVaildateControllerRoute(namespace, target)
}

type baseControllerResult struct {
	Code int
	Data interface{}
	Msg  string
}

func (this *BaseController) jsonRender(result baseControllerResult) {
	resultString, err := encoding.EncodeJson(result)
	if err != nil {
		panic(err)
	}
	this.Ctx.WriteString(string(resultString))
}

func (this *BaseController) AutoRender(returnValue interface{}, viewname string) {
	result := baseControllerResult{}
	resultError, ok := returnValue.(language.Exception)
	if ok {
		//带错误码的error
		result.Code = resultError.GetCode()
		result.Msg = resultError.GetMessage()
		result.Data = nil
	} else {
		//正常返回
		result.Code = 0
		result.Data = returnValue
		result.Msg = ""
	}
	this.Ctx.Output.Header("Cache-Control", "private, no-store, no-cache, must-revalidate, max-age=0")
	this.Ctx.Output.Header("Cache-Control", "post-check=0, pre-check=0")
	this.Ctx.Output.Header("Pragma", "no-cache")

	if viewname == "json" {
		this.jsonRender(result)
	} else {
		panic("不合法的viewName " + viewname)
	}
}
