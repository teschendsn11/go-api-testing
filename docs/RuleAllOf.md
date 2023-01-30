# RuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Description of the rule. | [optional] 
**Subsections** | Pointer to [**[]APIReference**](APIReference.md) | List of sections for each subheading underneath the rule in the SRD. | [optional] 

## Methods

### NewRuleAllOf

`func NewRuleAllOf() *RuleAllOf`

NewRuleAllOf instantiates a new RuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleAllOfWithDefaults

`func NewRuleAllOfWithDefaults() *RuleAllOf`

NewRuleAllOfWithDefaults instantiates a new RuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *RuleAllOf) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *RuleAllOf) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *RuleAllOf) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *RuleAllOf) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetSubsections

`func (o *RuleAllOf) GetSubsections() []APIReference`

GetSubsections returns the Subsections field if non-nil, zero value otherwise.

### GetSubsectionsOk

`func (o *RuleAllOf) GetSubsectionsOk() (*[]APIReference, bool)`

GetSubsectionsOk returns a tuple with the Subsections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsections

`func (o *RuleAllOf) SetSubsections(v []APIReference)`

SetSubsections sets Subsections field to given value.

### HasSubsections

`func (o *RuleAllOf) HasSubsections() bool`

HasSubsections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


