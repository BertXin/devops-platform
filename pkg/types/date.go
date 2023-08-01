package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormat+`"`, string(data), time.Local)
	*d = Date{now}
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormat)+2)
	b = append(b, '"')
	b = d.Time.AppendFormat(b, dateFormat)
	b = append(b, '"')
	return b, nil
}
func (d Date) String() string {
	return d.Time.Format(dateFormat)
}
func (d Date) Value() (driver.Value, error) {

	var zeroTime time.Time
	if d.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return d.Time, nil

}

func (d *Date) Scan(src interface{}) error {
	value, ok := src.(time.Time)
	if ok {
		*d = Date{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}
