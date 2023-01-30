# RaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Speed** | Pointer to **float32** | Base move speed for this race (in feet per round). | [optional] 
**AbilityBonuses** | Pointer to [**[]AbilityBonus**](AbilityBonus.md) | Racial bonuses to ability scores. | [optional] 
**Alignment** | Pointer to **string** | Flavor description of likely alignments this race takes. | [optional] 
**Age** | Pointer to **string** | Flavor description of possible ages for this race. | [optional] 
**Size** | Pointer to **string** | Size class of this race. | [optional] 
**SizeDescription** | Pointer to **string** | Flavor description of height and weight for this race. | [optional] 
**StartingProficiencies** | Pointer to [**[]APIReference**](APIReference.md) | Starting proficiencies for all new characters of this race. | [optional] 
**StartingProficiencyOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Languages** | Pointer to [**[]APIReference**](APIReference.md) | Starting languages for all new characters of this race. | [optional] 
**LanguageDesc** | Pointer to **string** | Flavor description of the languages this race knows. | [optional] 
**Traits** | Pointer to [**[]APIReference**](APIReference.md) | Racial traits that provide benefits to its members. | [optional] 
**Subraces** | Pointer to [**[]APIReference**](APIReference.md) | All possible subraces that this race includes. | [optional] 

## Methods

### NewRaceAllOf

`func NewRaceAllOf() *RaceAllOf`

NewRaceAllOf instantiates a new RaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaceAllOfWithDefaults

`func NewRaceAllOfWithDefaults() *RaceAllOf`

NewRaceAllOfWithDefaults instantiates a new RaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpeed

`func (o *RaceAllOf) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *RaceAllOf) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *RaceAllOf) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *RaceAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetAbilityBonuses

`func (o *RaceAllOf) GetAbilityBonuses() []AbilityBonus`

GetAbilityBonuses returns the AbilityBonuses field if non-nil, zero value otherwise.

### GetAbilityBonusesOk

`func (o *RaceAllOf) GetAbilityBonusesOk() (*[]AbilityBonus, bool)`

GetAbilityBonusesOk returns a tuple with the AbilityBonuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityBonuses

`func (o *RaceAllOf) SetAbilityBonuses(v []AbilityBonus)`

SetAbilityBonuses sets AbilityBonuses field to given value.

### HasAbilityBonuses

`func (o *RaceAllOf) HasAbilityBonuses() bool`

HasAbilityBonuses returns a boolean if a field has been set.

### GetAlignment

`func (o *RaceAllOf) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *RaceAllOf) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *RaceAllOf) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *RaceAllOf) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetAge

`func (o *RaceAllOf) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RaceAllOf) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RaceAllOf) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *RaceAllOf) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetSize

`func (o *RaceAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RaceAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RaceAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *RaceAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeDescription

`func (o *RaceAllOf) GetSizeDescription() string`

GetSizeDescription returns the SizeDescription field if non-nil, zero value otherwise.

### GetSizeDescriptionOk

`func (o *RaceAllOf) GetSizeDescriptionOk() (*string, bool)`

GetSizeDescriptionOk returns a tuple with the SizeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeDescription

`func (o *RaceAllOf) SetSizeDescription(v string)`

SetSizeDescription sets SizeDescription field to given value.

### HasSizeDescription

`func (o *RaceAllOf) HasSizeDescription() bool`

HasSizeDescription returns a boolean if a field has been set.

### GetStartingProficiencies

`func (o *RaceAllOf) GetStartingProficiencies() []APIReference`

GetStartingProficiencies returns the StartingProficiencies field if non-nil, zero value otherwise.

### GetStartingProficienciesOk

`func (o *RaceAllOf) GetStartingProficienciesOk() (*[]APIReference, bool)`

GetStartingProficienciesOk returns a tuple with the StartingProficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencies

`func (o *RaceAllOf) SetStartingProficiencies(v []APIReference)`

SetStartingProficiencies sets StartingProficiencies field to given value.

### HasStartingProficiencies

`func (o *RaceAllOf) HasStartingProficiencies() bool`

HasStartingProficiencies returns a boolean if a field has been set.

### GetStartingProficiencyOptions

`func (o *RaceAllOf) GetStartingProficiencyOptions() Choice`

GetStartingProficiencyOptions returns the StartingProficiencyOptions field if non-nil, zero value otherwise.

### GetStartingProficiencyOptionsOk

`func (o *RaceAllOf) GetStartingProficiencyOptionsOk() (*Choice, bool)`

GetStartingProficiencyOptionsOk returns a tuple with the StartingProficiencyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencyOptions

`func (o *RaceAllOf) SetStartingProficiencyOptions(v Choice)`

SetStartingProficiencyOptions sets StartingProficiencyOptions field to given value.

### HasStartingProficiencyOptions

`func (o *RaceAllOf) HasStartingProficiencyOptions() bool`

HasStartingProficiencyOptions returns a boolean if a field has been set.

### GetLanguages

`func (o *RaceAllOf) GetLanguages() []APIReference`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *RaceAllOf) GetLanguagesOk() (*[]APIReference, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *RaceAllOf) SetLanguages(v []APIReference)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *RaceAllOf) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguageDesc

`func (o *RaceAllOf) GetLanguageDesc() string`

GetLanguageDesc returns the LanguageDesc field if non-nil, zero value otherwise.

### GetLanguageDescOk

`func (o *RaceAllOf) GetLanguageDescOk() (*string, bool)`

GetLanguageDescOk returns a tuple with the LanguageDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageDesc

`func (o *RaceAllOf) SetLanguageDesc(v string)`

SetLanguageDesc sets LanguageDesc field to given value.

### HasLanguageDesc

`func (o *RaceAllOf) HasLanguageDesc() bool`

HasLanguageDesc returns a boolean if a field has been set.

### GetTraits

`func (o *RaceAllOf) GetTraits() []APIReference`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *RaceAllOf) GetTraitsOk() (*[]APIReference, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *RaceAllOf) SetTraits(v []APIReference)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *RaceAllOf) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetSubraces

`func (o *RaceAllOf) GetSubraces() []APIReference`

GetSubraces returns the Subraces field if non-nil, zero value otherwise.

### GetSubracesOk

`func (o *RaceAllOf) GetSubracesOk() (*[]APIReference, bool)`

GetSubracesOk returns a tuple with the Subraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubraces

`func (o *RaceAllOf) SetSubraces(v []APIReference)`

SetSubraces sets Subraces field to given value.

### HasSubraces

`func (o *RaceAllOf) HasSubraces() bool`

HasSubraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


