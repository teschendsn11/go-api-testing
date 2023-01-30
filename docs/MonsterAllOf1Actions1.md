# MonsterAllOf1Actions1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Desc** | Pointer to **string** |  | [optional] 
**ActionOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Actions** | Pointer to [**[]MonsterAllOf1Actions**](MonsterAllOf1Actions.md) |  | [optional] 
**Options** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**MultiattackType** | Pointer to **string** |  | [optional] 
**AttackBonus** | Pointer to **float32** |  | [optional] 
**Dc** | Pointer to [**DC**](DC.md) |  | [optional] 
**Attacks** | Pointer to [**[]MonsterAllOf1Attacks**](MonsterAllOf1Attacks.md) |  | [optional] 
**Damage** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewMonsterAllOf1Actions1

`func NewMonsterAllOf1Actions1() *MonsterAllOf1Actions1`

NewMonsterAllOf1Actions1 instantiates a new MonsterAllOf1Actions1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterAllOf1Actions1WithDefaults

`func NewMonsterAllOf1Actions1WithDefaults() *MonsterAllOf1Actions1`

NewMonsterAllOf1Actions1WithDefaults instantiates a new MonsterAllOf1Actions1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonsterAllOf1Actions1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterAllOf1Actions1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterAllOf1Actions1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonsterAllOf1Actions1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesc

`func (o *MonsterAllOf1Actions1) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *MonsterAllOf1Actions1) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *MonsterAllOf1Actions1) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *MonsterAllOf1Actions1) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetActionOptions

`func (o *MonsterAllOf1Actions1) GetActionOptions() Choice`

GetActionOptions returns the ActionOptions field if non-nil, zero value otherwise.

### GetActionOptionsOk

`func (o *MonsterAllOf1Actions1) GetActionOptionsOk() (*Choice, bool)`

GetActionOptionsOk returns a tuple with the ActionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOptions

`func (o *MonsterAllOf1Actions1) SetActionOptions(v Choice)`

SetActionOptions sets ActionOptions field to given value.

### HasActionOptions

`func (o *MonsterAllOf1Actions1) HasActionOptions() bool`

HasActionOptions returns a boolean if a field has been set.

### GetActions

`func (o *MonsterAllOf1Actions1) GetActions() []MonsterAllOf1Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MonsterAllOf1Actions1) GetActionsOk() (*[]MonsterAllOf1Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MonsterAllOf1Actions1) SetActions(v []MonsterAllOf1Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MonsterAllOf1Actions1) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetOptions

`func (o *MonsterAllOf1Actions1) GetOptions() Choice`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MonsterAllOf1Actions1) GetOptionsOk() (*Choice, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MonsterAllOf1Actions1) SetOptions(v Choice)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MonsterAllOf1Actions1) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetMultiattackType

`func (o *MonsterAllOf1Actions1) GetMultiattackType() string`

GetMultiattackType returns the MultiattackType field if non-nil, zero value otherwise.

### GetMultiattackTypeOk

`func (o *MonsterAllOf1Actions1) GetMultiattackTypeOk() (*string, bool)`

GetMultiattackTypeOk returns a tuple with the MultiattackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiattackType

`func (o *MonsterAllOf1Actions1) SetMultiattackType(v string)`

SetMultiattackType sets MultiattackType field to given value.

### HasMultiattackType

`func (o *MonsterAllOf1Actions1) HasMultiattackType() bool`

HasMultiattackType returns a boolean if a field has been set.

### GetAttackBonus

`func (o *MonsterAllOf1Actions1) GetAttackBonus() float32`

GetAttackBonus returns the AttackBonus field if non-nil, zero value otherwise.

### GetAttackBonusOk

`func (o *MonsterAllOf1Actions1) GetAttackBonusOk() (*float32, bool)`

GetAttackBonusOk returns a tuple with the AttackBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBonus

`func (o *MonsterAllOf1Actions1) SetAttackBonus(v float32)`

SetAttackBonus sets AttackBonus field to given value.

### HasAttackBonus

`func (o *MonsterAllOf1Actions1) HasAttackBonus() bool`

HasAttackBonus returns a boolean if a field has been set.

### GetDc

`func (o *MonsterAllOf1Actions1) GetDc() DC`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *MonsterAllOf1Actions1) GetDcOk() (*DC, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *MonsterAllOf1Actions1) SetDc(v DC)`

SetDc sets Dc field to given value.

### HasDc

`func (o *MonsterAllOf1Actions1) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetAttacks

`func (o *MonsterAllOf1Actions1) GetAttacks() []MonsterAllOf1Attacks`

GetAttacks returns the Attacks field if non-nil, zero value otherwise.

### GetAttacksOk

`func (o *MonsterAllOf1Actions1) GetAttacksOk() (*[]MonsterAllOf1Attacks, bool)`

GetAttacksOk returns a tuple with the Attacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttacks

`func (o *MonsterAllOf1Actions1) SetAttacks(v []MonsterAllOf1Attacks)`

SetAttacks sets Attacks field to given value.

### HasAttacks

`func (o *MonsterAllOf1Actions1) HasAttacks() bool`

HasAttacks returns a boolean if a field has been set.

### GetDamage

`func (o *MonsterAllOf1Actions1) GetDamage() []interface{}`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *MonsterAllOf1Actions1) GetDamageOk() (*[]interface{}, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *MonsterAllOf1Actions1) SetDamage(v []interface{})`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *MonsterAllOf1Actions1) HasDamage() bool`

HasDamage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


