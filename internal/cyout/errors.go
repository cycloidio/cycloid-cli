package cyout

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// apiErrorInfo is satisfied by *middleware.APIResponseError.
// Defined locally to avoid importing cmd/cycloid/middleware from internal/.
type apiErrorInfo interface {
	error
	HTTPStatusCode() int
	HTTPRequestMethod() string
	HTTPRequestPath() string
	HTTPRequestBody() []byte
	HTTPResponseBody() []byte
}

// apiErrorPayloader is optionally implemented by errors that carry a parsed payload.
type apiErrorPayloader interface {
	GetPayload() *models.ErrorPayload
}

var (
	styleHeader    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))  // bold red
	styleDim       = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))             // dim gray
	styleEndpoint  = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))            // yellow
	styleCode      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("8"))  // bold dim
	styleCmdHint   = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Italic(true) // dim italic
)

// printError writes a formatted error block to cmd's stderr.
// API errors get a rich display with status, method, path, payload details, and request ID.
// Local errors get a one-line command hint showing what was invoked.
func printError(cmd *cobra.Command, err error) {
	w := cmd.ErrOrStderr()
	if ae, ok := err.(apiErrorInfo); ok {
		printAPIError(w, ae)
		return
	}
	printLocalError(w, cmd)
}

func printAPIError(w io.Writer, ae apiErrorInfo) {
	// Header: "API Error 422 — POST /organizations/myorg/projects"
	status := ae.HTTPStatusCode()
	method := ae.HTTPRequestMethod()
	path := ae.HTTPRequestPath()

	header := fmt.Sprintf("%s %s %s",
		styleHeader.Render(fmt.Sprintf("API Error %d", status)),
		styleDim.Render("—"),
		styleEndpoint.Render(method+" "+path),
	)
	fmt.Fprintln(w, header)

	// Optional: sanitized request body
	if body := ae.HTTPRequestBody(); len(body) > 0 {
		fmt.Fprintf(w, "  %s %s\n", styleDim.Render("Body:"), string(body))
	}

	// Error details from parsed payload, or raw response body as fallback
	var requestID string
	if pe, ok := ae.(apiErrorPayloader); ok {
		if p := pe.GetPayload(); p != nil && len(p.Errors) > 0 {
			for _, item := range p.Errors {
				if item == nil {
					continue
				}
				printErrorItem(w, item)
			}
			requestID = p.RequestID
		}
	}

	if requestID == "" {
		// No structured payload — fall back to raw response body
		raw := strings.TrimSpace(string(ae.HTTPResponseBody()))
		if raw != "" {
			fmt.Fprintf(w, "  %s\n", raw)
		}
	}

	if requestID != "" {
		fmt.Fprintf(w, "  %s\n", styleDim.Render("Request-ID: "+requestID))
	}

	fmt.Fprintln(w)
}

func printErrorItem(w io.Writer, item *models.ErrorDetailsItem) {
	code := ""
	if item.Code != nil {
		code = styleCode.Render("["+*item.Code+"]") + " "
	}

	msg := ""
	if item.Message != nil {
		msg = *item.Message
	}

	fmt.Fprintf(w, "  %s%s\n", code, msg)

	for _, d := range item.Details {
		if d != "" {
			fmt.Fprintf(w, "      %s %s\n", styleDim.Render("·"), d)
		}
	}
}

func printLocalError(w io.Writer, cmd *cobra.Command) {
	// Build "cy project get --org myorg" hint from the command path and set flags
	var parts []string
	parts = append(parts, cmd.CommandPath())
	cmd.Flags().Visit(func(f *pflag.Flag) {
		parts = append(parts, "--"+f.Name+"="+f.Value.String())
	})

	hint := strings.Join(parts, " ")
	if hint != "" {
		fmt.Fprintln(w, styleCmdHint.Render(hint))
	}
}
