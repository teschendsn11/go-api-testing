# \SubracesApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSubracesIndexGet**](SubracesApi.md#ApiSubracesIndexGet) | **Get** /api/subraces/{index} | Get a subrace by index.
[**ApiSubracesIndexProficienciesGet**](SubracesApi.md#ApiSubracesIndexProficienciesGet) | **Get** /api/subraces/{index}/proficiencies | Get proficiences available for a subrace.
[**ApiSubracesIndexTraitsGet**](SubracesApi.md#ApiSubracesIndexTraitsGet) | **Get** /api/subraces/{index}/traits | Get traits available for a subrace.



## ApiSubracesIndexGet

> Subrace ApiSubracesIndexGet(ctx, index).Execute()

Get a subrace by index.



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
    index := "hill-dwarf" // string | The `index` of the subrace to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubracesApi.ApiSubracesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubracesApi.ApiSubracesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubracesIndexGet`: Subrace
    fmt.Fprintf(os.Stdout, "Response from `SubracesApi.ApiSubracesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subrace to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubracesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subrace**](Subrace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubracesIndexProficienciesGet

> APIReferenceList ApiSubracesIndexProficienciesGet(ctx, index).Execute()

Get proficiences available for a subrace.

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
    index := "hill-dwarf" // string | The `index` of the subrace to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubracesApi.ApiSubracesIndexProficienciesGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubracesApi.ApiSubracesIndexProficienciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubracesIndexProficienciesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `SubracesApi.ApiSubracesIndexProficienciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subrace to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubracesIndexProficienciesGetRequest struct via the builder pattern


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


## ApiSubracesIndexTraitsGet

> APIReferenceList ApiSubracesIndexTraitsGet(ctx, index).Execute()

Get traits available for a subrace.

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
    index := "hill-dwarf" // string | The `index` of the subrace to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubracesApi.ApiSubracesIndexTraitsGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubracesApi.ApiSubracesIndexTraitsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSubracesIndexTraitsGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `SubracesApi.ApiSubracesIndexTraitsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the subrace to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubracesIndexTraitsGetRequest struct via the builder pattern


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

