package random

import (
	crand "crypto/rand"
	"unsafe"
)

const (
	rs6Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

func AlphaNumeric(n int) string {
	b := make([]byte, n)
	if _, err := crand.Read(b); err != nil {
		panic(err)
	}
	for i := 0; i < n; {
		idx := int(b[i] & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i++
		} else {
			if _, err := crand.Read(b[i : i+1]); err != nil {
				panic(err)
			}
		}
	}
	return *(*string)(unsafe.Pointer(&b))
}
