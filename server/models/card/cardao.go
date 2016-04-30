package card

import (
	"github.com/fishedee/language"
	"mymanager/models/common"
)

type CardAoModel struct {
	common.BaseModel
	CardDb CardDbModel
}

func (this *CardAoModel) Search(userId int, where Card, limit common.CommonPage) Cards {
	where.UserId = userId
	return this.CardDb.Search(where, limit)
}

func (this *CardAoModel) Get(userId int, cardId int) Card {
	cardInfo := this.CardDb.Get(cardId)
	if cardInfo.UserId != userId {
		language.Throw(1, "你没有该权限")
	}
	return cardInfo
}

func (this *CardAoModel) Del(userId int, cardId int) {
	this.Get(userId, cardId)

	this.CardDb.Del(cardId)

	this.Queue.Publish("card_del", cardId)
}

func (this *CardAoModel) Add(userId int, card Card) {
	card.UserId = userId
	this.CardDb.Add(card)
}

func (this *CardAoModel) Mod(userId int, cardId int, cardInfo Card) {
	this.Get(userId, cardId)

	cardInfo.UserId = userId
	this.CardDb.Mod(cardId, cardInfo)
}
