package common

import (
	"net/url"
	"strings"
)

// PipelineBuildConsoleURL returns a deep link to a pipeline job build in the Cycloid console.
// base must be non-empty (e.g. https://console.cycloid.io). Path segments are escaped.
func PipelineBuildConsoleURL(base, org, project, env, component, pipeline, job, buildID string) (string, bool) {
	b := strings.TrimSpace(base)
	if b == "" {
		return "", false
	}
	b = strings.TrimSuffix(b, "/")
	parts := []string{
		b,
		"organizations", url.PathEscape(org),
		"projects", url.PathEscape(project),
		"environments", url.PathEscape(env),
		"components", url.PathEscape(component),
		"pipelines", url.PathEscape(pipeline),
		"jobs", url.PathEscape(job),
		"builds", url.PathEscape(buildID),
	}
	return strings.Join(parts, "/"), true
}
