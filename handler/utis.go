package handler

import (
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}
	length := uint64(len(alphabet))
	var encodeBuilder strings.Builder
	encodeBuilder.Grow(10)

	for number > 0 {
		remainder := number % length
		encodeBuilder.WriteByte(alphabet[remainder])
		number = number / length

	}
	encoded := encodeBuilder.String()
	runes := []rune(encoded)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
