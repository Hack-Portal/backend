package utils

func ContainsArr(value, target []int) bool {
	for _, t := range target {
		for _, v := range value {
			if t == v {
				return true
			}
		}
	}
	return false
}
