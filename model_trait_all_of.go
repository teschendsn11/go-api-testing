/*
D&D 5e API

# Introduction  Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API! This documentation should help you familiarize yourself with the resources available and how to consume them with HTTP requests. Read through the getting started section before you dive in. Most of your problems should be solved just by reading through it.  ## Getting Started  Let's make our first API request to the D&D 5th Edition API!  Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/) to make an API request for a resource. You can also scroll through the definitions below and send requests directly from the endpoint documentation!  For example, if you paste and run this `curl` command: ```bash curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\" ```  We should see a result containing details about the Charisma ability score: ```bash {   \"index\": \"cha\",   \"name\": \"CHA\",   \"full_name\": \"Charisma\",   \"desc\": [     \"Charisma measures your ability to interact effectively with others. It       includes such factors as confidence and eloquence, and it can represent       a charming or commanding personality.\",     \"A Charisma check might arise when you try to influence or entertain       others, when you try to make an impression or tell a convincing lie,       or when you are navigating a tricky social situation. The Deception,       Intimidation, Performance, and Persuasion skills reflect aptitude in       certain kinds of Charisma checks.\"   ],   \"skills\": [     {       \"name\": \"Deception\",       \"index\": \"deception\",       \"url\": \"/api/skills/deception\"     },     {       \"name\": \"Intimidation\",       \"index\": \"intimidation\",       \"url\": \"/api/skills/intimidation\"     },     {       \"name\": \"Performance\",       \"index\": \"performance\",       \"url\": \"/api/skills/performance\"     },     {       \"name\": \"Persuasion\",       \"index\": \"persuasion\",       \"url\": \"/api/skills/persuasion\"     }   ],   \"url\": \"/api/ability-scores/cha\" } ```  ## Authentication  The dnd5eapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just `GET`-ing the data. If you find a mistake in the data, feel free to [message us](https://discord.gg/TQuYTv7).  ## GraphQL  This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered by querying the endpoint with the Apollo sandbox explorer.  ## Schemas  Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.  ### `APIReference` Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`. ``` APIReference {   index     string   name      string   url       string } ``` <hr>  ### `DC` Represents a difficulty check. ``` DC {   dc_type       APIReference   dc_value      number   success_type  \"none\" | \"half\" | \"other\" } ``` <hr>  ### `Damage` Represents damage. ``` Damage {   damage_type     APIReference   damage_dice     string } ``` <hr>  ### `Choice` Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15)) ``` Choice {   desc      string   choose    number   type      string   from      OptionSet } ``` <hr>  ### `OptionSet` The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute. - `options_array`   - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen. - `equipment_category`   - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen. - `resource_list`   - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen. <hr>  ### `Option` When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below. - `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.   - `item` (APIReference): A reference to the chosen item. - `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.   - `action_name` (string): The name of the action, according to its `name` attribute.   - `count` (number | string): The number of times this action can be repeated if this option is chosen.   - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic. - `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.   - `items` (array): An array of Option objects. All of them must be taken if the option is chosen. - `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.   - `choice` (Choice): The Choice to resolve. - `string` - A terminal option. Contains a reference to a string.   - `string` (string): The string. - `ideal` - A terminal option. Contains information about an ideal.   - `desc` (string): A description of the ideal.   - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal. - `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.   - `count` (number): Count.   - `of` (ApiReference): Thing being referenced. - `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.   - `ability_score` (ApiReference): Ability score being referenced.   - `minimum_score` (number): The minimum score required to satisfy the prerequisite. - `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus   - `ability_score` (ApiReference): Ability score being referenced   - `bonus` (number): The bonus being applied to the ability score - `breath` - A terminal option: Contains a reference to information about a breath attack.   - `name` (string): Name of the breath.   - `dc` (DC): Difficulty check of the breath attack.   - `damage` ([Damage]): Damage dealt by the breath attack, if any. - `damage` - A terminal option. Contains information about damage.   - `damage_type` (ApiReference): Reference to type of damage.   - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").   - `notes` (string): Information regarding the damage.  ## FAQ  ### What is the SRD? The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)  ### What is the OGL? The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)  ### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it? Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.  ### Can this API be self hosted? Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.  #### Can I publish is on <insert platform>? Is this free use? Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.  # Status Page  The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/  # Chat  Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!  # Contribute  This API is built from two repositories.   - The repo containing the data lives here: https://github.com/bagelbits/5e-database   - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api  This is a evolving API and having fresh ideas are always welcome! You can open an issue in either repo, open a PR for changes, or just discuss with other users in this discord. 

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TraitAllOf struct for TraitAllOf
type TraitAllOf struct {
	// List of `Races` that have access to the trait.
	Races []APIReference `json:"races,omitempty"`
	// List of `Subraces` that have access to the trait.
	Subraces []APIReference `json:"subraces,omitempty"`
	// List of `Proficiencies` this trait grants.
	Proficiencies []APIReference `json:"proficiencies,omitempty"`
	ProficiencyChoices *Choice `json:"proficiency_choices,omitempty"`
	LanguageOptions *Choice `json:"language_options,omitempty"`
	// Information specific to this trait
	TraitSpecific NullableOneOfChoiceChoiceobject `json:"trait_specific,omitempty"`
}

// NewTraitAllOf instantiates a new TraitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraitAllOf() *TraitAllOf {
	this := TraitAllOf{}
	return &this
}

// NewTraitAllOfWithDefaults instantiates a new TraitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraitAllOfWithDefaults() *TraitAllOf {
	this := TraitAllOf{}
	return &this
}

// GetRaces returns the Races field value if set, zero value otherwise.
func (o *TraitAllOf) GetRaces() []APIReference {
	if o == nil || isNil(o.Races) {
		var ret []APIReference
		return ret
	}
	return o.Races
}

// GetRacesOk returns a tuple with the Races field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitAllOf) GetRacesOk() ([]APIReference, bool) {
	if o == nil || isNil(o.Races) {
    return nil, false
	}
	return o.Races, true
}

// HasRaces returns a boolean if a field has been set.
func (o *TraitAllOf) HasRaces() bool {
	if o != nil && !isNil(o.Races) {
		return true
	}

	return false
}

// SetRaces gets a reference to the given []APIReference and assigns it to the Races field.
func (o *TraitAllOf) SetRaces(v []APIReference) {
	o.Races = v
}

// GetSubraces returns the Subraces field value if set, zero value otherwise.
func (o *TraitAllOf) GetSubraces() []APIReference {
	if o == nil || isNil(o.Subraces) {
		var ret []APIReference
		return ret
	}
	return o.Subraces
}

// GetSubracesOk returns a tuple with the Subraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitAllOf) GetSubracesOk() ([]APIReference, bool) {
	if o == nil || isNil(o.Subraces) {
    return nil, false
	}
	return o.Subraces, true
}

// HasSubraces returns a boolean if a field has been set.
func (o *TraitAllOf) HasSubraces() bool {
	if o != nil && !isNil(o.Subraces) {
		return true
	}

	return false
}

// SetSubraces gets a reference to the given []APIReference and assigns it to the Subraces field.
func (o *TraitAllOf) SetSubraces(v []APIReference) {
	o.Subraces = v
}

// GetProficiencies returns the Proficiencies field value if set, zero value otherwise.
func (o *TraitAllOf) GetProficiencies() []APIReference {
	if o == nil || isNil(o.Proficiencies) {
		var ret []APIReference
		return ret
	}
	return o.Proficiencies
}

// GetProficienciesOk returns a tuple with the Proficiencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitAllOf) GetProficienciesOk() ([]APIReference, bool) {
	if o == nil || isNil(o.Proficiencies) {
    return nil, false
	}
	return o.Proficiencies, true
}

// HasProficiencies returns a boolean if a field has been set.
func (o *TraitAllOf) HasProficiencies() bool {
	if o != nil && !isNil(o.Proficiencies) {
		return true
	}

	return false
}

// SetProficiencies gets a reference to the given []APIReference and assigns it to the Proficiencies field.
func (o *TraitAllOf) SetProficiencies(v []APIReference) {
	o.Proficiencies = v
}

// GetProficiencyChoices returns the ProficiencyChoices field value if set, zero value otherwise.
func (o *TraitAllOf) GetProficiencyChoices() Choice {
	if o == nil || isNil(o.ProficiencyChoices) {
		var ret Choice
		return ret
	}
	return *o.ProficiencyChoices
}

// GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitAllOf) GetProficiencyChoicesOk() (*Choice, bool) {
	if o == nil || isNil(o.ProficiencyChoices) {
    return nil, false
	}
	return o.ProficiencyChoices, true
}

// HasProficiencyChoices returns a boolean if a field has been set.
func (o *TraitAllOf) HasProficiencyChoices() bool {
	if o != nil && !isNil(o.ProficiencyChoices) {
		return true
	}

	return false
}

// SetProficiencyChoices gets a reference to the given Choice and assigns it to the ProficiencyChoices field.
func (o *TraitAllOf) SetProficiencyChoices(v Choice) {
	o.ProficiencyChoices = &v
}

// GetLanguageOptions returns the LanguageOptions field value if set, zero value otherwise.
func (o *TraitAllOf) GetLanguageOptions() Choice {
	if o == nil || isNil(o.LanguageOptions) {
		var ret Choice
		return ret
	}
	return *o.LanguageOptions
}

// GetLanguageOptionsOk returns a tuple with the LanguageOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitAllOf) GetLanguageOptionsOk() (*Choice, bool) {
	if o == nil || isNil(o.LanguageOptions) {
    return nil, false
	}
	return o.LanguageOptions, true
}

// HasLanguageOptions returns a boolean if a field has been set.
func (o *TraitAllOf) HasLanguageOptions() bool {
	if o != nil && !isNil(o.LanguageOptions) {
		return true
	}

	return false
}

// SetLanguageOptions gets a reference to the given Choice and assigns it to the LanguageOptions field.
func (o *TraitAllOf) SetLanguageOptions(v Choice) {
	o.LanguageOptions = &v
}

// GetTraitSpecific returns the TraitSpecific field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TraitAllOf) GetTraitSpecific() OneOfChoiceChoiceobject {
	if o == nil || isNil(o.TraitSpecific.Get()) {
		var ret OneOfChoiceChoiceobject
		return ret
	}
	return *o.TraitSpecific.Get()
}

// GetTraitSpecificOk returns a tuple with the TraitSpecific field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TraitAllOf) GetTraitSpecificOk() (*OneOfChoiceChoiceobject, bool) {
	if o == nil {
    return nil, false
	}
	return o.TraitSpecific.Get(), o.TraitSpecific.IsSet()
}

// HasTraitSpecific returns a boolean if a field has been set.
func (o *TraitAllOf) HasTraitSpecific() bool {
	if o != nil && o.TraitSpecific.IsSet() {
		return true
	}

	return false
}

// SetTraitSpecific gets a reference to the given NullableOneOfChoiceChoiceobject and assigns it to the TraitSpecific field.
func (o *TraitAllOf) SetTraitSpecific(v OneOfChoiceChoiceobject) {
	o.TraitSpecific.Set(&v)
}
// SetTraitSpecificNil sets the value for TraitSpecific to be an explicit nil
func (o *TraitAllOf) SetTraitSpecificNil() {
	o.TraitSpecific.Set(nil)
}

// UnsetTraitSpecific ensures that no value is present for TraitSpecific, not even an explicit nil
func (o *TraitAllOf) UnsetTraitSpecific() {
	o.TraitSpecific.Unset()
}

func (o TraitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Races) {
		toSerialize["races"] = o.Races
	}
	if !isNil(o.Subraces) {
		toSerialize["subraces"] = o.Subraces
	}
	if !isNil(o.Proficiencies) {
		toSerialize["proficiencies"] = o.Proficiencies
	}
	if !isNil(o.ProficiencyChoices) {
		toSerialize["proficiency_choices"] = o.ProficiencyChoices
	}
	if !isNil(o.LanguageOptions) {
		toSerialize["language_options"] = o.LanguageOptions
	}
	if o.TraitSpecific.IsSet() {
		toSerialize["trait_specific"] = o.TraitSpecific.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTraitAllOf struct {
	value *TraitAllOf
	isSet bool
}

func (v NullableTraitAllOf) Get() *TraitAllOf {
	return v.value
}

func (v *NullableTraitAllOf) Set(val *TraitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTraitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTraitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraitAllOf(val *TraitAllOf) *NullableTraitAllOf {
	return &NullableTraitAllOf{value: val, isSet: true}
}

func (v NullableTraitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


