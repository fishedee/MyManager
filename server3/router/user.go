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
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		userInfo := &user.User{
			Name:c.Query("name"),
			Type:config.Atoi(c.Query("type")),
		};
		pageIndex := config.Atoi(c.Query("pageIndex"));
		pageSize := config.Atoi(c.Query("pageSize"));

		return user.UserAo.Search(userInfo,pageIndex,pageSize);
	}));

	router.GET("/get", view.Json( func(c *gin.Context)(interface{},error){
		error := checkAdmin(c);
		if error != nil{
			return nil,error;
		}

		return user.UserAo.Get(
			config.Atoi(c.Query("userId")),
		);
	}));

	router.POST("/add", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		userInfo := &user.User{
			Name:c.PostForm("name"),
			Type:config.Atoi(c.PostForm("type")),
			Password:c.PostForm("password"),
		}
		return nil,user.UserAo.Add(userInfo);
	}));

	router.POST("/del", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		return nil,user.UserAo.Del(
			config.Atoi(c.PostForm("userId")),
		);
	}));

	router.POST("/modType", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		return nil,user.UserAo.ModType(
			config.Atoi(c.PostForm("userId")),
			config.Atoi(c.PostForm("type")),
		);
	}));

	router.POST("/modPassword", view.Json( func(c *gin.Context)(interface{},error) {
		error := checkAdmin(c);
		if error != nil {
			return nil,error;
		}

		return nil,user.UserAo.ModPassword(
			config.Atoi(c.PostForm("userId")),
			c.PostForm("password"),
		);
	}));

	router.POST("/modMyPassword", view.Json( func(c *gin.Context)(interface{},error) {
		userinfo,error := user.LoginAo.IsLogin(c);
		if error != nil{
			return nil,error;
		}

		return nil,user.UserAo.ModPasswordByOld(
			userinfo.UserId,
			c.PostForm("oldPassword"),
			c.PostForm("newPassword"),
		);
	}));
}
