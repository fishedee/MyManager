package main

import (
	"github.com/fishedee/web"
	_ "mymanager/routers"
)

//go:generate fishgen ^./models/.*(ao|db)\.go$
func main() {
	web.Run()
}
