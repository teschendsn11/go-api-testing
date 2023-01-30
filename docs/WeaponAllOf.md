# WeaponAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**WeaponCategory** | Pointer to **string** | The category of weapon this falls into. | [optional] 
**WeaponRange** | Pointer to **string** | Whether this is a Melee or Ranged weapon. | [optional] 
**CategoryRange** | Pointer to **string** | A combination of weapon_category and weapon_range. | [optional] 
**Range** | Pointer to [**WeaponAllOfRange**](WeaponAllOfRange.md) |  | [optional] 
**Damage** | Pointer to [**Damage**](Damage.md) |  | [optional] 
**TwoHandedDamage** | Pointer to [**Damage**](Damage.md) |  | [optional] 
**Properties** | Pointer to [**[]APIReference**](APIReference.md) | A list of the properties this weapon has. | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Weight** | Pointer to **float32** | How much the equipment weighs. | [optional] 

## Methods

### NewWeaponAllOf

`func NewWeaponAllOf() *WeaponAllOf`

NewWeaponAllOf instantiates a new WeaponAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeaponAllOfWithDefaults

`func NewWeaponAllOfWithDefaults() *WeaponAllOf`

NewWeaponAllOfWithDefaults instantiates a new WeaponAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentCategory

`func (o *WeaponAllOf) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *WeaponAllOf) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *WeaponAllOf) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *WeaponAllOf) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetWeaponCategory

`func (o *WeaponAllOf) GetWeaponCategory() string`

GetWeaponCategory returns the WeaponCategory field if non-nil, zero value otherwise.

### GetWeaponCategoryOk

`func (o *WeaponAllOf) GetWeaponCategoryOk() (*string, bool)`

GetWeaponCategoryOk returns a tuple with the WeaponCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponCategory

`func (o *WeaponAllOf) SetWeaponCategory(v string)`

SetWeaponCategory sets WeaponCategory field to given value.

### HasWeaponCategory

`func (o *WeaponAllOf) HasWeaponCategory() bool`

HasWeaponCategory returns a boolean if a field has been set.

### GetWeaponRange

`func (o *WeaponAllOf) GetWeaponRange() string`

GetWeaponRange returns the WeaponRange field if non-nil, zero value otherwise.

### GetWeaponRangeOk

`func (o *WeaponAllOf) GetWeaponRangeOk() (*string, bool)`

GetWeaponRangeOk returns a tuple with the WeaponRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaponRange

`func (o *WeaponAllOf) SetWeaponRange(v string)`

SetWeaponRange sets WeaponRange field to given value.

### HasWeaponRange

`func (o *WeaponAllOf) HasWeaponRange() bool`

HasWeaponRange returns a boolean if a field has been set.

### GetCategoryRange

`func (o *WeaponAllOf) GetCategoryRange() string`

GetCategoryRange returns the CategoryRange field if non-nil, zero value otherwise.

### GetCategoryRangeOk

`func (o *WeaponAllOf) GetCategoryRangeOk() (*string, bool)`

GetCategoryRangeOk returns a tuple with the CategoryRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryRange

`func (o *WeaponAllOf) SetCategoryRange(v string)`

SetCategoryRange sets CategoryRange field to given value.

### HasCategoryRange

`func (o *WeaponAllOf) HasCategoryRange() bool`

HasCategoryRange returns a boolean if a field has been set.

### GetRange

`func (o *WeaponAllOf) GetRange() WeaponAllOfRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *WeaponAllOf) GetRangeOk() (*WeaponAllOfRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *WeaponAllOf) SetRange(v WeaponAllOfRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *WeaponAllOf) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetDamage

`func (o *WeaponAllOf) GetDamage() Damage`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *WeaponAllOf) GetDamageOk() (*Damage, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *WeaponAllOf) SetDamage(v Damage)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *WeaponAllOf) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetTwoHandedDamage

`func (o *WeaponAllOf) GetTwoHandedDamage() Damage`

GetTwoHandedDamage returns the TwoHandedDamage field if non-nil, zero value otherwise.

### GetTwoHandedDamageOk

`func (o *WeaponAllOf) GetTwoHandedDamageOk() (*Damage, bool)`

GetTwoHandedDamageOk returns a tuple with the TwoHandedDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoHandedDamage

`func (o *WeaponAllOf) SetTwoHandedDamage(v Damage)`

SetTwoHandedDamage sets TwoHandedDamage field to given value.

### HasTwoHandedDamage

`func (o *WeaponAllOf) HasTwoHandedDamage() bool`

HasTwoHandedDamage returns a boolean if a field has been set.

### GetProperties

`func (o *WeaponAllOf) GetProperties() []APIReference`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WeaponAllOf) GetPropertiesOk() (*[]APIReference, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WeaponAllOf) SetProperties(v []APIReference)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WeaponAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCost

`func (o *WeaponAllOf) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *WeaponAllOf) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *WeaponAllOf) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *WeaponAllOf) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetWeight

`func (o *WeaponAllOf) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WeaponAllOf) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WeaponAllOf) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WeaponAllOf) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


