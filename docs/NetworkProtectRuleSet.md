# NetworkProtectRuleSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAction** | [**NetworkProtectRuleSetDefaultAction**](NetworkProtectRuleSetDefaultAction.md) |  | 
**Rules** | [**[]NetworkProtectRuleResponse**](NetworkProtectRuleResponse.md) | Rules in evaluation order with read-only timestamps. | 

## Methods

### NewNetworkProtectRuleSet

`func NewNetworkProtectRuleSet(defaultAction NetworkProtectRuleSetDefaultAction, rules []NetworkProtectRuleResponse, ) *NetworkProtectRuleSet`

NewNetworkProtectRuleSet instantiates a new NetworkProtectRuleSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleSetWithDefaults

`func NewNetworkProtectRuleSetWithDefaults() *NetworkProtectRuleSet`

NewNetworkProtectRuleSetWithDefaults instantiates a new NetworkProtectRuleSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAction

`func (o *NetworkProtectRuleSet) GetDefaultAction() NetworkProtectRuleSetDefaultAction`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *NetworkProtectRuleSet) GetDefaultActionOk() (*NetworkProtectRuleSetDefaultAction, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *NetworkProtectRuleSet) SetDefaultAction(v NetworkProtectRuleSetDefaultAction)`

SetDefaultAction sets DefaultAction field to given value.


### GetRules

`func (o *NetworkProtectRuleSet) GetRules() []NetworkProtectRuleResponse`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NetworkProtectRuleSet) GetRulesOk() (*[]NetworkProtectRuleResponse, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NetworkProtectRuleSet) SetRules(v []NetworkProtectRuleResponse)`

SetRules sets Rules field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


