# WafFirewallVersionResponseDataAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether a specific firewall version is currently deployed. | [optional] [readonly] 
**ActiveRulesFastlyBlockCount** | Pointer to **int32** | The number of active Fastly rules set to block. | [optional] [readonly] 
**ActiveRulesFastlyLogCount** | Pointer to **int32** | The number of active Fastly rules set to log. | [optional] [readonly] 
**ActiveRulesFastlyScoreCount** | Pointer to **int32** | The number of active Fastly rules set to score. | [optional] [readonly] 
**ActiveRulesOwaspBlockCount** | Pointer to **int32** | The number of active OWASP rules set to block. | [optional] [readonly] 
**ActiveRulesOwaspLogCount** | Pointer to **int32** | The number of active OWASP rules set to log. | [optional] [readonly] 
**ActiveRulesOwaspScoreCount** | Pointer to **int32** | The number of active OWASP rules set to score. | [optional] [readonly] 
**ActiveRulesTrustwaveBlockCount** | Pointer to **int32** | The number of active Trustwave rules set to block. | [optional] [readonly] 
**ActiveRulesTrustwaveLogCount** | Pointer to **int32** | The number of active Trustwave rules set to log. | [optional] [readonly] 
**LastDeploymentStatus** | Pointer to **NullableString** | The status of the last deployment of this firewall version. | [optional] [readonly] 
**DeployedAt** | Pointer to **string** | Time-stamp (GMT) indicating when the firewall version was last deployed. | [optional] [readonly] 
**Error** | Pointer to **string** | Contains error message if the firewall version fails to deploy. | [optional] [readonly] 

## Methods

### NewWafFirewallVersionResponseDataAttributesAllOf

`func NewWafFirewallVersionResponseDataAttributesAllOf() *WafFirewallVersionResponseDataAttributesAllOf`

NewWafFirewallVersionResponseDataAttributesAllOf instantiates a new WafFirewallVersionResponseDataAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionResponseDataAttributesAllOfWithDefaults

`func NewWafFirewallVersionResponseDataAttributesAllOfWithDefaults() *WafFirewallVersionResponseDataAttributesAllOf`

NewWafFirewallVersionResponseDataAttributesAllOfWithDefaults instantiates a new WafFirewallVersionResponseDataAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActiveRulesFastlyBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCount() int32`

GetActiveRulesFastlyBlockCount returns the ActiveRulesFastlyBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyBlockCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCountOk() (*int32, bool)`

GetActiveRulesFastlyBlockCountOk returns a tuple with the ActiveRulesFastlyBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesFastlyBlockCount(v int32)`

SetActiveRulesFastlyBlockCount sets ActiveRulesFastlyBlockCount field to given value.

### HasActiveRulesFastlyBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesFastlyBlockCount() bool`

HasActiveRulesFastlyBlockCount returns a boolean if a field has been set.

### GetActiveRulesFastlyLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyLogCount() int32`

GetActiveRulesFastlyLogCount returns the ActiveRulesFastlyLogCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyLogCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyLogCountOk() (*int32, bool)`

GetActiveRulesFastlyLogCountOk returns a tuple with the ActiveRulesFastlyLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesFastlyLogCount(v int32)`

SetActiveRulesFastlyLogCount sets ActiveRulesFastlyLogCount field to given value.

### HasActiveRulesFastlyLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesFastlyLogCount() bool`

HasActiveRulesFastlyLogCount returns a boolean if a field has been set.

### GetActiveRulesFastlyScoreCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCount() int32`

GetActiveRulesFastlyScoreCount returns the ActiveRulesFastlyScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyScoreCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCountOk() (*int32, bool)`

GetActiveRulesFastlyScoreCountOk returns a tuple with the ActiveRulesFastlyScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyScoreCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesFastlyScoreCount(v int32)`

SetActiveRulesFastlyScoreCount sets ActiveRulesFastlyScoreCount field to given value.

### HasActiveRulesFastlyScoreCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesFastlyScoreCount() bool`

HasActiveRulesFastlyScoreCount returns a boolean if a field has been set.

### GetActiveRulesOwaspBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCount() int32`

