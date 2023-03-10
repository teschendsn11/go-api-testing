# \RacesApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRacesIndexGet**](RacesApi.md#ApiRacesIndexGet) | **Get** /api/races/{index} | Get a race by index.
[**ApiRacesIndexProficienciesGet**](RacesApi.md#ApiRacesIndexProficienciesGet) | **Get** /api/races/{index}/proficiencies | Get proficiencies available for a race.
[**ApiRacesIndexSubracesGet**](RacesApi.md#ApiRacesIndexSubracesGet) | **Get** /api/races/{index}/subraces | Get subraces available for a race.
[**ApiRacesIndexTraitsGet**](RacesApi.md#ApiRacesIndexTraitsGet) | **Get** /api/races/{index}/traits | Get traits available for a race.



## ApiRacesIndexGet

> Race ApiRacesIndexGet(ctx, index).Execute()

Get a race by index.



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
    index := "elf" // string | The `index` of the race to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacesApi.ApiRacesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacesApi.ApiRacesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRacesIndexGet`: Race
    fmt.Fprintf(os.Stdout, "Response from `RacesApi.ApiRacesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the race to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRacesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Race**](Race.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRacesIndexProficienciesGet

> APIReferenceList ApiRacesIndexProficienciesGet(ctx, index).Execute()

Get proficiencies available for a race.

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
    index := "elf" // string | The `index` of the race to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacesApi.ApiRacesIndexProficienciesGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacesApi.ApiRacesIndexProficienciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRacesIndexProficienciesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `RacesApi.ApiRacesIndexProficienciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the race to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRacesIndexProficienciesGetRequest struct via the builder pattern


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


## ApiRacesIndexSubracesGet

> APIReferenceList ApiRacesIndexSubracesGet(ctx, index).Execute()

Get subraces available for a race.

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
    index := "elf" // string | The `index` of the race to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacesApi.ApiRacesIndexSubracesGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacesApi.ApiRacesIndexSubracesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRacesIndexSubracesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `RacesApi.ApiRacesIndexSubracesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the race to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRacesIndexSubracesGetRequest struct via the builder pattern


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


## ApiRacesIndexTraitsGet

> APIReferenceList ApiRacesIndexTraitsGet(ctx, index).Execute()

Get traits available for a race.

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
    index := "elf" // string | The `index` of the race to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacesApi.ApiRacesIndexTraitsGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacesApi.ApiRacesIndexTraitsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRacesIndexTraitsGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `RacesApi.ApiRacesIndexTraitsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the race to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRacesIndexTraitsGetRequest struct via the builder pattern


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

