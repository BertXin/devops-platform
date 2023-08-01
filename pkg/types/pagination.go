package types

import (
	"strconv"
)


type Pagination struct {
	PageSize int `json:"page_size" form:"page_size"`
	Page     int `json:"page" form:"page"`
}

func (p *Pagination) Limit() int {
	if p.PageSize <= 0 {
		return 10
	}
	return p.PageSize
}

func (p *Pagination) Offset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit()
}

func (p *Pagination) PaginationCondition() string {
	return " limit " + strconv.Itoa(p.Offset()) + "," + strconv.Itoa(p.Limit())
}
