package file

import (
	. "github.com/fishedee/util"
	. "server2/models/common"
	"testing"
)

type UploadAoTest struct {
	BaseTest
	UploadAo UploadAoModel
}

func (this *UploadAoTest) TestFile() {
	testCase := []string{
		"",
		"helloworld",
		"你好 dd",
	}
	for _, singleTestCase := range testCase {
		url := this.UploadAo.UploadFile([]byte(singleTestCase))
		var data string
		err := DefaultAjaxPool.Get(&Ajax{
			Url:          url,
			ResponseData: &data,
		})
		this.AssertEqual(err == nil, true, singleTestCase)
		this.AssertEqual(data, singleTestCase, singleTestCase)
	}
}

func TestUpload(t *testing.T) {
	InitTest(t, &UploadAoTest{})
}
