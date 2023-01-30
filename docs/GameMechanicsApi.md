# \GameMechanicsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiConditionsIndexGet**](GameMechanicsApi.md#ApiConditionsIndexGet) | **Get** /api/conditions/{index} | Get a condition by index.
[**ApiDamageTypesIndexGet**](GameMechanicsApi.md#ApiDamageTypesIndexGet) | **Get** /api/damage-types/{index} | Get a damage type by index.
[**ApiMagicSchoolsIndexGet**](GameMechanicsApi.md#ApiMagicSchoolsIndexGet) | **Get** /api/magic-schools/{index} | Get a magic school by index.



## ApiConditionsIndexGet

> Condition ApiConditionsIndexGet(ctx, index).Execute()

Get a condition by index.



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
    index := "blinded" // string | The `index` of the condition to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameMechanicsApi.ApiConditionsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameMechanicsApi.ApiConditionsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiConditionsIndexGet`: Condition
    fmt.Fprintf(os.Stdout, "Response from `GameMechanicsApi.ApiConditionsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the condition to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiConditionsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Condition**](Condition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDamageTypesIndexGet

> DamageType ApiDamageTypesIndexGet(ctx, index).Execute()

Get a damage type by index.



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
    index := "acid" // string | The `index` of the damage type to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameMechanicsApi.ApiDamageTypesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameMechanicsApi.ApiDamageTypesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDamageTypesIndexGet`: DamageType
    fmt.Fprintf(os.Stdout, "Response from `GameMechanicsApi.ApiDamageTypesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the damage type to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDamageTypesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DamageType**](DamageType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMagicSchoolsIndexGet

> MagicSchool ApiMagicSchoolsIndexGet(ctx, index).Execute()

Get a magic school by index.



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
    index := "abjuration" // string | The `index` of the magic school to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameMechanicsApi.ApiMagicSchoolsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameMechanicsApi.ApiMagicSchoolsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMagicSchoolsIndexGet`: MagicSchool
    fmt.Fprintf(os.Stdout, "Response from `GameMechanicsApi.ApiMagicSchoolsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the magic school to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMagicSchoolsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MagicSchool**](MagicSchool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

