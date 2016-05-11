package main

import (
	"github.com/fishedee/web"
	_ "mymanager/routers"
)

//go:generate fishgen -force ^.*(ao|db|_testing|controller)\.go$
func main() {
	web.Run()
}
