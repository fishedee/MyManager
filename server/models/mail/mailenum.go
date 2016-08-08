package mail

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
)

var MailQueueEnum struct {
	EnumStructString
	TASK_SEND string `enum:"/mail/_send,邮件发送"`
}

func init() {
	InitEnumStructString(&MailQueueEnum)
	InitDaemon(func(this *MailAoModel) {
		this.Queue.Consume(MailQueueEnum.TASK_SEND, (*MailAoModel).sendInner)
	})
}
