# PathChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathId** | Pointer to **string** | Alphanumeric string identifying the path. Stable across versions of the routing config. | [optional] [readonly] 
**Path** | Pointer to **string** | The current path pattern. | [optional] 
**OldPath** | Pointer to **string** | The previous path pattern, if it changed. | [optional] 
**RulesAdded** | Pointer to [**[]RuleResponse**](RuleResponse.md) | Rules that were added to this path. | [optional] 
**RulesChanged** | Pointer to [**[]RuleChange**](RuleChange.md) | Rules that were modified on this path. | [optional] 
**RulesDeleted** | Pointer to [**[]RuleResponse**](RuleResponse.md) | Rules that were removed from this path. | [optional] 

## Methods

### NewPathChange

`func NewPathChange() *PathChange`

NewPathChange instantiates a new PathChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathChangeWithDefaults

`func NewPathChangeWithDefaults() *PathChange`

NewPathChangeWithDefaults instantiates a new PathChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPathId

`func (o *PathChange) GetPathId() string`

GetPathId returns the PathId field if non-nil, zero value otherwise.

### GetPathIdOk

`func (o *PathChange) GetPathIdOk() (*string, bool)`

GetPathIdOk returns a tuple with the PathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathId

`func (o *PathChange) SetPathId(v string)`

SetPathId sets PathId field to given value.

### HasPathId

`func (o *PathChange) HasPathId() bool`

HasPathId returns a boolean if a field has been set.

### GetPath

`func (o *PathChange) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathChange) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathChange) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathChange) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetOldPath

`func (o *PathChange) GetOldPath() string`

GetOldPath returns the OldPath field if non-nil, zero value otherwise.

### GetOldPathOk

`func (o *PathChange) GetOldPathOk() (*string, bool)`

GetOldPathOk returns a tuple with the OldPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPath

`func (o *PathChange) SetOldPath(v string)`

SetOldPath sets OldPath field to given value.

### HasOldPath

`func (o *PathChange) HasOldPath() bool`

HasOldPath returns a boolean if a field has been set.

### GetRulesAdded

`func (o *PathChange) GetRulesAdded() []RuleResponse`

GetRulesAdded returns the RulesAdded field if non-nil, zero value otherwise.

### GetRulesAddedOk

`func (o *PathChange) GetRulesAddedOk() (*[]RuleResponse, bool)`

GetRulesAddedOk returns a tuple with the RulesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesAdded

`func (o *PathChange) SetRulesAdded(v []RuleResponse)`

SetRulesAdded sets RulesAdded field to given value.

### HasRulesAdded

`func (o *PathChange) HasRulesAdded() bool`

HasRulesAdded returns a boolean if a field has been set.

### GetRulesChanged

`func (o *PathChange) GetRulesChanged() []RuleChange`

GetRulesChanged returns the RulesChanged field if non-nil, zero value otherwise.

### GetRulesChangedOk

`func (o *PathChange) GetRulesChangedOk() (*[]RuleChange, bool)`

GetRulesChangedOk returns a tuple with the RulesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesChanged

`func (o *PathChange) SetRulesChanged(v []RuleChange)`

SetRulesChanged sets RulesChanged field to given value.

### HasRulesChanged

`func (o *PathChange) HasRulesChanged() bool`

HasRulesChanged returns a boolean if a field has been set.

### GetRulesDeleted

`func (o *PathChange) GetRulesDeleted() []RuleResponse`

GetRulesDeleted returns the RulesDeleted field if non-nil, zero value otherwise.

### GetRulesDeletedOk

`func (o *PathChange) GetRulesDeletedOk() (*[]RuleResponse, bool)`

GetRulesDeletedOk returns a tuple with the RulesDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesDeleted

`func (o *PathChange) SetRulesDeleted(v []RuleResponse)`

SetRulesDeleted sets RulesDeleted field to given value.

### HasRulesDeleted

`func (o *PathChange) HasRulesDeleted() bool`

HasRulesDeleted returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


