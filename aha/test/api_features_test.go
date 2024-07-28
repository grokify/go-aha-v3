/*
Aha.io API

Testing FeaturesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package aha

import (
	"context"
	"testing"

	openapiclient "github.com/grokify/go-aha-v3/aha"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_aha_FeaturesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeaturesAPIService GetFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesAPI.GetFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService GetFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.GetFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService GetReleaseFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var releaseId string

		resp, httpRes, err := apiClient.FeaturesAPI.GetReleaseFeatures(context.Background(), releaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService UpdateFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesAPI.UpdateFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}