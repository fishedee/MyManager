package account

/*
import (
	"fmt"
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/card"
	. "mymanager/models/category"
	. "mymanager/models/common"
)

type AccountStatisticAoModel struct {
	BaseModel
	AccountDb  AccountDbModel
	CardAo     CardAoModel
	CategoryAo CategoryAoModel
}

func (this *AccountStatisticAoModel) GetWeekTypeStatistic(userId int) []AccountStatistic {
	result := this.AccountDb.GetWeekTypStatisticByUser(userId)
	typeMap := AccountTypeEnum.Entrys()
	for key, singleResult := range result {
		singleResult.Name = fmt.Sprintf(
			"%4d年%2d月",
			singleResult.Year,
			singleResult.Week,
		)
		singleResult.TypeName = typeMap[singleResult.Type]
		result[key] = singleResult
	}
	return result
}

func (this *AccountStatisticAoModel) GetWeekTypeStatisticDetail(userId int, year int, week int, accountType int) []AccountStatisticDetail {
	result := this.AccountDb.GetWeekTypeStatisticDetailByUser(userId, year, week, accountType)

	category := this.CategoryAo.Search(userId, Category{}, CommonPage{}).Data
	categoryMap := ArrayColumnMap(category, "CategoryId").(map[int]Category)

	totalMoney := 0
	for _, singleResult := range result {
		totalMoney += singleResult.Money
	}
	for key, singleResult := range result {
		var categoryName string
		category, ok := categoryMap[singleResult.CategoryId]
		if !ok {
			categoryName = "无分类"
		} else {
			categoryName = category.Name
		}
		singleResult.CategoryName = categoryName
		singleResult.Precent = Sprintf("%.2f", float64(singleResult.Money)/float64(totalMoney)*100)
		result[key] = singleResult
	}
	return result
}

func (this *AccountStatisticAoModel) GetWeekCardStatistic(userId int) []AccountStatistic {
	result := this.AccountDb.GetWeekCardStatisticByUser(userId)

	card := this.CardAo.Search(userId, Card{}, CommonPage{}).Data
	cardMap := ArrayColumnMap(card, "CardId").(map[int]Card)

	for key, singleResult := range result {
		singleResult.Name = fmt.Sprintf(
			"%4d年%2d月",
			singleResult.Year,
			singleResult.Week,
		)
		singleResult.CardName = cardMap[singleResult.CardId]
		result[key] = singleResult
	}
	return result
}

func (this *AccountStatisticAoModel) GetWeekCardStatisticDetail(userId int, year int, week int, cardId int) []AccountStatisticDetail {
	result := this.AccountDb.GetWeekCardStatisticDetailByUser(userId, year, week, cardId)

	typeMap := AccountTypeEnum.Entrys()

	totalMoney := 0
	for _, singleResult := range result {
		totalMoney += singleResult.Money
	}
	for key, singleResult := range result {
		singleResult.TypeName = typeMap[singleResult.Type]
		singleResult.Precent = Sprintf("%.2f", float64(singleResult.Money)/float64(totalMoney)*100)
		result[key] = singleResult
	}
	return result
}
*/
