# GearAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**GearCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Weight** | Pointer to **float32** | How much the equipment weighs. | [optional] 

## Methods

### NewGearAllOf

`func NewGearAllOf() *GearAllOf`

NewGearAllOf instantiates a new GearAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGearAllOfWithDefaults

`func NewGearAllOfWithDefaults() *GearAllOf`

NewGearAllOfWithDefaults instantiates a new GearAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentCategory

`func (o *GearAllOf) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *GearAllOf) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *GearAllOf) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *GearAllOf) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetGearCategory

`func (o *GearAllOf) GetGearCategory() APIReference`

GetGearCategory returns the GearCategory field if non-nil, zero value otherwise.

### GetGearCategoryOk

`func (o *GearAllOf) GetGearCategoryOk() (*APIReference, bool)`

GetGearCategoryOk returns a tuple with the GearCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearCategory

`func (o *GearAllOf) SetGearCategory(v APIReference)`

SetGearCategory sets GearCategory field to given value.

### HasGearCategory

`func (o *GearAllOf) HasGearCategory() bool`

HasGearCategory returns a boolean if a field has been set.

### GetCost

`func (o *GearAllOf) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GearAllOf) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GearAllOf) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GearAllOf) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetWeight

`func (o *GearAllOf) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GearAllOf) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GearAllOf) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GearAllOf) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


