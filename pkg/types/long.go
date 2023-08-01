package types

import (
	"strconv"
)

type Long int64

func (l *Long) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	length := len(str)

	if str[0] == '"' && str[length-1] == '"' {
		str = str[1 : length-1]
	}

	v, err := strconv.ParseInt(str, 10, 64)
	*l = Long(v)
	return
}
func (l Long) MarshalJSON() ([]byte, error) {
	str := "\"" + l.String() + "\""
	return []byte(str), nil
}
func (l Long) String() string {
	return strconv.FormatInt(int64(l), 10)
}

func (l *Long) In(values []Long) bool {
	exists := false
	for _, v := range values {
		if v == *l {
			exists = true
		}
	}
	return exists
}
func StringToLong(value string) (Long, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return Long(result), nil
}
