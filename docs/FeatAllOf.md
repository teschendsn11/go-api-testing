# FeatAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prerequisites** | Pointer to [**[]Prerequisite**](Prerequisite.md) | An object of APIReferences to ability scores and minimum scores. | [optional] 

## Methods

### NewFeatAllOf

`func NewFeatAllOf() *FeatAllOf`

NewFeatAllOf instantiates a new FeatAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatAllOfWithDefaults

`func NewFeatAllOfWithDefaults() *FeatAllOf`

NewFeatAllOfWithDefaults instantiates a new FeatAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrerequisites

`func (o *FeatAllOf) GetPrerequisites() []Prerequisite`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *FeatAllOf) GetPrerequisitesOk() (*[]Prerequisite, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *FeatAllOf) SetPrerequisites(v []Prerequisite)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *FeatAllOf) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


