package e2e

import "os"

var (
	CY_TEST_EMAIL    = "test@org.local.tld"
	CY_TEST_PASSWORD = "p4ssw0rd"
	CY_TEST_ORG      = "e2eOrg"
)

func init() {
	email := os.Getenv("CY_TEST_EMAIL")
	if len(email) > 0 {
		CY_TEST_EMAIL = email
	}

	password := os.Getenv("CY_TEST_PASSWORD")
	if len(password) > 0 {
		CY_TEST_PASSWORD = password
	}

	org := os.Getenv("CY_TEST_ORG")
	if len(org) > 0 {
		CY_TEST_ORG = org
	}
}
