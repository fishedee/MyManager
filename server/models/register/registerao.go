package register

import (
	"errors"
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
	. "mymanager/models/mail"
	"strconv"
	"strings"
	"time"
)

type RegisterAoModel struct {
	Model
	RegisterDb RegisterDbModel
	MailAo     MailAoModel
}

func (this *RegisterAoModel) Search(userId int, where Register, limit CommonPage) Registers {
	where.UserId = userId
	return this.RegisterDb.Search(where, limit)
}

func (this *RegisterAoModel) Get(userId int, registerId int) Register {
	registerInfo := this.RegisterDb.Get(registerId)
	if registerInfo.UserId != userId {
		Throw(1, "你没有权利查看或编辑等操作")
	}
	return registerInfo
}

func (this *RegisterAoModel) Del(userId int, registerId int) {
	this.Get(userId, registerId)

	this.RegisterDb.Del(registerId)
}

func (this *RegisterAoModel) Add(userId int, register Register) {
	register.UserId = userId
	register.HaveDealType = RegisterHaveDealType.NO
	register.HaveDealResult = ""
	this.RegisterDb.Add(register)
}

func (this *RegisterAoModel) Mod(userId int, registerId int, registerInfo Register) {
	this.Get(userId, registerId)

	registerInfo.UserId = userId
	registerInfo.HaveDealType = 0
	registerInfo.HaveDealResult = ""
	this.RegisterDb.Mod(registerId, registerInfo)
}

func (this *RegisterAoModel) getAllNeedRegister() []Register {
	now := time.Now()
	result := this.RegisterDb.GetAll()
	newResult := []Register{}
	for _, singleResult := range result {
		if singleResult.EndTime.Before(now) {
			continue
		}
		newResult = append(newResult, singleResult)
	}
	return newResult
}

func (this *RegisterAoModel) getNeedRegisterByDate(date string, data []Register) []Register {
	result := []Register{}
	for _, singleData := range data {
		beginDate := singleData.BeginTime.Format("2006-01-02")
		endDate := singleData.EndTime.Format("2006-01-02")
		if date >= beginDate && date <= endDate {
			result = append(result, singleData)
		}
	}
	return result
}

func (this *RegisterAoModel) getDoctorRegister(appId string, openId string, branchCode string, deptCode string, date string) ([]RegisterResult, error) {
	var byteResult []byte
	err := DefaultAjaxPool.Get(&Ajax{
		Url: "http://cs7.yx129.net/registrationfy/loadResInfo",
		Data: map[string]string{
			"appId":      appId,
			"openId":     openId,
			"branchCode": branchCode,
			"deptCode":   deptCode,
			"beginDate":  date,
			"endDate":    date,
		},
		ResponseData: &byteResult,
	})
	if err != nil {
		return nil, err
	}

	var errorResult struct {
		Content string
		Type    string
	}
	err = DecodeJson(byteResult, &errorResult)
	if err == nil && errorResult.Type == "error" {
		return nil, errors.New(errorResult.Content)
	}

	var doctorResult struct {
		Doctorls []RegisterResult
	}
	var successResult struct {
		Content interface{}
		Type    string
	}
	successResult.Content = &doctorResult
	err = DecodeJson(byteResult, &successResult)
	if err != nil {
		return nil, err
	}
	if successResult.Type != "success" {
		return nil, errors.New("奇怪的返回值 " + string(byteResult))
	}
	return doctorResult.Doctorls, nil
}

func (this *RegisterAoModel) getDoctorRegisterByDate(date string) ([]RegisterResult, error) {
	result, err := this.getDoctorRegister(
		"2015070900161782",
		"20880005665009381053113102715391",
		"1",
		"0760", //3100是产科，0760是神经内科
		date,
	)
	if err != nil {
		if err.Error() != "查询his号源信息出错" {
			this.Log.Error("查询挂号失败%v", err)
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (this *RegisterAoModel) getMatchDoctor(date string, doctor []RegisterResult, need Register) []RegisterMatch {
	result := []RegisterMatch{}
	for _, singleDoctor := range doctor {
		if singleDoctor.LeftCount == 0 {
			continue
		}
		if strings.Index(singleDoctor.DoctorName, "助产门诊") != -1 {
			continue
		}
		for _, singleTime := range singleDoctor.ListTi {
			if singleTime.LeftCount == 0 {
				continue
			}
			doctorBeginTime, err := time.ParseInLocation("2006-01-02 15:04:05", date+" "+singleTime.BeginTime, time.Local)
			if err != nil {
				panic(err)
			}
			doctorBeginTime = doctorBeginTime.Truncate(time.Second)
			doctorEndTime, err := time.ParseInLocation("2006-01-02 15:04:05", date+" "+singleTime.EndTime, time.Local)
			if err != nil {
				panic(err)
			}
			doctorEndTime = doctorEndTime.Truncate(time.Second)

			if need.BeginTime.After(doctorBeginTime) ||
				need.EndTime.Before(doctorEndTime) {
				continue
			}
			result = append(result, RegisterMatch{
				Need:   need,
				Doctor: singleDoctor,
				Time:   singleTime,
			})
		}
	}
	return result
}

func (this *RegisterAoModel) notify(date string, need Register, match []RegisterMatch) {
	notifyResult := "亲爱的" + need.Name + "妹子，佛山市一产科有号：<br/>"
	notifyResult += "你预定的时间为：" + need.BeginTime.Format("2006-01-02 15:04:05") + " 至 " + need.EndTime.Format("2006-01-02 15:04:05") + "<br/>"
	notifyResult += "可用的科室有：<br/>"
	for _, singleMatch := range match {
		notifyResult += "日期：" + date + " " + singleMatch.Time.BeginTime + " - " + singleMatch.Time.EndTime + "，科室：" + singleMatch.Doctor.DeptName + "，医生：" + singleMatch.Doctor.DoctorName + "，剩余数量：" + strconv.Itoa(singleMatch.Time.LeftCount) + "<br/>"
	}

	notifyTitle := "恭喜你，佛山市一产科有号！"
	this.MailAo.Send([]string{need.Mail}, notifyTitle, notifyResult)
}

func (this *RegisterAoModel) checkDoctor() {
	now := time.Now()
	need := this.getAllNeedRegister()
	if len(need) == 0 {
		return
	}

	for i := 0; i != 14; i++ {
		date := now.AddDate(0, 0, i).Format("2006-01-02")
		dateNeed := this.getNeedRegisterByDate(date, need)
		if len(dateNeed) == 0 {
			continue
		}
		data, err := this.getDoctorRegisterByDate(date)
		if err != nil {
			panic(err)
		}

		for _, singleDateNeed := range dateNeed {
			matchNeed := this.getMatchDoctor(date, data, singleDateNeed)
			this.Log.Debug("匹配的医生有:%v", matchNeed)
			if len(matchNeed) != 0 {
				this.notify(date, singleDateNeed, matchNeed)
			}
			if singleDateNeed.NeedDealType == RegisterNeedDealType.YES &&
				singleDateNeed.HaveDealType == RegisterNeedDealType.NO {
				//下单
			}
		}

	}
}

func (this *RegisterAoModel) dealDoctor(date time.Time, need Register) {
	//检查是否需要继续下单

	//检查是否有该名字

	//while{
	//获取当天的挂号记录

	//匹配然后下单
	//}
}

func init() {
	InitDaemon(func(this *RegisterAoModel) {
		this.Timer.Cron("*/30 * * * *", (*RegisterAoModel).checkDoctor)
	})
}
