# \ClassResourceListsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClassesIndexFeaturesGet**](ClassResourceListsApi.md#ApiClassesIndexFeaturesGet) | **Get** /api/classes/{index}/features | Get features available for a class.
[**ApiClassesIndexProficienciesGet**](ClassResourceListsApi.md#ApiClassesIndexProficienciesGet) | **Get** /api/classes/{index}/proficiencies | Get proficiencies available for a class.
[**ApiClassesIndexSpellsGet**](ClassResourceListsApi.md#ApiClassesIndexSpellsGet) | **Get** /api/classes/{index}/spells | Get spells available for a class.
[**ApiClassesIndexSubclassesGet**](ClassResourceListsApi.md#ApiClassesIndexSubclassesGet) | **Get** /api/classes/{index}/subclasses | Get subclasses available for a class.



## ApiClassesIndexFeaturesGet

> APIReferenceList ApiClassesIndexFeaturesGet(ctx, index).Execute()

Get features available for a class.

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
    index := "paladin" // string | The `index` of the class to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassResourceListsApi.ApiClassesIndexFeaturesGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassResourceListsApi.ApiClassesIndexFeaturesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexFeaturesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `ClassResourceListsApi.ApiClassesIndexFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexFeaturesGetRequest struct via the builder pattern


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


## ApiClassesIndexProficienciesGet

> APIReferenceList ApiClassesIndexProficienciesGet(ctx, index).Execute()

Get proficiencies available for a class.

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
    index := "paladin" // string | The `index` of the class to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassResourceListsApi.ApiClassesIndexProficienciesGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassResourceListsApi.ApiClassesIndexProficienciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexProficienciesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `ClassResourceListsApi.ApiClassesIndexProficienciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexProficienciesGetRequest struct via the builder pattern


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


## ApiClassesIndexSpellsGet

> APIReferenceList ApiClassesIndexSpellsGet(ctx, index).Execute()

Get spells available for a class.

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
    index := "paladin" // string | The `index` of the class to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassResourceListsApi.ApiClassesIndexSpellsGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassResourceListsApi.ApiClassesIndexSpellsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexSpellsGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `ClassResourceListsApi.ApiClassesIndexSpellsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexSpellsGetRequest struct via the builder pattern


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


## ApiClassesIndexSubclassesGet

> APIReferenceList ApiClassesIndexSubclassesGet(ctx, index).Execute()

Get subclasses available for a class.

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
    index := "paladin" // string | The `index` of the class to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassResourceListsApi.ApiClassesIndexSubclassesGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassResourceListsApi.ApiClassesIndexSubclassesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexSubclassesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `ClassResourceListsApi.ApiClassesIndexSubclassesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexSubclassesGetRequest struct via the builder pattern


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

