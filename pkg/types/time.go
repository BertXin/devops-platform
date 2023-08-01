package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {

	str := string(data)
	if str == "null" {
		*t = Time{*new(time.Time)}
		return nil
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, str, time.Local)
	*t = Time{now}
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}
func (t Time) String() string {
	return t.Time.Format(TimeFormat)
}
func (t Time) Value() (driver.Value, error) {

	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil

}

func (t *Time) Scan(src interface{}) error {
	value, ok := src.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}
func (t *Time) Interval(u *Time) string {
	duration := time.Nanosecond
	if u != nil {
		duration = t.Time.Sub(u.Time)
	}

	if duration < 0 {
		duration *= -1
	}
	duration /= time.Second
	duration *= time.Second

	return duration.String()
}
