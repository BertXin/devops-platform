package enum

import (
	"strconv"
)

// Status 状态枚举
type Status int

const (
	// StatusDisabled 禁用
	StatusDisabled Status = 0
	// StatusEnabled 启用
	StatusEnabled Status = 1
)

// ValidStatus 验证状态是否有效
func (s Status) ValidStatus() bool {
	return s == StatusDisabled || s == StatusEnabled
}

// String 返回状态的字符串表示
func (s Status) String() string {
	switch s {
	case StatusDisabled:
		return "禁用"
	case StatusEnabled:
		return "启用"
	default:
		return "未知状态"
	}
}

// MarshalJSON 自定义JSON序列化
func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(`{"code":` + strconv.Itoa(int(s)) + `,"desc":"` + s.String() + `"}`), nil
}
