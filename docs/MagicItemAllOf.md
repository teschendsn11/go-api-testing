# MagicItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Rarity** | Pointer to [**MagicItemAllOfRarity**](MagicItemAllOfRarity.md) |  | [optional] 
**Variants** | Pointer to [**[]APIReference**](APIReference.md) |  | [optional] 
**Variant** | Pointer to **bool** | Whether this is a variant or not | [optional] 

## Methods

### NewMagicItemAllOf

`func NewMagicItemAllOf() *MagicItemAllOf`

NewMagicItemAllOf instantiates a new MagicItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagicItemAllOfWithDefaults

`func NewMagicItemAllOfWithDefaults() *MagicItemAllOf`

NewMagicItemAllOfWithDefaults instantiates a new MagicItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentCategory

`func (o *MagicItemAllOf) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *MagicItemAllOf) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *MagicItemAllOf) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *MagicItemAllOf) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetRarity

`func (o *MagicItemAllOf) GetRarity() MagicItemAllOfRarity`

GetRarity returns the Rarity field if non-nil, zero value otherwise.

### GetRarityOk

`func (o *MagicItemAllOf) GetRarityOk() (*MagicItemAllOfRarity, bool)`

GetRarityOk returns a tuple with the Rarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarity

`func (o *MagicItemAllOf) SetRarity(v MagicItemAllOfRarity)`

SetRarity sets Rarity field to given value.

### HasRarity

`func (o *MagicItemAllOf) HasRarity() bool`

HasRarity returns a boolean if a field has been set.

### GetVariants

`func (o *MagicItemAllOf) GetVariants() []APIReference`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *MagicItemAllOf) GetVariantsOk() (*[]APIReference, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *MagicItemAllOf) SetVariants(v []APIReference)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *MagicItemAllOf) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetVariant

`func (o *MagicItemAllOf) GetVariant() bool`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *MagicItemAllOf) GetVariantOk() (*bool, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *MagicItemAllOf) SetVariant(v bool)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *MagicItemAllOf) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


