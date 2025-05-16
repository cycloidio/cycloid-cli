package middleware_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/client/organizations"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func ptrStr(s string) *string { return &s }

func TestNewApiError(t *testing.T) {
	t.Run("SuccessWith_ErrorPayloader", func(t *testing.T) {
		err := &organizations.CreateOrgUnprocessableEntity{
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

		aerr := middleware.NewAPIError(err)
		apierr := aerr.(*middleware.APIError)
		assert.Equal(t, "POST", apierr.HTTPMethod)
		assert.Equal(t, "/organizations", apierr.URL)
		assert.Equal(t, "422", apierr.HTTPCode)
		assert.Equal(t, "createOrgUnprocessableEntity", apierr.APIAction)

		assert.Equal(t, "A 422 error was returned on \"createOrgUnprocessableEntity\" call with message: the error that actually returned the BE", aerr.Error())
	})
	t.Run("SuccessWhenNo_ErrPayloader", func(t *testing.T) {
		err := fmt.Errorf("std error")
		aerr := middleware.NewAPIError(err)
		_, ok := aerr.(*middleware.APIError)
		assert.False(t, ok)
		assert.Equal(t, "std error", aerr.Error())
	})
	t.Run("SuccessWhenNil", func(t *testing.T) {
		aerr := middleware.NewAPIError(nil)
		_, ok := aerr.(*middleware.APIError)
		assert.False(t, ok)
		assert.Nil(t, aerr)
	})
}
