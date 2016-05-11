package card

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type CardAoModel interface {
	Search(userId int, where Card, limit CommonPage) (_fishgen1 Cards)
	Search_WithError(userId int, where Card, limit CommonPage) (_fishgen1 Cards, _fishgenErr Exception)
	Get(userId int, cardId int) (_fishgen1 Card)
	Get_WithError(userId int, cardId int) (_fishgen1 Card, _fishgenErr Exception)
	Del(userId int, cardId int)
	Del_WithError(userId int, cardId int) (_fishgenErr Exception)
	Add(userId int, card Card)
	Add_WithError(userId int, card Card) (_fishgenErr Exception)
	Mod(userId int, cardId int, cardInfo Card)
	Mod_WithError(userId int, cardId int, cardInfo Card) (_fishgenErr Exception)
}

type CardAoTest interface {
	InitEmpty()
	InitSample()
	TestBasic()
}

type CardDbModel interface {
	Search(where Card, limit CommonPage) (_fishgen1 Cards)
	Search_WithError(where Card, limit CommonPage) (_fishgen1 Cards, _fishgenErr Exception)
	Get(cardId int) (_fishgen1 Card)
	Get_WithError(cardId int) (_fishgen1 Card, _fishgenErr Exception)
	Del(cardId int)
	Del_WithError(cardId int) (_fishgenErr Exception)
	Add(card Card)
	Add_WithError(card Card) (_fishgenErr Exception)
	Mod(cardId int, card Card)
	Mod_WithError(cardId int, card Card) (_fishgenErr Exception)
}

func (this *cardAoModel) Search_WithError(userId int, where Card, limit CommonPage) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(userId, where, limit)
	return
}

func (this *cardAoModel) Get_WithError(userId int, cardId int) (_fishgen1 Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId, cardId)
	return
}

func (this *cardAoModel) Del_WithError(userId int, cardId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId, cardId)
	return
}

func (this *cardAoModel) Add_WithError(userId int, card Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(userId, card)
	return
}

func (this *cardAoModel) Mod_WithError(userId int, cardId int, cardInfo Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(userId, cardId, cardInfo)
	return
}

func (this *cardDbModel) Search_WithError(where Card, limit CommonPage) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *cardDbModel) Get_WithError(cardId int) (_fishgen1 Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(cardId)
	return
}

func (this *cardDbModel) Del_WithError(cardId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(cardId)
	return
}

func (this *cardDbModel) Add_WithError(card Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(card)
	return
}

func (this *cardDbModel) Mod_WithError(cardId int, card Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(cardId, card)
	return
}
func init() {
	v0 := CardAoModel(&cardAoModel{})
	InitModel(&v0)
	v1 := CardAoTest(&cardAoTest{})
	InitTest(&v1)
	v2 := CardDbModel(&cardDbModel{})
	InitModel(&v2)
}
