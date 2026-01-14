package resolvers

import (
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {
	type args struct {
		params map[string][]string
		data   any
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		wantErr bool
	}{
		{
			"NullResultShouldErr",
			args{
				params: map[string][]string{
					"key": {".does.not.exists"},
				},
				data: nil,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Query(tt.args.params, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
