// ANSI theme for human watch output only. Not related to cyargs project --color.
// See docs/pipeline-build-watch-output.md to disable or remove the whole watch formatting layer.
package buildwatch

// StreamTheme groups ANSI SGR sequences for --watch human output.
// Change DefaultStreamTheme to adjust the look; keep Reset for correct TTY behavior.
type StreamTheme struct {
	Reset string

	BuildLabel string
	BuildText  string

	TaskLabel string
	TaskText  string

	LogStdoutPrefix string
	LogStderrPrefix string

	ResourceLabel string

	Dim string
}

// DefaultStreamTheme is used when stdout is a TTY and output mode is human.
var DefaultStreamTheme = StreamTheme{
	Reset:           "\033[0m",
	BuildLabel:      "\033[1;35m",
	BuildText:       "\033[35m",
	TaskLabel:       "\033[1;36m",
	TaskText:        "\033[36m",
	LogStdoutPrefix: "\033[32m",
	LogStderrPrefix: "\033[33m",
	ResourceLabel:   "\033[90m",
	Dim:             "\033[2m",
}
