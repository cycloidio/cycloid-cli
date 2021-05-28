package e2e

import "os"

var (
	CY_TEST_ROOT_API_KEY = "my secret api key"
	CY_TEST_ROOT_ORG     = "cycloidio"
	// Note, this url should be accessible by Cycloid API
	CY_TEST_GIT_CR_URL    = "Url of the git repository used as config repository"
	CY_TEST_GIT_CR_BRANCH = "master"
)

func init() {
	apiKey := os.Getenv("CY_TEST_ROOT_API_KEY")
	if len(apiKey) > 0 {
		CY_TEST_ROOT_API_KEY = apiKey
	}

	org := os.Getenv("CY_TEST_ROOT_ORG")
	if len(org) > 0 {
		CY_TEST_ROOT_ORG = org
	}

	gitUrl := os.Getenv("CY_TEST_GIT_CR_URL")
	if len(gitUrl) > 0 {
		CY_TEST_GIT_CR_URL = gitUrl
	}

	gitBranch := os.Getenv("CY_TEST_GIT_CR_BRANCH")
	if len(gitBranch) > 0 {
		CY_TEST_GIT_CR_BRANCH = gitBranch
	}
}
