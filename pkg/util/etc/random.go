package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ランダムな文字列を返す
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func Random(n int) int {
	return rand.Intn(n) + 1
}

// ランダムなEmailを返す
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(10))
}

// 先頭から5文字削除する
func Remove5Strings(strings string) string {
	slicedString := []rune(strings)
	return string(slicedString[5:])
}

// ランダム量を返す
func RandomSelection(n, length int) []int32 {
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
