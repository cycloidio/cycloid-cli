package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// requestLogger formats and writes HTTP debug logs to stderr.
// Update the Log* methods here to change the output format for all requests.
type requestLogger struct{}

// logRequest logs an outgoing HTTP request to stderr.
func (l *requestLogger) logRequest(req *http.Request, body []byte) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "\n[DEBUG] → %s %s\n", req.Method, req.URL.String())
	sb.WriteString("Request Headers:\n")
	for k, vs := range req.Header {
		v := strings.Join(vs, ", ")
		if strings.EqualFold(k, "Authorization") {
			v = redactAuth(v)
		}
		fmt.Fprintf(&sb, "  %s: %s\n", k, v)
	}
	if len(body) > 0 {
		sb.WriteString("Request Body:\n")
		sb.WriteString(prettyJSON(body))
		sb.WriteString("\n")
	}
	fmt.Fprint(os.Stderr, sb.String())
}

// logResponse logs an incoming HTTP response to stderr.
func (l *requestLogger) logResponse(resp *http.Response, body []byte, elapsed time.Duration) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "[DEBUG] ← %s (%dms)\n", resp.Status, elapsed.Milliseconds())
	sb.WriteString("Response Headers:\n")
	for k, vs := range resp.Header {
		fmt.Fprintf(&sb, "  %s: %s\n", k, strings.Join(vs, ", "))
	}
	if len(body) > 0 {
		sb.WriteString("Response Body:\n")
		sb.WriteString(prettyJSON(body))
		sb.WriteString("\n")
	}
	fmt.Fprint(os.Stderr, sb.String())
}

// httpDebugLogger is the package-level logger used by GenericRequest.
var httpDebugLogger = &requestLogger{}

// redactAuth masks all but the last 5 characters of a Bearer token so the
// caller can verify which key is in use without exposing the full secret.
func redactAuth(value string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(value, prefix) {
		token := value[len(prefix):]
		if len(token) > 5 {
			return prefix + "***" + token[len(token)-5:]
		}
		return prefix + "***"
	}
	if len(value) > 5 {
		return "***" + value[len(value)-5:]
	}
	return "***"
}

// prettyJSON returns an indented JSON string, or the raw bytes if indenting fails.
func prettyJSON(data []byte) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		return string(data)
	}
	return buf.String()
}