GetActiveRulesOwaspBlockCount returns the ActiveRulesOwaspBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspBlockCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCountOk() (*int32, bool)`

GetActiveRulesOwaspBlockCountOk returns a tuple with the ActiveRulesOwaspBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesOwaspBlockCount(v int32)`

SetActiveRulesOwaspBlockCount sets ActiveRulesOwaspBlockCount field to given value.

### HasActiveRulesOwaspBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesOwaspBlockCount() bool`

HasActiveRulesOwaspBlockCount returns a boolean if a field has been set.

### GetActiveRulesOwaspLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspLogCount() int32`

GetActiveRulesOwaspLogCount returns the ActiveRulesOwaspLogCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspLogCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspLogCountOk() (*int32, bool)`

GetActiveRulesOwaspLogCountOk returns a tuple with the ActiveRulesOwaspLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesOwaspLogCount(v int32)`

SetActiveRulesOwaspLogCount sets ActiveRulesOwaspLogCount field to given value.

### HasActiveRulesOwaspLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesOwaspLogCount() bool`

HasActiveRulesOwaspLogCount returns a boolean if a field has been set.

### GetActiveRulesOwaspScoreCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCount() int32`

GetActiveRulesOwaspScoreCount returns the ActiveRulesOwaspScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspScoreCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCountOk() (*int32, bool)`

GetActiveRulesOwaspScoreCountOk returns a tuple with the ActiveRulesOwaspScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspScoreCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesOwaspScoreCount(v int32)`

SetActiveRulesOwaspScoreCount sets ActiveRulesOwaspScoreCount field to given value.

### HasActiveRulesOwaspScoreCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesOwaspScoreCount() bool`

HasActiveRulesOwaspScoreCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCount() int32`

GetActiveRulesTrustwaveBlockCount returns the ActiveRulesTrustwaveBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveBlockCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCountOk() (*int32, bool)`

GetActiveRulesTrustwaveBlockCountOk returns a tuple with the ActiveRulesTrustwaveBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesTrustwaveBlockCount(v int32)`

SetActiveRulesTrustwaveBlockCount sets ActiveRulesTrustwaveBlockCount field to given value.

### HasActiveRulesTrustwaveBlockCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesTrustwaveBlockCount() bool`

HasActiveRulesTrustwaveBlockCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCount() int32`

GetActiveRulesTrustwaveLogCount returns the ActiveRulesTrustwaveLogCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveLogCountOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCountOk() (*int32, bool)`

GetActiveRulesTrustwaveLogCountOk returns a tuple with the ActiveRulesTrustwaveLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesTrustwaveLogCount(v int32)`

SetActiveRulesTrustwaveLogCount sets ActiveRulesTrustwaveLogCount field to given value.

### HasActiveRulesTrustwaveLogCount

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesTrustwaveLogCount() bool`

HasActiveRulesTrustwaveLogCount returns a boolean if a field has been set.

### GetLastDeploymentStatus

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetLastDeploymentStatus() string`

GetLastDeploymentStatus returns the LastDeploymentStatus field if non-nil, zero value otherwise.

### GetLastDeploymentStatusOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetLastDeploymentStatusOk() (*string, bool)`

GetLastDeploymentStatusOk returns a tuple with the LastDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentStatus

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetLastDeploymentStatus(v string)`

SetLastDeploymentStatus sets LastDeploymentStatus field to given value.

### HasLastDeploymentStatus

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasLastDeploymentStatus() bool`

HasLastDeploymentStatus returns a boolean if a field has been set.

### SetLastDeploymentStatusNil

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetLastDeploymentStatusNil(b bool)`

 SetLastDeploymentStatusNil sets the value for LastDeploymentStatus to be an explicit nil

### UnsetLastDeploymentStatus
`func (o *WafFirewallVersionResponseDataAttributesAllOf) UnsetLastDeploymentStatus()`

UnsetLastDeploymentStatus ensures that no value is present for LastDeploymentStatus, not even an explicit nil
### GetDeployedAt

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetDeployedAt() string`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetDeployedAtOk() (*string, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetDeployedAt(v string)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetError

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WafFirewallVersionResponseDataAttributesAllOf) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WafFirewallVersionResponseDataAttributesAllOf) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WafFirewallVersionResponseDataAttributesAllOf) HasError() bool`

HasError returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
