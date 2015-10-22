package main

import (
    "github.com/gin-gonic/gin"
    "./router"
    "runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

    server := gin.Default()

    loginRouter := server.Group("/login");
    router.SetLoginRouter(loginRouter);
    userRouter := server.Group("/user");
    router.SetUserRoute(userRouter);

    server.Run(":3001")
}