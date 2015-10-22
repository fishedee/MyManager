package router

import (
    "github.com/gin-gonic/gin"
   	"../view"
   	"../service/user"
)

func SetLoginRouter(router *gin.RouterGroup){

	router.GET("/islogin", view.Json( func(c *gin.Context)(interface{},error){
		return user.LoginAo.IsLogin(c);
	}));

	router.POST("/checkin", view.Json( func(c *gin.Context)(interface{},error){
		return nil,user.LoginAo.Login(
			c,
			c.PostForm("name"),
			c.PostForm("password"),
		);
	}));

	router.GET("/checkout", view.Json( func(c *gin.Context)(interface{},error) {
		return nil,user.LoginAo.Logout(c);
	}));
}