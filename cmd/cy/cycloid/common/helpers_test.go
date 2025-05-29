package common

import "testing"

func TestUpdateMapField(t *testing.T) {
	type Args struct {
		field string
		value string
		m     map[string]map[string]map[string]any
	}
	tests := []struct {
		name    string
		args    Args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "StringDoubleQuote",
			args: Args{
				field: "section.group.string",
				value: `"my-string"`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"string": "my-string"}},
				},
			},
			wantErr: false,
		},
		{
			name: "StringSimpleQuote",
			args: Args{
				field: "section.group.string",
				value: `'my-string'`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"string": "my-string"}},
				},
			},
			wantErr: false,
		},
		{
			name: "StringNoQuote",
			args: Args{
				field: "section.group.string",
				value: `my-string`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"string": "my-string"}},
				},
			},
			wantErr: false,
		},
		{
			name: "Int",
			args: Args{
				field: "section.group.int",
				value: `1`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"int": 1}},
				},
			},
			wantErr: false,
		},
		{
			name: "IntAsString",
			args: Args{
				field: "section.group.string",
				value: `"1"`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"string": "1"}},
				},
			},
			wantErr: false,
		},
		{
			name: "Array",
			args: Args{
				field: "section.group.array",
				value: `["hello", "world"]`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"array": []string{"hello", "world"}}},
				},
			},
			wantErr: false,
		},
		{
			name: "ArraySingleValue",
			args: Args{
				field: "section.group.array",
				value: `["hi"]`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"array": []string{"hi"}}},
				},
			},
			wantErr: false,
		},
		{
			name: "Map",
			args: Args{
				field: "section.group.array",
				value: `{"hello": "world"}`,
				m: map[string]map[string]map[string]any{
					"section": {"group": {"array": map[string]string{"hello": "world"}}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateMapField(tt.args.field, tt.args.value, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMapField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
