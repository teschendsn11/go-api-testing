# SubclassAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**SubclassFlavor** | Pointer to **string** | Lore-friendly flavor text for a classes respective subclass. | [optional] 
**SubclassLevels** | Pointer to **string** | Resource url that shows the subclass level progression. | [optional] 
**Spells** | Pointer to [**[]SubclassAllOfSpells**](SubclassAllOfSpells.md) |  | [optional] 

## Methods

### NewSubclassAllOf

`func NewSubclassAllOf() *SubclassAllOf`

NewSubclassAllOf instantiates a new SubclassAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubclassAllOfWithDefaults

`func NewSubclassAllOfWithDefaults() *SubclassAllOf`

NewSubclassAllOfWithDefaults instantiates a new SubclassAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *SubclassAllOf) GetClass() APIReference`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *SubclassAllOf) GetClassOk() (*APIReference, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *SubclassAllOf) SetClass(v APIReference)`

SetClass sets Class field to given value.

### HasClass

`func (o *SubclassAllOf) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSubclassFlavor

`func (o *SubclassAllOf) GetSubclassFlavor() string`

GetSubclassFlavor returns the SubclassFlavor field if non-nil, zero value otherwise.

### GetSubclassFlavorOk

`func (o *SubclassAllOf) GetSubclassFlavorOk() (*string, bool)`

GetSubclassFlavorOk returns a tuple with the SubclassFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclassFlavor

`func (o *SubclassAllOf) SetSubclassFlavor(v string)`

SetSubclassFlavor sets SubclassFlavor field to given value.

### HasSubclassFlavor

`func (o *SubclassAllOf) HasSubclassFlavor() bool`

HasSubclassFlavor returns a boolean if a field has been set.

### GetSubclassLevels

`func (o *SubclassAllOf) GetSubclassLevels() string`

GetSubclassLevels returns the SubclassLevels field if non-nil, zero value otherwise.

### GetSubclassLevelsOk

`func (o *SubclassAllOf) GetSubclassLevelsOk() (*string, bool)`

GetSubclassLevelsOk returns a tuple with the SubclassLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclassLevels

`func (o *SubclassAllOf) SetSubclassLevels(v string)`

SetSubclassLevels sets SubclassLevels field to given value.

### HasSubclassLevels

`func (o *SubclassAllOf) HasSubclassLevels() bool`

HasSubclassLevels returns a boolean if a field has been set.

### GetSpells

`func (o *SubclassAllOf) GetSpells() []SubclassAllOfSpells`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *SubclassAllOf) GetSpellsOk() (*[]SubclassAllOfSpells, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *SubclassAllOf) SetSpells(v []SubclassAllOfSpells)`

SetSpells sets Spells field to given value.

### HasSpells

`func (o *SubclassAllOf) HasSpells() bool`

HasSpells returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


