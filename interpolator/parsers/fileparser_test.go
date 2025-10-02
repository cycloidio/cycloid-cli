package parsers_test

import (
	"testing"

	"github.com/cycloidio/cycloid-cli/interpolator/parsers"
	"github.com/cycloidio/cycloid-cli/interpolator/resolvers"
	"github.com/cycloidio/cycloid-cli/interpolator/resolvers/mockresolver"
)

func TestFileParserOk(t *testing.T) {
}

func TestReplaceFile(t *testing.T) {
	mock := mockresolver.NewMockResolverWithDefault()

	type args struct {
		resolver resolvers.ResourceResolver
		file     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"simpleInterpolationOk",
			args{
				resolver: mock,
				file:     `ssh: cy://simple/string`,
			},
			"ssh: simple",
			false,
		},
		{
			"simpleInterpolationNotOk",
			args{
				resolver: mock,
				file:     `ssh: cy://doesnot/exists`,
			},
			`ssh: cy://doesnot/exists`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsers.ReplaceFile(tt.args.resolver, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReplaceFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReplaceFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
