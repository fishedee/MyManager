package account

import (
	. "github.com/fishedee/web"
	. "mymanager/models/card"
	. "mymanager/models/category"
	. "mymanager/models/common"
	. "mymanager/models/user"
	"time"
)

type AccountAoTest struct {
	Test
	AccountAo          AccountAoModel
	AccountStatisticAo AccountStatisticAoModel
	UserAoTest         UserAoTest
	CategoryAoTest     CategoryAoTest
	CategoryAo         CategoryAoModel
	CardAoTest         CardAoTest
	CardAo             CardAoModel
}

func (this *AccountAoTest) InitEmpty() {
	var where Account
	limit := CommonPage{
		PageSize:  1000,
		PageIndex: 0,
	}
	Search := this.AccountAo.Search(10001, where, limit)
	for _, v := range Search.Data {
		this.AccountAo.Del(v.UserId, v.AccountId)
	}
}

func (this *AccountAoTest) testAdd() (Account, Account) {
	account1 := Account{
		AccountId:  10001,
		UserId:     10001,
		Name:       "中了彩票",
		Money:      2000,
		Remark:     "二等奖哦！",
		CategoryId: 10001,
		CardId:     10001,
		Type:       AccountTypeEnum.TYPE_IN,
	}

	account2 := Account{
		AccountId:  10002,
		UserId:     10001,
		Name:       "自助餐",
		Money:      399,
		Remark:     "哥顿自助餐",
		CategoryId: 10001,
		CardId:     10001,
		Type:       AccountTypeEnum.TYPE_OUT,
	}

	//添加卡片
	this.AccountAo.Add(10001, account1)
	this.AccountAo.Add(10001, account2)

	return account1, account2
}

func (this *AccountAoTest) testSearch(account1 Account, account2 Account) {
	//查看账务
	accountData := this.AccountAo.Get(10001, 10001)
	this.AssertEqual(accountData, account1)

	//搜索所有分类
	var where Account
	limit := CommonPage{
		PageSize:  99999,
		PageIndex: 0,
	}

	accountSearchData := this.AccountAo.Search(10001, where, limit)

	this.AssertEqual(accountSearchData.Count, 2)
	this.AssertEqual(accountSearchData.Data, []Account{account1, account2})

	//查看不属于他的分类
	_, err2 := this.AccountAo.Get_WithError(10002, 10001)
	this.AssertError(err2, 1, "你没有权利查看或编辑等操作")

	//查看不存在的分类
	_, err3 := this.AccountAo.Get_WithError(10001, 88888)
	this.AssertError(err3, 1, "该88888账务不存在")
}

func (this *AccountAoTest) testMod(account Account) Account {
	//修改账务
	account.Name = "小肥牛自助餐"
	account.Money = 88
	account.Remark = "这个更加便宜"
	this.AccountAo.Mod(10001, account.AccountId, account)
	accountData2 := this.AccountAo.Get(10001, 10001)
	this.AssertEqual(accountData2, account)

	//修改不属于他的账务
	err4 := this.AccountAo.Mod_WithError(10002, 10001, Account{
		AccountId:  10001,
		UserId:     10002,
		Name:       "小排档",
		Money:      50,
		Remark:     "吃粥",
		CategoryId: 10002,
		CardId:     10002,
		Type:       AccountTypeEnum.TYPE_OUT,
	})
	this.AssertError(err4, 1, "你没有权利查看或编辑等操作")

	//修改不存在的账务
	err5 := this.AccountAo.Mod_WithError(10001, 99999, Account{
		AccountId:  99999,
		UserId:     10002,
		Name:       "小排档",
		Money:      50,
		Remark:     "吃粥",
		CategoryId: 10002,
		CardId:     10002,
		Type:       AccountTypeEnum.TYPE_OUT,
	})
	this.AssertError(err5, 1, "该99999账务不存在")

	return account
}

func (this *AccountAoTest) testDel() {
	//删除不属于他的账务
	err6 := this.AccountAo.Del_WithError(10002, 10001)
	this.AssertError(err6, 1, "你没有权利查看或编辑等操作")

	//删除不存在的账务
	err7 := this.AccountAo.Del_WithError(10001, 77777777)
	this.AssertError(err7, 1, "该77777777账务不存在")

	//删除category类型，账务上的category为0
	this.CategoryAo.Del(10001, 10001)
	accountData := this.AccountAo.Get(10001, 10001)
	this.AssertEqual(accountData.CategoryId, 0)

	//删除card类型，账务上的cardId为0
	this.CardAo.Del(10001, 10001)
	accountData2 := this.AccountAo.Get(10001, 10001)
	this.AssertEqual(accountData2.CardId, 0)

	//删除账务
	this.AccountAo.Del(10001, 10001)
	_, err8 := this.AccountAo.Get_WithError(10001, 10001)
	this.AssertError(err8, 1, "该10001账务不存在")

}

