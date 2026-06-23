package version

import (
	"runtime"
)

// Build information. Populated at build-time.
var (
	Version     string
	Revision    string
	Branch      string
	BuildOrigin string
	BuildDate   string
	GoVersion   = runtime.Version()
)
