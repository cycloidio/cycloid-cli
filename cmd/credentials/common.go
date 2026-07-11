package credentials

import (
	"regexp"
)

func pathFromCanonical(canonical string) string {
	re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
	safePath := re.ReplaceAllString(canonical, "-")
	return safePath
}
