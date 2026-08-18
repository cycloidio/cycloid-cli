package engine

// KnownKeys returns the set of top-level template variables the Cycloid
// interpolator recognizes: every scalar context key (org, project, env,
// component, repo coordinates, …) plus the dynamic collection keys. The
// offline render wrapper uses this to tell a *known-but-unset* variable
// (rendered as a placeholder) apart from an *unknown* one (a likely typo,
// surfaced as a warning).
func KnownKeys() map[string]struct{} {
	keys := map[string]struct{}{
		"env_vars":              {},
		"environment_vars":      {},
		"env_providers":         {},
		"environment_providers": {},
	}
	// org is the first interpolatorEntity (iota 0) and currentUserUsername the
	// last; iterate the full enum and collect each snake-cased name.
	for e := org; e <= currentUserUsername; e++ {
		keys[e.String()] = struct{}{}
	}
	return keys
}
