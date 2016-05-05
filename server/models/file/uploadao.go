package file

import (
	. "github.com/fishedee/sdk"
	. "github.com/fishedee/web"
	"io/ioutil"
)

const (
	qiniuAccessKey = "dLlfL4vQDb0iD7g7ppj1jQe0xNKnYVle4uf2m7CR"
	qiniuSecretKey = "3Rc-ryiGWBKeeXnU1uGspPBxJp3iAuXoLCk11y3y"
	bucketName     = "blog"
	host           = "http://image.fishedee.com/"
)

type UploadAoModel struct {
	Model
}

func (this *UploadAoModel) UploadFile(data []byte) string {
	qiniuSdk := QiniuSdk{
		AccessKey: qiniuAccessKey,
		SecretKey: qiniuSecretKey,
	}
	fileNameHash, err := qiniuSdk.UploadString(bucketName, data)
	if err != nil {
		panic(err)
	}
	return host + fileNameHash
}

func (this *UploadAoModel) UploadFileFromLocal(fileAddress string) string {
	data, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		panic(err)
	}
	return this.UploadFile(data)
}
