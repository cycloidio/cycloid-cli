package testcfg

import (
	"math/rand"
	"strings"
)

func RandomCanonical(baseName string) string {
	var size = 4
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return strings.ToLower(baseName) + "-" + string(b)
}
