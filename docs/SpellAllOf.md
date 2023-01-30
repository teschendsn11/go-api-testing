# SpellAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HigherLevel** | Pointer to **[]string** | List of descriptions for casting the spell at higher levels. | [optional] 
**Range** | Pointer to **string** | Range of the spell, usually expressed in feet. | [optional] 
**Components** | Pointer to **[]string** | List of shorthand for required components of the spell. V: verbal S: somatic M: material  | [optional] 
**Material** | Pointer to **string** | Material component for the spell to be cast. | [optional] 
**AreaOfEffect** | Pointer to [**AreaOfEffect**](AreaOfEffect.md) |  | [optional] 
**Ritual** | Pointer to **bool** | Determines if a spell can be cast in a 10-min(in-game) ritual. | [optional] 
**Duration** | Pointer to **string** | How long the spell effect lasts. | [optional] 
**Concentration** | Pointer to **bool** | Determines if a spell needs concentration to persist. | [optional] 
**CastingTime** | Pointer to **string** | How long it takes for the spell to activate. | [optional] 
**Level** | Pointer to **float32** | Level of the spell. | [optional] 
**AttackType** | Pointer to **string** | Attack type of the spell. | [optional] 
**Damage** | Pointer to [**NullableOneOfDamageAtCharacterLevelDamageAtSlotLevel**](oneOf&lt;DamageAtCharacterLevel,DamageAtSlotLevel&gt;.md) |  | [optional] 
**School** | Pointer to [**APIReference**](APIReference.md) |  | [optional] 
**Classes** | Pointer to [**[]APIReference**](APIReference.md) | List of classes that are able to learn the spell. | [optional] 
**Subclasses** | Pointer to [**[]APIReference**](APIReference.md) | List of subclasses that have access to the spell. | [optional] 

## Methods

### NewSpellAllOf

`func NewSpellAllOf() *SpellAllOf`

NewSpellAllOf instantiates a new SpellAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellAllOfWithDefaults

`func NewSpellAllOfWithDefaults() *SpellAllOf`

NewSpellAllOfWithDefaults instantiates a new SpellAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHigherLevel

`func (o *SpellAllOf) GetHigherLevel() []string`

GetHigherLevel returns the HigherLevel field if non-nil, zero value otherwise.

### GetHigherLevelOk

`func (o *SpellAllOf) GetHigherLevelOk() (*[]string, bool)`

GetHigherLevelOk returns a tuple with the HigherLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherLevel

`func (o *SpellAllOf) SetHigherLevel(v []string)`

SetHigherLevel sets HigherLevel field to given value.

### HasHigherLevel

`func (o *SpellAllOf) HasHigherLevel() bool`

HasHigherLevel returns a boolean if a field has been set.

### GetRange

`func (o *SpellAllOf) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SpellAllOf) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SpellAllOf) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *SpellAllOf) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetComponents

`func (o *SpellAllOf) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SpellAllOf) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SpellAllOf) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SpellAllOf) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetMaterial

`func (o *SpellAllOf) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *SpellAllOf) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *SpellAllOf) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *SpellAllOf) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetAreaOfEffect

`func (o *SpellAllOf) GetAreaOfEffect() AreaOfEffect`

GetAreaOfEffect returns the AreaOfEffect field if non-nil, zero value otherwise.

### GetAreaOfEffectOk

`func (o *SpellAllOf) GetAreaOfEffectOk() (*AreaOfEffect, bool)`

GetAreaOfEffectOk returns a tuple with the AreaOfEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaOfEffect

`func (o *SpellAllOf) SetAreaOfEffect(v AreaOfEffect)`

SetAreaOfEffect sets AreaOfEffect field to given value.

### HasAreaOfEffect

`func (o *SpellAllOf) HasAreaOfEffect() bool`

HasAreaOfEffect returns a boolean if a field has been set.

### GetRitual

`func (o *SpellAllOf) GetRitual() bool`

GetRitual returns the Ritual field if non-nil, zero value otherwise.

### GetRitualOk

`func (o *SpellAllOf) GetRitualOk() (*bool, bool)`

GetRitualOk returns a tuple with the Ritual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRitual

`func (o *SpellAllOf) SetRitual(v bool)`

SetRitual sets Ritual field to given value.

### HasRitual

`func (o *SpellAllOf) HasRitual() bool`

HasRitual returns a boolean if a field has been set.

### GetDuration

`func (o *SpellAllOf) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SpellAllOf) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SpellAllOf) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SpellAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetConcentration

