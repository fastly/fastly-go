# RuleChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | Alphanumeric string identifying the rule. Stable across versions of the routing config. | [optional] [readonly] 
**OldAction** | Pointer to [**Action**](Action.md) |  | [optional] 
**NewAction** | Pointer to [**Action**](Action.md) |  | [optional] 

## Methods

### NewRuleChange

`func NewRuleChange() *RuleChange`

NewRuleChange instantiates a new RuleChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleChangeWithDefaults

`func NewRuleChangeWithDefaults() *RuleChange`

NewRuleChangeWithDefaults instantiates a new RuleChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *RuleChange) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleChange) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleChange) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RuleChange) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetOldAction

`func (o *RuleChange) GetOldAction() Action`

GetOldAction returns the OldAction field if non-nil, zero value otherwise.

### GetOldActionOk

`func (o *RuleChange) GetOldActionOk() (*Action, bool)`

GetOldActionOk returns a tuple with the OldAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAction

`func (o *RuleChange) SetOldAction(v Action)`

SetOldAction sets OldAction field to given value.

### HasOldAction

`func (o *RuleChange) HasOldAction() bool`

HasOldAction returns a boolean if a field has been set.

### GetNewAction

`func (o *RuleChange) GetNewAction() Action`

GetNewAction returns the NewAction field if non-nil, zero value otherwise.

### GetNewActionOk

`func (o *RuleChange) GetNewActionOk() (*Action, bool)`

GetNewActionOk returns a tuple with the NewAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAction

`func (o *RuleChange) SetNewAction(v Action)`

SetNewAction sets NewAction field to given value.

### HasNewAction

`func (o *RuleChange) HasNewAction() bool`

HasNewAction returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


