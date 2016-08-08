package mail

import (
	. "github.com/fishedee/language"
)

func (this *MailAoModel) Send_WithError(to []string, subject string, body string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Send(to, subject, body)
	return
}
