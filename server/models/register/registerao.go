package register

import (
	. "github.com/fishedee/util"
	. "github.com/fishedee/sdk"
	. "github.com/fishedee/web"
	. "github.com/fishedee/encoding"
	"strings"
	"errors"
	"strconv"
	"time"
)

type RegisterAoModel struct {
	Model
}

func (this *RegisterAoModel) getDoctorRegister(appId string,openId string,branchCode string,deptCode string,date string) ([]RegisterResult,error) {
	var byteResult []byte
	err := DefaultAjaxPool.Get(&Ajax{
		Url:"http://cs7.yx129.net/registrationfy/loadResInfo",
		Data:map[string]string{
			"appId":appId,
			"openId":openId,
			"branchCode":branchCode,
			"deptCode":deptCode,
			"beginDate":date,
			"endDate":date,
		},
		ResponseData:&byteResult,
	})
	if err != nil{
		return nil,err
	}

	var errorResult struct{
		Content string
		Type string
	}
	err = DecodeJson(byteResult,&errorResult)
	if err == nil && errorResult.Type == "error"{
		return nil,errors.New(errorResult.Content)
	}

	var doctorResult struct{
		Doctorls []RegisterResult
	}
	var successResult struct{
		Content interface{}
		Type string
	}
	successResult.Content = &doctorResult
	err = DecodeJson(byteResult,&successResult)
	if err != nil {
		return nil,err
	}
	if successResult.Type != "success"{
		return nil,errors.New("奇怪的返回值 "+string(byteResult))
	}
	return doctorResult.Doctorls,nil
}

func (this *RegisterAoModel) checkValidDoctor(data []RegisterResult)[]RegisterResult{
	result := []RegisterResult{}

	for _,singleData := range data{
		if strings.Index(singleData.DoctorName,"助产门诊") != -1 {
			continue
		}
		if singleData.LeftCount <= 0{
			continue
		}
		result = append(result,singleData)
	}
	return result
}

func (this *RegisterAoModel) notify(date string,data []RegisterResult){
	if len(data) == 0{
		return
	}

	notifyResult := "恭喜你，目前佛山市第一人民医院以下医生有号：<br/>"
	for _,singleData := range data{
		notifyResult += "日期："+date+"，科室："+singleData.DeptName+"，医生："+singleData.DoctorName+"，剩余数量："+strconv.Itoa(singleData.LeftCount)+"<br/>"
	}
	
	smtp := &SmtpSdk{
		Host:"smtp.163.com:25",
	}
	err := smtp.Send(SmtpSdkMailAuth{
		UserName:"15018749403@163.com",
		Password:"9616966",
	},SmtpSdkMail{
		From:"15018749403@163.com",
		To:[]string{"306766045@qq.com"},
		Subject:"佛山市一挂号",
		Body:notifyResult,
	})
	if err != nil{
		panic(err)
	}
}

func (this *RegisterAoModel) checkDoctor(){
	now := time.Now()

	for i := 0 ; i != 14 ; i++{
		date := now.AddDate(0,0,i).Format("2006-01-02")
		data,err := this.getDoctorRegister(
			"2015070900161782",
			"20880005665009381053113102715391",
			"1",
			"3100",
			date,
		)
		if err != nil{
			if err.Error() != "查询his号源信息出错"{
				this.Log.Error("查询挂号失败%v",err)
			}
			continue
		}
		this.Log.Debug("查询挂号时间%v,信息为：%v",date,data)

		data = this.checkValidDoctor(data)

		this.notify(date,data)
	}
}

func (this *RegisterAoModel) addDoctor(){

}

func init(){
	/*
	InitDaemon(func(this *RegisterAoModel){
		this.Timer.Cron("30 * * * *",(*RegisterAoModel).checkDoctor)
	})
	*/
}
