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

// Multiclassing `Multiclassing` 
type Multiclassing struct {
	// List of prerequisites that must be met.
	Prerequisites []Prerequisite `json:"prerequisites,omitempty"`
	// List of choices of prerequisites to meet for.
	PrerequisiteOptions []Choice `json:"prerequisite_options,omitempty"`
	// List of proficiencies available when multiclassing.
	Proficiencies []APIReference `json:"proficiencies,omitempty"`
	// List of choices of proficiencies that are given when multiclassing.
	ProficiencyChoices []Choice `json:"proficiency_choices,omitempty"`
}

// NewMulticlassing instantiates a new Multiclassing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMulticlassing() *Multiclassing {
	this := Multiclassing{}
	return &this
}

// NewMulticlassingWithDefaults instantiates a new Multiclassing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMulticlassingWithDefaults() *Multiclassing {
	this := Multiclassing{}
	return &this
}

// GetPrerequisites returns the Prerequisites field value if set, zero value otherwise.
func (o *Multiclassing) GetPrerequisites() []Prerequisite {
	if o == nil || isNil(o.Prerequisites) {
		var ret []Prerequisite
		return ret
	}
	return o.Prerequisites
}

// GetPrerequisitesOk returns a tuple with the Prerequisites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Multiclassing) GetPrerequisitesOk() ([]Prerequisite, bool) {
	if o == nil || isNil(o.Prerequisites) {
    return nil, false
	}
	return o.Prerequisites, true
}

// HasPrerequisites returns a boolean if a field has been set.
func (o *Multiclassing) HasPrerequisites() bool {
	if o != nil && !isNil(o.Prerequisites) {
		return true
	}

	return false
}

// SetPrerequisites gets a reference to the given []Prerequisite and assigns it to the Prerequisites field.
func (o *Multiclassing) SetPrerequisites(v []Prerequisite) {
	o.Prerequisites = v
}

// GetPrerequisiteOptions returns the PrerequisiteOptions field value if set, zero value otherwise.
func (o *Multiclassing) GetPrerequisiteOptions() []Choice {
	if o == nil || isNil(o.PrerequisiteOptions) {
		var ret []Choice
		return ret
	}
	return o.PrerequisiteOptions
}

// GetPrerequisiteOptionsOk returns a tuple with the PrerequisiteOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Multiclassing) GetPrerequisiteOptionsOk() ([]Choice, bool) {
	if o == nil || isNil(o.PrerequisiteOptions) {
    return nil, false
	}
	return o.PrerequisiteOptions, true
}

// HasPrerequisiteOptions returns a boolean if a field has been set.
func (o *Multiclassing) HasPrerequisiteOptions() bool {
	if o != nil && !isNil(o.PrerequisiteOptions) {
		return true
	}

	return false
}

// SetPrerequisiteOptions gets a reference to the given []Choice and assigns it to the PrerequisiteOptions field.
func (o *Multiclassing) SetPrerequisiteOptions(v []Choice) {
	o.PrerequisiteOptions = v
}

// GetProficiencies returns the Proficiencies field value if set, zero value otherwise.
func (o *Multiclassing) GetProficiencies() []APIReference {
	if o == nil || isNil(o.Proficiencies) {
		var ret []APIReference
		return ret
	}
	return o.Proficiencies
}

// GetProficienciesOk returns a tuple with the Proficiencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Multiclassing) GetProficienciesOk() ([]APIReference, bool) {
	if o == nil || isNil(o.Proficiencies) {
    return nil, false
	}
	return o.Proficiencies, true
}

// HasProficiencies returns a boolean if a field has been set.
func (o *Multiclassing) HasProficiencies() bool {
	if o != nil && !isNil(o.Proficiencies) {
		return true
	}

	return false
}

// SetProficiencies gets a reference to the given []APIReference and assigns it to the Proficiencies field.
func (o *Multiclassing) SetProficiencies(v []APIReference) {
	o.Proficiencies = v
}

// GetProficiencyChoices returns the ProficiencyChoices field value if set, zero value otherwise.
func (o *Multiclassing) GetProficiencyChoices() []Choice {
	if o == nil || isNil(o.ProficiencyChoices) {
		var ret []Choice
		return ret
	}
	return o.ProficiencyChoices
}

// GetProficiencyChoicesOk returns a tuple with the ProficiencyChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Multiclassing) GetProficiencyChoicesOk() ([]Choice, bool) {
	if o == nil || isNil(o.ProficiencyChoices) {
    return nil, false
	}
	return o.ProficiencyChoices, true
}

// HasProficiencyChoices returns a boolean if a field has been set.
func (o *Multiclassing) HasProficiencyChoices() bool {
	if o != nil && !isNil(o.ProficiencyChoices) {
		return true
	}

	return false
}

// SetProficiencyChoices gets a reference to the given []Choice and assigns it to the ProficiencyChoices field.
func (o *Multiclassing) SetProficiencyChoices(v []Choice) {
	o.ProficiencyChoices = v
}

func (o Multiclassing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Prerequisites) {
		toSerialize["prerequisites"] = o.Prerequisites
	}
	if !isNil(o.PrerequisiteOptions) {
		toSerialize["prerequisite_options"] = o.PrerequisiteOptions
	}
	if !isNil(o.Proficiencies) {
		toSerialize["proficiencies"] = o.Proficiencies
	}
	if !isNil(o.ProficiencyChoices) {
		toSerialize["proficiency_choices"] = o.ProficiencyChoices
	}
	return json.Marshal(toSerialize)
}

type NullableMulticlassing struct {
	value *Multiclassing
	isSet bool
}

func (v NullableMulticlassing) Get() *Multiclassing {
	return v.value
}

func (v *NullableMulticlassing) Set(val *Multiclassing) {
	v.value = val
	v.isSet = true
}

func (v NullableMulticlassing) IsSet() bool {
	return v.isSet
}

func (v *NullableMulticlassing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMulticlassing(val *Multiclassing) *NullableMulticlassing {
	return &NullableMulticlassing{value: val, isSet: true}
}

func (v NullableMulticlassing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMulticlassing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


