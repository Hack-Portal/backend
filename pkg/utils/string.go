package utils

import (
	"strconv"
	"strings"
)

func StrToIntArr(v string) ([]int, error) {
	parts := strings.Split(v, ",")
	arr := make([]int, 0, len(parts))
	for _, val := range parts {
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		arr = append(arr, v)
	}
	return arr, nil
}
