package utils

import (
	"strconv"
	"strings"
)

func StrToIntArr(v string) ([]int32, error) {
	parts := strings.Split(v, ",")
	arr := make([]int32, 0, len(parts))
	for _, val := range parts {
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		arr = append(arr, int32(v))
	}
	return arr, nil
}
