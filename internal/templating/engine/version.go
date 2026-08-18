package engine

import "strconv"

// Version is the Cycloid service-catalog interpolation version. It is a local,
// dependency-free copy of youdeploy-http-api/services/youdeploy/svccat/version
// Version, carrying only what the interpolator needs. On the CLI→backend merge
// this file is deleted and the backend type imported directly (see VENDORED.md).
type Version uint8

const (
	V1 Version = iota + 1 // 1
	V2                    // 2
	V3                    // 3
	V4                    // 4

	// Latest is the default used by the offline templating engine.
	Latest = V4
	// LatestNotDeprecated is the oldest non-deprecated version.
	LatestNotDeprecated = V2
)

// IsNewInterpolation reports whether the version uses the Go text/template
// engine (V3+) rather than the deprecated regex string-replace path.
func (v Version) IsNewInterpolation() bool { return v > V2 }

// IsDeprecatedVersion reports whether the version predates LatestNotDeprecated.
func (v Version) IsDeprecatedVersion() bool { return v < LatestNotDeprecated }

// IsAVersion reports whether v is a known version value.
func (v Version) IsAVersion() bool { return v >= V1 && v <= V4 }

// String renders the version as its numeric form, matching the backend's
// snake/linecomment enumer output ("1".."4").
func (v Version) String() string { return strconv.Itoa(int(v)) }
