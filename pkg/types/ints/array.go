package ints

import (
	"database/sql/driver"
	"fmt"
)

type Array []int

func (a *Array) Scan(src interface{}) (err error) {
	if value, ok := src.(string); ok {
		*a, err = StringToIntArray(value)
	} else if value, ok := src.([]byte); ok {
		*a, err = StringToIntArray(string(value))
	} else {
		err = fmt.Errorf("can not convert %v to ints.Array", value)
	}
	return
}

func (a Array) Value() (driver.Value, error) {
	return IntArrayToString(a), nil
}