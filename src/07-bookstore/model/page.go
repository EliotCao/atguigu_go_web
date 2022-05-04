package model

//Page 结构
type Page struct {
	Book        []*Book
	PageNo      int64
	PageSize    int64
	TotalPageNo int64
	TotalRecord int64
	MinPrice    int64
	MaxPrice    int64
	IsLogin     bool
	Username    string
}

func (p Page) HasPrev() bool {
	return p.PageNo > 1
}

func (p Page) HasNext() bool {
	return p.PageNo < p.TotalPageNo
}

func (p Page) GetPrevPageNo() int64 {
	if p.HasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}
