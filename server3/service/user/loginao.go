package user

import (
   	"../../config"
   	"errors"
   	"github.com/gin-gonic/gin"
)

type LoginAoData struct{

}

var LoginAo LoginAoData;

func (this *LoginAoData) IsLogin(c* gin.Context)(*User,error){
	session, error := config.Session.Get(c);
	if error != nil {
		return nil,error
	}

	userId , ok := session["userId"];
	userIdInt , ok2 := userId.(int);
	if( ok && ok2 && userIdInt >= 1000 ){
		user , error := UserDb.Get(userIdInt);
		if error != nil{
			return nil,error;
		}
		return user,nil;
	}else{
		return nil,errors.New("帐号未登录")
	}
}

func (this *LoginAoData) Logout(c* gin.Context)(error){
	return config.Session.Destroy(c);
}

func (this *LoginAoData) Login(c* gin.Context,name string,password string)(error){
	users,error := UserDb.GetByNameAndPass(name,config.Sha1(password));
	if error != nil{
		return error;
	}

	if len(users) == 0{
		return errors.New("账号或密码错误")
	}

	return config.Session.Set(c,"userId",users[0].UserId);
}