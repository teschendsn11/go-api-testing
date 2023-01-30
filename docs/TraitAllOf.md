# TraitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Races** | Pointer to [**[]APIReference**](APIReference.md) | List of &#x60;Races&#x60; that have access to the trait. | [optional] 
**Subraces** | Pointer to [**[]APIReference**](APIReference.md) | List of &#x60;Subraces&#x60; that have access to the trait. | [optional] 
**Proficiencies** | Pointer to [**[]APIReference**](APIReference.md) | List of &#x60;Proficiencies&#x60; this trait grants. | [optional] 
**ProficiencyChoices** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**LanguageOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**TraitSpecific** | Pointer to [**NullableOneOfChoiceChoiceobject**](oneOf&lt;Choice,Choice,object&gt;.md) | Information specific to this trait | [optional] 

## Methods

### NewTraitAllOf

`func NewTraitAllOf() *TraitAllOf`

NewTraitAllOf instantiates a new TraitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraitAllOfWithDefaults

`func NewTraitAllOfWithDefaults() *TraitAllOf`

NewTraitAllOfWithDefaults instantiates a new TraitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaces

`func (o *TraitAllOf) GetRaces() []APIReference`

GetRaces returns the Races field if non-nil, zero value otherwise.

### GetRacesOk

`func (o *TraitAllOf) GetRacesOk() (*[]APIReference, bool)`

GetRacesOk returns a tuple with the Races field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaces

`func (o *TraitAllOf) SetRaces(v []APIReference)`

SetRaces sets Races field to given value.

### HasRaces

`func (o *TraitAllOf) HasRaces() bool`

HasRaces returns a boolean if a field has been set.

### GetSubraces

`func (o *TraitAllOf) GetSubraces() []APIReference`

GetSubraces returns the Subraces field if non-nil, zero value otherwise.

### GetSubracesOk

`func (o *TraitAllOf) GetSubracesOk() (*[]APIReference, bool)`

GetSubracesOk returns a tuple with the Subraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubraces

`func (o *TraitAllOf) SetSubraces(v []APIReference)`

SetSubraces sets Subraces field to given value.

### HasSubraces

`func (o *TraitAllOf) HasSubraces() bool`

HasSubraces returns a boolean if a field has been set.

### GetProficiencies

`func (o *TraitAllOf) GetProficiencies() []APIReference`

GetProficiencies returns the Proficiencies field if non-nil, zero value otherwise.

### GetProficienciesOk

`func (o *TraitAllOf) GetProficienciesOk() (*[]APIReference, bool)`

GetProficienciesOk returns a tuple with the Proficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencies

`func (o *TraitAllOf) SetProficiencies(v []APIReference)`

SetProficiencies sets Proficiencies field to given value.

### HasProficiencies

`func (o *TraitAllOf) HasProficiencies() bool`

HasProficiencies returns a boolean if a field has been set.

### GetProficiencyChoices

`func (o *TraitAllOf) GetProficiencyChoices() Choice`

GetProficiencyChoices returns the ProficiencyChoices field if non-nil, zero value otherwise.

### GetProficiencyChoicesOk

`func (o *TraitAllOf) GetProficiencyChoicesOk() (*Choice, bool)`

GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyChoices

`func (o *TraitAllOf) SetProficiencyChoices(v Choice)`

SetProficiencyChoices sets ProficiencyChoices field to given value.

### HasProficiencyChoices

`func (o *TraitAllOf) HasProficiencyChoices() bool`

HasProficiencyChoices returns a boolean if a field has been set.

### GetLanguageOptions

`func (o *TraitAllOf) GetLanguageOptions() Choice`

GetLanguageOptions returns the LanguageOptions field if non-nil, zero value otherwise.

### GetLanguageOptionsOk

`func (o *TraitAllOf) GetLanguageOptionsOk() (*Choice, bool)`

GetLanguageOptionsOk returns a tuple with the LanguageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOptions

`func (o *TraitAllOf) SetLanguageOptions(v Choice)`

SetLanguageOptions sets LanguageOptions field to given value.

### HasLanguageOptions

`func (o *TraitAllOf) HasLanguageOptions() bool`

HasLanguageOptions returns a boolean if a field has been set.

### GetTraitSpecific

`func (o *TraitAllOf) GetTraitSpecific() OneOfChoiceChoiceobject`

GetTraitSpecific returns the TraitSpecific field if non-nil, zero value otherwise.

### GetTraitSpecificOk

`func (o *TraitAllOf) GetTraitSpecificOk() (*OneOfChoiceChoiceobject, bool)`

GetTraitSpecificOk returns a tuple with the TraitSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitSpecific

`func (o *TraitAllOf) SetTraitSpecific(v OneOfChoiceChoiceobject)`

SetTraitSpecific sets TraitSpecific field to given value.

### HasTraitSpecific

`func (o *TraitAllOf) HasTraitSpecific() bool`

HasTraitSpecific returns a boolean if a field has been set.

### SetTraitSpecificNil

`func (o *TraitAllOf) SetTraitSpecificNil(b bool)`

 SetTraitSpecificNil sets the value for TraitSpecific to be an explicit nil

### UnsetTraitSpecific
`func (o *TraitAllOf) UnsetTraitSpecific()`

UnsetTraitSpecific ensures that no value is present for TraitSpecific, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


