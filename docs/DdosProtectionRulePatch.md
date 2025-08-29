# DdosProtectionRulePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action types for a rule. Supported action values are default, block, log, off. The default action value follows the current protection mode of the associated service. | [optional] [default to "default"]

## Methods

### NewDdosProtectionRulePatch

`func NewDdosProtectionRulePatch() *DdosProtectionRulePatch`

NewDdosProtectionRulePatch instantiates a new DdosProtectionRulePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionRulePatchWithDefaults

`func NewDdosProtectionRulePatchWithDefaults() *DdosProtectionRulePatch`

NewDdosProtectionRulePatchWithDefaults instantiates a new DdosProtectionRulePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DdosProtectionRulePatch) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DdosProtectionRulePatch) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DdosProtectionRulePatch) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DdosProtectionRulePatch) HasAction() bool`

HasAction returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
