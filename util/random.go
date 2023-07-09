package util

import (
	"fmt"
	"math/rand"
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
