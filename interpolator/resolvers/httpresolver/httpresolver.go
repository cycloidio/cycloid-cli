package httpresolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/interpolator/resources"
	"github.com/itchyny/gojq"
)

type HTTPResolverOption func(*HTTPResolver) error

func NewHTTPResolver(options ...HTTPResolverOption) (*HTTPResolver, error) {
	resolver := &HTTPResolver{
		client: http.DefaultClient,
	}

	for _, option := range options {
		err := option(resolver)
		if err != nil {
			return nil, fmt.Errorf("failed to configure the HTTP Resolver: %s", err.Error())
		}
	}

	return resolver, nil
}

type HTTPResolver struct {
	client *http.Client
}

func (r HTTPResolver) Resolve(ref *resources.Reference) ([]any, error) {
	// We could need to make the cli config retrieval generic when tenant feature
	// will be implemented
	api := common.NewAPI()
	apiKey := api.GetToken(nil)

	request, err := http.NewRequest(
		http.MethodGet,
		api.Config.URL+ref.Path+"?"+ref.Params.Encode(),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API URL %q: %s", api.Config.URL, err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+apiKey)

	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch resource with ref '%s': %s", ref.Path, err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err.Error())
	}

	var apiResponse *APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %s", err.Error())
	}

	switch {
	case response.StatusCode >= 200 && response.StatusCode < 300:
		data, err := resolvers.Query(ref.Params, apiResponse.Data)
		if err != nil {
			return nil, err
		}

		return data, nil
	default:
		var details = make([]string, len(apiResponse.Errors))
		for index, apiErr := range apiResponse.Errors {
			details[index] = apiErr.String()
		}

		return nil, fmt.Errorf("failed to request '%s' to API: %s", ref.Path, apiResponse.Errors.Error())
	}
}

func (r HTTPResolver) query(params map[string][]string, data any) ([]any, error) {
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

		if err, ok := v.(error); ok {
			if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
				break
			}
			queryErr = fmt.Errorf("%s: %s", queryErr.Error(), err.Error())
		}

		outData = append(outData, v)
	}
	if queryErr != nil {
		return nil, fmt.Errorf("key query has reported an error: %s", err.Error())
	}

	return outData, nil
}
