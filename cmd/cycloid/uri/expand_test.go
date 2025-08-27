package uri

import (
	"testing"
)

func TestIsIgnored(t *testing.T) {
	type args struct {
		path    string
		ignores []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"TestSimpleRelPathOk",
			args{
				"ignore-me/",
				[]string{
					"ignore-me/",
				},
			},
			true,
			false,
		},
		{
			"TestSimpleAbsOk",
			args{
				"/home/toto",
				[]string{
					"/home/toto",
				},
			},
			true,
			false,
		},
		{
			"TestAbsPatternOk",
			args{
				"/home/toto/something",
				[]string{
					"/home/toto/*",
				},
			},
			true,
			false,
		},
		{
			"TestAbsPatternNestedOk",
			args{
				"/home/toto/something/else",
				[]string{
					"/home/toto/*",
				},
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsIgnored(tt.args.path, tt.args.ignores)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsIgnored() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsIgnored() = %v, want %v", got, tt.want)
			}
		})
	}
}
