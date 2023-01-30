# EquipmentPackAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**GearCategory** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Contents** | Pointer to [**[]APIReference**](APIReference.md) | The list of adventuring gear in the pack. | [optional] 

## Methods

### NewEquipmentPackAllOf

`func NewEquipmentPackAllOf() *EquipmentPackAllOf`

NewEquipmentPackAllOf instantiates a new EquipmentPackAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPackAllOfWithDefaults

`func NewEquipmentPackAllOfWithDefaults() *EquipmentPackAllOf`

NewEquipmentPackAllOfWithDefaults instantiates a new EquipmentPackAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentCategory

`func (o *EquipmentPackAllOf) GetEquipmentCategory() APIReference`

GetEquipmentCategory returns the EquipmentCategory field if non-nil, zero value otherwise.

### GetEquipmentCategoryOk

`func (o *EquipmentPackAllOf) GetEquipmentCategoryOk() (*APIReference, bool)`

GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentCategory

`func (o *EquipmentPackAllOf) SetEquipmentCategory(v APIReference)`

SetEquipmentCategory sets EquipmentCategory field to given value.

### HasEquipmentCategory

`func (o *EquipmentPackAllOf) HasEquipmentCategory() bool`

HasEquipmentCategory returns a boolean if a field has been set.

### GetGearCategory

`func (o *EquipmentPackAllOf) GetGearCategory() APIReference`

GetGearCategory returns the GearCategory field if non-nil, zero value otherwise.

### GetGearCategoryOk

`func (o *EquipmentPackAllOf) GetGearCategoryOk() (*APIReference, bool)`

GetGearCategoryOk returns a tuple with the GearCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearCategory

`func (o *EquipmentPackAllOf) SetGearCategory(v APIReference)`

SetGearCategory sets GearCategory field to given value.

### HasGearCategory

`func (o *EquipmentPackAllOf) HasGearCategory() bool`

HasGearCategory returns a boolean if a field has been set.

### GetCost

`func (o *EquipmentPackAllOf) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *EquipmentPackAllOf) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *EquipmentPackAllOf) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *EquipmentPackAllOf) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetContents

`func (o *EquipmentPackAllOf) GetContents() []APIReference`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *EquipmentPackAllOf) GetContentsOk() (*[]APIReference, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *EquipmentPackAllOf) SetContents(v []APIReference)`

SetContents sets Contents field to given value.

### HasContents

`func (o *EquipmentPackAllOf) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


