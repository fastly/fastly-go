# PathWithRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to [**PathResponse**](PathResponse.md) |  | [optional] 
**Rules** | Pointer to [**[]RuleResponse**](RuleResponse.md) |  | [optional] 

## Methods

### NewPathWithRules

`func NewPathWithRules() *PathWithRules`

NewPathWithRules instantiates a new PathWithRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathWithRulesWithDefaults

`func NewPathWithRulesWithDefaults() *PathWithRules`

NewPathWithRulesWithDefaults instantiates a new PathWithRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PathWithRules) GetPath() PathResponse`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathWithRules) GetPathOk() (*PathResponse, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathWithRules) SetPath(v PathResponse)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathWithRules) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRules

`func (o *PathWithRules) GetRules() []RuleResponse`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PathWithRules) GetRulesOk() (*[]RuleResponse, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PathWithRules) SetRules(v []RuleResponse)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PathWithRules) HasRules() bool`

HasRules returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


