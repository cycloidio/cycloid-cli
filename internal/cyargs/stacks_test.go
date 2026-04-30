package cyargs_test

import (
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

// stubMiddleware implements only ListStackVersions; all other methods panic.
type stubMiddleware struct {
	middleware.Middleware
	versions []*middleware.StackVersion
	err      error
}

func (s *stubMiddleware) ListStackVersions(org, ref string, _ ...middleware.LHSFilter) ([]*middleware.StackVersion, *http.Response, error) {
	return s.versions, nil, s.err
}

func newCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cyargs.AddStackVersionFlags(cmd)
	return cmd
}

func makeVersion(name, vType, commitHash string, isLatest bool) *middleware.StackVersion {
	return &middleware.StackVersion{
		Name:       ptr.Ptr(name),
		Type:       ptr.Ptr(vType),
		CommitHash: ptr.Ptr(commitHash),
		IsLatest:   ptr.Ptr(isLatest),
	}
}

func TestResolveStackVersionArg(t *testing.T) {
	tagVersion := makeVersion("v1.0.0", "tag", "aabbccdd1234567", true)
	branchVersion := makeVersion("main", "branch", "deadbeef1234567", false)
	ambigTag := makeVersion("release", "tag", "1111111111111111", false)
	ambigBranch := makeVersion("release", "branch", "2222222222222222", false)

	canned := []*middleware.StackVersion{tagVersion, branchVersion, ambigTag, ambigBranch}

	m := &stubMiddleware{versions: canned}
	const org = "myorg"
	const stackRef = "myorg:mystack"

	cases := []struct {
		name       string
		args       []string
		wantTag    string
		wantBranch string
		wantHash   string
		wantErr    string
	}{
		{
			name:    "empty returns empty triple",
			args:    nil,
			wantTag: "", wantBranch: "", wantHash: "",
		},
		{
			name:    "tag prefix resolved client-side",
			args:    []string{"--stack-version", "tag:v1.0.0"},
			wantTag: "v1.0.0",
		},
		{
			name:       "branch prefix resolved client-side",
			args:       []string{"--stack-version", "branch:main"},
			wantBranch: "main",
		},
		{
			name:     "sha prefix resolved client-side",
			args:     []string{"--stack-version", "sha:deadbeef"},
			wantHash: "deadbeef",
		},
		{
			name:     "commit prefix alias resolved client-side",
			args:     []string{"--stack-version", "commit:deadbeef"},
			wantHash: "deadbeef",
		},
		{
			name:    "bare value resolves to tag (tag wins over branch)",
			args:    []string{"--stack-version", "v1.0.0"},
			wantTag: "v1.0.0",
		},
		{
			name:       "bare value resolves to branch when no tag matches",
			args:       []string{"--stack-version", "main"},
			wantBranch: "main",
		},
		{
			name:     "bare short hash resolves to commit",
			args:     []string{"--stack-version", "aabbccd"},
			wantHash: "aabbccdd1234567",
		},
		{
			name:    "collision between tag and branch errors with hint",
			args:    []string{"--stack-version", "release"},
			wantErr: "ambiguous version",
		},
		{
			name:    "unknown bare value returns not-found error",
			args:    []string{"--stack-version", "nonexistent"},
			wantErr: "not found",
		},
		{
			name:    "short hash under 7 chars is not matched as commit",
			args:    []string{"--stack-version", "aabbcc"},
			wantErr: "not found",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := newCmd()
			if tc.args != nil {
				require.NoError(t, cmd.ParseFlags(tc.args))
			}

			tag, branch, hash, err := cyargs.ResolveStackVersionArg(cmd, m, org, stackRef)

			if tc.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.wantTag, tag)
			assert.Equal(t, tc.wantBranch, branch)
			assert.Equal(t, tc.wantHash, hash)
		})
	}

	t.Run("legacy --stack-tag bypasses resolution", func(t *testing.T) {
		noCallMiddleware := &stubMiddleware{} // panics if ListStackVersions called
		cmd := newCmd()
		require.NoError(t, cmd.ParseFlags([]string{"--stack-tag", "v2.0.0"}))

		tag, branch, hash, err := cyargs.ResolveStackVersionArg(cmd, noCallMiddleware, org, stackRef)
		require.NoError(t, err)
		assert.Equal(t, "v2.0.0", tag)
		assert.Empty(t, branch)
		assert.Empty(t, hash)
	})

	t.Run("legacy --stack-branch bypasses resolution", func(t *testing.T) {
		noCallMiddleware := &stubMiddleware{}
		cmd := newCmd()
		require.NoError(t, cmd.ParseFlags([]string{"--stack-branch", "dev"}))

		tag, branch, hash, err := cyargs.ResolveStackVersionArg(cmd, noCallMiddleware, org, stackRef)
		require.NoError(t, err)
		assert.Empty(t, tag)
		assert.Equal(t, "dev", branch)
		assert.Empty(t, hash)
	})

	t.Run("tag prefix with no API call (noCallMiddleware should not panic)", func(t *testing.T) {
		noCallMiddleware := &stubMiddleware{}
		cmd := newCmd()
		require.NoError(t, cmd.ParseFlags([]string{"--stack-version", "tag:v1.0.0"}))

		tag, _, _, err := cyargs.ResolveStackVersionArg(cmd, noCallMiddleware, org, stackRef)
		require.NoError(t, err)
		assert.Equal(t, "v1.0.0", tag)
	})

	t.Run("bare value without stackRef returns helpful error", func(t *testing.T) {
		cmd := newCmd()
		require.NoError(t, cmd.ParseFlags([]string{"--stack-version", "main"}))

		_, _, _, err := cyargs.ResolveStackVersionArg(cmd, m, org, "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "stack reference")
	})
}
