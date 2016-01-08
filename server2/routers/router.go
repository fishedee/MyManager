package routers

import (
	. "mymanager/controllers"
)

func init() {
	InitRoute("user", &UserController{})
	InitRoute("login", &LoginController{})
	InitRoute("category", &CategoryController{})
	InitRoute("card", &CardController{})
}
