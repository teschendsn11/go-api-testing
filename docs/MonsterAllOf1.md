# MonsterAllOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | The image url of the monster. | [optional] 
**Size** | Pointer to **string** | The size of the monster ranging from Tiny to Gargantuan.\&quot; | [optional] 
**Type** | Pointer to **string** | The type of monster. | [optional] 
**Subtype** | Pointer to **string** | The sub-category of a monster used for classification of monsters.\&quot; | [optional] 
**Alignments** | Pointer to **string** | A creature&#39;s general moral and personal attitudes. | [optional] 
**ArmorClass** | Pointer to **float32** | The difficulty for a player to successfully deal damage to a monster. | [optional] 
**HitPoints** | Pointer to **float32** | The hit points of a monster determine how much damage it is able to take before it can be defeated. | [optional] 
**HitDice** | Pointer to **string** | The hit die of a monster can be used to make a version of the same monster whose hit points are determined by the roll of the die. For example: A monster with 2d6 would have its hit points determine by rolling a 6 sided die twice. | [optional] 
**HitPointsRoll** | Pointer to **string** | The roll for determining a monster&#39;s hit points, which consists of the hit dice (e.g. 18d10) and the modifier determined by its Constitution (e.g. +36). For example, 18d10+36 | [optional] 
**Actions** | Pointer to [**[]MonsterAllOf1Actions1**](MonsterAllOf1Actions1.md) | A list of actions that are available to the monster to take during combat. | [optional] 
**LegendaryActions** | Pointer to **[]map[string]interface{}** | A list of legendary actions that are available to the monster to take during combat. | [optional] 
**ChallengeRating** | Pointer to **float32** | A monster&#39;s challenge rating is a guideline number that says when a monster becomes an appropriate challenge against the party&#39;s average level. For example. A group of 4 players with an average level of 4 would have an appropriate combat challenge against a monster with a challenge rating of 4 but a monster with a challenge rating of 8 against the same group of players would pose a significant threat. | [optional] 
**ConditionImmunities** | Pointer to [**[]APIReference**](APIReference.md) | A list of conditions that a monster is immune to. | [optional] 
**DamageImmunities** | Pointer to **[]string** | A list of damage types that a monster will take double damage from. | [optional] 
**DamageResistances** | Pointer to **[]string** | A list of damage types that a monster will take half damage from. | [optional] 
**DamageVulnerabilities** | Pointer to **[]string** | A list of damage types that a monster will take double damage from. | [optional] 
**Forms** | Pointer to [**[]APIReference**](APIReference.md) | List of other related monster entries that are of the same form. Only applicable to Lycanthropes that have multiple forms. | [optional] 
**Languages** | Pointer to **string** | The languages a monster is able to speak. | [optional] 
**Proficiencies** | Pointer to [**[]MonsterAllOf1Proficiencies**](MonsterAllOf1Proficiencies.md) | A list of proficiencies of a monster. | [optional] 
**Reactions** | Pointer to **[]map[string]interface{}** | A list of reactions that is available to the monster to take during combat. | [optional] 
**Senses** | Pointer to **interface{}** | Monsters typically have a passive perception but they might also have other senses to detect players. | [optional] 
**SpecialAbilities** | Pointer to [**[]MonsterAllOf1SpecialAbilities**](MonsterAllOf1SpecialAbilities.md) | A list of the monster&#39;s special abilities. | [optional] 
**Speed** | Pointer to [**MonsterAllOf1Speed**](MonsterAllOf1Speed.md) |  | [optional] 
**Xp** | Pointer to **float32** | The number of experience points (XP) a monster is worth is based on its challenge rating. | [optional] 

## Methods

### NewMonsterAllOf1

`func NewMonsterAllOf1() *MonsterAllOf1`

NewMonsterAllOf1 instantiates a new MonsterAllOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterAllOf1WithDefaults

`func NewMonsterAllOf1WithDefaults() *MonsterAllOf1`

