package common

// PageResult 分页结果
type PageResult struct {
	// List 列表数据
	List interface{} `json:"list"`
	// Total 总数
	Total int64 `json:"total"`
	// Page 当前页码
	Page int `json:"page"`
	// Size 每页大小
	Size int `json:"size"`
}
