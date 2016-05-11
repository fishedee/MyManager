package card

import (
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	. "mymanager/models/user"
	"time"
)

type cardAoTest struct {
	Test
	CardAo     CardAoModel
	UserAo     UserAoModel
	UserAoTest UserAoTest
}

func (this *cardAoTest) InitEmpty() {
	var where Card
	limit := CommonPage{
		PageSize:  1000,
		PageIndex: 0,
	}
	Search := this.CardAo.Search(10001, where, limit)
	for _, v := range Search.Data {
		this.CardAo.Del(v.UserId, v.CardId)
	}
}

func (this *cardAoTest) InitSample() {
	this.InitEmpty()

	nowTime := time.Now().Truncate(time.Second)

	card10001 := Card{
		CardId:     10001,
		UserId:     10001,
		Name:       "银行卡",
		Bank:       "中国工商银行",
		Card:       "2348235923",
		Money:      11,
		Remark:     "哈哈",
		CreateTime: nowTime,
		ModifyTime: nowTime,
	}

	card10002 := Card{
		CardId:     10002,
		UserId:     10001,
		Name:       "信用卡",
		Bank:       "中国农业银行",
		Card:       "8989839489",
		Money:      99,
		Remark:     "哈哈",
		CreateTime: nowTime,
		ModifyTime: nowTime,
	}

	//添加卡片
	this.CardAo.Add(10001, card10001)
	this.CardAo.Add(10001, card10002)
}
func (this *cardAoTest) testAdd() (Card, Card) {
	card10001 := Card{
		CardId: 10001,
		UserId: 10001,
		Name:   "银行卡",
		Bank:   "中国工商银行",
		Card:   "2348235923",
		Money:  11,
		Remark: "哈哈",
	}

	card10002 := Card{
		CardId: 10003,
		UserId: 10001,
		Name:   "信用卡",
		Bank:   "中国农业银行",
		Card:   "8989839489",
		Money:  99,
		Remark: "哈哈",
	}

	//添加卡片
	this.CardAo.Add(10001, card10001)
	this.CardAo.Add(10001, card10002)

	return card10001, card10002
}
func (this *cardAoTest) testSearch(card10001 Card, card10002 Card) {
	//查看卡片
	cardData := this.CardAo.Get(10001, 10001)
	this.AssertEqual(cardData, card10001)

	//搜索所有卡片
	var where Card
	limit := CommonPage{
		PageSize:  99999,
		PageIndex: 0,
	}
	cardSearchData := this.CardAo.Search(10001, where, limit)
	cardDiffData := []Card{card10001, card10002}

	this.AssertEqual(cardSearchData.Count, 2)
	this.AssertEqual(cardSearchData.Data, cardDiffData)

	//查看不属于他的卡片
	_, err2 := this.CardAo.Get_WithError(10002, 10001)
	this.AssertError(err2, 1, "你没有权利查看或编辑等操作")

	//查看不存在的卡片
	_, err3 := this.CardAo.Get_WithError(10001, 88888888)
	this.AssertError(err3, 1, "该88888888银行卡不存在")

}
func (this *cardAoTest) testMod(card10001 Card) Card {
	//修改卡片
	card10001.Bank = "顺德农商银行"
	card10001.Card = "888888888"
	this.CardAo.Mod(10001, card10001.CardId, card10001)
	cardData2 := this.CardAo.Get(10001, 10001)
	this.AssertEqual(cardData2, card10001)

	//修改不属于他的卡片
	err4 := this.CardAo.Mod_WithError(10002, 10001, Card{
		CardId: 10001,
		UserId: 10001,
		Name:   "银行卡",
		Bank:   "顺德农商银行",
		Card:   "88888888888",
		Money:  0,
		Remark: "yes",
	})
	this.AssertError(err4, 1, "你没有权利查看或编辑等操作")

	//修改不存在的卡片
	err5 := this.CardAo.Mod_WithError(10001, 99999, Card{
		CardId: 99999,
		UserId: 10001,
		Name:   "银行卡",
		Bank:   "顺德农商银行",
		Card:   "88888888888",
		Money:  0,
		Remark: "yes",
	})
	this.AssertError(err5, 1, "该99999银行卡不存在")

	return card10001
}

func (this *cardAoTest) testDel() {
	//删除不属于他的卡片
	err6 := this.CardAo.Del_WithError(10002, 10001)
	this.AssertError(err6, 1, "你没有权利查看或编辑等操作")

	//删除不存在的卡片
	err7 := this.CardAo.Del_WithError(10001, 77777777)
	this.AssertError(err7, 1, "该77777777银行卡不存在")

	//删除卡片
	this.CardAo.Del(10001, 10001)
	_, err8 := this.CardAo.Get_WithError(10001, 10001)
	this.AssertError(err8, 1, "该10001银行卡不存在")

}
func (this *cardAoTest) TestBasic() {
	this.UserAoTest.InitSample()
	this.InitEmpty()

	card10001, card10002 := this.testAdd()
	this.testSearch(card10001, card10002)
	card10001 = this.testMod(card10001)
	this.testDel()

}
