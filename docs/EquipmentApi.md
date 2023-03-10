# \EquipmentApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEquipmentCategoriesIndexGet**](EquipmentApi.md#ApiEquipmentCategoriesIndexGet) | **Get** /api/equipment-categories/{index} | Get an equipment category by index.
[**ApiEquipmentIndexGet**](EquipmentApi.md#ApiEquipmentIndexGet) | **Get** /api/equipment/{index} | Get an equipment item by index.
[**ApiMagicItemsIndexGet**](EquipmentApi.md#ApiMagicItemsIndexGet) | **Get** /api/magic-items/{index} | Get a magic item by index.
[**ApiWeaponPropertiesIndexGet**](EquipmentApi.md#ApiWeaponPropertiesIndexGet) | **Get** /api/weapon-properties/{index} | Get a weapon property by index.



## ApiEquipmentCategoriesIndexGet

> EquipmentCategory ApiEquipmentCategoriesIndexGet(ctx, index).Execute()

Get an equipment category by index.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    index := "waterborne-vehicles" // string | The `index` of the equipment category score to get.  Available values can be found in the resource list for this endpoint. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentApi.ApiEquipmentCategoriesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.ApiEquipmentCategoriesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEquipmentCategoriesIndexGet`: EquipmentCategory
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.ApiEquipmentCategoriesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the equipment category score to get.  Available values can be found in the resource list for this endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEquipmentCategoriesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EquipmentCategory**](EquipmentCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEquipmentIndexGet

> Equipment ApiEquipmentIndexGet(ctx, index).Execute()

Get an equipment item by index.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    index := "club" // string | The `index` of the equipment to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `equipment`. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentApi.ApiEquipmentIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.ApiEquipmentIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEquipmentIndexGet`: Equipment
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.ApiEquipmentIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the equipment to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;equipment&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEquipmentIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Equipment**](Equipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMagicItemsIndexGet

> MagicItem ApiMagicItemsIndexGet(ctx, index).Execute()

Get a magic item by index.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    index := "adamantine-armor" // string | The `index` of the magic item to get.  Available values can be found in the resource list for this endpoint. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentApi.ApiMagicItemsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.ApiMagicItemsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMagicItemsIndexGet`: MagicItem
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.ApiMagicItemsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the magic item to get.  Available values can be found in the resource list for this endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMagicItemsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MagicItem**](MagicItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiWeaponPropertiesIndexGet

> WeaponProperty ApiWeaponPropertiesIndexGet(ctx, index).Execute()

Get a weapon property by index.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    index := "ammunition" // string | The `index` of the weapon property to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentApi.ApiWeaponPropertiesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.ApiWeaponPropertiesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiWeaponPropertiesIndexGet`: WeaponProperty
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.ApiWeaponPropertiesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the weapon property to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiWeaponPropertiesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WeaponProperty**](WeaponProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

