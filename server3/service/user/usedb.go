package user

import (
	"../../config"
	"errors"
)

type UserDbData struct{

}

var UserDb = &UserDbData{};

func (this *UserDbData) Search(user *User,pageIndex int,pageSize int)(*Users,error){
	db := config.DB;
	if( user.Name != ""){
		db = db.Where("name like ?","%"+user.Name+"%");
	}
	if( user.Type != 0 ){
		db = db.Where("type = ?",user.Type);
	}

	var users Users;
	var error error;
	//error = db.Limit(pageSize).Offset(pageIndex).Find(&users.Data).Error;
	error = db.Find(&users.Data).Error;
	if error != nil{
		return nil,error;
	}

	error = db.Model(user).Count(&users.Count).Error;
	if error != nil{
		return nil,error;
	}
	return &users,error
}

func (this *UserDbData) Get(userId int)(*User,error){
	var user User;
	error := config.DB.Where(&User{UserId:userId}).Find(&user).Error;
	return &user,error;
}

func (this *UserDbData) Del(userId int)(error){
	return config.DB.Delete(&User{UserId:userId}).Error;
}
 
func (this *UserDbData) Add(user *User)(error){
	return config.DB.Create(&user).Error;
}

func (this *UserDbData) Mod(userId int,user* User)(error){
	return config.DB.Where(&User{UserId:userId}).Updates(&user).Error;
}

func (this *UserDbData) ModPassword(userId int,oldPassword string,newPassword string)(error){
	db := config.DB.Where(&User{Password:oldPassword}).Updates(&User{Password:newPassword});
	if db.Error != nil{
		return db.Error;
	}else if db.RowsAffected == 0 {
		return errors.New("原密码错误")
	}else{
		return nil;
	}
}

func (this *UserDbData) GetByIdAndPass(userId int,password string)([]*User,error){
	var users []*User;
	error := config.DB.Where(&User{UserId:userId,Password:password}).Find(&users).Error;
	return users,error;
}

func (this *UserDbData) GetByNameAndPass(name string,password string)([]*User,error){
	var users []*User;
	error := config.DB.Where(&User{Name:name,Password:password}).Find(&users).Error;
	return users,error;
}

func (this* UserDbData) GetByName(name string)([]*User,error){
	var users []*User;
	error := config.DB.Where(&User{Name:name}).Find(&users).Error;
	return users,error;
}