package register

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
)

func (this *RegisterAoModel) Search_WithError(userId int, where Register, limit CommonPage) (_fishgen1 Registers, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(userId, where, limit)
	return
}

func (this *RegisterAoModel) Get_WithError(userId int, registerId int) (_fishgen1 Register, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId, registerId)
	return
}

func (this *RegisterAoModel) Del_WithError(userId int, registerId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId, registerId)
	return
}

func (this *RegisterAoModel) Add_WithError(userId int, register Register) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(userId, register)
	return
}

func (this *RegisterAoModel) Mod_WithError(userId int, registerId int, registerInfo Register) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(userId, registerId, registerInfo)
	return
}

func (this *RegisterDbModel) Search_WithError(where Register, limit CommonPage) (_fishgen1 Registers, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *RegisterDbModel) Get_WithError(registerId int) (_fishgen1 Register, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(registerId)
	return
}

func (this *RegisterDbModel) Del_WithError(registerId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(registerId)
	return
}

func (this *RegisterDbModel) Add_WithError(register Register) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(register)
	return
}

func (this *RegisterDbModel) Mod_WithError(registerId int, register Register) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(registerId, register)
	return
}
