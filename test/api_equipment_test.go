/*
D&D 5e API

Testing EquipmentApiService

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

func Test_openapi_EquipmentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test EquipmentApiService ApiEquipmentCategoriesIndexGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string

        resp, httpRes, err := apiClient.EquipmentApi.ApiEquipmentCategoriesIndexGet(context.Background(), index).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test EquipmentApiService ApiEquipmentIndexGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string

        resp, httpRes, err := apiClient.EquipmentApi.ApiEquipmentIndexGet(context.Background(), index).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test EquipmentApiService ApiMagicItemsIndexGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string

        resp, httpRes, err := apiClient.EquipmentApi.ApiMagicItemsIndexGet(context.Background(), index).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test EquipmentApiService ApiWeaponPropertiesIndexGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var index string

        resp, httpRes, err := apiClient.EquipmentApi.ApiWeaponPropertiesIndexGet(context.Background(), index).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
