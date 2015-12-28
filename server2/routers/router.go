package routers

import (
	. "mymanager/controllers"
)

func init() {
	InitRoute("user", &UserController{})
	InitRoute("category", &CategoryController{})
}
