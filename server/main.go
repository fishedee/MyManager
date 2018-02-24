package main

import (
	"github.com/fishedee/app/router"
	"github.com/fishedee/web"
	_ "mymanager/routers"
)

//go:generate fishgen ^./models/.*(ao|db)\.go$
func main() {
	factory := router.NewRouterFactory()
	factory.Static("/", "../static/build")
	web.RunAppRouter(factory)
}
