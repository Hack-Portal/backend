package utils

import (
	"strconv"
	"strings"
)

func StringToArrayInt32(base string) (result []int32, err error) {
	base = strings.Replace(base, "[", "", -1)
	base = strings.Replace(base, "]", "", -1)
	bases := strings.Split(base, ",")
	for _, b := range bases {
		r, err := strconv.Atoi(b)
		if err != nil {
			return result, err
		}
		result = append(result, int32(r))
	}
	return
}

func StringToArray(base string) []string {
	return strings.Split(base, ",")
}

func CheckDiff(b, a string) bool {
	if len(a) != 0 {
		if b != a {
			return true
		}
	}
	return false
}
