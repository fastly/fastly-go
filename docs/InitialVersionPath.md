# InitialVersionPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The URL path pattern, beginning with `/`. Maximum 2048 characters. | 
**Rules** | Pointer to [**[]RuleCreate**](RuleCreate.md) | The rules to create on this path. | [optional] 

## Methods

### NewInitialVersionPath

`func NewInitialVersionPath(path string, ) *InitialVersionPath`

NewInitialVersionPath instantiates a new InitialVersionPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialVersionPathWithDefaults

`func NewInitialVersionPathWithDefaults() *InitialVersionPath`

NewInitialVersionPathWithDefaults instantiates a new InitialVersionPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *InitialVersionPath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InitialVersionPath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InitialVersionPath) SetPath(v string)`

SetPath sets Path field to given value.


### GetRules

`func (o *InitialVersionPath) GetRules() []RuleCreate`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InitialVersionPath) GetRulesOk() (*[]RuleCreate, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InitialVersionPath) SetRules(v []RuleCreate)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InitialVersionPath) HasRules() bool`

HasRules returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


