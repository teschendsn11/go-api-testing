# MonsterAllOf1Spellcasting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ability** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Dc** | Pointer to **float32** |  | [optional] 
**Modifier** | Pointer to **float32** |  | [optional] 
**ComponentsRequired** | Pointer to **[]string** |  | [optional] 
**School** | Pointer to **string** |  | [optional] 
**Slots** | Pointer to **map[string]float32** |  | [optional] 
**Spells** | Pointer to [**[]MonsterAllOf1SpellcastingSpells**](MonsterAllOf1SpellcastingSpells.md) |  | [optional] 

## Methods

### NewMonsterAllOf1Spellcasting

`func NewMonsterAllOf1Spellcasting() *MonsterAllOf1Spellcasting`

NewMonsterAllOf1Spellcasting instantiates a new MonsterAllOf1Spellcasting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterAllOf1SpellcastingWithDefaults

`func NewMonsterAllOf1SpellcastingWithDefaults() *MonsterAllOf1Spellcasting`

NewMonsterAllOf1SpellcastingWithDefaults instantiates a new MonsterAllOf1Spellcasting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbility

`func (o *MonsterAllOf1Spellcasting) GetAbility() APIReference`

GetAbility returns the Ability field if non-nil, zero value otherwise.

### GetAbilityOk

`func (o *MonsterAllOf1Spellcasting) GetAbilityOk() (*APIReference, bool)`

GetAbilityOk returns a tuple with the Ability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbility

`func (o *MonsterAllOf1Spellcasting) SetAbility(v APIReference)`

SetAbility sets Ability field to given value.

### HasAbility

`func (o *MonsterAllOf1Spellcasting) HasAbility() bool`

HasAbility returns a boolean if a field has been set.

### GetDc

`func (o *MonsterAllOf1Spellcasting) GetDc() float32`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *MonsterAllOf1Spellcasting) GetDcOk() (*float32, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *MonsterAllOf1Spellcasting) SetDc(v float32)`

SetDc sets Dc field to given value.

### HasDc

`func (o *MonsterAllOf1Spellcasting) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetModifier

`func (o *MonsterAllOf1Spellcasting) GetModifier() float32`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *MonsterAllOf1Spellcasting) GetModifierOk() (*float32, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *MonsterAllOf1Spellcasting) SetModifier(v float32)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *MonsterAllOf1Spellcasting) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### GetComponentsRequired

`func (o *MonsterAllOf1Spellcasting) GetComponentsRequired() []string`

GetComponentsRequired returns the ComponentsRequired field if non-nil, zero value otherwise.

### GetComponentsRequiredOk

`func (o *MonsterAllOf1Spellcasting) GetComponentsRequiredOk() (*[]string, bool)`

GetComponentsRequiredOk returns a tuple with the ComponentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentsRequired

`func (o *MonsterAllOf1Spellcasting) SetComponentsRequired(v []string)`

SetComponentsRequired sets ComponentsRequired field to given value.

### HasComponentsRequired

`func (o *MonsterAllOf1Spellcasting) HasComponentsRequired() bool`

HasComponentsRequired returns a boolean if a field has been set.

### GetSchool

`func (o *MonsterAllOf1Spellcasting) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *MonsterAllOf1Spellcasting) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *MonsterAllOf1Spellcasting) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *MonsterAllOf1Spellcasting) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetSlots

`func (o *MonsterAllOf1Spellcasting) GetSlots() map[string]float32`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *MonsterAllOf1Spellcasting) GetSlotsOk() (*map[string]float32, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *MonsterAllOf1Spellcasting) SetSlots(v map[string]float32)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *MonsterAllOf1Spellcasting) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetSpells

`func (o *MonsterAllOf1Spellcasting) GetSpells() []MonsterAllOf1SpellcastingSpells`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *MonsterAllOf1Spellcasting) GetSpellsOk() (*[]MonsterAllOf1SpellcastingSpells, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *MonsterAllOf1Spellcasting) SetSpells(v []MonsterAllOf1SpellcastingSpells)`

SetSpells sets Spells field to given value.

### HasSpells

`func (o *MonsterAllOf1Spellcasting) HasSpells() bool`

HasSpells returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


