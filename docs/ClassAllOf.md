# ClassAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HitDie** | Pointer to **float32** | Hit die of the class. (ex: 12 &#x3D;&#x3D; 1d12). | [optional] 
**ClassLevels** | Pointer to **string** | URL of the level resource for the class. | [optional] 
**MultiClassing** | Pointer to [**Multiclassing**](Multiclassing.md) |  | [optional] 
**Spellcasting** | Pointer to [**Spellcasting**](Spellcasting.md) |  | [optional] 
**Spells** | Pointer to **string** | URL of the spell resource list for the class. | [optional] 
**StartingEquipment** | Pointer to [**[]ClassAllOfStartingEquipment**](ClassAllOfStartingEquipment.md) | List of equipment and their quantities all players of the class start with. | [optional] 
**StartingEquipmentOptions** | Pointer to [**[]Choice**](Choice.md) | List of choices of starting equipment. | [optional] 
**ProficiencyChoices** | Pointer to [**[]Choice**](Choice.md) | List of choices of starting proficiencies. | [optional] 
**Proficiencies** | Pointer to [**[]APIReference**](APIReference.md) | List of starting proficiencies for all new characters of this class. | [optional] 
**SavingThrows** | Pointer to [**[]APIReference**](APIReference.md) | Saving throws the class is proficient in. | [optional] 
**Subclasses** | Pointer to [**[]APIReference**](APIReference.md) | List of all possible subclasses this class can specialize in. | [optional] 

## Methods

### NewClassAllOf

`func NewClassAllOf() *ClassAllOf`

NewClassAllOf instantiates a new ClassAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassAllOfWithDefaults

`func NewClassAllOfWithDefaults() *ClassAllOf`

NewClassAllOfWithDefaults instantiates a new ClassAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHitDie

`func (o *ClassAllOf) GetHitDie() float32`

GetHitDie returns the HitDie field if non-nil, zero value otherwise.

### GetHitDieOk

`func (o *ClassAllOf) GetHitDieOk() (*float32, bool)`

GetHitDieOk returns a tuple with the HitDie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitDie

`func (o *ClassAllOf) SetHitDie(v float32)`

SetHitDie sets HitDie field to given value.

### HasHitDie

`func (o *ClassAllOf) HasHitDie() bool`

HasHitDie returns a boolean if a field has been set.

### GetClassLevels

`func (o *ClassAllOf) GetClassLevels() string`

GetClassLevels returns the ClassLevels field if non-nil, zero value otherwise.

### GetClassLevelsOk

`func (o *ClassAllOf) GetClassLevelsOk() (*string, bool)`

GetClassLevelsOk returns a tuple with the ClassLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassLevels

`func (o *ClassAllOf) SetClassLevels(v string)`

SetClassLevels sets ClassLevels field to given value.

### HasClassLevels

`func (o *ClassAllOf) HasClassLevels() bool`

HasClassLevels returns a boolean if a field has been set.

### GetMultiClassing

`func (o *ClassAllOf) GetMultiClassing() Multiclassing`

GetMultiClassing returns the MultiClassing field if non-nil, zero value otherwise.

### GetMultiClassingOk

`func (o *ClassAllOf) GetMultiClassingOk() (*Multiclassing, bool)`

GetMultiClassingOk returns a tuple with the MultiClassing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiClassing

`func (o *ClassAllOf) SetMultiClassing(v Multiclassing)`

SetMultiClassing sets MultiClassing field to given value.

### HasMultiClassing

`func (o *ClassAllOf) HasMultiClassing() bool`

HasMultiClassing returns a boolean if a field has been set.

### GetSpellcasting

`func (o *ClassAllOf) GetSpellcasting() Spellcasting`

GetSpellcasting returns the Spellcasting field if non-nil, zero value otherwise.

### GetSpellcastingOk

`func (o *ClassAllOf) GetSpellcastingOk() (*Spellcasting, bool)`

GetSpellcastingOk returns a tuple with the Spellcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellcasting

`func (o *ClassAllOf) SetSpellcasting(v Spellcasting)`

SetSpellcasting sets Spellcasting field to given value.

### HasSpellcasting

`func (o *ClassAllOf) HasSpellcasting() bool`

HasSpellcasting returns a boolean if a field has been set.

### GetSpells

