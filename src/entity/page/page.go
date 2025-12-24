package page

type QueryPage struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func Defaults() QueryPage {
	return QueryPage{PageNum: 1, PageSize: 10}
}

func (p QueryPage) SetPageNum(pageNum int) {
	p.PageNum = pageNum
}

func (p QueryPage) SetPageSize(pageSize int) {
	p.PageSize = pageSize
}
