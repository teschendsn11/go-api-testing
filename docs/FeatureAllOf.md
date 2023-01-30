# FeatureAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **float32** | The level this feature is gained. | [optional] 
**Class** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Subclass** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Parent** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Prerequisites** | Pointer to [**[]AnyOfobjectobjectobject**](AnyOfobjectobjectobject.md) | The prerequisites for this feature. | [optional] 
**FeatureSpecific** | Pointer to **map[string]interface{}** | Information specific to this feature. | [optional] 

## Methods

### NewFeatureAllOf

`func NewFeatureAllOf() *FeatureAllOf`

NewFeatureAllOf instantiates a new FeatureAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureAllOfWithDefaults

`func NewFeatureAllOfWithDefaults() *FeatureAllOf`

NewFeatureAllOfWithDefaults instantiates a new FeatureAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *FeatureAllOf) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *FeatureAllOf) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *FeatureAllOf) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *FeatureAllOf) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetClass

`func (o *FeatureAllOf) GetClass() APIReference`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *FeatureAllOf) GetClassOk() (*APIReference, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *FeatureAllOf) SetClass(v APIReference)`

SetClass sets Class field to given value.

### HasClass

`func (o *FeatureAllOf) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSubclass

`func (o *FeatureAllOf) GetSubclass() APIReference`

GetSubclass returns the Subclass field if non-nil, zero value otherwise.

### GetSubclassOk

`func (o *FeatureAllOf) GetSubclassOk() (*APIReference, bool)`

GetSubclassOk returns a tuple with the Subclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclass

`func (o *FeatureAllOf) SetSubclass(v APIReference)`

SetSubclass sets Subclass field to given value.

### HasSubclass

`func (o *FeatureAllOf) HasSubclass() bool`

HasSubclass returns a boolean if a field has been set.

### GetParent

`func (o *FeatureAllOf) GetParent() APIReference`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FeatureAllOf) GetParentOk() (*APIReference, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FeatureAllOf) SetParent(v APIReference)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FeatureAllOf) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrerequisites

`func (o *FeatureAllOf) GetPrerequisites() []AnyOfobjectobjectobject`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *FeatureAllOf) GetPrerequisitesOk() (*[]AnyOfobjectobjectobject, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *FeatureAllOf) SetPrerequisites(v []AnyOfobjectobjectobject)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *FeatureAllOf) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.

### GetFeatureSpecific

`func (o *FeatureAllOf) GetFeatureSpecific() map[string]interface{}`

GetFeatureSpecific returns the FeatureSpecific field if non-nil, zero value otherwise.

### GetFeatureSpecificOk

`func (o *FeatureAllOf) GetFeatureSpecificOk() (*map[string]interface{}, bool)`

GetFeatureSpecificOk returns a tuple with the FeatureSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSpecific

`func (o *FeatureAllOf) SetFeatureSpecific(v map[string]interface{})`

SetFeatureSpecific sets FeatureSpecific field to given value.

### HasFeatureSpecific

`func (o *FeatureAllOf) HasFeatureSpecific() bool`

HasFeatureSpecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