`func (o *ClassAllOf) GetSpells() string`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *ClassAllOf) GetSpellsOk() (*string, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *ClassAllOf) SetSpells(v string)`

SetSpells sets Spells field to given value.

### HasSpells

`func (o *ClassAllOf) HasSpells() bool`

HasSpells returns a boolean if a field has been set.

### GetStartingEquipment

`func (o *ClassAllOf) GetStartingEquipment() []ClassAllOfStartingEquipment`

GetStartingEquipment returns the StartingEquipment field if non-nil, zero value otherwise.

### GetStartingEquipmentOk

`func (o *ClassAllOf) GetStartingEquipmentOk() (*[]ClassAllOfStartingEquipment, bool)`

GetStartingEquipmentOk returns a tuple with the StartingEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipment

`func (o *ClassAllOf) SetStartingEquipment(v []ClassAllOfStartingEquipment)`

SetStartingEquipment sets StartingEquipment field to given value.

### HasStartingEquipment

`func (o *ClassAllOf) HasStartingEquipment() bool`

HasStartingEquipment returns a boolean if a field has been set.

### GetStartingEquipmentOptions

`func (o *ClassAllOf) GetStartingEquipmentOptions() []Choice`

GetStartingEquipmentOptions returns the StartingEquipmentOptions field if non-nil, zero value otherwise.

### GetStartingEquipmentOptionsOk

`func (o *ClassAllOf) GetStartingEquipmentOptionsOk() (*[]Choice, bool)`

GetStartingEquipmentOptionsOk returns a tuple with the StartingEquipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingEquipmentOptions

`func (o *ClassAllOf) SetStartingEquipmentOptions(v []Choice)`

SetStartingEquipmentOptions sets StartingEquipmentOptions field to given value.

### HasStartingEquipmentOptions

`func (o *ClassAllOf) HasStartingEquipmentOptions() bool`

HasStartingEquipmentOptions returns a boolean if a field has been set.

### GetProficiencyChoices

`func (o *ClassAllOf) GetProficiencyChoices() []Choice`

GetProficiencyChoices returns the ProficiencyChoices field if non-nil, zero value otherwise.

### GetProficiencyChoicesOk

`func (o *ClassAllOf) GetProficiencyChoicesOk() (*[]Choice, bool)`

GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyChoices

`func (o *ClassAllOf) SetProficiencyChoices(v []Choice)`

SetProficiencyChoices sets ProficiencyChoices field to given value.

### HasProficiencyChoices

`func (o *ClassAllOf) HasProficiencyChoices() bool`

HasProficiencyChoices returns a boolean if a field has been set.

### GetProficiencies

`func (o *ClassAllOf) GetProficiencies() []APIReference`

GetProficiencies returns the Proficiencies field if non-nil, zero value otherwise.

### GetProficienciesOk

`func (o *ClassAllOf) GetProficienciesOk() (*[]APIReference, bool)`

GetProficienciesOk returns a tuple with the Proficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencies

`func (o *ClassAllOf) SetProficiencies(v []APIReference)`

SetProficiencies sets Proficiencies field to given value.

### HasProficiencies

`func (o *ClassAllOf) HasProficiencies() bool`

HasProficiencies returns a boolean if a field has been set.

### GetSavingThrows

`func (o *ClassAllOf) GetSavingThrows() []APIReference`

GetSavingThrows returns the SavingThrows field if non-nil, zero value otherwise.

### GetSavingThrowsOk

`func (o *ClassAllOf) GetSavingThrowsOk() (*[]APIReference, bool)`

GetSavingThrowsOk returns a tuple with the SavingThrows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingThrows

`func (o *ClassAllOf) SetSavingThrows(v []APIReference)`

SetSavingThrows sets SavingThrows field to given value.

### HasSavingThrows

`func (o *ClassAllOf) HasSavingThrows() bool`

HasSavingThrows returns a boolean if a field has been set.

### GetSubclasses

`func (o *ClassAllOf) GetSubclasses() []APIReference`

GetSubclasses returns the Subclasses field if non-nil, zero value otherwise.

### GetSubclassesOk

`func (o *ClassAllOf) GetSubclassesOk() (*[]APIReference, bool)`

GetSubclassesOk returns a tuple with the Subclasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclasses

`func (o *ClassAllOf) SetSubclasses(v []APIReference)`

SetSubclasses sets Subclasses field to given value.

### HasSubclasses

`func (o *ClassAllOf) HasSubclasses() bool`

HasSubclasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


