package file

import (
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
)

type UploadAoModel interface {
	UploadFile(data []byte) (_fishgen1 string)
	UploadFile_WithError(data []byte) (_fishgen1 string, _fishgenErr Exception)
	UploadFileFromLocal(fileAddress string) (_fishgen1 string)
	UploadFileFromLocal_WithError(fileAddress string) (_fishgen1 string, _fishgenErr Exception)
}

type UploadAoTest interface {
	TestFile()
}

func (this *uploadAoModel) UploadFile_WithError(data []byte) (_fishgen1 string, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.UploadFile(data)
	return
}

func (this *uploadAoModel) UploadFileFromLocal_WithError(fileAddress string) (_fishgen1 string, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.UploadFileFromLocal(fileAddress)
	return
}

func init() {
	v0 := UploadAoModel(&uploadAoModel{})
	InitModel(&v0)
	v1 := UploadAoTest(&uploadAoTest{})
	InitTest(&v1)
}
