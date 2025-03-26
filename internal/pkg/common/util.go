package common

import (
	"time"
)

// GetCurrentMillis 获取当前时间毫秒时间戳
func GetCurrentMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetCurrentSeconds 获取当前时间秒时间戳
func GetCurrentSeconds() int64 {
	return time.Now().Unix()
}

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ParseTime 解析时间字符串
func ParseTime(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
}

// FormatDate 格式化日期
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// ParseDate 解析日期字符串
func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", s, time.Local)
}

// GetStartOfDay 获取一天的开始时间
func GetStartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// GetEndOfDay 获取一天的结束时间
func GetEndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999000000, t.Location())
}

// GetStartOfMonth 获取一个月的开始时间
func GetStartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// GetEndOfMonth 获取一个月的结束时间
func GetEndOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	nextMonth := month + 1
	nextMonthYear := year
	if nextMonth > 12 {
		nextMonth = 1
		nextMonthYear++
	}
	nextMonthFirstDay := time.Date(nextMonthYear, nextMonth, 1, 0, 0, 0, 0, t.Location())
	return nextMonthFirstDay.Add(-time.Nanosecond)
}
