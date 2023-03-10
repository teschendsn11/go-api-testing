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

// MonsterAllOf struct for MonsterAllOf
type MonsterAllOf struct {
	// A monster's ability to charm or intimidate a player.
	Charisma *float32 `json:"charisma,omitempty"`
	// How sturdy a monster is.\"
	Constitution *float32 `json:"constitution,omitempty"`
	// The monster's ability for swift movement or stealth
	Dexterity *float32 `json:"dexterity,omitempty"`
	// The monster's ability to outsmart a player.
	Intelligence *float32 `json:"intelligence,omitempty"`
	// How hard a monster can hit a player.
	Strength *float32 `json:"strength,omitempty"`
	// A monster's ability to ascertain the player's plan.
	Wisdom *float32 `json:"wisdom,omitempty"`
}

// NewMonsterAllOf instantiates a new MonsterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonsterAllOf() *MonsterAllOf {
	this := MonsterAllOf{}
	return &this
}

// NewMonsterAllOfWithDefaults instantiates a new MonsterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonsterAllOfWithDefaults() *MonsterAllOf {
	this := MonsterAllOf{}
	return &this
}

// GetCharisma returns the Charisma field value if set, zero value otherwise.
func (o *MonsterAllOf) GetCharisma() float32 {
	if o == nil || isNil(o.Charisma) {
		var ret float32
		return ret
	}
	return *o.Charisma
}

// GetCharismaOk returns a tuple with the Charisma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonsterAllOf) GetCharismaOk() (*float32, bool) {
	if o == nil || isNil(o.Charisma) {
    return nil, false
	}
	return o.Charisma, true
}

// HasCharisma returns a boolean if a field has been set.
func (o *MonsterAllOf) HasCharisma() bool {
	if o != nil && !isNil(o.Charisma) {
		return true
	}

	return false
}

// SetCharisma gets a reference to the given float32 and assigns it to the Charisma field.
func (o *MonsterAllOf) SetCharisma(v float32) {
	o.Charisma = &v
}

// GetConstitution returns the Constitution field value if set, zero value otherwise.
func (o *MonsterAllOf) GetConstitution() float32 {
	if o == nil || isNil(o.Constitution) {
		var ret float32
		return ret
	}
	return *o.Constitution
}

// GetConstitutionOk returns a tuple with the Constitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonsterAllOf) GetConstitutionOk() (*float32, bool) {
	if o == nil || isNil(o.Constitution) {
    return nil, false
	}
	return o.Constitution, true
}

// HasConstitution returns a boolean if a field has been set.
func (o *MonsterAllOf) HasConstitution() bool {
	if o != nil && !isNil(o.Constitution) {
		return true
	}

	return false
}

// SetConstitution gets a reference to the given float32 and assigns it to the Constitution field.
func (o *MonsterAllOf) SetConstitution(v float32) {
	o.Constitution = &v
}

// GetDexterity returns the Dexterity field value if set, zero value otherwise.
func (o *MonsterAllOf) GetDexterity() float32 {
	if o == nil || isNil(o.Dexterity) {
		var ret float32
		return ret
	}
	return *o.Dexterity
}

// GetDexterityOk returns a tuple with the Dexterity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonsterAllOf) GetDexterityOk() (*float32, bool) {
	if o == nil || isNil(o.Dexterity) {
    return nil, false
	}
	return o.Dexterity, true
}

// HasDexterity returns a boolean if a field has been set.
func (o *MonsterAllOf) HasDexterity() bool {
	if o != nil && !isNil(o.Dexterity) {
		return true
	}

	return false
}

// SetDexterity gets a reference to the given float32 and assigns it to the Dexterity field.
func (o *MonsterAllOf) SetDexterity(v float32) {
	o.Dexterity = &v
}

// GetIntelligence returns the Intelligence field value if set, zero value otherwise.
func (o *MonsterAllOf) GetIntelligence() float32 {
	if o == nil || isNil(o.Intelligence) {
		var ret float32
		return ret
	}
	return *o.Intelligence
}

// GetIntelligenceOk returns a tuple with the Intelligence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonsterAllOf) GetIntelligenceOk() (*float32, bool) {
	if o == nil || isNil(o.Intelligence) {
    return nil, false
	}
	return o.Intelligence, true
}

// HasIntelligence returns a boolean if a field has been set.
func (o *MonsterAllOf) HasIntelligence() bool {
	if o != nil && !isNil(o.Intelligence) {
		return true
	}

	return false
}

// SetIntelligence gets a reference to the given float32 and assigns it to the Intelligence field.
func (o *MonsterAllOf) SetIntelligence(v float32) {
	o.Intelligence = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *MonsterAllOf) GetStrength() float32 {
	if o == nil || isNil(o.Strength) {
		var ret float32
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonsterAllOf) GetStrengthOk() (*float32, bool) {
	if o == nil || isNil(o.Strength) {
    return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *MonsterAllOf) HasStrength() bool {
	if o != nil && !isNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given float32 and assigns it to the Strength field.
func (o *MonsterAllOf) SetStrength(v float32) {
	o.Strength = &v
}

// GetWisdom returns the Wisdom field value if set, zero value otherwise.
func (o *MonsterAllOf) GetWisdom() float32 {
	if o == nil || isNil(o.Wisdom) {
		var ret float32
		return ret
	}
	return *o.Wisdom
}

// GetWisdomOk returns a tuple with the Wisdom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonsterAllOf) GetWisdomOk() (*float32, bool) {
	if o == nil || isNil(o.Wisdom) {
    return nil, false
	}
	return o.Wisdom, true
}

// HasWisdom returns a boolean if a field has been set.
func (o *MonsterAllOf) HasWisdom() bool {
	if o != nil && !isNil(o.Wisdom) {
		return true
	}

	return false
}

// SetWisdom gets a reference to the given float32 and assigns it to the Wisdom field.
func (o *MonsterAllOf) SetWisdom(v float32) {
	o.Wisdom = &v
}

func (o MonsterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Charisma) {
		toSerialize["charisma"] = o.Charisma
	}
	if !isNil(o.Constitution) {
		toSerialize["constitution"] = o.Constitution
	}
	if !isNil(o.Dexterity) {
		toSerialize["dexterity"] = o.Dexterity
	}
	if !isNil(o.Intelligence) {
		toSerialize["intelligence"] = o.Intelligence
	}
	if !isNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	if !isNil(o.Wisdom) {
		toSerialize["wisdom"] = o.Wisdom
	}
	return json.Marshal(toSerialize)
}

type NullableMonsterAllOf struct {
	value *MonsterAllOf
	isSet bool
}

func (v NullableMonsterAllOf) Get() *MonsterAllOf {
	return v.value
}

func (v *NullableMonsterAllOf) Set(val *MonsterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMonsterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMonsterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonsterAllOf(val *MonsterAllOf) *NullableMonsterAllOf {
	return &NullableMonsterAllOf{value: val, isSet: true}
}

func (v NullableMonsterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonsterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


