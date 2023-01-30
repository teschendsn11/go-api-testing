# SubraceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Description of the subrace. | [optional] 
**Race** | Pointer to [**APIReference**](APIReference.md) | Parent race for the subrace. | [optional] 
**AbilityBonuses** | Pointer to [**[]AbilityBonus**](AbilityBonus.md) | Additional ability bonuses for the subrace. | [optional] 
**StartingProficiencies** | Pointer to [**[]APIReference**](APIReference.md) | Starting proficiencies for all new characters of the subrace. | [optional] 
**Languages** | Pointer to [**[]APIReference**](APIReference.md) | Starting languages for all new characters of the subrace. | [optional] 
**LanguageOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**RacialTraits** | Pointer to [**[]APIReference**](APIReference.md) | List of traits that for the subrace. | [optional] 

## Methods

### NewSubraceAllOf

`func NewSubraceAllOf() *SubraceAllOf`

NewSubraceAllOf instantiates a new SubraceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubraceAllOfWithDefaults

`func NewSubraceAllOfWithDefaults() *SubraceAllOf`

NewSubraceAllOfWithDefaults instantiates a new SubraceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *SubraceAllOf) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SubraceAllOf) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SubraceAllOf) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *SubraceAllOf) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetRace

`func (o *SubraceAllOf) GetRace() APIReference`

GetRace returns the Race field if non-nil, zero value otherwise.

### GetRaceOk

`func (o *SubraceAllOf) GetRaceOk() (*APIReference, bool)`

GetRaceOk returns a tuple with the Race field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRace

`func (o *SubraceAllOf) SetRace(v APIReference)`

SetRace sets Race field to given value.

### HasRace

`func (o *SubraceAllOf) HasRace() bool`

HasRace returns a boolean if a field has been set.

### GetAbilityBonuses

`func (o *SubraceAllOf) GetAbilityBonuses() []AbilityBonus`

GetAbilityBonuses returns the AbilityBonuses field if non-nil, zero value otherwise.

### GetAbilityBonusesOk

`func (o *SubraceAllOf) GetAbilityBonusesOk() (*[]AbilityBonus, bool)`

GetAbilityBonusesOk returns a tuple with the AbilityBonuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityBonuses

`func (o *SubraceAllOf) SetAbilityBonuses(v []AbilityBonus)`

SetAbilityBonuses sets AbilityBonuses field to given value.

### HasAbilityBonuses

`func (o *SubraceAllOf) HasAbilityBonuses() bool`

HasAbilityBonuses returns a boolean if a field has been set.

### GetStartingProficiencies

`func (o *SubraceAllOf) GetStartingProficiencies() []APIReference`

GetStartingProficiencies returns the StartingProficiencies field if non-nil, zero value otherwise.

### GetStartingProficienciesOk

`func (o *SubraceAllOf) GetStartingProficienciesOk() (*[]APIReference, bool)`

GetStartingProficienciesOk returns a tuple with the StartingProficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencies

`func (o *SubraceAllOf) SetStartingProficiencies(v []APIReference)`

SetStartingProficiencies sets StartingProficiencies field to given value.

### HasStartingProficiencies

`func (o *SubraceAllOf) HasStartingProficiencies() bool`

HasStartingProficiencies returns a boolean if a field has been set.

### GetLanguages

`func (o *SubraceAllOf) GetLanguages() []APIReference`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *SubraceAllOf) GetLanguagesOk() (*[]APIReference, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *SubraceAllOf) SetLanguages(v []APIReference)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *SubraceAllOf) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguageOptions

`func (o *SubraceAllOf) GetLanguageOptions() Choice`

GetLanguageOptions returns the LanguageOptions field if non-nil, zero value otherwise.

### GetLanguageOptionsOk

`func (o *SubraceAllOf) GetLanguageOptionsOk() (*Choice, bool)`

GetLanguageOptionsOk returns a tuple with the LanguageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOptions

`func (o *SubraceAllOf) SetLanguageOptions(v Choice)`

SetLanguageOptions sets LanguageOptions field to given value.

### HasLanguageOptions

`func (o *SubraceAllOf) HasLanguageOptions() bool`

HasLanguageOptions returns a boolean if a field has been set.

### GetRacialTraits

`func (o *SubraceAllOf) GetRacialTraits() []APIReference`

GetRacialTraits returns the RacialTraits field if non-nil, zero value otherwise.

### GetRacialTraitsOk

`func (o *SubraceAllOf) GetRacialTraitsOk() (*[]APIReference, bool)`

GetRacialTraitsOk returns a tuple with the RacialTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacialTraits

`func (o *SubraceAllOf) SetRacialTraits(v []APIReference)`

SetRacialTraits sets RacialTraits field to given value.

### HasRacialTraits

`func (o *SubraceAllOf) HasRacialTraits() bool`

HasRacialTraits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


