package router

import (
    "github.com/gin-gonic/gin"
    "errors"
   	"../view"
   	"../service/user"
   	"../config"
)

func checkAdmin(c *gin.Context)(error){
	userinfo,error := user.LoginAo.IsLogin(c);
	if error != nil{
		return error;
	}

	if userinfo.Type != user.Type.ADMIN{
		return errors.New("你没有权限执行此操作");
	}

	return nil;
}
func SetUserRoute(router *gin.RouterGroup){

	router.GET("/search", view.Json( func(c *gin.Context)(interface{},error){
		/*
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}
		*/

		userInfo := &user.User{
			Name:c.Query("Name"),
			Type:config.Atoi(c.Query("Type")),
		};
		pageIndex := config.Atoi(c.Query("PageIndex"));
		pageSize := config.Atoi(c.Query("PageSize"));

		return user.UserAo.Search(userInfo,pageIndex,pageSize);
	}));

	router.POST("/get", view.Json( func(c *gin.Context)(interface{},error){
		error := checkAdmin(c);
		if error != nil{
			return nil,error;
		}

		return user.UserAo.Get(
			config.Atoi(c.Query("UserId")),
		);
	}));

	router.GET("/add", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		userInfo := &user.User{
			Name:c.PostForm("Name"),
			Type:config.Atoi(c.PostForm("Type")),
			Password:c.PostForm("Password"),
		}
		return nil,user.UserAo.Add(userInfo);
	}));

	router.GET("/del", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		return nil,user.UserAo.Del(
			config.Atoi(c.PostForm("UserId")),
		);
	}));

	router.GET("/modType", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		return nil,user.UserAo.ModType(
			config.Atoi(c.PostForm("UserId")),
			config.Atoi(c.PostForm("Type")),
		);
	}));

	router.GET("/modPassword", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		return nil,user.UserAo.ModPassword(
			config.Atoi(c.PostForm("UserId")),
			c.PostForm("Password"),
		);
	}));

	router.GET("/modMyPassword", view.Json( func(c *gin.Context)(interface{},error) {
		userinfo,error := user.LoginAo.IsLogin(c);
		if error != nil{
			return nil,error;
		}

		return nil,user.UserAo.ModPasswordByOld(
			userinfo.UserId,
			c.PostForm("OldPassword"),
			c.PostForm("NewPassword"),
		);
	}));
}
