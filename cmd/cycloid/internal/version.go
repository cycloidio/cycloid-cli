package internal

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
)

func Error(out io.Writer, msg string) {
	switch viper.GetString("verbosity") {
	case "info", "debug", "warning":
		// This is still dirty, we should detect if the current
		// terminal is able to support colors
		// But that would be for another PR.
		fmt.Fprintf(out, "\033[1;31merror:\033[0m %s", msg)
	default:
	}
}

func Warning(out io.Writer, msg string) {
	switch viper.GetString("verbosity") {
	case "info", "debug", "warning":
		// This is still dirty, we should detect if the current
		// terminal is able to support colors
		// But that would be for another PR.
		fmt.Fprintf(out, "\033[1;35mwarning:\033[0m %s", msg)
	default:
	}
}

func Debug(msg ...any) {
	switch viper.GetString("verbosity") {
	case "debug":
		// This is still dirty, we should detect if the current
		// terminal is able to support colors
		// But that would be for another PR.
		fmt.Fprintf(os.Stderr, "\033[1;34mdebug:\033[0m ")
		fmt.Fprintln(os.Stderr, msg...)
	default:
	}
}
