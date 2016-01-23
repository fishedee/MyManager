package account

import (
	"fmt"
	. "github.com/fishedee/language"
	. "mymanager/models/card"
	. "mymanager/models/category"
	. "mymanager/models/common"
)

type AccountStatisticAoModel struct {
	BaseModel
	AccountDb  AccountStatisticDbModel
	CardAo     CardAoModel
	CategoryAo CategoryAoModel
}

func (this *AccountStatisticAoModel) GetWeekTypeStatistic(userId int) []AccountStatistic {
	statistic := this.AccountDb.GetWeekTypStatisticByUser(userId)

	statisticYearWeekTypeMap := ArrayColumnMap(statistic, "Year", "Week", "Type").(map[int]map[int]map[int]AccountStatistic)
	statisticYearWeekSort := ArrayColumnSort(statistic, "Year", "Week").([]AccountStatistic)
	statisticYearWeekSort = ArrayColumnUnique(statisticYearWeekSort, "Year", "Week").([]AccountStatistic)

	result := []AccountStatistic{}
	for _, singleStatistic := range statisticYearWeekSort {
		year := singleStatistic.Year
		week := singleStatistic.Week
		for singleType, singleTypeName := range AccountTypeEnum.Entrys() {
			singleData := statisticYearWeekTypeMap[year][week][singleType]

			singleResult := AccountStatistic{
				Year:     year,
				Week:     week,
				Type:     singleType,
				TypeName: singleTypeName,
				Money:    singleData.Money,
				Name: fmt.Sprintf(
					"%4d年%2d月",
					year,
					week,
				),
			}
			result = append(result, singleResult)
		}
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
		singleResult.Precent = fmt.Sprintf("%.2f", float64(singleResult.Money)/float64(totalMoney)*100)
		result[key] = singleResult
	}
	return result
}

func (this *AccountStatisticAoModel) GetWeekCardStatistic(userId int) []AccountStatistic {
	statistic := this.AccountDb.GetWeekCardStatisticByUser(userId)
	statisticYearWeekCardMap := ArrayColumnMap(statistic, "Year", "Week", "CardId").(map[int]map[int]map[int]AccountStatistic)
	statisticYearWeekSort := ArrayColumnSort(statistic, "Year", "Week").([]AccountStatistic)
	statisticYearWeekSort = ArrayColumnUnique(statisticYearWeekSort, "Year", "Week").([]AccountStatistic)

	card := this.CardAo.Search(userId, Card{}, CommonPage{}).Data

	result := []AccountStatistic{}
	for _, singleStatistic := range statisticYearWeekSort {
		year := singleStatistic.Year
		week := singleStatistic.Week
		for key, singleCard := range card {
			singleData := statisticYearWeekCardMap[year][week][singleCard.CardId]
			singleCard.Money += singleData.Money
			card[key] = singleCard

			singleResult := AccountStatistic{
				Year:     year,
				Week:     week,
				CardId:   singleCard.CardId,
				CardName: singleCard.Name,
				Money:    singleCard.Money,
				Name: fmt.Sprintf(
					"%4d年%2d月",
					year,
					week,
				),
			}
			result = append(result, singleResult)
		}
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
		singleResult.Precent = fmt.Sprintf("%.2f", float64(singleResult.Money)/float64(totalMoney)*100)
		result[key] = singleResult
	}
	return result
}
