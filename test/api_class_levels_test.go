/*
D&D 5e API

Testing ClassLevelsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_ClassLevelsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ClassLevelsApiService ApiClassesIndexLevelsClassLevelFeaturesGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string
        var classLevel float32

        resp, httpRes, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsClassLevelFeaturesGet(context.Background(), index, classLevel).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClassLevelsApiService ApiClassesIndexLevelsClassLevelGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string
        var classLevel float32

        resp, httpRes, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsClassLevelGet(context.Background(), index, classLevel).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClassLevelsApiService ApiClassesIndexLevelsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string

        resp, httpRes, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsGet(context.Background(), index).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClassLevelsApiService ApiClassesIndexLevelsSpellLevelSpellsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string
        var spellLevel float32

        resp, httpRes, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsSpellLevelSpellsGet(context.Background(), index, spellLevel).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