func (this *AccountAoTest) testStatisticsAdd() {

	nowTime := time.Date(2016, 4, 21, 13, 0, 0, 0, time.Local).Truncate(time.Second)
	oldTime := nowTime.AddDate(0, -1, 0).Truncate(time.Second)

	accountAddData := []Account{
		Account{
			UserId:     10001,
			Name:       "中了彩票",
			Money:      2000,
			Remark:     "二等奖哦！",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: oldTime,
			ModifyTime: oldTime,
		},

		Account{
			UserId:     10001,
			Name:       "自助餐",
			Money:      399,
			Remark:     "哥顿自助餐",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_OUT,
			CreateTime: oldTime,
			ModifyTime: oldTime,
		},
		Account{
			UserId:     10001,
			Name:       "工资",
			Money:      1000,
			Remark:     "哈哈公司",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_TRANSFER_IN,
			CreateTime: oldTime,
			ModifyTime: oldTime,
		},
		Account{
			UserId:     10001,
			Name:       "淘宝",
			Money:      200,
			Remark:     "购物",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_TRANSFER_OUT,
			CreateTime: oldTime,
			ModifyTime: oldTime,
		},
		Account{
			UserId:     10001,
			Name:       "收到还钱",
			Money:      616,
			Remark:     "还钱",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_BORROW_IN,
			CreateTime: oldTime,
			ModifyTime: oldTime,
		},
		Account{
			UserId:     10001,
			Name:       "借钱给朋友",
			Money:      363,
			Remark:     "借钱",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_BORROW_OUT,
			CreateTime: oldTime,
			ModifyTime: oldTime,
		},
		Account{
			UserId:     10001,
			Name:       "中了彩票",
			Money:      3000,
			Remark:     "二等奖哦！",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},

		Account{
			UserId:     10001,
			Name:       "自助餐",
			Money:      99,
			Remark:     "哥顿自助餐",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_OUT,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "工资",
			Money:      5000,
			Remark:     "哈哈公司",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_TRANSFER_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "淘宝",
			Money:      300,
			Remark:     "购物",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_TRANSFER_OUT,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "收到还钱",
			Money:      666,
			Remark:     "还钱",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_BORROW_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "借钱给朋友",
			Money:      333,
			Remark:     "借钱",
			CategoryId: 10001,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_BORROW_OUT,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "捡到钱了",
			Money:      100,
			Remark:     "真幸运",
			CategoryId: 10002,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "微信红包",
			Money:      30,
			Remark:     "红包",
			CategoryId: 10002,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "捡到钱哦1",
			Money:      50,
			Remark:     "幸运",
			CategoryId: 10003,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "捡到钱哦2",
			Money:      510,
			Remark:     "幸运",
			CategoryId: 10003,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "捡到钱哦3",
			Money:      150,
			Remark:     "幸运",
			CategoryId: 10003,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
		Account{
			UserId:     10001,
			Name:       "捡到钱哦4",
			Money:      520,
			Remark:     "幸运",
			CategoryId: 10003,
			CardId:     10001,
			Type:       AccountTypeEnum.TYPE_IN,
			CreateTime: nowTime,
			ModifyTime: nowTime,
		},
	}

	//添加账务
	for index, singleAddData := range accountAddData {
		singleAddData.AccountId = 10001 + index
		this.AccountAo.Add(10001, singleAddData)
		this.DB.Where("AccountId = ?", singleAddData.AccountId).NoAutoTime().Cols("CreateTime,ModifyTime").Update(&singleAddData)
	}
}

func (this *AccountAoTest) testGetWeekTypeStatistic() []AccountStatistic {
	getWeekTypeStatistic := this.AccountStatisticAo.GetWeekTypeStatistic(10001)

	this.AssertEqual(getWeekTypeStatistic, []AccountStatistic{
		AccountStatistic{CardId: 0, CardName: "", Money: 4360, Name: "2016年16周", Type: 1, TypeName: "收入", Week: 16, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 99, Name: "2016年16周", Type: 2, TypeName: "支出", Week: 16, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 5000, Name: "2016年16周", Type: 3, TypeName: "转账收入", Week: 16, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 300, Name: "2016年16周", Type: 4, TypeName: "转账支出", Week: 16, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 666, Name: "2016年16周", Type: 5, TypeName: "借还款收入", Week: 16, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 333, Name: "2016年16周", Type: 6, TypeName: "借还款支出", Week: 16, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 2000, Name: "2016年12周", Type: 1, TypeName: "收入", Week: 12, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 399, Name: "2016年12周", Type: 2, TypeName: "支出", Week: 12, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 1000, Name: "2016年12周", Type: 3, TypeName: "转账收入", Week: 12, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 200, Name: "2016年12周", Type: 4, TypeName: "转账支出", Week: 12, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 616, Name: "2016年12周", Type: 5, TypeName: "借还款收入", Week: 12, Year: 2016},
		AccountStatistic{CardId: 0, CardName: "", Money: 363, Name: "2016年12周", Type: 6, TypeName: "借还款支出", Week: 12, Year: 2016},
	})
	return getWeekTypeStatistic

}
func (this *AccountAoTest) testGetWeekDetailTypeStatistic(getWeekTypeStatistic []AccountStatistic) {
	getWeekDetailTypeStatistic := this.AccountStatisticAo.GetWeekTypeStatisticDetail(
		10001,
		getWeekTypeStatistic[0].Year,
		getWeekTypeStatistic[0].Week,
		getWeekTypeStatistic[0].Type,
	)

	this.AssertEqual(getWeekDetailTypeStatistic, []AccountStatisticDetail{
		AccountStatisticDetail{CategoryId: 10001, CategoryName: "生活用品", Type: 0, TypeName: "", Money: 3000, Precent: "68.81"},
		AccountStatisticDetail{CategoryId: 10002, CategoryName: "化妆品", Type: 0, TypeName: "", Money: 130, Precent: "2.98"},
		AccountStatisticDetail{CategoryId: 10003, CategoryName: "其他", Type: 0, TypeName: "", Money: 1230, Precent: "28.21"},
	})

}
func (this *AccountAoTest) testGetWeekCardStatistic() []AccountStatistic {
	getWeekCardStatistic := this.AccountStatisticAo.GetWeekCardStatistic(10001)

	this.AssertEqual(getWeekCardStatistic, []AccountStatistic{
		AccountStatistic{CardId: 10001, CardName: "银行卡", Money: 11959, Name: "2016年16周", Type: 0, TypeName: "", Week: 16, Year: 2016},
		AccountStatistic{CardId: 10002, CardName: "信用卡", Money: 99, Name: "2016年16周", Type: 0, TypeName: "", Week: 16, Year: 2016},
		AccountStatistic{CardId: 10001, CardName: "银行卡", Money: 2665, Name: "2016年12周", Type: 0, TypeName: "", Week: 12, Year: 2016},
		AccountStatistic{CardId: 10002, CardName: "信用卡", Money: 99, Name: "2016年12周", Type: 0, TypeName: "", Week: 12, Year: 2016},
	})

	return getWeekCardStatistic

}
func (this *AccountAoTest) testGetWeekDetailCardStatistic(getWeekCardStatistic []AccountStatistic) {

	getWeekDetailCardStatistic := this.AccountStatisticAo.GetWeekCardStatisticDetail(
		10001,
		getWeekCardStatistic[0].Year,
		getWeekCardStatistic[0].Week,
		getWeekCardStatistic[0].CardId,
	)

	this.AssertEqual(getWeekDetailCardStatistic, []AccountStatisticDetail{
		AccountStatisticDetail{CategoryId: 0, CategoryName: "", Type: 1, TypeName: "收入", Money: 4360, Precent: "40.53"},
		AccountStatisticDetail{CategoryId: 0, CategoryName: "", Type: 2, TypeName: "支出", Money: 99, Precent: "0.92"},
		AccountStatisticDetail{CategoryId: 0, CategoryName: "", Type: 3, TypeName: "转账收入", Money: 5000, Precent: "46.48"},
		AccountStatisticDetail{CategoryId: 0, CategoryName: "", Type: 4, TypeName: "转账支出", Money: 300, Precent: "2.79"},
		AccountStatisticDetail{CategoryId: 0, CategoryName: "", Type: 5, TypeName: "借还款收入", Money: 666, Precent: "6.19"},
		AccountStatisticDetail{CategoryId: 0, CategoryName: "", Type: 6, TypeName: "借还款支出", Money: 333, Precent: "3.10"},
	})
}

func (this *AccountAoTest) testStatistics() {
	getWeekTypeStatistic := this.testGetWeekTypeStatistic()
	this.testGetWeekDetailTypeStatistic(getWeekTypeStatistic)

	getWeekCardStatistic := this.testGetWeekCardStatistic()
	this.testGetWeekDetailCardStatistic(getWeekCardStatistic)
}

func (this *AccountAoTest) TestBasic() {
	this.UserAoTest.InitSample()
	this.CategoryAoTest.InitSample()
	this.CardAoTest.InitSample()
	this.InitEmpty()

	account1, account2 := this.testAdd()
	this.testSearch(account1, account2)
	account1 = this.testMod(account1)
	this.testDel()
}

func (this *AccountAoTest) TestStatistic() {
	this.UserAoTest.InitSample()
	this.CategoryAoTest.InitSample()
	this.CardAoTest.InitSample()
	this.InitEmpty()

	this.testStatisticsAdd()
	this.testStatistics()
}

func init() {
	InitTest(&AccountAoTest{})
}
