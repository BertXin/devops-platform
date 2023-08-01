package ints

import (
	"strconv"
	"strings"
)

func StringToIntArray(value string) ([]int, error) {
	if value == "" {
		return []int{}, nil
	}
	valueStringArray := strings.Split(value, ",")
	result := make([]int, len(valueStringArray))
	for i, v := range valueStringArray {
		r, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

func IntArrayToString(value []int) string {

	if value == nil {
		return ""
	}
	length := len(value)
	if length == 0 {
		return ""
	}
	valueStringArray := make([]string, length)
	for i, v := range value {
		valueStringArray[i] = strconv.Itoa(v)
	}
	return strings.Join(valueStringArray, ",")

}
