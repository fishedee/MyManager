package common

type CommonPage struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

var (
	CommonAllPage = CommonPage{-1, -1}
)
