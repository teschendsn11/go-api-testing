# \ClassLevelsApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClassesIndexLevelsClassLevelFeaturesGet**](ClassLevelsApi.md#ApiClassesIndexLevelsClassLevelFeaturesGet) | **Get** /api/classes/{index}/levels/{class_level}/features | Get features available to a class at the requested level.
[**ApiClassesIndexLevelsClassLevelGet**](ClassLevelsApi.md#ApiClassesIndexLevelsClassLevelGet) | **Get** /api/classes/{index}/levels/{class_level} | Get level resource for a class and level.
[**ApiClassesIndexLevelsGet**](ClassLevelsApi.md#ApiClassesIndexLevelsGet) | **Get** /api/classes/{index}/levels | Get all level resources for a class.
[**ApiClassesIndexLevelsSpellLevelSpellsGet**](ClassLevelsApi.md#ApiClassesIndexLevelsSpellLevelSpellsGet) | **Get** /api/classes/{index}/levels/{spell_level}/spells | Get spells of the requested level available to the class.



## ApiClassesIndexLevelsClassLevelFeaturesGet

> APIReferenceList ApiClassesIndexLevelsClassLevelFeaturesGet(ctx, index, classLevel).Execute()

Get features available to a class at the requested level.

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
    classLevel := float32(3) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsClassLevelFeaturesGet(context.Background(), index, classLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassLevelsApi.ApiClassesIndexLevelsClassLevelFeaturesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexLevelsClassLevelFeaturesGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `ClassLevelsApi.ApiClassesIndexLevelsClassLevelFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 
**classLevel** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexLevelsClassLevelFeaturesGetRequest struct via the builder pattern


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


## ApiClassesIndexLevelsClassLevelGet

> ClassLevel ApiClassesIndexLevelsClassLevelGet(ctx, index, classLevel).Execute()

Get level resource for a class and level.

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
    classLevel := float32(3) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsClassLevelGet(context.Background(), index, classLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassLevelsApi.ApiClassesIndexLevelsClassLevelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexLevelsClassLevelGet`: ClassLevel
    fmt.Fprintf(os.Stdout, "Response from `ClassLevelsApi.ApiClassesIndexLevelsClassLevelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 
**classLevel** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexLevelsClassLevelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClassLevel**](ClassLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClassesIndexLevelsGet

> []ClassLevel ApiClassesIndexLevelsGet(ctx, index).Subclass(subclass).Execute()

Get all level resources for a class.

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
    subclass := "berserker" // string | Adds subclasses for class to the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsGet(context.Background(), index).Subclass(subclass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassLevelsApi.ApiClassesIndexLevelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexLevelsGet`: []ClassLevel
    fmt.Fprintf(os.Stdout, "Response from `ClassLevelsApi.ApiClassesIndexLevelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexLevelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subclass** | **string** | Adds subclasses for class to the response | 

### Return type

[**[]ClassLevel**](ClassLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClassesIndexLevelsSpellLevelSpellsGet

> APIReferenceList ApiClassesIndexLevelsSpellLevelSpellsGet(ctx, index, spellLevel).Execute()

Get spells of the requested level available to the class.

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
    spellLevel := float32(4) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassLevelsApi.ApiClassesIndexLevelsSpellLevelSpellsGet(context.Background(), index, spellLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassLevelsApi.ApiClassesIndexLevelsSpellLevelSpellsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiClassesIndexLevelsSpellLevelSpellsGet`: APIReferenceList
    fmt.Fprintf(os.Stdout, "Response from `ClassLevelsApi.ApiClassesIndexLevelsSpellLevelSpellsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the class to get.  | 
**spellLevel** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClassesIndexLevelsSpellLevelSpellsGetRequest struct via the builder pattern


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

