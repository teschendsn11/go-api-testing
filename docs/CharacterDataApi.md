# \CharacterDataApi

All URIs are relative to *https://www.dnd5eapi.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAbilityScoresIndexGet**](CharacterDataApi.md#ApiAbilityScoresIndexGet) | **Get** /api/ability-scores/{index} | Get an ability score by index.
[**ApiAlignmentsIndexGet**](CharacterDataApi.md#ApiAlignmentsIndexGet) | **Get** /api/alignments/{index} | Get an alignment by index.
[**ApiBackgroundsIndexGet**](CharacterDataApi.md#ApiBackgroundsIndexGet) | **Get** /api/backgrounds/{index} | Get a background by index.
[**ApiLanguagesIndexGet**](CharacterDataApi.md#ApiLanguagesIndexGet) | **Get** /api/languages/{index} | Get a language by index.
[**ApiProficienciesIndexGet**](CharacterDataApi.md#ApiProficienciesIndexGet) | **Get** /api/proficiencies/{index} | Get a proficiency by index.
[**ApiSkillsIndexGet**](CharacterDataApi.md#ApiSkillsIndexGet) | **Get** /api/skills/{index} | Get a skill by index.



## ApiAbilityScoresIndexGet

> AbilityScore ApiAbilityScoresIndexGet(ctx, index).Execute()

Get an ability score by index.



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
    index := "cha" // string | The `index` of the ability score to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterDataApi.ApiAbilityScoresIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterDataApi.ApiAbilityScoresIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAbilityScoresIndexGet`: AbilityScore
    fmt.Fprintf(os.Stdout, "Response from `CharacterDataApi.ApiAbilityScoresIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the ability score to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAbilityScoresIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AbilityScore**](AbilityScore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAlignmentsIndexGet

> Alignment ApiAlignmentsIndexGet(ctx, index).Execute()

Get an alignment by index.



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
    index := "chaotic-neutral" // string | The `index` of the alignment to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterDataApi.ApiAlignmentsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterDataApi.ApiAlignmentsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAlignmentsIndexGet`: Alignment
    fmt.Fprintf(os.Stdout, "Response from `CharacterDataApi.ApiAlignmentsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the alignment to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAlignmentsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Alignment**](Alignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackgroundsIndexGet

> Background ApiBackgroundsIndexGet(ctx, index).Execute()

Get a background by index.



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
    index := "acolyte" // string | The `index` of the background to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterDataApi.ApiBackgroundsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterDataApi.ApiBackgroundsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackgroundsIndexGet`: Background
    fmt.Fprintf(os.Stdout, "Response from `CharacterDataApi.ApiBackgroundsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the background to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackgroundsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Background**](Background.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiLanguagesIndexGet

> Language ApiLanguagesIndexGet(ctx, index).Execute()

Get a language by index.



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
    index := "abyssal" // string | The `index` of the language to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterDataApi.ApiLanguagesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterDataApi.ApiLanguagesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiLanguagesIndexGet`: Language
    fmt.Fprintf(os.Stdout, "Response from `CharacterDataApi.ApiLanguagesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the language to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiLanguagesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Language**](Language.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProficienciesIndexGet

> Proficiency ApiProficienciesIndexGet(ctx, index).Execute()

Get a proficiency by index.



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
    index := "medium-armor" // string | The `index` of the proficiency to get.  Available values can be found in the [`ResourceList`](#get-/api/-endpoint-) for `proficiencies`. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterDataApi.ApiProficienciesIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterDataApi.ApiProficienciesIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProficienciesIndexGet`: Proficiency
    fmt.Fprintf(os.Stdout, "Response from `CharacterDataApi.ApiProficienciesIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the proficiency to get.  Available values can be found in the [&#x60;ResourceList&#x60;](#get-/api/-endpoint-) for &#x60;proficiencies&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProficienciesIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Proficiency**](Proficiency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSkillsIndexGet

> Skill ApiSkillsIndexGet(ctx, index).Execute()

Get a skill by index.



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
    index := "nature" // string | The `index` of the skill to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterDataApi.ApiSkillsIndexGet(context.Background(), index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterDataApi.ApiSkillsIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSkillsIndexGet`: Skill
    fmt.Fprintf(os.Stdout, "Response from `CharacterDataApi.ApiSkillsIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | The &#x60;index&#x60; of the skill to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSkillsIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Skill**](Skill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

