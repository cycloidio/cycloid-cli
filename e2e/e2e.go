package e2e

import (
	"log"
	"math/rand"
	"os"
)

var (
	// Note, this url should be accessible by Cycloid API
	CY_TEST_GIT_CR_URL    = "git@github.com:cycloidio/cycloid-cli-test-catalog.git"
	CY_TEST_GIT_CR_BRANCH = "master"
	CY_API_URL            = "https://api-cli-test.staging.cycloid.io/"
	CY_TEST_ROOT_ORG      = "cycloid"
	CY_TEST_API_KEY       = ""
)

func init() {
	apiKey, ok := os.LookupEnv("CY_TEST_API_KEY")
	if !ok {
		log.Fatal("Missing API Key, set one with CY_TEST_API_KEY env var.")
	}
	CY_TEST_API_KEY = apiKey

	org := os.Getenv("CY_TEST_ROOT_ORG")
	if len(org) > 0 {
		CY_TEST_ROOT_ORG = org
	}

	gitBranch := os.Getenv("CY_TEST_GIT_CR_BRANCH")
	if len(gitBranch) > 0 {
		CY_TEST_GIT_CR_BRANCH = gitBranch
	}

	apiUrl, ok := os.LookupEnv("CY_API_URL")
	if ok {
		CY_API_URL = apiUrl
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
