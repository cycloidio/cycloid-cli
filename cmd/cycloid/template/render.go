package template

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/internal/templating"
)

func NewRenderCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "render [flags]",
		Short: "Render templates offline with Cycloid interpolation",
		Long: `Render one or more templates locally using the Cycloid interpolation engine,
with no backend call. Context variables are layered, lowest precedence first:

  --context-file  JSON or YAML file
  stdin           piped JSON object (when no template is read from stdin)
  --context       raw JSON object string
  --set           key=value pairs (dotted keys nest); highest precedence

Variables referenced by a template but not provided render as the literal
"<placeholder:$name>" when they are known Cycloid variables, or are reported as
warnings when unknown.`,
		Example: `
  # render a file with a couple of variables
  cy template render -f main.tf.tpl --set project=my-app --set env=prod

  # pull-once-iterate-locally: real context from a file, tweak one var
  cy template render -f main.tf.tpl --context-file ctx.yaml --set env_vars.region=eu-west-1

  # render from stdin context, template from a directory, JSON output
  cat ctx.json | cy template render --dir ./templates -o json`,
		RunE: runRender,
		Args: cobra.NoArgs,
	}
	cyargs.AddTemplateRenderFlags(cmd)
	return cmd
}

func runRender(cmd *cobra.Command, _ []string) error {
	// Step 1: all flags first.
	files, err := cyargs.GetTemplateFiles(cmd)
	if err != nil {
		return err
	}
	dir, err := cyargs.GetTemplateDir(cmd)
	if err != nil {
		return err
	}
	ctxFile, err := cyargs.GetTemplateContextFile(cmd)
	if err != nil {
		return err
	}
	ctxStr, err := cyargs.GetTemplateContextString(cmd)
	if err != nil {
		return err
	}
	sets, err := cyargs.GetTemplateSet(cmd)
	if err != nil {
		return err
	}

	// Step 2: resolve stdin once. It feeds a "-" template if requested,
	// otherwise it is treated as a piped JSON context.
	wantStdinTemplate := false
	for _, f := range files {
		if f == "-" {
			wantStdinTemplate = true
		}
	}
	var stdinData []byte
	if hasStdinData() {
		stdinData, err = io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read stdin: %w", err)
		}
	}

	// Step 3: build the layered context (ascending precedence).
	ctx := templating.Context{}
	if ctxFile != "" {
		fileCtx, err := templating.LoadContextFile(ctxFile)
		if err != nil {
			return err
		}
		templating.Merge(ctx, fileCtx)
	}
	if !wantStdinTemplate && len(stdinData) > 0 {
		stdinCtx, err := templating.ParseContextString(string(stdinData))
		if err != nil {
			return fmt.Errorf("stdin: %w", err)
		}
		templating.Merge(ctx, stdinCtx)
	}
	if ctxStr != "" {
		strCtx, err := templating.ParseContextString(ctxStr)
		if err != nil {
			return err
		}
		templating.Merge(ctx, strCtx)
	}
	if len(sets) > 0 {
		setCtx, err := templating.ParseSet(sets)
		if err != nil {
			return err
		}
		templating.Merge(ctx, setCtx)
	}

	// Step 4: gather templates.
	type tmpl struct{ name, content string }
	var tmpls []tmpl
	for _, f := range files {
		if f == "-" {
			tmpls = append(tmpls, tmpl{name: "stdin", content: string(stdinData)})
			continue
		}
		content, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("failed to read template %q: %w", f, err)
		}
		tmpls = append(tmpls, tmpl{name: f, content: string(content)})
	}
	if dir != "" {
		err = filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read template %q: %w", path, err)
			}
			tmpls = append(tmpls, tmpl{name: path, content: string(content)})
			return nil
		})
		if err != nil {
			return err
		}
	}
	if len(tmpls) == 0 {
		return fmt.Errorf("no template provided: pass --file, --dir, or pipe a template with --file -")
	}

	// Step 5: render and report.
	reports := make([]templating.Report, 0, len(tmpls))
	failed := 0
	for _, t := range tmpls {
		r := templating.Render(t.name, t.content, ctx)
		if r.Error != "" {
			failed++
		}
		reports = append(reports, r)
	}

	// A single template prints a single object; multiple print a list.
	var out any = reports
	if len(reports) == 1 {
		out = reports[0]
	}
	if printErr := cyout.Print(cmd, out, nil, ""); printErr != nil {
		return printErr
	}
	if failed > 0 {
		return fmt.Errorf("%d of %d template(s) failed to render", failed, len(tmpls))
	}
	return nil
}

// hasStdinData reports whether stdin is a pipe or redirected file (i.e. has
// data) rather than an interactive terminal.
func hasStdinData() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice == 0
}