NewMonsterAllOf1WithDefaults instantiates a new MonsterAllOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *MonsterAllOf1) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MonsterAllOf1) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MonsterAllOf1) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *MonsterAllOf1) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSize

`func (o *MonsterAllOf1) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MonsterAllOf1) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MonsterAllOf1) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MonsterAllOf1) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *MonsterAllOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonsterAllOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonsterAllOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonsterAllOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *MonsterAllOf1) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MonsterAllOf1) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MonsterAllOf1) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *MonsterAllOf1) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetAlignments

`func (o *MonsterAllOf1) GetAlignments() string`

GetAlignments returns the Alignments field if non-nil, zero value otherwise.

### GetAlignmentsOk

`func (o *MonsterAllOf1) GetAlignmentsOk() (*string, bool)`

GetAlignmentsOk returns a tuple with the Alignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignments

`func (o *MonsterAllOf1) SetAlignments(v string)`

SetAlignments sets Alignments field to given value.

### HasAlignments

`func (o *MonsterAllOf1) HasAlignments() bool`

HasAlignments returns a boolean if a field has been set.

### GetArmorClass

`func (o *MonsterAllOf1) GetArmorClass() float32`

GetArmorClass returns the ArmorClass field if non-nil, zero value otherwise.

### GetArmorClassOk

`func (o *MonsterAllOf1) GetArmorClassOk() (*float32, bool)`

GetArmorClassOk returns a tuple with the ArmorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmorClass

`func (o *MonsterAllOf1) SetArmorClass(v float32)`

SetArmorClass sets ArmorClass field to given value.

### HasArmorClass

`func (o *MonsterAllOf1) HasArmorClass() bool`

HasArmorClass returns a boolean if a field has been set.

### GetHitPoints

`func (o *MonsterAllOf1) GetHitPoints() float32`

GetHitPoints returns the HitPoints field if non-nil, zero value otherwise.

### GetHitPointsOk

`func (o *MonsterAllOf1) GetHitPointsOk() (*float32, bool)`

GetHitPointsOk returns a tuple with the HitPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitPoints

`func (o *MonsterAllOf1) SetHitPoints(v float32)`

SetHitPoints sets HitPoints field to given value.

### HasHitPoints

`func (o *MonsterAllOf1) HasHitPoints() bool`

HasHitPoints returns a boolean if a field has been set.

### GetHitDice

`func (o *MonsterAllOf1) GetHitDice() string`

GetHitDice returns the HitDice field if non-nil, zero value otherwise.

### GetHitDiceOk

`func (o *MonsterAllOf1) GetHitDiceOk() (*string, bool)`

GetHitDiceOk returns a tuple with the HitDice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitDice

`func (o *MonsterAllOf1) SetHitDice(v string)`

SetHitDice sets HitDice field to given value.

### HasHitDice

`func (o *MonsterAllOf1) HasHitDice() bool`

HasHitDice returns a boolean if a field has been set.

### GetHitPointsRoll

`func (o *MonsterAllOf1) GetHitPointsRoll() string`

GetHitPointsRoll returns the HitPointsRoll field if non-nil, zero value otherwise.

### GetHitPointsRollOk

`func (o *MonsterAllOf1) GetHitPointsRollOk() (*string, bool)`

GetHitPointsRollOk returns a tuple with the HitPointsRoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitPointsRoll

`func (o *MonsterAllOf1) SetHitPointsRoll(v string)`

SetHitPointsRoll sets HitPointsRoll field to given value.

### HasHitPointsRoll

`func (o *MonsterAllOf1) HasHitPointsRoll() bool`

HasHitPointsRoll returns a boolean if a field has been set.

### GetActions

`func (o *MonsterAllOf1) GetActions() []MonsterAllOf1Actions1`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MonsterAllOf1) GetActionsOk() (*[]MonsterAllOf1Actions1, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MonsterAllOf1) SetActions(v []MonsterAllOf1Actions1)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MonsterAllOf1) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLegendaryActions

`func (o *MonsterAllOf1) GetLegendaryActions() []map[string]interface{}`

GetLegendaryActions returns the LegendaryActions field if non-nil, zero value otherwise.

### GetLegendaryActionsOk

`func (o *MonsterAllOf1) GetLegendaryActionsOk() (*[]map[string]interface{}, bool)`

GetLegendaryActionsOk returns a tuple with the LegendaryActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendaryActions

`func (o *MonsterAllOf1) SetLegendaryActions(v []map[string]interface{})`

SetLegendaryActions sets LegendaryActions field to given value.

### HasLegendaryActions

`func (o *MonsterAllOf1) HasLegendaryActions() bool`

HasLegendaryActions returns a boolean if a field has been set.

### GetChallengeRating

`func (o *MonsterAllOf1) GetChallengeRating() float32`

GetChallengeRating returns the ChallengeRating field if non-nil, zero value otherwise.

### GetChallengeRatingOk

`func (o *MonsterAllOf1) GetChallengeRatingOk() (*float32, bool)`

GetChallengeRatingOk returns a tuple with the ChallengeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeRating

`func (o *MonsterAllOf1) SetChallengeRating(v float32)`

SetChallengeRating sets ChallengeRating field to given value.

### HasChallengeRating

`func (o *MonsterAllOf1) HasChallengeRating() bool`

HasChallengeRating returns a boolean if a field has been set.

### GetConditionImmunities

`func (o *MonsterAllOf1) GetConditionImmunities() []APIReference`

GetConditionImmunities returns the ConditionImmunities field if non-nil, zero value otherwise.

### GetConditionImmunitiesOk

`func (o *MonsterAllOf1) GetConditionImmunitiesOk() (*[]APIReference, bool)`

GetConditionImmunitiesOk returns a tuple with the ConditionImmunities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionImmunities

`func (o *MonsterAllOf1) SetConditionImmunities(v []APIReference)`

SetConditionImmunities sets ConditionImmunities field to given value.

### HasConditionImmunities

`func (o *MonsterAllOf1) HasConditionImmunities() bool`

HasConditionImmunities returns a boolean if a field has been set.

### GetDamageImmunities

`func (o *MonsterAllOf1) GetDamageImmunities() []string`

GetDamageImmunities returns the DamageImmunities field if non-nil, zero value otherwise.

### GetDamageImmunitiesOk

`func (o *MonsterAllOf1) GetDamageImmunitiesOk() (*[]string, bool)`

GetDamageImmunitiesOk returns a tuple with the DamageImmunities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageImmunities

`func (o *MonsterAllOf1) SetDamageImmunities(v []string)`

SetDamageImmunities sets DamageImmunities field to given value.

### HasDamageImmunities

`func (o *MonsterAllOf1) HasDamageImmunities() bool`

HasDamageImmunities returns a boolean if a field has been set.

### GetDamageResistances

`func (o *MonsterAllOf1) GetDamageResistances() []string`

GetDamageResistances returns the DamageResistances field if non-nil, zero value otherwise.

### GetDamageResistancesOk

`func (o *MonsterAllOf1) GetDamageResistancesOk() (*[]string, bool)`

GetDamageResistancesOk returns a tuple with the DamageResistances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageResistances

`func (o *MonsterAllOf1) SetDamageResistances(v []string)`

SetDamageResistances sets DamageResistances field to given value.

### HasDamageResistances

`func (o *MonsterAllOf1) HasDamageResistances() bool`

HasDamageResistances returns a boolean if a field has been set.

### GetDamageVulnerabilities

`func (o *MonsterAllOf1) GetDamageVulnerabilities() []string`

GetDamageVulnerabilities returns the DamageVulnerabilities field if non-nil, zero value otherwise.

### GetDamageVulnerabilitiesOk

`func (o *MonsterAllOf1) GetDamageVulnerabilitiesOk() (*[]string, bool)`

GetDamageVulnerabilitiesOk returns a tuple with the DamageVulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageVulnerabilities

`func (o *MonsterAllOf1) SetDamageVulnerabilities(v []string)`

SetDamageVulnerabilities sets DamageVulnerabilities field to given value.

### HasDamageVulnerabilities

`func (o *MonsterAllOf1) HasDamageVulnerabilities() bool`

HasDamageVulnerabilities returns a boolean if a field has been set.

### GetForms

`func (o *MonsterAllOf1) GetForms() []APIReference`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *MonsterAllOf1) GetFormsOk() (*[]APIReference, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *MonsterAllOf1) SetForms(v []APIReference)`

