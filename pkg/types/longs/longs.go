package longs

import (
	"devops-platform/pkg/types"
	"strconv"
	"strings"
)

func StringToLong(value string) (types.Long, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return types.Long(result), nil
}

func StringToLongArray(value string) ([]types.Long, error) {
	if value == "" {
		return make([]types.Long, 0), nil
	}
	valueStringArray := strings.Split(value, ",")
	result := make([]types.Long, len(valueStringArray))
	for i, v := range valueStringArray {
		l, err := StringToLong(v)
		if err != nil {
			return nil, err
		}
		result[i] = l
	}
	return result, nil
}

func LongArrayToString(value []types.Long) string {
	if value == nil {
		return ""
	}
	length := len(value)
	if length == 0 {
		return ""
	}
	valueStringArray := make([]string, length)
	for i, v := range value {
		valueStringArray[i] = v.String()
	}
	return strings.Join(valueStringArray, ",")

}

func LongArrayToMap(value []types.Long) (valueMap map[types.Long]bool) {

	if value == nil {
		return
	}
	length := len(value)
	if length == 0 {
		return
	}
	valueMap = make(map[types.Long]bool)
	for _, v := range value {
		valueMap[v] = true
	}
	return

}
func MapToLongArray(valueMap map[types.Long]bool) (value []types.Long) {

	if valueMap == nil {
		return
	}
	length := len(valueMap)
	if length == 0 {
		return
	}
	value = make([]types.Long, length)
	for k, v := range valueMap {
		if v {
			value = append(value, k)
		}
	}
	return

}

func AppendNotExists(values []types.Long, value types.Long) []types.Long {

	length := len(values)
	if length == 0 {
		return []types.Long{value}
	}

	if !value.In(values) {
		values = append(values, value)
	}
	return values
}

func AppendArrayNotExists(values []types.Long, value []types.Long) []types.Long {

	length := len(values)
	if length == 0 {
		return value
	}
	for _, v := range value {
		values = AppendNotExists(values, v)
	}

	return values
}
