package card

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
)

func (this *CardAoModel) Search_WithError(userId int, where Card, limit CommonPage) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(userId, where, limit)
	return
}

func (this *CardAoModel) Get_WithError(userId int, cardId int) (_fishgen1 Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(userId, cardId)
	return
}

func (this *CardAoModel) Del_WithError(userId int, cardId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(userId, cardId)
	return
}

func (this *CardAoModel) Add_WithError(userId int, card Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(userId, card)
	return
}

func (this *CardAoModel) Mod_WithError(userId int, cardId int, cardInfo Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(userId, cardId, cardInfo)
	return
}

func (this *CardDbModel) Search_WithError(where Card, limit CommonPage) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *CardDbModel) Get_WithError(cardId int) (_fishgen1 Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(cardId)
	return
}

func (this *CardDbModel) Del_WithError(cardId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Del(cardId)
	return
}

func (this *CardDbModel) Add_WithError(card Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(card)
	return
}

func (this *CardDbModel) Mod_WithError(cardId int, card Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(cardId, card)
	return
}
