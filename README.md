# Go API client for openapi

# Introduction

Welcome to the dnd5eapi, the Dungeons & Dragons 5th Edition API!
This documentation should help you familiarize yourself with the resources
available and how to consume them with HTTP requests. Read through the getting
started section before you dive in. Most of your problems should be solved
just by reading through it.

## Getting Started

Let's make our first API request to the D&D 5th Edition API!

Open up a terminal and use [curl](http://curl.haxx.se/) or [httpie](http://httpie.org/)
to make an API request for a resource. You can also scroll through the
definitions below and send requests directly from the endpoint documentation!

For example, if you paste and run this `curl` command:
```bash
curl -X GET \"https://www.dnd5eapi.co/api/ability-scores/cha\" -H \"Accept: application/json\"
```

We should see a result containing details about the Charisma ability score:
```bash
{
  \"index\": \"cha\",
  \"name\": \"CHA\",
  \"full_name\": \"Charisma\",
  \"desc\": [
    \"Charisma measures your ability to interact effectively with others. It
      includes such factors as confidence and eloquence, and it can represent
      a charming or commanding personality.\",
    \"A Charisma check might arise when you try to influence or entertain
      others, when you try to make an impression or tell a convincing lie,
      or when you are navigating a tricky social situation. The Deception,
      Intimidation, Performance, and Persuasion skills reflect aptitude in
      certain kinds of Charisma checks.\"
  ],
  \"skills\": [
    {
      \"name\": \"Deception\",
      \"index\": \"deception\",
      \"url\": \"/api/skills/deception\"
    },
    {
      \"name\": \"Intimidation\",
      \"index\": \"intimidation\",
      \"url\": \"/api/skills/intimidation\"
    },
    {
      \"name\": \"Performance\",
      \"index\": \"performance\",
      \"url\": \"/api/skills/performance\"
    },
    {
      \"name\": \"Persuasion\",
      \"index\": \"persuasion\",
      \"url\": \"/api/skills/persuasion\"
    }
  ],
  \"url\": \"/api/ability-scores/cha\"
}
```

## Authentication

The dnd5eapi is a completely open API. No authentication is required to query
and get data. This also means that we've limited what you can do to just
`GET`-ing the data. If you find a mistake in the data, feel free to
[message us](https://discord.gg/TQuYTv7).

## GraphQL

This API supports [GraphQL](https://graphql.org/). The GraphQL URL for this API
is `https://www.dnd5eapi.co/graphql`. Most of your questions regarding the GraphQL schema can be answered
by querying the endpoint with the Apollo sandbox explorer.

## Schemas

Definitions of all schemas will be accessible in a future update. Two of the most common schemas are described here.

### `APIReference`
Represents a minimal representation of a resource. The detailed representation of the referenced resource can be retrieved by making a request to the referenced `URL`.
```
APIReference {
  index     string
  name      string
  url       string
}
```
<hr>

### `DC`
Represents a difficulty check.
```
DC {
  dc_type       APIReference
  dc_value      number
  success_type  \"none\" | \"half\" | \"other\"
}
```
<hr>

### `Damage`
Represents damage.
```
Damage {
  damage_type     APIReference
  damage_dice     string
}
```
<hr>

### `Choice`
Represents a choice made by a player. Commonly seen related to decisions made during character creation or combat (e.g.: the description of the cleric class, under **Proficiencies**, states \"Skills: Choose two from History, Insight, Medicine, Persuasion, and Religion\" [[SRD p15]](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf#page=15))
```
Choice {
  desc      string
  choose    number
  type      string
  from      OptionSet
}
```
<hr>

### `OptionSet`
The OptionSet structure provides the options to be chosen from, or sufficient data to fetch and interpret the options. All OptionSets have an `option_set_type` attribute that indicates the structure of the object that contains the options. The possible values are `options_array`, `equipment_category`, and `reference_list`. Other attributes on the OptionSet depend on the value of this attribute.
- `options_array`
  - `options` (array): An array of Option objects. Each item in the array represents an option that can be chosen.
- `equipment_category`
  - `equipment_category` (APIReference): A reference to an EquipmentCategory. Each item in the EquipmentCategory's `equipment` array represents one option that can be chosen.
- `resource_list`
  - `resource_list_url` (string): A reference (by URL) to a collection in the database. The URL may include query parameters. Each item in the resulting ResourceList's `results` array represents one option that can be chosen.
<hr>

### `Option`
When the options are given in an `options_array`, each item in the array inherits from the Option structure. All Options have an `option_type` attribute that indicates the structure of the option. The value of this attribute indicates how the option should be handled, and each type has different attributes. The possible values and their corresponding attributes are listed below.
- `reference` - A terminal option. Contains a reference to a Document that can be added to the list of options chosen.
  - `item` (APIReference): A reference to the chosen item.
- `action` - A terminal option. Contains information describing an action, for use within Multiattack actions.
  - `action_name` (string): The name of the action, according to its `name` attribute.
  - `count` (number | string): The number of times this action can be repeated if this option is chosen.
  - `type` (string = `\"melee\" | \"ranged\" | \"ability\" | \"magic\"`, optional): For attack actions that can be either melee, ranged, abilities, or magic.
- `multiple` - When this option is chosen, all of its child options are chosen, and must be resolved the same way as a normal option.
  - `items` (array): An array of Option objects. All of them must be taken if the option is chosen.
- `choice` - A nested choice. If this option is chosen, the Choice structure contained within must be resolved like a normal Choice structure, and the results are the chosen options.
  - `choice` (Choice): The Choice to resolve.
- `string` - A terminal option. Contains a reference to a string.
  - `string` (string): The string.
- `ideal` - A terminal option. Contains information about an ideal.
  - `desc` (string): A description of the ideal.
  - `alignments` (ApiReference[]): A list of alignments of those who might follow the ideal.
- `counted_reference` - A terminal option. Contains a reference to something else in the API along with a count.
  - `count` (number): Count.
  - `of` (ApiReference): Thing being referenced.
- `score_prerequisite` - A terminal option. Contains a reference to an ability score and a minimum score.
  - `ability_score` (ApiReference): Ability score being referenced.
  - `minimum_score` (number): The minimum score required to satisfy the prerequisite.
- `ability_bonus` - A terminal option. Contains a reference to an ability score and a bonus
  - `ability_score` (ApiReference): Ability score being referenced
  - `bonus` (number): The bonus being applied to the ability score
- `breath` - A terminal option: Contains a reference to information about a breath attack.
  - `name` (string): Name of the breath.
  - `dc` (DC): Difficulty check of the breath attack.
  - `damage` ([Damage]): Damage dealt by the breath attack, if any.
- `damage` - A terminal option. Contains information about damage.
  - `damage_type` (ApiReference): Reference to type of damage.
  - `damage_dice` (string): Damage expressed in dice (e.g. \"13d6\").
  - `notes` (string): Information regarding the damage.

## FAQ

### What is the SRD?
The SRD, or Systems Reference Document, contains guidelines for publishing content under the OGL. This allows for some of the data for D&D 5e to be open source. The API only covers data that can be found in the SRD. [Here's a link to the full text of the SRD.](https://media.wizards.com/2016/downloads/DND/SRD-OGL_V5.1.pdf)

### What is the OGL?
The Open Game License (OGL) is a public copyright license by Wizards of the Coast that may be used by tabletop role-playing game developers to grant permission to modify, copy, and redistribute some of the content designed for their games, notably game mechanics. However, they must share-alike copies and derivative works. [More information about the OGL can be found here.](https://en.wikipedia.org/wiki/Open_Game_License)

### A monster, spell, subclass, etc. is missing from the API / Database. Can I add it?
Please check if the data is within the SRD. If it is, feel free to open an issue or PR to add it yourself. Otherwise, due to legal reasons, we cannot add it.

### Can this API be self hosted?
Yes it can! You can also host the data yourself if you don't want to use the API at all. You can also make changes and add extra data if you like. However, it is up to you to merge in new changes to the data and API.

#### Can I publish is on <insert platform>? Is this free use?
Yes, you can. The API itself is under the [MIT license](https://opensource.org/licenses/MIT), and the underlying data accessible via the API is supported under the SRD and OGL.

# Status Page

The status page for the API can be found here: https://5e-bits.github.io/dnd-uptime/

# Chat

Come hang out with us [on Discord](https://discord.gg/TQuYTv7)!

# Contribute

This API is built from two repositories.
  - The repo containing the data lives here: https://github.com/bagelbits/5e-database
  - The repo with the API implementation lives here: https://github.com/bagelbits/5e-srd-api

This is a evolving API and having fresh ideas are always welcome! You can
open an issue in either repo, open a PR for changes, or just discuss with
other users in this discord.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/5e-bits](https://github.com/5e-bits)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://www.dnd5eapi.co*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CharacterDataApi* | [**ApiAbilityScoresIndexGet**](docs/CharacterDataApi.md#apiabilityscoresindexget) | **Get** /api/ability-scores/{index} | Get an ability score by index.
*CharacterDataApi* | [**ApiAlignmentsIndexGet**](docs/CharacterDataApi.md#apialignmentsindexget) | **Get** /api/alignments/{index} | Get an alignment by index.
*CharacterDataApi* | [**ApiBackgroundsIndexGet**](docs/CharacterDataApi.md#apibackgroundsindexget) | **Get** /api/backgrounds/{index} | Get a background by index.
*CharacterDataApi* | [**ApiLanguagesIndexGet**](docs/CharacterDataApi.md#apilanguagesindexget) | **Get** /api/languages/{index} | Get a language by index.
*CharacterDataApi* | [**ApiProficienciesIndexGet**](docs/CharacterDataApi.md#apiproficienciesindexget) | **Get** /api/proficiencies/{index} | Get a proficiency by index.
*CharacterDataApi* | [**ApiSkillsIndexGet**](docs/CharacterDataApi.md#apiskillsindexget) | **Get** /api/skills/{index} | Get a skill by index.
*ClassApi* | [**ApiClassesIndexGet**](docs/ClassApi.md#apiclassesindexget) | **Get** /api/classes/{index} | Get a class by index.
*ClassApi* | [**ApiClassesIndexMultiClassingGet**](docs/ClassApi.md#apiclassesindexmulticlassingget) | **Get** /api/classes/{index}/multi-classing | Get multiclassing resource for a class.
*ClassApi* | [**ApiClassesIndexSpellcastingGet**](docs/ClassApi.md#apiclassesindexspellcastingget) | **Get** /api/classes/{index}/spellcasting | Get spellcasting info for a class.
*ClassLevelsApi* | [**ApiClassesIndexLevelsClassLevelFeaturesGet**](docs/ClassLevelsApi.md#apiclassesindexlevelsclasslevelfeaturesget) | **Get** /api/classes/{index}/levels/{class_level}/features | Get features available to a class at the requested level.
*ClassLevelsApi* | [**ApiClassesIndexLevelsClassLevelGet**](docs/ClassLevelsApi.md#apiclassesindexlevelsclasslevelget) | **Get** /api/classes/{index}/levels/{class_level} | Get level resource for a class and level.
*ClassLevelsApi* | [**ApiClassesIndexLevelsGet**](docs/ClassLevelsApi.md#apiclassesindexlevelsget) | **Get** /api/classes/{index}/levels | Get all level resources for a class.
*ClassLevelsApi* | [**ApiClassesIndexLevelsSpellLevelSpellsGet**](docs/ClassLevelsApi.md#apiclassesindexlevelsspelllevelspellsget) | **Get** /api/classes/{index}/levels/{spell_level}/spells | Get spells of the requested level available to the class.
*ClassResourceListsApi* | [**ApiClassesIndexFeaturesGet**](docs/ClassResourceListsApi.md#apiclassesindexfeaturesget) | **Get** /api/classes/{index}/features | Get features available for a class.
*ClassResourceListsApi* | [**ApiClassesIndexProficienciesGet**](docs/ClassResourceListsApi.md#apiclassesindexproficienciesget) | **Get** /api/classes/{index}/proficiencies | Get proficiencies available for a class.
*ClassResourceListsApi* | [**ApiClassesIndexSpellsGet**](docs/ClassResourceListsApi.md#apiclassesindexspellsget) | **Get** /api/classes/{index}/spells | Get spells available for a class.
*ClassResourceListsApi* | [**ApiClassesIndexSubclassesGet**](docs/ClassResourceListsApi.md#apiclassesindexsubclassesget) | **Get** /api/classes/{index}/subclasses | Get subclasses available for a class.
*CommonApi* | [**ApiEndpointGet**](docs/CommonApi.md#apiendpointget) | **Get** /api/{endpoint} | Get list of all available resources for an endpoint.
*CommonApi* | [**ApiGet**](docs/CommonApi.md#apiget) | **Get** /api | Get all resource URLs.
*EquipmentApi* | [**ApiEquipmentCategoriesIndexGet**](docs/EquipmentApi.md#apiequipmentcategoriesindexget) | **Get** /api/equipment-categories/{index} | Get an equipment category by index.
*EquipmentApi* | [**ApiEquipmentIndexGet**](docs/EquipmentApi.md#apiequipmentindexget) | **Get** /api/equipment/{index} | Get an equipment item by index.
*EquipmentApi* | [**ApiMagicItemsIndexGet**](docs/EquipmentApi.md#apimagicitemsindexget) | **Get** /api/magic-items/{index} | Get a magic item by index.
*EquipmentApi* | [**ApiWeaponPropertiesIndexGet**](docs/EquipmentApi.md#apiweaponpropertiesindexget) | **Get** /api/weapon-properties/{index} | Get a weapon property by index.
*FeatsApi* | [**ApiFeatsIndexGet**](docs/FeatsApi.md#apifeatsindexget) | **Get** /api/feats/{index} | Get a feat by index.
*FeaturesApi* | [**ApiFeaturesIndexGet**](docs/FeaturesApi.md#apifeaturesindexget) | **Get** /api/features/{index} | Get a feature by index.
*GameMechanicsApi* | [**ApiConditionsIndexGet**](docs/GameMechanicsApi.md#apiconditionsindexget) | **Get** /api/conditions/{index} | Get a condition by index.
*GameMechanicsApi* | [**ApiDamageTypesIndexGet**](docs/GameMechanicsApi.md#apidamagetypesindexget) | **Get** /api/damage-types/{index} | Get a damage type by index.
*GameMechanicsApi* | [**ApiMagicSchoolsIndexGet**](docs/GameMechanicsApi.md#apimagicschoolsindexget) | **Get** /api/magic-schools/{index} | Get a magic school by index.
*MonstersApi* | [**ApiMonstersGet**](docs/MonstersApi.md#apimonstersget) | **Get** /api/monsters | Get list of monsters with optional filtering
*MonstersApi* | [**ApiMonstersIndexGet**](docs/MonstersApi.md#apimonstersindexget) | **Get** /api/monsters/{index} | Get monster by index.
*RacesApi* | [**ApiRacesIndexGet**](docs/RacesApi.md#apiracesindexget) | **Get** /api/races/{index} | Get a race by index.
*RacesApi* | [**ApiRacesIndexProficienciesGet**](docs/RacesApi.md#apiracesindexproficienciesget) | **Get** /api/races/{index}/proficiencies | Get proficiencies available for a race.
*RacesApi* | [**ApiRacesIndexSubracesGet**](docs/RacesApi.md#apiracesindexsubracesget) | **Get** /api/races/{index}/subraces | Get subraces available for a race.
*RacesApi* | [**ApiRacesIndexTraitsGet**](docs/RacesApi.md#apiracesindextraitsget) | **Get** /api/races/{index}/traits | Get traits available for a race.
*RulesApi* | [**ApiRuleSectionsIndexGet**](docs/RulesApi.md#apirulesectionsindexget) | **Get** /api/rule-sections/{index} | Get a rule section by index.
*RulesApi* | [**ApiRulesIndexGet**](docs/RulesApi.md#apirulesindexget) | **Get** /api/rules/{index} | Get a rule by index.
*SpellsApi* | [**ApiSpellsGet**](docs/SpellsApi.md#apispellsget) | **Get** /api/spells | Get list of spells with optional filtering.
*SpellsApi* | [**ApiSpellsIndexGet**](docs/SpellsApi.md#apispellsindexget) | **Get** /api/spells/{index} | Get a spell by index.
*SubclassesApi* | [**ApiSubclassesIndexFeaturesGet**](docs/SubclassesApi.md#apisubclassesindexfeaturesget) | **Get** /api/subclasses/{index}/features | Get features available for a subclass.
*SubclassesApi* | [**ApiSubclassesIndexGet**](docs/SubclassesApi.md#apisubclassesindexget) | **Get** /api/subclasses/{index} | Get a subclass by index.
*SubclassesApi* | [**ApiSubclassesIndexLevelsGet**](docs/SubclassesApi.md#apisubclassesindexlevelsget) | **Get** /api/subclasses/{index}/levels | Get all level resources for a subclass.
*SubclassesApi* | [**ApiSubclassesIndexLevelsSubclassLevelFeaturesGet**](docs/SubclassesApi.md#apisubclassesindexlevelssubclasslevelfeaturesget) | **Get** /api/subclasses/{index}/levels/{subclass_level}/features | Get features of the requested spell level available to the class.
*SubclassesApi* | [**ApiSubclassesIndexLevelsSubclassLevelGet**](docs/SubclassesApi.md#apisubclassesindexlevelssubclasslevelget) | **Get** /api/subclasses/{index}/levels/{subclass_level} | Get level resources for a subclass and level.
*SubracesApi* | [**ApiSubracesIndexGet**](docs/SubracesApi.md#apisubracesindexget) | **Get** /api/subraces/{index} | Get a subrace by index.
*SubracesApi* | [**ApiSubracesIndexProficienciesGet**](docs/SubracesApi.md#apisubracesindexproficienciesget) | **Get** /api/subraces/{index}/proficiencies | Get proficiences available for a subrace.
*SubracesApi* | [**ApiSubracesIndexTraitsGet**](docs/SubracesApi.md#apisubracesindextraitsget) | **Get** /api/subraces/{index}/traits | Get traits available for a subrace.
*TraitsApi* | [**ApiTraitsIndexGet**](docs/TraitsApi.md#apitraitsindexget) | **Get** /api/traits/{index} | Get a trait by index.


## Documentation For Models

 - [APIReference](docs/APIReference.md)
 - [APIReferenceList](docs/APIReferenceList.md)
 - [AbilityBonus](docs/AbilityBonus.md)
 - [AbilityScore](docs/AbilityScore.md)
 - [AbilityScoreAllOf](docs/AbilityScoreAllOf.md)
 - [Alignment](docs/Alignment.md)
 - [AlignmentAllOf](docs/AlignmentAllOf.md)
 - [AreaOfEffect](docs/AreaOfEffect.md)
 - [Armor](docs/Armor.md)
 - [ArmorAllOf](docs/ArmorAllOf.md)
 - [Background](docs/Background.md)
 - [BackgroundAllOf](docs/BackgroundAllOf.md)
 - [BackgroundAllOfFeature](docs/BackgroundAllOfFeature.md)
 - [Choice](docs/Choice.md)
 - [Class](docs/Class.md)
 - [ClassAllOf](docs/ClassAllOf.md)
 - [ClassAllOfStartingEquipment](docs/ClassAllOfStartingEquipment.md)
 - [ClassLevel](docs/ClassLevel.md)
 - [ClassLevelClassSpecific](docs/ClassLevelClassSpecific.md)
 - [ClassLevelClassSpecificAnyOf](docs/ClassLevelClassSpecificAnyOf.md)
 - [ClassLevelClassSpecificAnyOf1](docs/ClassLevelClassSpecificAnyOf1.md)
 - [ClassLevelClassSpecificAnyOf10](docs/ClassLevelClassSpecificAnyOf10.md)
 - [ClassLevelClassSpecificAnyOf11](docs/ClassLevelClassSpecificAnyOf11.md)
 - [ClassLevelClassSpecificAnyOf2](docs/ClassLevelClassSpecificAnyOf2.md)
 - [ClassLevelClassSpecificAnyOf3](docs/ClassLevelClassSpecificAnyOf3.md)
 - [ClassLevelClassSpecificAnyOf4](docs/ClassLevelClassSpecificAnyOf4.md)
 - [ClassLevelClassSpecificAnyOf5](docs/ClassLevelClassSpecificAnyOf5.md)
 - [ClassLevelClassSpecificAnyOf5MartialArts](docs/ClassLevelClassSpecificAnyOf5MartialArts.md)
 - [ClassLevelClassSpecificAnyOf6](docs/ClassLevelClassSpecificAnyOf6.md)
 - [ClassLevelClassSpecificAnyOf7](docs/ClassLevelClassSpecificAnyOf7.md)
 - [ClassLevelClassSpecificAnyOf8](docs/ClassLevelClassSpecificAnyOf8.md)
 - [ClassLevelClassSpecificAnyOf9](docs/ClassLevelClassSpecificAnyOf9.md)
 - [ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner](docs/ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner.md)
 - [Condition](docs/Condition.md)
 - [Cost](docs/Cost.md)
 - [DC](docs/DC.md)
 - [Damage](docs/Damage.md)
 - [DamageAtCharacterLevel](docs/DamageAtCharacterLevel.md)
 - [DamageAtSlotLevel](docs/DamageAtSlotLevel.md)
 - [DamageType](docs/DamageType.md)
 - [Equipment](docs/Equipment.md)
 - [EquipmentCategory](docs/EquipmentCategory.md)
 - [EquipmentCategoryAllOf](docs/EquipmentCategoryAllOf.md)
 - [EquipmentPack](docs/EquipmentPack.md)
 - [EquipmentPackAllOf](docs/EquipmentPackAllOf.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Feat](docs/Feat.md)
 - [FeatAllOf](docs/FeatAllOf.md)
 - [Feature](docs/Feature.md)
 - [FeatureAllOf](docs/FeatureAllOf.md)
 - [Gear](docs/Gear.md)
 - [GearAllOf](docs/GearAllOf.md)
 - [Language](docs/Language.md)
 - [LanguageAllOf](docs/LanguageAllOf.md)
 - [MagicItem](docs/MagicItem.md)
 - [MagicItemAllOf](docs/MagicItemAllOf.md)
 - [MagicItemAllOfRarity](docs/MagicItemAllOfRarity.md)
 - [MagicSchool](docs/MagicSchool.md)
 - [MagicSchoolAllOf](docs/MagicSchoolAllOf.md)
 - [Monster](docs/Monster.md)
 - [MonsterAllOf](docs/MonsterAllOf.md)
 - [MonsterAllOf1](docs/MonsterAllOf1.md)
 - [MonsterAllOf1Actions](docs/MonsterAllOf1Actions.md)
 - [MonsterAllOf1Actions1](docs/MonsterAllOf1Actions1.md)
 - [MonsterAllOf1Attacks](docs/MonsterAllOf1Attacks.md)
 - [MonsterAllOf1Proficiencies](docs/MonsterAllOf1Proficiencies.md)
 - [MonsterAllOf1SpecialAbilities](docs/MonsterAllOf1SpecialAbilities.md)
 - [MonsterAllOf1Speed](docs/MonsterAllOf1Speed.md)
 - [MonsterAllOf1Spellcasting](docs/MonsterAllOf1Spellcasting.md)
 - [MonsterAllOf1SpellcastingSpells](docs/MonsterAllOf1SpellcastingSpells.md)
 - [MonsterAllOf1Usage](docs/MonsterAllOf1Usage.md)
 - [Multiclassing](docs/Multiclassing.md)
 - [Option](docs/Option.md)
 - [OptionOneOf](docs/OptionOneOf.md)
 - [OptionOneOf1](docs/OptionOneOf1.md)
 - [OptionOneOf10](docs/OptionOneOf10.md)
 - [OptionOneOf2](docs/OptionOneOf2.md)
 - [OptionOneOf3](docs/OptionOneOf3.md)
 - [OptionOneOf4](docs/OptionOneOf4.md)
 - [OptionOneOf5](docs/OptionOneOf5.md)
 - [OptionOneOf6](docs/OptionOneOf6.md)
 - [OptionOneOf7](docs/OptionOneOf7.md)
 - [OptionOneOf8](docs/OptionOneOf8.md)
 - [OptionOneOf9](docs/OptionOneOf9.md)
 - [OptionSet](docs/OptionSet.md)
 - [OptionSetOneOf](docs/OptionSetOneOf.md)
 - [OptionSetOneOf1](docs/OptionSetOneOf1.md)
 - [OptionSetOneOf2](docs/OptionSetOneOf2.md)
 - [Prerequisite](docs/Prerequisite.md)
 - [PrerequisiteAbilityScore](docs/PrerequisiteAbilityScore.md)
 - [Proficiency](docs/Proficiency.md)
 - [ProficiencyAllOf](docs/ProficiencyAllOf.md)
 - [Race](docs/Race.md)
 - [RaceAllOf](docs/RaceAllOf.md)
 - [ResourceDescription](docs/ResourceDescription.md)
 - [Rule](docs/Rule.md)
 - [RuleAllOf](docs/RuleAllOf.md)
 - [RuleSection](docs/RuleSection.md)
 - [RuleSectionAllOf](docs/RuleSectionAllOf.md)
 - [Skill](docs/Skill.md)
 - [SkillAllOf](docs/SkillAllOf.md)
 - [Spell](docs/Spell.md)
 - [SpellAllOf](docs/SpellAllOf.md)
 - [SpellPrerequisite](docs/SpellPrerequisite.md)
 - [SpellPrerequisiteAllOf](docs/SpellPrerequisiteAllOf.md)
 - [Spellcasting](docs/Spellcasting.md)
 - [SpellcastingInfoInner](docs/SpellcastingInfoInner.md)
 - [SpellcastingSpellcastingAbility](docs/SpellcastingSpellcastingAbility.md)
 - [Subclass](docs/Subclass.md)
 - [SubclassAllOf](docs/SubclassAllOf.md)
 - [SubclassAllOfSpells](docs/SubclassAllOfSpells.md)
 - [SubclassLevel](docs/SubclassLevel.md)
 - [SubclassLevelResource](docs/SubclassLevelResource.md)
 - [SubclassLevelSpellcasting](docs/SubclassLevelSpellcasting.md)
 - [Subrace](docs/Subrace.md)
 - [SubraceAllOf](docs/SubraceAllOf.md)
 - [Trait](docs/Trait.md)
 - [TraitAllOf](docs/TraitAllOf.md)
 - [Weapon](docs/Weapon.md)
 - [WeaponAllOf](docs/WeaponAllOf.md)
 - [WeaponAllOfRange](docs/WeaponAllOfRange.md)
 - [WeaponProperty](docs/WeaponProperty.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



