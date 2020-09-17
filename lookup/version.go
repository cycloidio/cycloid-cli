package lookup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type AppVersionResp struct {
	Data *AppVersion `json:"data"`
}
type AppVersion struct {
	Branch   string `json:"branch"`
	Revision string `json:"revision"`
	Version  string `json:"version"`
}

func GetAPIVersion() (*AppVersion, error) {
	// Because of https://github.com/golang/go/issues/27751
	// When a Go plugin import a lib. And The static go binary import the same lib with a different version,
	// Basically in our case, we can't use swagger generated client in the plugin and the static.
	// you end up with this kind of error "plugin was built with a different version of package"
	// So we decided to not use swagger lib to get /version endpoint.
	// This is also the reason why -trim is enabled during the build of the CLI. It prevent this error but use the version of
	// the lib from the static build (even if the version from the plugin is newer)

	var (
		versionResp *AppVersionResp
		host        = "http-api-stoplight.cycloid.io"
		scheme      = "https"
		basePath    = "/"
	)

	rawApiUrl := viper.GetString("api-url")
	apiUrl, err := url.Parse(rawApiUrl)
	if err != nil {
		return nil, err
	}

	if apiUrl.Host != "" {
		host = apiUrl.Host
	}
	if apiUrl.Scheme != "" {
		scheme = apiUrl.Scheme
	}
	if apiUrl.Path != "" {
		basePath = apiUrl.Path
	}
	basePath = strings.Trim(basePath, "/")

	// Query the API version endpoint
	resp, err := http.Get(fmt.Sprintf("%s://%s/%s", scheme, host, path.Join(basePath, "version")))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarchal the json
	if err := json.Unmarshal(body, &versionResp); err != nil {
		return nil, err
	}

	return versionResp.Data, nil
}
