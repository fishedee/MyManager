package routers

import (
	"mymanager/controllers"
)

func init() {
	controllers.InitRoute("user", &controllers.UserController{})
	controllers.InitRoute("login", &controllers.LoginController{})
	controllers.InitRoute("category", &controllers.CategoryController{})
	controllers.InitRoute("card", &controllers.CardController{})
	controllers.InitRoute("account", &controllers.AccountController{})
	controllers.InitRoute("blog", &controllers.BlogController{})
}
