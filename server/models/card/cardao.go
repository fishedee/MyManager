package card

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type CardAoModel struct {
	Model
	CardDb CardDbModel
}

func (this *CardAoModel) Search(userId int, where Card, limit CommonPage) Cards {
	where.UserId = userId
	return this.CardDb.Search(where, limit)
}

func (this *CardAoModel) Get(userId int, cardId int) Card {
	cardInfo := this.CardDb.Get(cardId)
	if cardInfo.UserId != userId {
		Throw(1, "你没有权利查看或编辑等操作")
	}
	return cardInfo
}

func (this *CardAoModel) Del(userId int, cardId int) {
	this.Get(userId, cardId)

	this.CardDb.Del(cardId)

	this.Queue.Publish(CardQueueEnum.EVENT_DEL, cardId)
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
