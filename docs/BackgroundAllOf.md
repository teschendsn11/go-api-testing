# BackgroundAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartingProficiencies** | Pointer to [**[]APIReference**](APIReference.md) | Starting proficiencies for all new characters of this background. | [optional] 
**StartingEquipment** | Pointer to [**[]APIReference**](APIReference.md) | Starting equipment for all new characters of this background. | [optional] 
**StartingEquipmentOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**LanguageOptions** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Feature** | Pointer to [**BackgroundAllOfFeature**](BackgroundAllOfFeature.md) |  | [optional] 
**PersonalityTraits** | Pointer to **map[string]interface{}** | Choice of personality traits for this background. | [optional] 
**Ideals** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Bonds** | Pointer to [**Choice**](Choice.md) |  | [optional] 
**Flaws** | Pointer to [**Choice**](Choice.md) |  | [optional] 

## Methods

### NewBackgroundAllOf

`func NewBackgroundAllOf() *BackgroundAllOf`

NewBackgroundAllOf instantiates a new BackgroundAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundAllOfWithDefaults

`func NewBackgroundAllOfWithDefaults() *BackgroundAllOf`

NewBackgroundAllOfWithDefaults instantiates a new BackgroundAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartingProficiencies

`func (o *BackgroundAllOf) GetStartingProficiencies() []APIReference`

GetStartingProficiencies returns the StartingProficiencies field if non-nil, zero value otherwise.

### GetStartingProficienciesOk

`func (o *BackgroundAllOf) GetStartingProficienciesOk() (*[]APIReference, bool)`

GetStartingProficienciesOk returns a tuple with the StartingProficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingProficiencies

`func (o *BackgroundAllOf) SetStartingProficiencies(v []APIReference)`

SetStartingProficiencies sets StartingProficiencies field to given value.

### HasStartingProficiencies

`func (o *BackgroundAllOf) HasStartingProficiencies() bool`

HasStartingProficiencies returns a boolean if a field has been set.

### GetStartingEquipment

`func (o *BackgroundAllOf) GetStartingEquipment() []APIReference`

GetStartingEquipment returns the StartingEquipment field if non-nil, zero value otherwise.

### GetStartingEquipmentOk

`func (o *BackgroundAllOf) GetStartingEquipmentOk() (*[]APIReference, bool)`

GetStartingEquipmentOk returns a tuple with the StartingEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipment

`func (o *BackgroundAllOf) SetStartingEquipment(v []APIReference)`

SetStartingEquipment sets StartingEquipment field to given value.

### HasStartingEquipment

`func (o *BackgroundAllOf) HasStartingEquipment() bool`

HasStartingEquipment returns a boolean if a field has been set.

### GetStartingEquipmentOptions

`func (o *BackgroundAllOf) GetStartingEquipmentOptions() Choice`

GetStartingEquipmentOptions returns the StartingEquipmentOptions field if non-nil, zero value otherwise.

### GetStartingEquipmentOptionsOk

`func (o *BackgroundAllOf) GetStartingEquipmentOptionsOk() (*Choice, bool)`

GetStartingEquipmentOptionsOk returns a tuple with the StartingEquipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipmentOptions

`func (o *BackgroundAllOf) SetStartingEquipmentOptions(v Choice)`

SetStartingEquipmentOptions sets StartingEquipmentOptions field to given value.

### HasStartingEquipmentOptions

`func (o *BackgroundAllOf) HasStartingEquipmentOptions() bool`

HasStartingEquipmentOptions returns a boolean if a field has been set.

### GetLanguageOptions

`func (o *BackgroundAllOf) GetLanguageOptions() Choice`

GetLanguageOptions returns the LanguageOptions field if non-nil, zero value otherwise.

### GetLanguageOptionsOk

`func (o *BackgroundAllOf) GetLanguageOptionsOk() (*Choice, bool)`

GetLanguageOptionsOk returns a tuple with the LanguageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOptions

`func (o *BackgroundAllOf) SetLanguageOptions(v Choice)`

SetLanguageOptions sets LanguageOptions field to given value.

### HasLanguageOptions

`func (o *BackgroundAllOf) HasLanguageOptions() bool`

HasLanguageOptions returns a boolean if a field has been set.

### GetFeature

`func (o *BackgroundAllOf) GetFeature() BackgroundAllOfFeature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *BackgroundAllOf) GetFeatureOk() (*BackgroundAllOfFeature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *BackgroundAllOf) SetFeature(v BackgroundAllOfFeature)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *BackgroundAllOf) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetPersonalityTraits

`func (o *BackgroundAllOf) GetPersonalityTraits() map[string]interface{}`

GetPersonalityTraits returns the PersonalityTraits field if non-nil, zero value otherwise.

### GetPersonalityTraitsOk

`func (o *BackgroundAllOf) GetPersonalityTraitsOk() (*map[string]interface{}, bool)`

GetPersonalityTraitsOk returns a tuple with the PersonalityTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalityTraits

`func (o *BackgroundAllOf) SetPersonalityTraits(v map[string]interface{})`

SetPersonalityTraits sets PersonalityTraits field to given value.

### HasPersonalityTraits

`func (o *BackgroundAllOf) HasPersonalityTraits() bool`

HasPersonalityTraits returns a boolean if a field has been set.

### GetIdeals

`func (o *BackgroundAllOf) GetIdeals() Choice`

GetIdeals returns the Ideals field if non-nil, zero value otherwise.

### GetIdealsOk

`func (o *BackgroundAllOf) GetIdealsOk() (*Choice, bool)`

GetIdealsOk returns a tuple with the Ideals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeals

`func (o *BackgroundAllOf) SetIdeals(v Choice)`

SetIdeals sets Ideals field to given value.

### HasIdeals

`func (o *BackgroundAllOf) HasIdeals() bool`

HasIdeals returns a boolean if a field has been set.

### GetBonds

`func (o *BackgroundAllOf) GetBonds() Choice`

GetBonds returns the Bonds field if non-nil, zero value otherwise.

### GetBondsOk

`func (o *BackgroundAllOf) GetBondsOk() (*Choice, bool)`

GetBondsOk returns a tuple with the Bonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonds

`func (o *BackgroundAllOf) SetBonds(v Choice)`

SetBonds sets Bonds field to given value.

### HasBonds

`func (o *BackgroundAllOf) HasBonds() bool`

HasBonds returns a boolean if a field has been set.

### GetFlaws

`func (o *BackgroundAllOf) GetFlaws() Choice`

GetFlaws returns the Flaws field if non-nil, zero value otherwise.

### GetFlawsOk

`func (o *BackgroundAllOf) GetFlawsOk() (*Choice, bool)`

GetFlawsOk returns a tuple with the Flaws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaws

`func (o *BackgroundAllOf) SetFlaws(v Choice)`

SetFlaws sets Flaws field to given value.

### HasFlaws

`func (o *BackgroundAllOf) HasFlaws() bool`

HasFlaws returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


