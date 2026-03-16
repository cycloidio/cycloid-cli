// Human-readable formatting for pipeline build watch SSE events (cy pipeline build trigger --watch).
// Rollback: docs/pipeline-build-watch-output.md (quick disable in build_trigger.go or remove this package).
package buildwatch

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type lineFormat int

const (
	lineFormatHuman lineFormat = iota
	lineFormatRaw
)

type consumeOpts struct {
	format       lineFormat
	verbose      bool
	theme        StreamTheme
	stripLogANSI bool
}

var ansiEscapeSeq = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSISequences(s string) string {
	return ansiEscapeSeq.ReplaceAllString(s, "")
}

func shortOriginID(id string) string {
	if i := strings.Index(id, "/"); i >= 0 {
		return id[:i]
	}
	return id
}

func writeString(w io.Writer, s string) error {
	if sw, ok := w.(io.StringWriter); ok {
		_, err := sw.WriteString(s)
		return err
	}
	_, err := w.Write([]byte(s))
	return err
}

func formatAndWriteBuildEvent(w io.Writer, payload string, opts consumeOpts) error {
	if opts.format == lineFormatRaw {
		if err := writeString(w, payload); err != nil {
			return err
		}
		if !strings.HasSuffix(payload, "\n") {
			return writeString(w, "\n")
		}
		return nil
	}

	out, ok := formatHumanBuildEvent(payload, opts)
	if !ok {
		return nil
	}
	return writeString(w, out)
}

func formatHumanBuildEvent(payload string, opts consumeOpts) (string, bool) {
	var env eventEnvelope
	if err := json.Unmarshal([]byte(payload), &env); err != nil {
		return payload + "\n", true
	}

	t := opts.theme

	switch env.Event {
	case "status":
		var d eventDataStatus
		if err := json.Unmarshal(env.Data, &d); err != nil {
			return "", false
		}
		msg := fmt.Sprintf("status: %s", d.Status)
		return paintBuildLine(msg, t), true

	case "log":
		var d eventDataLog
		if err := json.Unmarshal(env.Data, &d); err != nil {
			return "", false
		}
		text := d.Payload
		if opts.stripLogANSI {
			text = stripANSISequences(text)
		}
		out := paintLogLines(d.Origin.Source, text, t)
		if out == "" {
			return "", false
		}
		return out, true

	case "initialize-task":
		if !opts.verbose {
			return "", false
		}
		var d eventDataTask
		if err := json.Unmarshal(env.Data, &d); err != nil {
			return "", false
		}
		id := shortOriginID(d.Origin.ID)
		return paintTaskLine(fmt.Sprintf("%s: initialize", id), t), true

	case "start-task":
		var d eventDataTask
		if err := json.Unmarshal(env.Data, &d); err != nil {
			return "", false
		}
		id := shortOriginID(d.Origin.ID)
		msg := fmt.Sprintf("%s: started", id)
		if opts.verbose {
			msg = fmt.Sprintf("%s: start-task", id)
		}
		return paintTaskLine(msg, t), true

	case "finish-task":
		var d eventDataTask
		if err := json.Unmarshal(env.Data, &d); err != nil {
			return "", false
		}
		id := shortOriginID(d.Origin.ID)
		msg := fmt.Sprintf("%s: finished (exit %d)", id, d.ExitStatus)
		return paintTaskLine(msg, t), true

	case "initialize-check", "start", "image-check", "image-get", "initialize-get", "start-get", "selected-worker":
		if !opts.verbose {
			return "", false
		}
		return paintResourceLine(env.Event, compactJSONData(env.Data), t), true

	case "finish":
		var d eventDataFinish
		if err := json.Unmarshal(env.Data, &d); err != nil {
			return "", false
		}
		if !opts.verbose {
			if d.Succeeded == nil || *d.Succeeded {
				return "", false
			}
		}
		detail := shortOriginID(d.Origin.ID)
		if d.Name != "" {
			detail = fmt.Sprintf("%s check %q", detail, d.Name)
		}
		if d.Succeeded != nil {
			detail = fmt.Sprintf("%s succeeded=%v", detail, *d.Succeeded)
		}
		return paintResourceLine(env.Event, detail, t), true

	case "finish-get":
		var d eventDataFinishGet
		if err := json.Unmarshal(env.Data, &d); err != nil {
			return "", false
		}
		if !opts.verbose && d.ExitStatus == 0 {
			return "", false
		}
		id := shortOriginID(d.Origin.ID)
		return paintResourceLine(env.Event, fmt.Sprintf("%s exit=%d", id, d.ExitStatus), t), true

	default:
		if !opts.verbose {
			return "", false
		}
		return paintResourceLine(env.Event, compactJSONData(env.Data), t), true
	}
}

func compactJSONData(raw json.RawMessage) string {
	s := strings.TrimSpace(string(raw))
	if len(s) > 120 {
		return s[:117] + "..."
	}
	return s
}

func paintBuildLine(msg string, t StreamTheme) string {
	if t.Reset == "" {
		return fmt.Sprintf("[build] %s\n", msg)
	}
	return fmt.Sprintf("%s[build]%s %s%s%s\n", t.BuildLabel, t.Reset, t.BuildText, msg, t.Reset)
}

func paintTaskLine(msg string, t StreamTheme) string {
	if t.Reset == "" {
		return fmt.Sprintf("[task] %s\n", msg)
	}
	return fmt.Sprintf("%s[task]%s %s%s%s\n", t.TaskLabel, t.Reset, t.TaskText, msg, t.Reset)
}

func paintResourceLine(event, detail string, t StreamTheme) string {
	if t.Reset == "" {
		return fmt.Sprintf("[res] %s %s\n", event, detail)
	}
	line := fmt.Sprintf("%s[res]%s %s%s:%s %s%s\n",
		t.ResourceLabel, t.Reset, t.Dim, event, t.Reset, detail, t.Reset)
	return line
}

func paintLogLines(source, payload string, t StreamTheme) string {
	label := "stdout"
	prefixColor := t.LogStdoutPrefix
	if source == "stderr" {
		label = "stderr"
		prefixColor = t.LogStderrPrefix
	}

	if strings.TrimSpace(payload) == "" {
		return ""
	}

	lines := strings.Split(strings.TrimRight(payload, "\r\n"), "\n")
	var b strings.Builder
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if t.Reset == "" {
			b.WriteString(fmt.Sprintf("[%s] %s\n", label, line))
			continue
		}
		b.WriteString(fmt.Sprintf("%s[%s]%s %s\n", prefixColor, label, t.Reset, line))
	}
	if b.Len() == 0 {
		return ""
	}
	return b.String()
}
