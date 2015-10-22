package user;

import (
	"../../config"
	"errors"
)

type UserAoData struct {

}

var UserAo UserAoData;

func (this* UserAoData) Search(user *User,pageIndex int ,pageSize int)(*Users,error){
	return UserDb.Search(user,pageIndex,pageSize);
}

func (this* UserAoData) Get(userId int)(*User,error){
	return UserDb.Get(userId)
}

func (this* UserAoData) Del(userId int)(error){
	return UserDb.Del(userId);
}

func (this* UserAoData) Add(user *User)(error){
	users,error := UserDb.GetByName(user.Name);
	if error != nil{
		return error;
	}
	if len(users) != 0{
		return errors.New("存在重复的用户名")
	}

	user.Password = config.Sha1(user.Password)
	return UserDb.Add(user);
}

func (this* UserAoData) ModType(userId int,Type int)(error){
	return UserDb.Mod(userId,&User{Type:Type});
}

func (this* UserAoData) ModPassword(userId int,password string)(error){
	return UserDb.Mod(userId,&User{Password:config.Sha1(password)});
}

func (this* UserAoData) ModPasswordByOld(userId int,oldPassword string,newPassword string)(error){
	users,error := UserDb.GetByIdAndPass(userId,config.Sha1(oldPassword));
	if error != nil{
		return error;
	}
	if len(users) == 0{
		return errors.New("原密码错误");
	}

	return UserDb.Mod(userId,&User{Password:config.Sha1(newPassword)});
}
