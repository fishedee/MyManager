package user

import (
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	"strconv"
)

type userAoTest struct {
	Test
	UserAo      UserAoModel
	UserLoginAo UserLoginAoModel
}

func (this *userAoTest) InitEmpty() {
	var where User
	limit := CommonPage{
		PageSize:  1000,
		PageIndex: 0,
	}
	Search := this.UserAo.Search(where, limit)
	for _, v := range Search.Data {
		this.UserAo.Del(v.UserId)
	}
}

func (this *userAoTest) InitSample() {
	this.InitEmpty()
	users := []User{
		User{
			UserId:   10001,
			Name:     "fish",
			Password: "123",
			Type:     UserTypeEnum.ADMIN,
		},
		User{
			UserId:   10002,
			Name:     "edward",
			Password: "456",
			Type:     UserTypeEnum.USER,
		},
	}
	for _, singleUser := range users {
		this.UserAo.Add(singleUser)
	}
}
func (this *userAoTest) testAdd(userFish *User, userEdward *User) {
	//添加用户
	*userFish = User{
		UserId:   10001,
		Name:     "fish",
		Password: "123",
		Type:     UserTypeEnum.ADMIN,
	}

	this.UserAo.Add(*userFish)
	userFish.Password = "40bd001563085fc35165329ea1ff5c5ecbdbbeef"

	//重复用户
	userFish2 := User{
		Name:     "fish",
		Password: "aaa",
		Type:     UserTypeEnum.USER,
	}
	err := this.UserAo.Add_WithError(userFish2)
	this.AssertError(err, 1, "存在重复的用户名")

	//添加多一个用户
	*userEdward = User{
		UserId:   10002,
		Name:     "edward",
		Password: "123",
		Type:     UserTypeEnum.USER,
	}
	this.UserAo.Add(*userEdward)
	userEdward.Password = "40bd001563085fc35165329ea1ff5c5ecbdbbeef"

}

func (this *userAoTest) testSearch(userFish User, userEdward User) {
	//搜索所有用户
	var where2 User
	limit2 := CommonPage{
		PageSize:  10,
		PageIndex: 0,
	}
	searchUserData := this.UserAo.Search(where2, limit2)
	userData := []User{userFish, userEdward}

	this.AssertEqual(searchUserData.Count, 2)
	this.AssertEqual(searchUserData.Data, userData)

	//根据UserId搜索用户
	getFishData := this.UserAo.Get(userFish.UserId)
	this.AssertEqual(userFish, getFishData)

	//根据UserName搜索用户
	getEdwardData := this.UserAo.GetByName(userEdward.Name)
	this.AssertEqual(userEdward, getEdwardData[0])

}

func (this *userAoTest) testmod(userEdward *User) {
	//修改类型
	this.UserAo.ModType(userEdward.UserId, 1)
	userEdward.Type = 1
	getEdwardData2 := this.UserAo.Get(userEdward.UserId)
	this.AssertEqual(*userEdward, getEdwardData2)

}

func (this *userAoTest) testPassword(userEdward User) {
	//验证用户旧密码 改为新密码
	this.UserAo.ModPasswordByOld(userEdward.UserId, "123", "666")
	this.UserLoginAo.Login("edward", "666")
	UserData := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData.Name, "edward")

	//输入错误旧密码，修改不成功
	err2 := this.UserAo.ModPasswordByOld_WithError(userEdward.UserId, "errorPassword", "hacker")
	this.AssertError(err2, 1, "密码不正确")

	//直接修改用户密码
	this.UserAo.ModPassword(userEdward.UserId, "888")
	this.UserLoginAo.Login(userEdward.Name, "888")
	UserData2 := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData2.Name, "edward")
}
func (this *userAoTest) testDel(userEdward User) {
	//删除用户
	this.UserAo.Del(userEdward.UserId)
	_, err3 := this.UserAo.Get_WithError(userEdward.UserId)
	this.AssertError(err3, 1, "该"+strconv.Itoa(userEdward.UserId)+"用户不存在")

	//验证密码
	this.UserAo.CheckMustVaildPassword("123", "40bd001563085fc35165329ea1ff5c5ecbdbbeef")
	//验证错误密码
	err4 := this.UserAo.CheckMustVaildPassword_WithError("milkbobo", "40bd001563085fc35165329ea1ff5c5ecbdbbeef")
	this.AssertError(err4, 1, "密码不正确")

}
func (this *userAoTest) TestBasic() {
	this.InitEmpty()

	userFish := User{}
	userEdward := User{}

	this.testAdd(&userFish, &userEdward)
	this.testSearch(userFish, userEdward)
	this.testmod(&userEdward)
	this.testPassword(userEdward)
	this.testDel(userEdward)

}
