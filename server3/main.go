package main

import (
    "github.com/gin-gonic/gin"
    "./router"
)

func main() {
    server := gin.Default()

    loginRouter := server.Group("/login");
    router.SetLoginRouter(loginRouter);
    userRouter := server.Group("/user");
    router.SetUserRoute(userRouter);

    server.Run(":30001")
}