package jq

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/itchyny/gojq"

	"github.com/cycloidio/cycloid-cli/printer"
)

// JQ implements the printer interface using a gojq expression.
// Output is raw (like jq -r): strings printed as-is, other values as pretty JSON.
type JQ struct {
	code *gojq.Code
}

// New compiles a jq expression and returns a JQ printer.
func New(expr string) (*JQ, error) {
	query, err := gojq.Parse(expr)
	if err != nil {
		return nil, fmt.Errorf("invalid jq expression %q: %w", expr, err)
	}
	code, err := gojq.Compile(query)
	if err != nil {
		return nil, fmt.Errorf("unable to compile jq expression %q: %w", expr, err)
	}
	return &JQ{code: code}, nil
}

// Print marshals obj to JSON, runs the jq expression, and writes raw output.
func (j *JQ) Print(obj interface{}, opt printer.Options, w io.Writer) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("unable to marshal to JSON: %w", err)
	}
	var v interface{}
	if err := json.Unmarshal(raw, &v); err != nil {
		return fmt.Errorf("unable to unmarshal JSON: %w", err)
	}

	iter := j.code.Run(v)
	for {
		val, ok := iter.Next()
		if !ok {
			break
		}
		if jqErr, ok := val.(error); ok {
			return jqErr
		}
		switch typed := val.(type) {
		case string:
			fmt.Fprintln(w, typed)
		case nil:
			fmt.Fprintln(w, "null")
		default:
			b, err := json.MarshalIndent(typed, "", "  ")
			if err != nil {
				return fmt.Errorf("unable to marshal jq result: %w", err)
			}
			fmt.Fprintln(w, string(b))
		}
	}
	return nil
}
