/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>
*/
package docker

import (
	"math/rand/v2"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func generateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
