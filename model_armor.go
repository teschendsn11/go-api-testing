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

// Armor `Armor` 
type Armor struct {
	// Resource index for shorthand searching.
	Index *string `json:"index,omitempty"`
	// Name of the referenced resource.
	Name *string `json:"name,omitempty"`
	// URL of the referenced resource.
	Url *string `json:"url,omitempty"`
	// Description of the resource.
	Desc []string `json:"desc,omitempty"`
	EquipmentCategory *APIReference `json:"equipment_category,omitempty"`
	// The category of armor this falls into.
	ArmorCategory *string `json:"armor_category,omitempty"`
	// Details on how to calculate armor class.
	ArmorClass *map[string]string `json:"armor_class,omitempty"`
	// Minimum STR required to use this armor.
	StrMinimum *float32 `json:"str_minimum,omitempty"`
	// Whether the armor gives disadvantage for Stealth.
	StealthDisadvantage *bool `json:"stealth_disadvantage,omitempty"`
	Cost *Cost `json:"cost,omitempty"`
	// How much the equipment weighs.
	Weight *float32 `json:"weight,omitempty"`
}

// NewArmor instantiates a new Armor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArmor() *Armor {
	this := Armor{}
	return &this
}

// NewArmorWithDefaults instantiates a new Armor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArmorWithDefaults() *Armor {
	this := Armor{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Armor) GetIndex() string {
	if o == nil || isNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetIndexOk() (*string, bool) {
	if o == nil || isNil(o.Index) {
    return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Armor) HasIndex() bool {
	if o != nil && !isNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *Armor) SetIndex(v string) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Armor) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Armor) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Armor) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Armor) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Armor) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Armor) SetUrl(v string) {
	o.Url = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *Armor) GetDesc() []string {
	if o == nil || isNil(o.Desc) {
		var ret []string
		return ret
	}
	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetDescOk() ([]string, bool) {
	if o == nil || isNil(o.Desc) {
    return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *Armor) HasDesc() bool {
	if o != nil && !isNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given []string and assigns it to the Desc field.
func (o *Armor) SetDesc(v []string) {
	o.Desc = v
}

// GetEquipmentCategory returns the EquipmentCategory field value if set, zero value otherwise.
func (o *Armor) GetEquipmentCategory() APIReference {
	if o == nil || isNil(o.EquipmentCategory) {
		var ret APIReference
		return ret
	}
	return *o.EquipmentCategory
}

// GetEquipmentCategoryOk returns a tuple with the EquipmentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetEquipmentCategoryOk() (*APIReference, bool) {
	if o == nil || isNil(o.EquipmentCategory) {
    return nil, false
	}
	return o.EquipmentCategory, true
}

// HasEquipmentCategory returns a boolean if a field has been set.
func (o *Armor) HasEquipmentCategory() bool {
	if o != nil && !isNil(o.EquipmentCategory) {
		return true
	}

	return false
}

// SetEquipmentCategory gets a reference to the given APIReference and assigns it to the EquipmentCategory field.
func (o *Armor) SetEquipmentCategory(v APIReference) {
	o.EquipmentCategory = &v
}

// GetArmorCategory returns the ArmorCategory field value if set, zero value otherwise.
func (o *Armor) GetArmorCategory() string {
	if o == nil || isNil(o.ArmorCategory) {
		var ret string
		return ret
	}
	return *o.ArmorCategory
}

// GetArmorCategoryOk returns a tuple with the ArmorCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetArmorCategoryOk() (*string, bool) {
	if o == nil || isNil(o.ArmorCategory) {
    return nil, false
	}
	return o.ArmorCategory, true
}

// HasArmorCategory returns a boolean if a field has been set.
func (o *Armor) HasArmorCategory() bool {
	if o != nil && !isNil(o.ArmorCategory) {
		return true
	}

	return false
}

// SetArmorCategory gets a reference to the given string and assigns it to the ArmorCategory field.
func (o *Armor) SetArmorCategory(v string) {
	o.ArmorCategory = &v
}

// GetArmorClass returns the ArmorClass field value if set, zero value otherwise.
func (o *Armor) GetArmorClass() map[string]string {
	if o == nil || isNil(o.ArmorClass) {
		var ret map[string]string
		return ret
	}
	return *o.ArmorClass
}

// GetArmorClassOk returns a tuple with the ArmorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetArmorClassOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ArmorClass) {
    return nil, false
	}
	return o.ArmorClass, true
}

