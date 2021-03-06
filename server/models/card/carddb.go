package card

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type CardDbModel struct {
	Model
}

func (this *CardDbModel) Search(where Card, limit CommonPage) Cards {
	db := this.DB.NewSession()

	if where.Name != "" {
		db = db.Where("name like ?", "%"+where.Name+"%")
	}
	if where.Remark != "" {
		db = db.Where("name like ?", "%"+where.Remark+"%")
	}
	if where.UserId != 0 {
		db = db.Where("userId = ? ", where.UserId)
	}

	data := []Card{}
	err := db.OrderBy("createTime desc").Limit(limit.PageSize, limit.PageIndex).Find(&data)
	if err != nil {
		panic(err)
	}

	count, err := db.Count(&where)
	if err != nil {
		panic(err)
	}

	return Cards{
		Count: int(count),
		Data:  data,
	}
}

func (this *CardDbModel) Get(cardId int) Card {
	var cards []Card
	err := this.DB.Where("cardId = ?", cardId).Find(&cards)
	if err != nil {
		panic(err)
	}
	if len(cards) == 0 {
		Throw(1, "该"+strconv.Itoa(cardId)+"银行卡不存在")
	}
	return cards[0]
}

func (this *CardDbModel) Del(cardId int) {
	_, err := this.DB.Where("cardId = ?", cardId).Delete(&Card{})
	if err != nil {
		panic(err)
	}
}

func (this *CardDbModel) Add(card Card) {
	_, err := this.DB.Insert(card)
	if err != nil {
		panic(err)
	}
}

func (this *CardDbModel) Mod(cardId int, card Card) {
	_, err := this.DB.Where("cardId = ?", cardId).Update(&card)
	if err != nil {
		panic(err)
	}
}
