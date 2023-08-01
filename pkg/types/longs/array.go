package longs

import (
	"database/sql/driver"
	"devops-platform/pkg/types"
	"fmt"
)

type Array []types.Long

func (a *Array) Scan(src interface{}) (err error) {
	if value, ok := src.(string); ok {
		*a, err = StringToLongArray(value)
	} else if value, ok := src.([]byte); ok {
		*a, err = StringToLongArray(string(value))
	} else {
		err = fmt.Errorf("can not convert %v to longs.Array", value)
	}
	return
}
func (a Array) Value() (driver.Value, error) {
	return LongArrayToString(a), nil
}
