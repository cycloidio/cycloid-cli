package resolvers

import (
	"errors"
	"fmt"

	"github.com/itchyny/gojq"
)

func Query(params map[string][]string, data any) ([]any, error) {
	var query *gojq.Query
	var err error
	if paths, ok := params["key"]; ok {
		query, err = gojq.Parse(paths[0])
		if err != nil {
			return nil, fmt.Errorf("invalid key parameter '%s': %s", paths[0], err.Error())
		}
	} else {
		query, err = gojq.Parse(".")
		if err != nil {
			return nil, fmt.Errorf("invalid key default parameter: %s", err.Error())
		}
	}

	var outData []any
	var queryErr error
	iter := query.Run(data)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}

		var gojqErr *gojq.HaltError
		if err, ok := v.(error); ok {
			if errors.As(err, &gojqErr) && err == nil {
				break
			}
			queryErr = fmt.Errorf("%w: %w", queryErr, err)
			continue
		}

		outData = append(outData, v)
	}
	if queryErr != nil {
		return nil, fmt.Errorf("key query has reported an error: %w", queryErr)
	}

	return outData, nil
}
