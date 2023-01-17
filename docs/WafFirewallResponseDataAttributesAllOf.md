# WafFirewallResponseDataAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**ActiveRulesFastlyBlockCount** | Pointer to **int32** | The number of active Fastly rules set to block on the active or latest firewall version. | [optional] [readonly] 
**ActiveRulesFastlyLogCount** | Pointer to **int32** | The number of active Fastly rules set to log on the active or latest firewall version. | [optional] [readonly] 
**ActiveRulesFastlyScoreCount** | Pointer to **int32** | The number of active Fastly rules set to score on the active or latest firewall version. | [optional] [readonly] 
**ActiveRulesOwaspBlockCount** | Pointer to **int32** | The number of active OWASP rules set to block on the active or latest firewall version. | [optional] [readonly] 
**ActiveRulesOwaspLogCount** | Pointer to **int32** | The number of active OWASP rules set to log on the active or latest firewall version. | [optional] [readonly] 
**ActiveRulesOwaspScoreCount** | Pointer to **int32** | The number of active OWASP rules set to score on the active or latest firewall version. | [optional] [readonly] 
**ActiveRulesTrustwaveBlockCount** | Pointer to **int32** | The number of active Trustwave rules set to block on the active or latest firewall version. | [optional] [readonly] 
**ActiveRulesTrustwaveLogCount** | Pointer to **int32** | The number of active Trustwave rules set to log on the active or latest firewall version. | [optional] [readonly] 

## Methods

### NewWafFirewallResponseDataAttributesAllOf

`func NewWafFirewallResponseDataAttributesAllOf() *WafFirewallResponseDataAttributesAllOf`

NewWafFirewallResponseDataAttributesAllOf instantiates a new WafFirewallResponseDataAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallResponseDataAttributesAllOfWithDefaults

`func NewWafFirewallResponseDataAttributesAllOfWithDefaults() *WafFirewallResponseDataAttributesAllOf`

NewWafFirewallResponseDataAttributesAllOfWithDefaults instantiates a new WafFirewallResponseDataAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceID

`func (o *WafFirewallResponseDataAttributesAllOf) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *WafFirewallResponseDataAttributesAllOf) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *WafFirewallResponseDataAttributesAllOf) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetActiveRulesFastlyBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCount() int32`

GetActiveRulesFastlyBlockCount returns the ActiveRulesFastlyBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyBlockCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCountOk() (*int32, bool)`

GetActiveRulesFastlyBlockCountOk returns a tuple with the ActiveRulesFastlyBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesFastlyBlockCount(v int32)`

SetActiveRulesFastlyBlockCount sets ActiveRulesFastlyBlockCount field to given value.

### HasActiveRulesFastlyBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesFastlyBlockCount() bool`

HasActiveRulesFastlyBlockCount returns a boolean if a field has been set.

### GetActiveRulesFastlyLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyLogCount() int32`

GetActiveRulesFastlyLogCount returns the ActiveRulesFastlyLogCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyLogCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyLogCountOk() (*int32, bool)`

GetActiveRulesFastlyLogCountOk returns a tuple with the ActiveRulesFastlyLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesFastlyLogCount(v int32)`

SetActiveRulesFastlyLogCount sets ActiveRulesFastlyLogCount field to given value.

### HasActiveRulesFastlyLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesFastlyLogCount() bool`

HasActiveRulesFastlyLogCount returns a boolean if a field has been set.

### GetActiveRulesFastlyScoreCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCount() int32`

GetActiveRulesFastlyScoreCount returns the ActiveRulesFastlyScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyScoreCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCountOk() (*int32, bool)`

GetActiveRulesFastlyScoreCountOk returns a tuple with the ActiveRulesFastlyScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyScoreCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesFastlyScoreCount(v int32)`

SetActiveRulesFastlyScoreCount sets ActiveRulesFastlyScoreCount field to given value.

### HasActiveRulesFastlyScoreCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesFastlyScoreCount() bool`

HasActiveRulesFastlyScoreCount returns a boolean if a field has been set.

### GetActiveRulesOwaspBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCount() int32`

GetActiveRulesOwaspBlockCount returns the ActiveRulesOwaspBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspBlockCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCountOk() (*int32, bool)`

GetActiveRulesOwaspBlockCountOk returns a tuple with the ActiveRulesOwaspBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesOwaspBlockCount(v int32)`

SetActiveRulesOwaspBlockCount sets ActiveRulesOwaspBlockCount field to given value.

### HasActiveRulesOwaspBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesOwaspBlockCount() bool`

HasActiveRulesOwaspBlockCount returns a boolean if a field has been set.

### GetActiveRulesOwaspLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspLogCount() int32`

GetActiveRulesOwaspLogCount returns the ActiveRulesOwaspLogCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspLogCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspLogCountOk() (*int32, bool)`

GetActiveRulesOwaspLogCountOk returns a tuple with the ActiveRulesOwaspLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesOwaspLogCount(v int32)`

SetActiveRulesOwaspLogCount sets ActiveRulesOwaspLogCount field to given value.

### HasActiveRulesOwaspLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesOwaspLogCount() bool`

HasActiveRulesOwaspLogCount returns a boolean if a field has been set.

### GetActiveRulesOwaspScoreCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCount() int32`

GetActiveRulesOwaspScoreCount returns the ActiveRulesOwaspScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspScoreCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCountOk() (*int32, bool)`

GetActiveRulesOwaspScoreCountOk returns a tuple with the ActiveRulesOwaspScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspScoreCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesOwaspScoreCount(v int32)`

SetActiveRulesOwaspScoreCount sets ActiveRulesOwaspScoreCount field to given value.

### HasActiveRulesOwaspScoreCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesOwaspScoreCount() bool`

HasActiveRulesOwaspScoreCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCount() int32`

GetActiveRulesTrustwaveBlockCount returns the ActiveRulesTrustwaveBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveBlockCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCountOk() (*int32, bool)`

GetActiveRulesTrustwaveBlockCountOk returns a tuple with the ActiveRulesTrustwaveBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesTrustwaveBlockCount(v int32)`

SetActiveRulesTrustwaveBlockCount sets ActiveRulesTrustwaveBlockCount field to given value.

### HasActiveRulesTrustwaveBlockCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesTrustwaveBlockCount() bool`

HasActiveRulesTrustwaveBlockCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCount() int32`

GetActiveRulesTrustwaveLogCount returns the ActiveRulesTrustwaveLogCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveLogCountOk

`func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCountOk() (*int32, bool)`

GetActiveRulesTrustwaveLogCountOk returns a tuple with the ActiveRulesTrustwaveLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesTrustwaveLogCount(v int32)`

SetActiveRulesTrustwaveLogCount sets ActiveRulesTrustwaveLogCount field to given value.

### HasActiveRulesTrustwaveLogCount

`func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesTrustwaveLogCount() bool`

HasActiveRulesTrustwaveLogCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
