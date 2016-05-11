package user

import (
	. "github.com/fishedee/web"
)

type userLoginAoTest struct {
	Test
	UserLoginAo UserLoginAoModel
	UserAoTest  UserAoTest
}

func (this *userLoginAoTest) TestBasic() {
	this.UserAoTest.InitSample()

	//没有登录
	this.UserLoginAo.IsLogin()
	_, err := this.UserLoginAo.CheckMustLogin_WithError()
	this.AssertError(err, 1, "用户未登录")
	// this.AssertEqual(err != nil, true)
	// this.AssertEqual(err.GetCode(), 1)
	// this.AssertEqual(err.GetMessage(), "用户未登录")

	//错误登录
	err2 := this.UserLoginAo.Login_WithError("edward", "123dd")
	this.AssertError(err2, 1, "密码不正确")

	//正确登录
	this.UserLoginAo.Login("edward", "456")
	UserData := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData.Name, "edward")
	this.UserLoginAo.Logout()
	UserData2 := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData2, User{})

	//测试管理员登录
	this.UserLoginAo.Login("fish", "123")
	UserData3 := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData3.Name, "fish")
	this.UserLoginAo.CheckMustAdmin()

	//登出
	this.UserLoginAo.Logout()
	UserData4 := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData4, User{})

	// //reset用法
	this.UserLoginAo.Login_WithError("edward", "456")
	UserData5 := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData5.Name, "edward")

	this.RequestReset()

	_, err6 := this.UserLoginAo.CheckMustLogin_WithError()
	this.AssertError(err6, 1, "用户未登录")

	this.UserLoginAo.Login_WithError("edward", "456")
	UserData7 := this.UserLoginAo.IsLogin()
	this.AssertEqual(UserData7.Name, "edward")

}
