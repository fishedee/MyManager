package file

import (
	. "github.com/fishedee/language"
)

func (this *UploadAoModel) UploadFile_WithError(data []byte) (_fishgen1 string, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.UploadFile(data)
	return
}

func (this *UploadAoModel) UploadFileFromLocal_WithError(fileAddress string) (_fishgen1 string, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.UploadFileFromLocal(fileAddress)
	return
}
