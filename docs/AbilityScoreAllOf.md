# AbilityScoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **string** | Full name of the ability score. | [optional] 
**Skills** | Pointer to [**[]APIReference**](APIReference.md) | List of skills that use this ability score. | [optional] 

## Methods

### NewAbilityScoreAllOf

`func NewAbilityScoreAllOf() *AbilityScoreAllOf`

NewAbilityScoreAllOf instantiates a new AbilityScoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbilityScoreAllOfWithDefaults

`func NewAbilityScoreAllOfWithDefaults() *AbilityScoreAllOf`

NewAbilityScoreAllOfWithDefaults instantiates a new AbilityScoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *AbilityScoreAllOf) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AbilityScoreAllOf) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AbilityScoreAllOf) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AbilityScoreAllOf) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetSkills

`func (o *AbilityScoreAllOf) GetSkills() []APIReference`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *AbilityScoreAllOf) GetSkillsOk() (*[]APIReference, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *AbilityScoreAllOf) SetSkills(v []APIReference)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *AbilityScoreAllOf) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


