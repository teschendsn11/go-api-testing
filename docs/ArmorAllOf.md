# ArmorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**ArmorCategory** | Pointer to **string** | The category of armor this falls into. | [optional] 
**ArmorClass** | Pointer to **map[string]string** | Details on how to calculate armor class. | [optional] 
**StrMinimum** | Pointer to **float32** | Minimum STR required to use this armor. | [optional] 
**StealthDisadvantage** | Pointer to **bool** | Whether the armor gives disadvantage for Stealth. | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Weight** | Pointer to **float32** | How much the equipment weighs. | [optional] 

## Methods

### NewArmorAllOf

`func NewArmorAllOf() *ArmorAllOf`

NewArmorAllOf instantiates a new ArmorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmorAllOfWithDefaults

`func NewArmorAllOfWithDefaults() *ArmorAllOf`

NewArmorAllOfWithDefaults instantiates a new ArmorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentCategory

`func (o *ArmorAllOf) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *ArmorAllOf) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *ArmorAllOf) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *ArmorAllOf) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetArmorCategory

`func (o *ArmorAllOf) GetArmorCategory() string`

GetArmorCategory returns the ArmorCategory field if non-nil, zero value otherwise.

### GetArmorCategoryOk

`func (o *ArmorAllOf) GetArmorCategoryOk() (*string, bool)`

GetArmorCategoryOk returns a tuple with the ArmorCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorCategory

`func (o *ArmorAllOf) SetArmorCategory(v string)`

SetArmorCategory sets ArmorCategory field to given value.

### HasArmorCategory

`func (o *ArmorAllOf) HasArmorCategory() bool`

HasArmorCategory returns a boolean if a field has been set.

### GetArmorClass

`func (o *ArmorAllOf) GetArmorClass() map[string]string`

GetArmorClass returns the ArmorClass field if non-nil, zero value otherwise.

### GetArmorClassOk

`func (o *ArmorAllOf) GetArmorClassOk() (*map[string]string, bool)`

GetArmorClassOk returns a tuple with the ArmorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorClass

`func (o *ArmorAllOf) SetArmorClass(v map[string]string)`

SetArmorClass sets ArmorClass field to given value.

### HasArmorClass

`func (o *ArmorAllOf) HasArmorClass() bool`

HasArmorClass returns a boolean if a field has been set.

### GetStrMinimum

`func (o *ArmorAllOf) GetStrMinimum() float32`

GetStrMinimum returns the StrMinimum field if non-nil, zero value otherwise.

### GetStrMinimumOk

`func (o *ArmorAllOf) GetStrMinimumOk() (*float32, bool)`

GetStrMinimumOk returns a tuple with the StrMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrMinimum

`func (o *ArmorAllOf) SetStrMinimum(v float32)`

SetStrMinimum sets StrMinimum field to given value.

### HasStrMinimum

`func (o *ArmorAllOf) HasStrMinimum() bool`

HasStrMinimum returns a boolean if a field has been set.

### GetStealthDisadvantage

`func (o *ArmorAllOf) GetStealthDisadvantage() bool`

GetStealthDisadvantage returns the StealthDisadvantage field if non-nil, zero value otherwise.

### GetStealthDisadvantageOk

`func (o *ArmorAllOf) GetStealthDisadvantageOk() (*bool, bool)`

GetStealthDisadvantageOk returns a tuple with the StealthDisadvantage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthDisadvantage

`func (o *ArmorAllOf) SetStealthDisadvantage(v bool)`

SetStealthDisadvantage sets StealthDisadvantage field to given value.

### HasStealthDisadvantage

`func (o *ArmorAllOf) HasStealthDisadvantage() bool`

HasStealthDisadvantage returns a boolean if a field has been set.

### GetCost

`func (o *ArmorAllOf) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ArmorAllOf) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ArmorAllOf) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ArmorAllOf) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetWeight

`func (o *ArmorAllOf) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ArmorAllOf) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ArmorAllOf) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ArmorAllOf) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