// HasArmorClass returns a boolean if a field has been set.
func (o *Armor) HasArmorClass() bool {
	if o != nil && !isNil(o.ArmorClass) {
		return true
	}

	return false
}

// SetArmorClass gets a reference to the given map[string]string and assigns it to the ArmorClass field.
func (o *Armor) SetArmorClass(v map[string]string) {
	o.ArmorClass = &v
}

// GetStrMinimum returns the StrMinimum field value if set, zero value otherwise.
func (o *Armor) GetStrMinimum() float32 {
	if o == nil || isNil(o.StrMinimum) {
		var ret float32
		return ret
	}
	return *o.StrMinimum
}

// GetStrMinimumOk returns a tuple with the StrMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetStrMinimumOk() (*float32, bool) {
	if o == nil || isNil(o.StrMinimum) {
    return nil, false
	}
	return o.StrMinimum, true
}

// HasStrMinimum returns a boolean if a field has been set.
func (o *Armor) HasStrMinimum() bool {
	if o != nil && !isNil(o.StrMinimum) {
		return true
	}

	return false
}

// SetStrMinimum gets a reference to the given float32 and assigns it to the StrMinimum field.
func (o *Armor) SetStrMinimum(v float32) {
	o.StrMinimum = &v
}

// GetStealthDisadvantage returns the StealthDisadvantage field value if set, zero value otherwise.
func (o *Armor) GetStealthDisadvantage() bool {
	if o == nil || isNil(o.StealthDisadvantage) {
		var ret bool
		return ret
	}
	return *o.StealthDisadvantage
}

// GetStealthDisadvantageOk returns a tuple with the StealthDisadvantage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetStealthDisadvantageOk() (*bool, bool) {
	if o == nil || isNil(o.StealthDisadvantage) {
    return nil, false
	}
	return o.StealthDisadvantage, true
}

// HasStealthDisadvantage returns a boolean if a field has been set.
func (o *Armor) HasStealthDisadvantage() bool {
	if o != nil && !isNil(o.StealthDisadvantage) {
		return true
	}

	return false
}

// SetStealthDisadvantage gets a reference to the given bool and assigns it to the StealthDisadvantage field.
func (o *Armor) SetStealthDisadvantage(v bool) {
	o.StealthDisadvantage = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *Armor) GetCost() Cost {
	if o == nil || isNil(o.Cost) {
		var ret Cost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetCostOk() (*Cost, bool) {
	if o == nil || isNil(o.Cost) {
    return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *Armor) HasCost() bool {
	if o != nil && !isNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given Cost and assigns it to the Cost field.
func (o *Armor) SetCost(v Cost) {
	o.Cost = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Armor) GetWeight() float32 {
	if o == nil || isNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Armor) GetWeightOk() (*float32, bool) {
	if o == nil || isNil(o.Weight) {
    return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Armor) HasWeight() bool {
	if o != nil && !isNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *Armor) SetWeight(v float32) {
	o.Weight = &v
}

func (o Armor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !isNil(o.EquipmentCategory) {
		toSerialize["equipment_category"] = o.EquipmentCategory
	}
	if !isNil(o.ArmorCategory) {
		toSerialize["armor_category"] = o.ArmorCategory
	}
	if !isNil(o.ArmorClass) {
		toSerialize["armor_class"] = o.ArmorClass
	}
	if !isNil(o.StrMinimum) {
		toSerialize["str_minimum"] = o.StrMinimum
	}
	if !isNil(o.StealthDisadvantage) {
		toSerialize["stealth_disadvantage"] = o.StealthDisadvantage
	}
	if !isNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !isNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableArmor struct {
	value *Armor
	isSet bool
}

func (v NullableArmor) Get() *Armor {
	return v.value
}

func (v *NullableArmor) Set(val *Armor) {
	v.value = val
	v.isSet = true
}

func (v NullableArmor) IsSet() bool {
	return v.isSet
}

func (v *NullableArmor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArmor(val *Armor) *NullableArmor {
	return &NullableArmor{value: val, isSet: true}
}

func (v NullableArmor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArmor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


