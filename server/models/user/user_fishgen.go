package user

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
)

func (this *UserAoModel) CheckMustVaildPassword_WithError(password string, passwordHash string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.CheckMustVaildPassword(password, passwordHash)
	return
}

func (this *UserAoModel) Search_WithError(user User, pageInfo CommonPage) (_fishgen1 Users, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(user, pageInfo)
	return
}

func (this *UserAoModel) Get_WithError(userId int) (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId)
	return
}

func (this *UserAoModel) GetByName_WithError(name string) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetByName(name)
	return
}

func (this *UserAoModel) Del_WithError(userId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId)
	return
}

func (this *UserAoModel) Add_WithError(user User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(user)
	return
}

func (this *UserAoModel) ModType_WithError(userId int, userType int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModType(userId, userType)
	return
}

func (this *UserAoModel) ModPassword_WithError(userId int, password string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModPassword(userId, password)
	return
}

func (this *UserAoModel) ModPasswordByOld_WithError(userId int, oldPassword string, newPassword string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.ModPasswordByOld(userId, oldPassword, newPassword)
	return
}

func (this *UserDbModel) Search_WithError(where User, limit CommonPage) (_fishgen1 Users, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *UserDbModel) Get_WithError(id int) (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(id)
	return
}

func (this *UserDbModel) GetByName_WithError(name string) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetByName(name)
	return
}

func (this *UserDbModel) Del_WithError(id int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(id)
	return
}

func (this *UserDbModel) Add_WithError(user User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(user)
	return
}

func (this *UserDbModel) Mod_WithError(id int, user User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(id, user)
	return
}

func (this *UserLoginAoModel) IsLogin_WithError() (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.IsLogin()
	return
}

func (this *UserLoginAoModel) CheckMustLogin_WithError() (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.CheckMustLogin()
	return
}

func (this *UserLoginAoModel) CheckMustAdmin_WithError() (_fishgen1 User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.CheckMustAdmin()
	return
}

func (this *UserLoginAoModel) Logout_WithError() (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Logout()
	return
}

func (this *UserLoginAoModel) Login_WithError(name string, password string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Login(name, password)
	return
}
