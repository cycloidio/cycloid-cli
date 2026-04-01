package middleware_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func ptrStr(s string) *string { return &s }

func TestAPIResponseError(t *testing.T) {
	t.Run("ErrorWithPayload", func(t *testing.T) {
		err := &middleware.APIResponseError{
			StatusCode: 422,
			Status:     "422 Unprocessable Entity",
			Payload: &models.ErrorPayload{
				Errors: []*models.ErrorDetailsItem{
					{
						Code:    ptrStr("code"),
						Details: []string{"some", "details"},
						Message: ptrStr("the error that actually returned the BE"),
					},
				},
			},
		}

		assert.Equal(t, 422, err.StatusCode)
		assert.Equal(t, "API error 422: the error that actually returned the BE", err.Error())
	})

	t.Run("ErrorWithoutPayload", func(t *testing.T) {
		err := &middleware.APIResponseError{
			StatusCode: 500,
			Status:     "500 Internal Server Error",
		}
		assert.Equal(t, "API error 500: 500 Internal Server Error", err.Error())
	})

	t.Run("ErrorWithoutPayloadFallbackToRawBodyAndPath", func(t *testing.T) {
		err := &middleware.APIResponseError{
			StatusCode: 422,
			Status:     "422 Unprocessable Entity",
			Path:       "/organizations/org/projects/project/environments/env/components",
			Body:       []byte("stack branch simple-terraform not found"),
		}
		assert.Equal(
			t,
			`API error 422 on "/organizations/org/projects/project/environments/env/components": stack branch simple-terraform not found`,
			err.Error(),
		)
	})

	t.Run("GetPayload", func(t *testing.T) {
		payload := &models.ErrorPayload{}
		err := &middleware.APIResponseError{
			StatusCode: http.StatusConflict,
			Payload:    payload,
		}
		assert.Equal(t, payload, err.GetPayload())
	})

	t.Run("ErrorWithoutPayloadFallbackToRawBodyWithoutPath", func(t *testing.T) {
		err := &middleware.APIResponseError{
			StatusCode: 422,
			Status:     "422 Unprocessable Entity",
			Body:       []byte("stack branch simple-terraform not found"),
		}
		assert.Equal(
			t,
			"API error 422: stack branch simple-terraform not found",
			err.Error(),
		)
	})
}
