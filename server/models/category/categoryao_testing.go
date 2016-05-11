package category

import (
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	. "mymanager/models/user"
)

type categoryAoTest struct {
	Test
	CategoryAo CategoryAoModel
	UserAo     UserAoModel
	UserAoTest UserAoTest
}

func (this *categoryAoTest) InitEmpty() {
	var where Category
	limit := CommonPage{
		PageSize:  1000,
		PageIndex: 0,
	}
	Search := this.CategoryAo.Search(10001, where, limit)
	for _, v := range Search.Data {
		this.CategoryAo.Del(v.UserId, v.CategoryId)
	}
}

func (this *categoryAoTest) InitSample() {
	this.InitEmpty()

	categorys := []Category{
		Category{
			CategoryId: 10001,
			UserId:     10001,
			Name:       "生活用品",
			Remark:     "牛奶",
		},
		Category{
			CategoryId: 10002,
			UserId:     10001,
			Name:       "化妆品",
			Remark:     "面膜",
		},
		Category{
			CategoryId: 10003,
			UserId:     10001,
			Name:       "其他",
			Remark:     "其他吧",
		},
	}

	for _, singleCategory := range categorys {
		this.CategoryAo.Add(10001, singleCategory)
	}
}

func (this *categoryAoTest) assertCategoryEqual(categorys []Category) {
	var where Category
	limit := CommonPage{
		PageSize:  99999,
		PageIndex: 0,
	}
	categorySearchData := this.CategoryAo.Search(10001, where, limit)
	this.AssertEqual(categorySearchData.Count, len(categorys))
	this.AssertEqual(categorySearchData.Data, categorys)
}

func (this *categoryAoTest) testAdd() []Category {
	categorys := []Category{
		{
			CategoryId: 10001,
			UserId:     10001,
			Name:       "生活用品",
			Remark:     "牛奶",
		},
		{
			CategoryId: 10002,
			UserId:     10001,
			Name:       "化妆品",
			Remark:     "面膜",
		},
	}

	for _, singleCategory := range categorys {
		this.CategoryAo.Add(10001, singleCategory)
	}
	this.assertCategoryEqual(categorys)
	return categorys
}

func (this *categoryAoTest) testGet(categorys []Category) {
	//查看分类
	categoryData := this.CategoryAo.Get(10001, 10001)
	this.AssertEqual(categoryData, categorys[0])

	//查看不属于他的分类
	_, err2 := this.CategoryAo.Get_WithError(10002, 10001)
	this.AssertError(err2, 1, "你没有权利查看或编辑等操作")

	//查看不存在的分类
	_, err3 := this.CategoryAo.Get_WithError(10001, 88888)
	this.AssertError(err3, 1, "该88888分类不存在")
}

func (this *categoryAoTest) testMod(categorys []Category) {
	//修改分类
	modCategory := Category{
		CategoryId: 10001,
		UserId:     10001,
		Name:       "旅游费用",
		Remark:     "出国",
	}
	this.CategoryAo.Mod(10001, modCategory.CategoryId, modCategory)
	categorys[0] = modCategory
	this.assertCategoryEqual(categorys)

	//修改不属于他的分类
	err4 := this.CategoryAo.Mod_WithError(10002, 10001, Category{
		CategoryId: 10001,
		UserId:     10001,
		Name:       "黑客",
		Remark:     "被黑了",
	})
	this.AssertError(err4, 1, "你没有权利查看或编辑等操作")

	//修改不存在的分类
	err5 := this.CategoryAo.Mod_WithError(10001, 99999, Category{
		CategoryId: 99999,
		UserId:     10001,
		Name:       "黑客",
		Remark:     "被黑了",
	})
	this.AssertError(err5, 1, "该99999分类不存在")
}

func (this *categoryAoTest) testDel(categorys []Category) {
	//删除不属于他的分类
	err6 := this.CategoryAo.Del_WithError(10002, 10001)
	this.AssertError(err6, 1, "你没有权利查看或编辑等操作")

	//删除不存在的分类
	err7 := this.CategoryAo.Del_WithError(10001, 99999)
	this.AssertError(err7, 1, "该99999分类不存在")

	//删除分类
	this.CategoryAo.Del(10001, 10001)
	categorys = categorys[1:]
	this.assertCategoryEqual(categorys)
}

func (this *categoryAoTest) TestBasic() {
	this.UserAoTest.InitSample()
	this.InitEmpty()

	data := this.testAdd()
	this.testGet(data)
	this.testMod(data)
	this.testDel(data)
}
