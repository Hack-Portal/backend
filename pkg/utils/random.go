package utils

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomInt(n int) int {
	return rand.Intn(n)
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[RandomInt(len(letters))]
	}
	return string(b)
}

func RandomIntArr(n, length int) []int32 {
	array := make([]int, length)
	// ランダムな値を生成
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(n) + 1
	}

	// 重複を除去
	result := make([]int32, 0)
	for _, value := range array {
		if !contains(result, int32(value)) {
			result = append(result, int32(value))
		}
	}

	return result
}

// 線形比較
func contains(array []int32, variable int32) bool {
	for _, value := range array {
		if value == variable {
			return true
		}
	}
	return false
}
