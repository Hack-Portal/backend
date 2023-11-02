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
