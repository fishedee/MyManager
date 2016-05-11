package user

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type UserAoModel interface {
	CheckMustVaildPassword(password string, passwordHash string)
	CheckMustVaildPassword_WithError(password string, passwordHash string) (_fishgenErr Exception)
	Search(user User, pageInfo CommonPage) (_fishgen1 Users)
	Search_WithError(user User, pageInfo CommonPage) (_fishgen1 Users, _fishgenErr Exception)
	Get(userId int) (_fishgen1 User)
	Get_WithError(userId int) (_fishgen1 User, _fishgenErr Exception)
	GetByName(name string) (_fishgen1 []User)
	GetByName_WithError(name string) (_fishgen1 []User, _fishgenErr Exception)
	Del(userId int)
	Del_WithError(userId int) (_fishgenErr Exception)
	Add(user User)
	Add_WithError(user User) (_fishgenErr Exception)
	ModType(userId int, userType int)
	ModType_WithError(userId int, userType int) (_fishgenErr Exception)
	ModPassword(userId int, password string)
	ModPassword_WithError(userId int, password string) (_fishgenErr Exception)
	ModPasswordByOld(userId int, oldPassword string, newPassword string)
	ModPasswordByOld_WithError(userId int, oldPassword string, newPassword string) (_fishgenErr Exception)
}

type UserAoTest interface {
	InitEmpty()
	InitSample()
	TestBasic()
}

type UserDbModel interface {
	Search(where User, limit CommonPage) (_fishgen1 Users)
	Search_WithError(where User, limit CommonPage) (_fishgen1 Users, _fishgenErr Exception)
	Get(id int) (_fishgen1 User)
	Get_WithError(id int) (_fishgen1 User, _fishgenErr Exception)
	GetByName(name string) (_fishgen1 []User)
	GetByName_WithError(name string) (_fishgen1 []User, _fishgenErr Exception)
	Del(id int)
	Del_WithError(id int) (_fishgenErr Exception)
	Add(user User)
	Add_WithError(user User) (_fishgenErr Exception)
	Mod(id int, user User)
	Mod_WithError(id int, user User) (_fishgenErr Exception)
}

type UserLoginAoModel interface {
	IsLogin() (_fishgen1 User)
	IsLogin_WithError() (_fishgen1 User, _fishgenErr Exception)
	CheckMustLogin() (_fishgen1 User)
	CheckMustLogin_WithError() (_fishgen1 User, _fishgenErr Exception)
	CheckMustAdmin() (_fishgen1 User)
	CheckMustAdmin_WithError() (_fishgen1 User, _fishgenErr Exception)
	Logout()
	Logout_WithError() (_fishgenErr Exception)
	Login(name string, password string)
	Login_WithError(name string, password string) (_fishgenErr Exception)
}

type UserLoginAoTest interface {
	TestBasic()
}

func (this *userAoModel) CheckMustVaildPassword_WithError(password string, passwordHash string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.CheckMustVaildPassword(password, passwordHash)
	return
}

func (this *userAoModel) Search_WithError(user User, pageInfo CommonPage) (_fishgen1 Users, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(user, pageInfo)
	return
}

func (this *userAoModel) Get_WithError(userId int) (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId)
	return
}

func (this *userAoModel) GetByName_WithError(name string) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetByName(name)
	return
}

func (this *userAoModel) Del_WithError(userId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId)
	return
}

func (this *userAoModel) Add_WithError(user User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(user)
	return
}

func (this *userAoModel) ModType_WithError(userId int, userType int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModType(userId, userType)
	return
}

func (this *userAoModel) ModPassword_WithError(userId int, password string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModPassword(userId, password)
	return
}

func (this *userAoModel) ModPasswordByOld_WithError(userId int, oldPassword string, newPassword string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModPasswordByOld(userId, oldPassword, newPassword)
	return
}

func (this *userDbModel) Search_WithError(where User, limit CommonPage) (_fishgen1 Users, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *userDbModel) Get_WithError(id int) (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(id)
	return
}

func (this *userDbModel) GetByName_WithError(name string) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetByName(name)
	return
}

func (this *userDbModel) Del_WithError(id int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(id)
	return
}

func (this *userDbModel) Add_WithError(user User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(user)
	return
}

func (this *userDbModel) Mod_WithError(id int, user User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(id, user)
	return
}

func (this *userLoginAoModel) IsLogin_WithError() (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.IsLogin()
	return
}

func (this *userLoginAoModel) CheckMustLogin_WithError() (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.CheckMustLogin()
	return
}

func (this *userLoginAoModel) CheckMustAdmin_WithError() (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.CheckMustAdmin()
	return
}

func (this *userLoginAoModel) Logout_WithError() (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Logout()
	return
}

func (this *userLoginAoModel) Login_WithError(name string, password string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Login(name, password)
	return
}

func init() {
	v0 := UserAoModel(&userAoModel{})
	InitModel(&v0)
	v1 := UserAoTest(&userAoTest{})
	InitTest(&v1)
	v2 := UserDbModel(&userDbModel{})
	InitModel(&v2)
	v3 := UserLoginAoModel(&userLoginAoModel{})
	InitModel(&v3)
	v4 := UserLoginAoTest(&userLoginAoTest{})
	InitTest(&v4)
}
