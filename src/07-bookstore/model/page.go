package model

//Page 结构
type Page struct {
	Book        []*Book
	PageNo      int64
	PageSize    int64
	TotalPageNo int64
	TotalRecord int64
}