`func (o *SpellAllOf) GetConcentration() bool`

GetConcentration returns the Concentration field if non-nil, zero value otherwise.

### GetConcentrationOk

`func (o *SpellAllOf) GetConcentrationOk() (*bool, bool)`

GetConcentrationOk returns a tuple with the Concentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcentration

`func (o *SpellAllOf) SetConcentration(v bool)`

SetConcentration sets Concentration field to given value.

### HasConcentration

`func (o *SpellAllOf) HasConcentration() bool`

HasConcentration returns a boolean if a field has been set.

### GetCastingTime

`func (o *SpellAllOf) GetCastingTime() string`

GetCastingTime returns the CastingTime field if non-nil, zero value otherwise.

### GetCastingTimeOk

`func (o *SpellAllOf) GetCastingTimeOk() (*string, bool)`

GetCastingTimeOk returns a tuple with the CastingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastingTime

`func (o *SpellAllOf) SetCastingTime(v string)`

SetCastingTime sets CastingTime field to given value.

### HasCastingTime

`func (o *SpellAllOf) HasCastingTime() bool`

HasCastingTime returns a boolean if a field has been set.

### GetLevel

`func (o *SpellAllOf) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SpellAllOf) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SpellAllOf) SetLevel(v float32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SpellAllOf) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetAttackType

`func (o *SpellAllOf) GetAttackType() string`

GetAttackType returns the AttackType field if non-nil, zero value otherwise.

### GetAttackTypeOk

`func (o *SpellAllOf) GetAttackTypeOk() (*string, bool)`

GetAttackTypeOk returns a tuple with the AttackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackType

`func (o *SpellAllOf) SetAttackType(v string)`

SetAttackType sets AttackType field to given value.

### HasAttackType

`func (o *SpellAllOf) HasAttackType() bool`

HasAttackType returns a boolean if a field has been set.

### GetDamage

`func (o *SpellAllOf) GetDamage() OneOfDamageAtCharacterLevelDamageAtSlotLevel`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *SpellAllOf) GetDamageOk() (*OneOfDamageAtCharacterLevelDamageAtSlotLevel, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *SpellAllOf) SetDamage(v OneOfDamageAtCharacterLevelDamageAtSlotLevel)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *SpellAllOf) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### SetDamageNil

`func (o *SpellAllOf) SetDamageNil(b bool)`

 SetDamageNil sets the value for Damage to be an explicit nil

### UnsetDamage
`func (o *SpellAllOf) UnsetDamage()`

UnsetDamage ensures that no value is present for Damage, not even an explicit nil
### GetSchool

`func (o *SpellAllOf) GetSchool() APIReference`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *SpellAllOf) GetSchoolOk() (*APIReference, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *SpellAllOf) SetSchool(v APIReference)`

SetSchool sets School field to given value.

### HasSchool

`func (o *SpellAllOf) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetClasses

`func (o *SpellAllOf) GetClasses() []APIReference`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *SpellAllOf) GetClassesOk() (*[]APIReference, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *SpellAllOf) SetClasses(v []APIReference)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *SpellAllOf) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetSubclasses

`func (o *SpellAllOf) GetSubclasses() []APIReference`

GetSubclasses returns the Subclasses field if non-nil, zero value otherwise.

### GetSubclassesOk

`func (o *SpellAllOf) GetSubclassesOk() (*[]APIReference, bool)`

GetSubclassesOk returns a tuple with the Subclasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclasses

`func (o *SpellAllOf) SetSubclasses(v []APIReference)`

SetSubclasses sets Subclasses field to given value.

### HasSubclasses

`func (o *SpellAllOf) HasSubclasses() bool`

HasSubclasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


