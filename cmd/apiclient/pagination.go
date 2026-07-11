package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

const (
	defaultPageSize = 100
	maxPages        = 1_000
)

type paginatedListResponse[T any] struct {
	Data       []T                `json:"data"`
	Pagination *models.Pagination `json:"pagination"`
}

// paginatedList fetches all pages for a paginated list endpoint.
func paginatedList[T any](m *apiClient, req Request, pageSize int) ([]T, *http.Response, error) {
	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	var (
		all      []T
		lastResp *http.Response
	)

	for page := 1; page <= maxPages; page++ {
		query := url.Values{}
		if req.Query != nil {
			encoded, err := encodeQuery(req.Query)
			if err != nil {
				return nil, lastResp, err
			}
			query = encoded
		}
		query.Set("page_index", strconv.Itoa(page))
		query.Set("page_size", strconv.Itoa(pageSize))

		pageReq := req
		pageReq.Query = query

		resp, body, err := m.genericRequestRaw(pageReq)
		lastResp = resp
		if err != nil {
			return nil, resp, err
		}

		var pageResult paginatedListResponse[T]
		if err := json.Unmarshal(body, &pageResult); err != nil {
			return nil, resp, fmt.Errorf("failed to decode paginated JSON response: %w", err)
		}

		all = append(all, pageResult.Data...)

		if len(pageResult.Data) == 0 {
			break
		}
		if pageResult.Pagination == nil || pageResult.Pagination.Total == nil {
			break
		}
		if uint64(len(all)) >= *pageResult.Pagination.Total {
			break
		}
	}

	return all, lastResp, nil
}
