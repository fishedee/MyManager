package mail

import (
	. "github.com/fishedee/sdk"
	. "github.com/fishedee/web"
)

type MailAoModel struct {
	Model
}

func (this *MailAoModel) Send(to []string, subject string, body string) {
	this.Queue.Produce(MailQueueEnum.TASK_SEND, to, subject, body)
}

func (this *MailAoModel) sendInner(to []string, subject string, body string) {
	smtp := &SmtpSdk{
		Host: "smtp.163.com:25",
	}
	err := smtp.Send(SmtpSdkMailAuth{
		UserName: "15018749403@163.com",
		Password: "9616966",
	}, SmtpSdkMail{
		From:    "15018749403@163.com",
		To:      to,
		Subject: subject,
		Body:    body,
	})
	if err != nil {
		panic(err)
	}
}
