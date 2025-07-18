package e2e

import (
	"log"
	"math/rand"
	"os"
)

var (
	// Note, this url should be accessible by Cycloid API
	CyTestCatalogRepoURL    = "git@github.com:cycloidio/cycloid-cli-test-catalog.git"
	CyTestCatalogRepoBranch = "master"
	CyTestConfigRepo        = "cycloid-template-catalog-config"
	TestAPIURL              = "https://api-cli-test.staging.cycloid.io/"
	TestRootOrg             = "cycloid"
	TestAPIKey              = ""
)

func init() {
	apiKey, ok := os.LookupEnv("CY_TEST_API_KEY")
	if !ok {
		log.Fatal("Missing API Key, set one with CY_TEST_API_KEY env var.")
	}
	TestAPIKey = apiKey

	org := os.Getenv("CY_TEST_ROOT_ORG")
	if len(org) > 0 {
		TestRootOrg = org
	}

	gitBranch := os.Getenv("CY_TEST_GIT_CR_BRANCH")
	if len(gitBranch) > 0 {
		CyTestCatalogRepoBranch = gitBranch
	}

	apiURL, ok := os.LookupEnv("CY_API_URL")
	if ok {
		TestAPIURL = apiURL
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
