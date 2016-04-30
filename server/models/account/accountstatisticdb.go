package account

import (
	"mymanager/models/common"
)

type AccountStatisticDbModel struct {
	common.BaseModel
}

func (this *AccountStatisticDbModel) GetWeekTypStatisticByUser(userId int) []AccountStatistic {
	var data []AccountStatistic
	err := this.DB.
		Select("DATE_FORMAT(createTime,'%x') as year,TRIM( LEADING '0' From DATE_FORMAT(createTime,'%v')) as week,type,SUM(money) as money").
		Alias("t_account").
		Where("userId = ?", userId).
		GroupBy("year,week,type").
		OrderBy("year desc,week desc").
		Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func (this *AccountStatisticDbModel) GetWeekTypeStatisticDetailByUser(userId int, year int, week int, accountType int) []AccountStatisticDetail {
	var data []AccountStatisticDetail
	err := this.DB.
		Select("categoryId , sum(money) as money").
		Alias("t_account").
		Where("DATE_FORMAT(createTime,'%x') = ? and TRIM( LEADING '0' From DATE_FORMAT(createTime,'%v')) = ? and userId = ? and type = ?", year, week, userId, accountType).
		GroupBy("categoryId").
		Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func (this *AccountStatisticDbModel) GetWeekCardStatisticByUser(userId int) []AccountStatistic {
	var data []AccountStatistic
	err := this.DB.
		Where("userId = ?", userId).
		Alias("t_account").
		Select("DATE_FORMAT(createTime,'%x') as year,TRIM( LEADING '0' From DATE_FORMAT(createTime,'%v')) as week,cardId,type,SUM(money) as money").
		GroupBy("year,week,cardId,type").
		OrderBy("year desc,week desc").
		Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func (this *AccountStatisticDbModel) GetWeekCardStatisticDetailByUser(userId int, year int, week int, cardId int) []AccountStatisticDetail {
	var data []AccountStatisticDetail
	err := this.DB.
		Select("type , sum(money) as money").
		Alias("t_account").
		Where("DATE_FORMAT(createTime,'%x') = ? and TRIM( LEADING '0' From DATE_FORMAT(createTime,'%v')) = ? and userId = ? and cardId = ?", year, week, userId, cardId).
		GroupBy("type").
		Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}
