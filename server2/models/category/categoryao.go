package category

import (
	"fmt"
	. "github.com/fishedee/web"
	. "mymanager/models/common"
)

type CategoryAoModel struct {
	BaseModel
}

func (this *CategoryAoModel) Produce() {
	this.Queue.Produce("message", "Hello World")
}

func (this *CategoryAoModel) Consume(result string) {
	fmt.Println(result)
}

func (this *CategoryAoModel) Publish() {
	this.Queue.Publish("message2", "Hello World2")
}

func (this *CategoryAoModel) Subscribe(result string) {
	fmt.Println(result)
}

func init() {
	InitDaemon(func(this *CategoryAoModel) {
		this.Queue.Consume("message", this.Consume)
		this.Queue.Consume("message", this.Consume)
		this.Queue.Subscribe("message2", this.Subscribe)
		this.Queue.Subscribe("message2", this.Subscribe)
	})
}
