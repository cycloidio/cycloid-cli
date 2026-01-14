package httpresolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/interpolator/resolvers"
	"github.com/cycloidio/cycloid-cli/interpolator/resources"
)

type HTTPResolverOption func(*HTTPResolver) error

func NewHTTPResolver(options ...HTTPResolverOption) (*HTTPResolver, error) {
	resolver := &HTTPResolver{
		client: http.DefaultClient,
	}

	for _, option := range options {
		err := option(resolver)
		if err != nil {
			return nil, fmt.Errorf("failed to configure the HTTP Resolver: %w", err)
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
		return nil, fmt.Errorf("failed to parse API URL %q: %w", api.Config.URL, err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+apiKey)

	client := http.DefaultClient
	client.Timeout = time.Second * 30

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch resource with ref %q: %w", ref.Path, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %q", err)
	}

	var apiResponse *APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %q", err)
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

		return nil, fmt.Errorf(
			"failed to request %q to API details: %s: %w",
			ref.Path, strings.Join(details, "\n"), apiResponse.Errors,
		)
	}
}
