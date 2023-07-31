package util

import "strings"

func EqualString(data1, data2 string) bool {
	return data1 == data2
}

func Equalint(data1, data2 int) bool {
	return data1 == data2
}
func StringLength(data string) int {
	return len(strings.TrimSpace(data))
}
