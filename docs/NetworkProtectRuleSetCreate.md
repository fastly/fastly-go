# NetworkProtectRuleSetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAction** | [**NetworkProtectRuleSetDefaultAction**](NetworkProtectRuleSetDefaultAction.md) |  | 
**Rules** | Pointer to [**[]NetworkProtectRule**](NetworkProtectRule.md) | Rules in evaluation order; the first match wins. | [optional] 

## Methods

### NewNetworkProtectRuleSetCreate

`func NewNetworkProtectRuleSetCreate(defaultAction NetworkProtectRuleSetDefaultAction, ) *NetworkProtectRuleSetCreate`

NewNetworkProtectRuleSetCreate instantiates a new NetworkProtectRuleSetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleSetCreateWithDefaults

`func NewNetworkProtectRuleSetCreateWithDefaults() *NetworkProtectRuleSetCreate`

NewNetworkProtectRuleSetCreateWithDefaults instantiates a new NetworkProtectRuleSetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAction

`func (o *NetworkProtectRuleSetCreate) GetDefaultAction() NetworkProtectRuleSetDefaultAction`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *NetworkProtectRuleSetCreate) GetDefaultActionOk() (*NetworkProtectRuleSetDefaultAction, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *NetworkProtectRuleSetCreate) SetDefaultAction(v NetworkProtectRuleSetDefaultAction)`

SetDefaultAction sets DefaultAction field to given value.


### GetRules

`func (o *NetworkProtectRuleSetCreate) GetRules() []NetworkProtectRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NetworkProtectRuleSetCreate) GetRulesOk() (*[]NetworkProtectRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NetworkProtectRuleSetCreate) SetRules(v []NetworkProtectRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *NetworkProtectRuleSetCreate) HasRules() bool`

HasRules returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


