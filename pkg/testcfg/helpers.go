package testcfg

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
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

func EnvDefault(envVar, defaultValue string) string {
	if value, ok := os.LookupEnv(envVar); ok {
		return value
	}

	return defaultValue
}

func FindRepoRoot() (string, error) {
	path := ".git"
	for i := 0; true; i++ {
		if pwd, err := os.Getwd(); pwd == "/" || err != nil {
			return "", fmt.Errorf("failed to find project root we reached /")
		}

		testPath := strings.Repeat("../", i) + path
		_, err := os.Stat(testPath)
		if err == nil {
			return filepath.Join(filepath.Dir(testPath)), nil
		}
	}

	return "", fmt.Errorf("failed for some reason")
}
