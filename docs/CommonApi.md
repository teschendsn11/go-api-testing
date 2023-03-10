# \CommonApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEndpointGet**](CommonApi.md#ApiEndpointGet) | **Get** /api/{endpoint} | Get list of all available resources for an endpoint.
[**ApiGet**](CommonApi.md#ApiGet) | **Get** /api | Get all resource URLs.



## ApiEndpointGet

> APIReferenceList ApiEndpointGet(ctx, endpoint).Execute()

Get list of all available resources for an endpoint.



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
    endpoint := "ability-scores" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommonApi.ApiEndpointGet(context.Background(), endpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.ApiEndpointGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEndpointGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `CommonApi.ApiEndpointGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpoint** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEndpointGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**APIReferenceList**](APIReferenceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGet

> map[string]string ApiGet(ctx).Execute()

Get all resource URLs.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommonApi.ApiGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.ApiGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGet`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `CommonApi.ApiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiGetRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

