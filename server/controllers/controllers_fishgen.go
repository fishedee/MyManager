package controllers

import (
	. "github.com/fishedee/web"
)

type AccountController interface {
	Search_Json() (_fishgen1 interface{})
	Get_Json() (_fishgen1 interface{})
	Add_Json()
	Del_Json()
	Mod_Json()
	GetType_Json() (_fishgen1 interface{})
	GetWeekTypeStatistic_Json() (_fishgen1 interface{})
	GetWeekDetailTypeStatistic_Json() (_fishgen1 interface{})
	GetWeekCardStatistic_Json() (_fishgen1 interface{})
	GetWeekDetailCardStatistic_Json() (_fishgen1 interface{})
}

type BaseController interface {
	AutoRender(returnValue interface{}, viewname string)
}

type BlogController interface {
	SearchAuto_Json() (_fishgen1 interface{})
	GetAuto_Json() (_fishgen1 interface{})
	AddAuto_Json()
	ModAuto_Json()
	DelAuto_Json()
	SearchTask_Json() (_fishgen1 interface{})
	GetTask_Json() (_fishgen1 interface{})
	AddTask_Json()
	RestartTask_Json()
}

type CardController interface {
	Search_Json() (_fishgen1 interface{})
	Get_Json() (_fishgen1 interface{})
	Add_Json()
	Del_Json()
	Mod_Json()
}

type CategoryController interface {
	Search_Json() (_fishgen1 interface{})
	Get_Json() (_fishgen1 interface{})
	Add_Json()
	Del_Json()
	Mod_Json()
}

type LoginController interface {
	Islogin_Json() (_fishgen1 interface{})
	Checkin_Json()
	Checkout_Json()
}

type UserController interface {
	GetType_Json() (_fishgen1 interface{})
	Search_Json() (_fishgen1 interface{})
	Get_Json() (_fishgen1 interface{})
	Add_Json()
	Del_Json()
	ModType_Json()
	ModPassword_Json()
	ModMyPassword_Json()
}

func init() {
	v0 := AccountController(&accountController{})
	InitController(&v0)
	v1 := BaseController(&baseController{})
	InitController(&v1)
	v2 := BlogController(&blogController{})
	InitController(&v2)
	v3 := CardController(&cardController{})
	InitController(&v3)
	v4 := CategoryController(&categoryController{})
	InitController(&v4)
	v5 := LoginController(&loginController{})
	InitController(&v5)
	v6 := UserController(&userController{})
	InitController(&v6)
}
