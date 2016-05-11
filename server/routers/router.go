package routers

import (
	. "github.com/fishedee/web"
	. "mymanager/controllers"
)

func init() {
	InitRoute("user", (*UserController)(nil))
	InitRoute("login", (*LoginController)(nil))
	InitRoute("category", (*CategoryController)(nil))
	InitRoute("card", (*CardController)(nil))
	InitRoute("account", (*AccountController)(nil))
	InitRoute("blog", (*BlogController)(nil))
}
