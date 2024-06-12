/*
@open-sauced/api.opensauced.pizza

Testing BlogSummaryServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"testing"

	openapiclient "github.com/open-sauced/go-api/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_BlogSummaryServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlogSummaryServiceAPIService GenerateBlogSummary", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogSummaryServiceAPI.GenerateBlogSummary(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
