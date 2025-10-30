/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>
*/
package utils

import (
	"math/rand/v2"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func GenerateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
