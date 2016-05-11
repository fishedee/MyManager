package user

import (
	"crypto/sha1"
	"fmt"
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	"io"
	. "mymanager/models/common"
)

type userAoModel struct {
	Model
	UserDb UserDbModel
}

func (this *userAoModel) CheckMustVaildPassword(password string, passwordHash string) {
	if this.getPasswordHash(password) != passwordHash {
		Throw(1, "密码不正确")
	}
}

func (this *userAoModel) getPasswordHash(password string) string {
	sha1er := sha1.New()
	io.WriteString(sha1er, password)
	dataHash := fmt.Sprintf("%x", sha1er.Sum(nil))
	return dataHash
}

func (this *userAoModel) Search(user User, pageInfo CommonPage) Users {
	return this.UserDb.Search(user, pageInfo)
}

func (this *userAoModel) Get(userId int) User {
	return this.UserDb.Get(userId)
}

func (this *userAoModel) GetByName(name string) []User {
	return this.UserDb.GetByName(name)
}

func (this *userAoModel) Del(userId int) {
	this.UserDb.Del(userId)
}

func (this *userAoModel) Add(user User) {
	//检查是否有重名
	userInfo := this.UserDb.GetByName(user.Name)
	if len(userInfo) != 0 {
		Throw(1, "存在重复的用户名")
	}

	//添加用户
	user.Password = this.getPasswordHash(user.Password)
	this.UserDb.Add(user)
}

func (this *userAoModel) ModType(userId int, userType int) {
	user := User{
		Type: userType,
	}
	this.UserDb.Mod(userId, user)
}

func (this *userAoModel) ModPassword(userId int, password string) {
	user := User{
		Password: this.getPasswordHash(password),
	}
	this.UserDb.Mod(userId, user)
}

func (this *userAoModel) ModPasswordByOld(userId int, oldPassword string, newPassword string) {
	//检查原密码是否正确
	userInfo := this.UserDb.Get(userId)
	this.CheckMustVaildPassword(oldPassword, userInfo.Password)

	//修改密码
	this.ModPassword(userId, newPassword)
}