SetForms sets Forms field to given value.

### HasForms

`func (o *MonsterAllOf1) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetLanguages

`func (o *MonsterAllOf1) GetLanguages() string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *MonsterAllOf1) GetLanguagesOk() (*string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *MonsterAllOf1) SetLanguages(v string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *MonsterAllOf1) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetProficiencies

`func (o *MonsterAllOf1) GetProficiencies() []MonsterAllOf1Proficiencies`

GetProficiencies returns the Proficiencies field if non-nil, zero value otherwise.

### GetProficienciesOk

`func (o *MonsterAllOf1) GetProficienciesOk() (*[]MonsterAllOf1Proficiencies, bool)`

GetProficienciesOk returns a tuple with the Proficiencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencies

`func (o *MonsterAllOf1) SetProficiencies(v []MonsterAllOf1Proficiencies)`

SetProficiencies sets Proficiencies field to given value.

### HasProficiencies

`func (o *MonsterAllOf1) HasProficiencies() bool`

HasProficiencies returns a boolean if a field has been set.

### GetReactions

`func (o *MonsterAllOf1) GetReactions() []map[string]interface{}`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *MonsterAllOf1) GetReactionsOk() (*[]map[string]interface{}, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *MonsterAllOf1) SetReactions(v []map[string]interface{})`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *MonsterAllOf1) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetSenses

`func (o *MonsterAllOf1) GetSenses() interface{}`

GetSenses returns the Senses field if non-nil, zero value otherwise.

### GetSensesOk

`func (o *MonsterAllOf1) GetSensesOk() (*interface{}, bool)`

GetSensesOk returns a tuple with the Senses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenses

`func (o *MonsterAllOf1) SetSenses(v interface{})`

SetSenses sets Senses field to given value.

### HasSenses

`func (o *MonsterAllOf1) HasSenses() bool`

HasSenses returns a boolean if a field has been set.

### GetSpecialAbilities

`func (o *MonsterAllOf1) GetSpecialAbilities() []MonsterAllOf1SpecialAbilities`

GetSpecialAbilities returns the SpecialAbilities field if non-nil, zero value otherwise.

### GetSpecialAbilitiesOk

`func (o *MonsterAllOf1) GetSpecialAbilitiesOk() (*[]MonsterAllOf1SpecialAbilities, bool)`

GetSpecialAbilitiesOk returns a tuple with the SpecialAbilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAbilities

`func (o *MonsterAllOf1) SetSpecialAbilities(v []MonsterAllOf1SpecialAbilities)`

SetSpecialAbilities sets SpecialAbilities field to given value.

### HasSpecialAbilities

`func (o *MonsterAllOf1) HasSpecialAbilities() bool`

HasSpecialAbilities returns a boolean if a field has been set.

### GetSpeed

`func (o *MonsterAllOf1) GetSpeed() MonsterAllOf1Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MonsterAllOf1) GetSpeedOk() (*MonsterAllOf1Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MonsterAllOf1) SetSpeed(v MonsterAllOf1Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MonsterAllOf1) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetXp

`func (o *MonsterAllOf1) GetXp() float32`

GetXp returns the Xp field if non-nil, zero value otherwise.

### GetXpOk

`func (o *MonsterAllOf1) GetXpOk() (*float32, bool)`

GetXpOk returns a tuple with the Xp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXp

`func (o *MonsterAllOf1) SetXp(v float32)`

SetXp sets Xp field to given value.

### HasXp

`func (o *MonsterAllOf1) HasXp() bool`

HasXp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


