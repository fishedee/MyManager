package brush

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
)

func (this *BrushAoModel) SearchTask_WithError(userId int, where BrushTask, limit CommonPage) (_fishgen1 BrushTasks, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.SearchTask(userId, where, limit)
	return
}

func (this *BrushAoModel) GetTask_WithError(userId int, brushTaskId int) (_fishgen1 BrushTask, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetTask(userId, brushTaskId)
	return
}

func (this *BrushAoModel) AddTask_WithError(userId int, data BrushTask) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddTask(userId, data)
	return
}

func (this *BrushAoModel) SearchCrawl_WithError(userId int, where BrushCrawl, limit CommonPage) (_fishgen1 BrushCrawls, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.SearchCrawl(userId, where, limit)
	return
}

func (this *BrushCrawlDbModel) Search_WithError(where BrushCrawl, limit CommonPage) (_fishgen1 BrushCrawls, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *BrushCrawlDbModel) GetByState_WithError(state int) (_fishgen1 []BrushCrawl, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetByState(state)
	return
}

func (this *BrushCrawlDbModel) Get_WithError(brushCrawlId int) (_fishgen1 BrushCrawl, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(brushCrawlId)
	return
}

func (this *BrushCrawlDbModel) Add_WithError(task BrushCrawl) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Add(task)
	return
}

func (this *BrushCrawlDbModel) Mod_WithError(brushCrawlId int, task BrushCrawl) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(brushCrawlId, task)
	return
}

func (this *BrushCrawlDbModel) IncrRetryNum_WithError(brushCrawlId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.IncrRetryNum(brushCrawlId)
	return
}

func (this *BrushProxyAoModel) GetMimvpProxy_WithError() (_fishgen1 string, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetMimvpProxy()
	return
}

func (this *BrushProxyAoModel) GetXiciProxy_WithError() (_fishgen1 string, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetXiciProxy()
	return
}

func (this *BrushTaskDbModel) Search_WithError(where BrushTask, limit CommonPage) (_fishgen1 BrushTasks, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Search(where, limit)
	return
}

func (this *BrushTaskDbModel) Get_WithError(brushTaskId int) (_fishgen1 BrushTask, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Get(brushTaskId)
	return
}

func (this *BrushTaskDbModel) GetByIds_WithError(brushTaskIds []int) (_fishgen1 []BrushTask, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetByIds(brushTaskIds)
	return
}

func (this *BrushTaskDbModel) Add_WithError(task BrushTask) (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.Add(task)
	return
}

func (this *BrushTaskDbModel) Mod_WithError(brushTaskId int, task BrushTask) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Mod(brushTaskId, task)
	return
}

func (this *BrushTaskDbModel) AddSuccessNum_WithError(brushTaskId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddSuccessNum(brushTaskId)
	return
}

func (this *BrushTaskDbModel) AddFailNum_WithError(brushTaskId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.AddFailNum(brushTaskId)
	return
}
