# MonsterAllOf1Speed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Walk** | Pointer to **string** | All creatures have a walking speed, simply called the monster’s speed. Creatures that have no form of ground-based locomotion have a walking speed of 0 feet. | [optional] 
**Burrow** | Pointer to **string** | A monster that has a burrowing speed can use that speed to move through sand, earth, mud, or ice. A monster can’t burrow through solid rock unless it has a special trait that allows it to do so. | [optional] 
**Climb** | Pointer to **string** | A monster that has a climbing speed can use all or part of its movement to move on vertical surfaces. The monster doesn’t need to spend extra movement to climb. | [optional] 
**Fly** | Pointer to **string** | A monster that has a flying speed can use all or part of its movement to fly. | [optional] 
**Swim** | Pointer to **string** | A monster that has a swimming speed doesn’t need to spend extra movement to swim. | [optional] 

## Methods

### NewMonsterAllOf1Speed

`func NewMonsterAllOf1Speed() *MonsterAllOf1Speed`

NewMonsterAllOf1Speed instantiates a new MonsterAllOf1Speed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterAllOf1SpeedWithDefaults

`func NewMonsterAllOf1SpeedWithDefaults() *MonsterAllOf1Speed`

NewMonsterAllOf1SpeedWithDefaults instantiates a new MonsterAllOf1Speed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalk

`func (o *MonsterAllOf1Speed) GetWalk() string`

GetWalk returns the Walk field if non-nil, zero value otherwise.

### GetWalkOk

`func (o *MonsterAllOf1Speed) GetWalkOk() (*string, bool)`

GetWalkOk returns a tuple with the Walk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalk

`func (o *MonsterAllOf1Speed) SetWalk(v string)`

SetWalk sets Walk field to given value.

### HasWalk

`func (o *MonsterAllOf1Speed) HasWalk() bool`

HasWalk returns a boolean if a field has been set.

### GetBurrow

`func (o *MonsterAllOf1Speed) GetBurrow() string`

GetBurrow returns the Burrow field if non-nil, zero value otherwise.

### GetBurrowOk

`func (o *MonsterAllOf1Speed) GetBurrowOk() (*string, bool)`

GetBurrowOk returns a tuple with the Burrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurrow

`func (o *MonsterAllOf1Speed) SetBurrow(v string)`

SetBurrow sets Burrow field to given value.

### HasBurrow

`func (o *MonsterAllOf1Speed) HasBurrow() bool`

HasBurrow returns a boolean if a field has been set.

### GetClimb

`func (o *MonsterAllOf1Speed) GetClimb() string`

GetClimb returns the Climb field if non-nil, zero value otherwise.

### GetClimbOk

`func (o *MonsterAllOf1Speed) GetClimbOk() (*string, bool)`

GetClimbOk returns a tuple with the Climb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimb

`func (o *MonsterAllOf1Speed) SetClimb(v string)`

SetClimb sets Climb field to given value.

### HasClimb

`func (o *MonsterAllOf1Speed) HasClimb() bool`

HasClimb returns a boolean if a field has been set.

### GetFly

`func (o *MonsterAllOf1Speed) GetFly() string`

GetFly returns the Fly field if non-nil, zero value otherwise.

### GetFlyOk

`func (o *MonsterAllOf1Speed) GetFlyOk() (*string, bool)`

GetFlyOk returns a tuple with the Fly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFly

`func (o *MonsterAllOf1Speed) SetFly(v string)`

SetFly sets Fly field to given value.

### HasFly

`func (o *MonsterAllOf1Speed) HasFly() bool`

HasFly returns a boolean if a field has been set.

### GetSwim

`func (o *MonsterAllOf1Speed) GetSwim() string`

GetSwim returns the Swim field if non-nil, zero value otherwise.

### GetSwimOk

`func (o *MonsterAllOf1Speed) GetSwimOk() (*string, bool)`

GetSwimOk returns a tuple with the Swim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwim

`func (o *MonsterAllOf1Speed) SetSwim(v string)`

SetSwim sets Swim field to given value.

### HasSwim

`func (o *MonsterAllOf1Speed) HasSwim() bool`

HasSwim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


