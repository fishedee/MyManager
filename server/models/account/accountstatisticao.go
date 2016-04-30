package account

import (
	"fmt"
	"github.com/fishedee/language"
	"mymanager/models/card"
	"mymanager/models/category"
	"mymanager/models/common"
)

type AccountStatisticAoModel struct {
	common.BaseModel
	AccountDb  AccountStatisticDbModel
	CardAo     card.CardAoModel
	CategoryAo category.CategoryAoModel
}

func (this *AccountStatisticAoModel) GetWeekTypeStatistic(userId int) []AccountStatistic {
	statistic := this.AccountDb.GetWeekTypStatisticByUser(userId)
	enums := AccountTypeEnum.Datas()

	return language.QueryGroup(statistic, "   Year    desc    ,   Week    desc", func(weekStatistic []AccountStatistic) []AccountStatistic {
		single := weekStatistic[0]
		result := language.QueryLeftJoin(enums, weekStatistic, "Id = Type", func(left language.EnumData, right AccountStatistic) AccountStatistic {
			return AccountStatistic{
				Year: single.Year,
				Week: single.Week,
				Name: fmt.Sprintf(
					"%4d年%02d周",
					single.Year,
					single.Week,
				),
				Type:     left.Id,
				TypeName: left.Name,
				Money:    right.Money,
			}
		}).([]AccountStatistic)
		return result
	}).([]AccountStatistic)
}

func (this *AccountStatisticAoModel) GetWeekTypeStatisticDetail(userId int, year int, week int, accountType int) []AccountStatisticDetail {
	statistic := this.AccountDb.GetWeekTypeStatisticDetailByUser(userId, year, week, accountType)
	categorys := this.CategoryAo.Search(userId, category.Category{}, common.CommonPage{}).Data

	totalMoney := language.QuerySum(language.QueryColumn(statistic, "Money")).(int)
	return language.QueryLeftJoin(statistic, categorys, "CategoryId = CategoryId", func(left AccountStatisticDetail, right category.Category) AccountStatisticDetail {
		if right.Name == "" {
			right.Name = "无分类"
		}
		left.CategoryName = right.Name
		left.Precent = fmt.Sprintf("%.2f", float64(left.Money)/float64(totalMoney)*100)
		return left
	}).([]AccountStatisticDetail)
}

func (this *AccountStatisticAoModel) GetWeekCardStatistic(userId int) []AccountStatistic {
	statistic := this.AccountDb.GetWeekCardStatisticByUser(userId)
	cards := this.CardAo.Search(userId, card.Card{}, common.CommonPage{}).Data

	statistic = language.QueryGroup(statistic, "Year desc,Week desc,CardId desc", func(weekStatistic []AccountStatistic) []AccountStatistic {
		sum := language.QuerySum(language.QuerySelect(weekStatistic, func(singleStatistic AccountStatistic) int {
			if singleStatistic.Type == AccountTypeEnum.TYPE_BORROW_IN ||
				singleStatistic.Type == AccountTypeEnum.TYPE_IN ||
				singleStatistic.Type == AccountTypeEnum.TYPE_TRANSFER_IN {
				return singleStatistic.Money
			} else {
				return -singleStatistic.Money
			}
		})).(int)
		left := weekStatistic[0]
		left.Money = sum
		return []AccountStatistic{left}
	}).([]AccountStatistic)

	cardMoney := map[int]int{}
	statistic = language.QueryGroup(statistic, "Year asc ,Week asc", func(weekStatistic []AccountStatistic) []AccountStatistic {
		single := weekStatistic[0]
		return language.QueryLeftJoin(cards, weekStatistic, "CardId = CardId", func(left card.Card, right AccountStatistic) AccountStatistic {
			currentMoney, ok := cardMoney[left.CardId]
			if !ok {
				currentMoney = left.Money
			}
			currentMoney += right.Money
			cardMoney[left.CardId] = currentMoney
			return AccountStatistic{
				Year: single.Year,
				Week: single.Week,
				Name: fmt.Sprintf(
					"%4d年%02d周",
					single.Year,
					single.Week,
				),
				CardId:   left.CardId,
				CardName: left.Name,
				Money:    currentMoney,
			}
		}).([]AccountStatistic)
	}).([]AccountStatistic)

	return language.QueryReverse(statistic).([]AccountStatistic)
}

func (this *AccountStatisticAoModel) GetWeekCardStatisticDetail(userId int, year int, week int, cardId int) []AccountStatisticDetail {
	statistic := this.AccountDb.GetWeekCardStatisticDetailByUser(userId, year, week, cardId)
	enums := AccountTypeEnum.Datas()

	totalMoney := language.QuerySum(language.QueryColumn(statistic, "Money")).(int)
	return language.QueryLeftJoin(statistic, enums, "Type = Id", func(left AccountStatisticDetail, right language.EnumData) AccountStatisticDetail {
		left.TypeName = right.Name
		left.Precent = fmt.Sprintf("%.2f", float64(left.Money)/float64(totalMoney)*100)
		return left
	}).([]AccountStatisticDetail)
}
