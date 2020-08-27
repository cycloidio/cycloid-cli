package common

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

var orgRe = regexp.MustCompile(`\(\$ organization_canonical \$\)`)
var envRe = regexp.MustCompile(`\(\$ environment \$\)`)
var projRe = regexp.MustCompile(`\(\$ project \$\)`)

func RequiredPersistentFlag(withFlag func(cmd *cobra.Command) string, cmd *cobra.Command) {
	flagName := withFlag(cmd)
	cmd.MarkPersistentFlagRequired(flagName)
}
func RequiredFlag(withFlag func(cmd *cobra.Command) string, cmd *cobra.Command) {
	flagName := withFlag(cmd)
	cmd.MarkFlagRequired(flagName)
}

type CycloidContext struct {
	Org     string
	Env     string
	Project string
}

func ReplaceCycloidVars(ctx CycloidContext, text []byte) []byte {
	if ctx.Org != "" {
		text = orgRe.ReplaceAll(text, []byte(ctx.Org))
	}
	if ctx.Env != "" {
		text = envRe.ReplaceAll(text, []byte(ctx.Env))
	}
	if ctx.Project != "" {
		text = projRe.ReplaceAll(text, []byte(ctx.Project))
	}
	return text
}

func ReplaceCycloidVarsString(ctx CycloidContext, text string) string {
	if ctx.Org != "" {
		text = orgRe.ReplaceAllString(text, ctx.Org)
	}
	if ctx.Env != "" {
		text = envRe.ReplaceAllString(text, ctx.Env)
	}
	if ctx.Project != "" {
		text = projRe.ReplaceAllString(text, ctx.Project)
	}
	return text
}

func IsInList(pattern string, list []string) bool {
	for _, x := range list {
		if x == pattern {
			return true
		}
	}
	return false
}

func GetPipelineName(project, env string) string {
	return fmt.Sprintf("%s-%s", project, env)
}
