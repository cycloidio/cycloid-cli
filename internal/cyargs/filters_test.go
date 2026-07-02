package cyargs_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

func TestGetLHSFilters(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []middleware.LHSFilter
		wantErr bool
	}{
		{
			name: "no filters",
			args: nil,
			want: []middleware.LHSFilter{},
		},
		{
			name: "single eq filter",
			args: []string{"--filter", "output_type[eq]=string"},
			want: []middleware.LHSFilter{{Attribute: "output_type", Condition: "eq", Value: "string"}},
		},
		{
			name: "multiple filters",
			args: []string{"--filter", "project_canonical[eq]=p", "--filter", "output_is_pinned[eq]=true"},
			want: []middleware.LHSFilter{
				{Attribute: "project_canonical", Condition: "eq", Value: "p"},
				{Attribute: "output_is_pinned", Condition: "eq", Value: "true"},
			},
		},
		{
			name: "rlike with regex metachars",
			args: []string{"--filter", "resources_name[rlike]=lhs-res-.*"},
			want: []middleware.LHSFilter{{Attribute: "resources_name", Condition: "rlike", Value: "lhs-res-.*"}},
		},
		{
			name: "value containing equals",
			args: []string{"--filter", "output_key[eq]=a=b"},
			want: []middleware.LHSFilter{{Attribute: "output_key", Condition: "eq", Value: "a=b"}},
		},
		{
			name:    "missing bracket",
			args:    []string{"--filter", "output_typeeq=string"},
			wantErr: true,
		},
		{
			name:    "missing closing bracket-eq",
			args:    []string{"--filter", "output_type[eq string"},
			wantErr: true,
		},
		{
			name:    "empty condition",
			args:    []string{"--filter", "output_type[]=string"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{Use: "test", RunE: func(*cobra.Command, []string) error { return nil }}
			cyargs.AddLHSFilterFlag(cmd)
			require.NoError(t, cmd.ParseFlags(tt.args))

			got, err := cyargs.GetLHSFilters(cmd)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
