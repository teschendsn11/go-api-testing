# ProficiencyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The general category of the proficiency. | [optional] 
**Classes** | Pointer to [**[]APIReference**](APIReference.md) | Classes that start with this proficiency. | [optional] 
**Races** | Pointer to [**[]APIReference**](APIReference.md) | Races that start with this proficiency. | [optional] 
**Reference** | Pointer to [**APIReference**](APIReference.md) | &#x60;APIReference&#x60; to the full description of the related resource.  | [optional] 

## Methods

### NewProficiencyAllOf

`func NewProficiencyAllOf() *ProficiencyAllOf`

NewProficiencyAllOf instantiates a new ProficiencyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProficiencyAllOfWithDefaults

`func NewProficiencyAllOfWithDefaults() *ProficiencyAllOf`

NewProficiencyAllOfWithDefaults instantiates a new ProficiencyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProficiencyAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProficiencyAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProficiencyAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProficiencyAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClasses

`func (o *ProficiencyAllOf) GetClasses() []APIReference`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *ProficiencyAllOf) GetClassesOk() (*[]APIReference, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *ProficiencyAllOf) SetClasses(v []APIReference)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *ProficiencyAllOf) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetRaces

`func (o *ProficiencyAllOf) GetRaces() []APIReference`

GetRaces returns the Races field if non-nil, zero value otherwise.

### GetRacesOk

`func (o *ProficiencyAllOf) GetRacesOk() (*[]APIReference, bool)`

GetRacesOk returns a tuple with the Races field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaces

`func (o *ProficiencyAllOf) SetRaces(v []APIReference)`

SetRaces sets Races field to given value.

### HasRaces

`func (o *ProficiencyAllOf) HasRaces() bool`

HasRaces returns a boolean if a field has been set.

### GetReference

`func (o *ProficiencyAllOf) GetReference() APIReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ProficiencyAllOf) GetReferenceOk() (*APIReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ProficiencyAllOf) SetReference(v APIReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ProficiencyAllOf) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


