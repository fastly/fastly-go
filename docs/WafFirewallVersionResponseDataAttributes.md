# WafFirewallVersionResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
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

### NewWafFirewallVersionResponseDataAttributes

`func NewWafFirewallVersionResponseDataAttributes() *WafFirewallVersionResponseDataAttributes`

NewWafFirewallVersionResponseDataAttributes instantiates a new WafFirewallVersionResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionResponseDataAttributesWithDefaults

`func NewWafFirewallVersionResponseDataAttributesWithDefaults() *WafFirewallVersionResponseDataAttributes`

NewWafFirewallVersionResponseDataAttributesWithDefaults instantiates a new WafFirewallVersionResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WafFirewallVersionResponseDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WafFirewallVersionResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WafFirewallVersionResponseDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WafFirewallVersionResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *WafFirewallVersionResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WafFirewallVersionResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *WafFirewallVersionResponseDataAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WafFirewallVersionResponseDataAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WafFirewallVersionResponseDataAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WafFirewallVersionResponseDataAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *WafFirewallVersionResponseDataAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *WafFirewallVersionResponseDataAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *WafFirewallVersionResponseDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WafFirewallVersionResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WafFirewallVersionResponseDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WafFirewallVersionResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *WafFirewallVersionResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WafFirewallVersionResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetActive

`func (o *WafFirewallVersionResponseDataAttributes) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WafFirewallVersionResponseDataAttributes) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WafFirewallVersionResponseDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActiveRulesFastlyBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyBlockCount() int32`

GetActiveRulesFastlyBlockCount returns the ActiveRulesFastlyBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyBlockCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyBlockCountOk() (*int32, bool)`

GetActiveRulesFastlyBlockCountOk returns a tuple with the ActiveRulesFastlyBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesFastlyBlockCount(v int32)`

SetActiveRulesFastlyBlockCount sets ActiveRulesFastlyBlockCount field to given value.

### HasActiveRulesFastlyBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesFastlyBlockCount() bool`

HasActiveRulesFastlyBlockCount returns a boolean if a field has been set.

### GetActiveRulesFastlyLogCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyLogCount() int32`

GetActiveRulesFastlyLogCount returns the ActiveRulesFastlyLogCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyLogCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyLogCountOk() (*int32, bool)`

GetActiveRulesFastlyLogCountOk returns a tuple with the ActiveRulesFastlyLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyLogCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesFastlyLogCount(v int32)`

SetActiveRulesFastlyLogCount sets ActiveRulesFastlyLogCount field to given value.

### HasActiveRulesFastlyLogCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesFastlyLogCount() bool`

HasActiveRulesFastlyLogCount returns a boolean if a field has been set.

### GetActiveRulesFastlyScoreCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyScoreCount() int32`

GetActiveRulesFastlyScoreCount returns the ActiveRulesFastlyScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyScoreCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyScoreCountOk() (*int32, bool)`

GetActiveRulesFastlyScoreCountOk returns a tuple with the ActiveRulesFastlyScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyScoreCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesFastlyScoreCount(v int32)`

SetActiveRulesFastlyScoreCount sets ActiveRulesFastlyScoreCount field to given value.

### HasActiveRulesFastlyScoreCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesFastlyScoreCount() bool`

HasActiveRulesFastlyScoreCount returns a boolean if a field has been set.

### GetActiveRulesOwaspBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspBlockCount() int32`

GetActiveRulesOwaspBlockCount returns the ActiveRulesOwaspBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspBlockCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspBlockCountOk() (*int32, bool)`

GetActiveRulesOwaspBlockCountOk returns a tuple with the ActiveRulesOwaspBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesOwaspBlockCount(v int32)`

SetActiveRulesOwaspBlockCount sets ActiveRulesOwaspBlockCount field to given value.

### HasActiveRulesOwaspBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesOwaspBlockCount() bool`

HasActiveRulesOwaspBlockCount returns a boolean if a field has been set.

### GetActiveRulesOwaspLogCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspLogCount() int32`

GetActiveRulesOwaspLogCount returns the ActiveRulesOwaspLogCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspLogCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspLogCountOk() (*int32, bool)`

GetActiveRulesOwaspLogCountOk returns a tuple with the ActiveRulesOwaspLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspLogCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesOwaspLogCount(v int32)`

SetActiveRulesOwaspLogCount sets ActiveRulesOwaspLogCount field to given value.

### HasActiveRulesOwaspLogCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesOwaspLogCount() bool`

HasActiveRulesOwaspLogCount returns a boolean if a field has been set.

### GetActiveRulesOwaspScoreCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspScoreCount() int32`

GetActiveRulesOwaspScoreCount returns the ActiveRulesOwaspScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspScoreCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspScoreCountOk() (*int32, bool)`

GetActiveRulesOwaspScoreCountOk returns a tuple with the ActiveRulesOwaspScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspScoreCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesOwaspScoreCount(v int32)`

SetActiveRulesOwaspScoreCount sets ActiveRulesOwaspScoreCount field to given value.

### HasActiveRulesOwaspScoreCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesOwaspScoreCount() bool`

HasActiveRulesOwaspScoreCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveBlockCount() int32`

GetActiveRulesTrustwaveBlockCount returns the ActiveRulesTrustwaveBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveBlockCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveBlockCountOk() (*int32, bool)`

GetActiveRulesTrustwaveBlockCountOk returns a tuple with the ActiveRulesTrustwaveBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesTrustwaveBlockCount(v int32)`

SetActiveRulesTrustwaveBlockCount sets ActiveRulesTrustwaveBlockCount field to given value.

### HasActiveRulesTrustwaveBlockCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesTrustwaveBlockCount() bool`

HasActiveRulesTrustwaveBlockCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveLogCount

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveLogCount() int32`

GetActiveRulesTrustwaveLogCount returns the ActiveRulesTrustwaveLogCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveLogCountOk

`func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveLogCountOk() (*int32, bool)`

GetActiveRulesTrustwaveLogCountOk returns a tuple with the ActiveRulesTrustwaveLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveLogCount

`func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesTrustwaveLogCount(v int32)`

SetActiveRulesTrustwaveLogCount sets ActiveRulesTrustwaveLogCount field to given value.

### HasActiveRulesTrustwaveLogCount

`func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesTrustwaveLogCount() bool`

HasActiveRulesTrustwaveLogCount returns a boolean if a field has been set.

### GetLastDeploymentStatus

`func (o *WafFirewallVersionResponseDataAttributes) GetLastDeploymentStatus() string`

GetLastDeploymentStatus returns the LastDeploymentStatus field if non-nil, zero value otherwise.

### GetLastDeploymentStatusOk

`func (o *WafFirewallVersionResponseDataAttributes) GetLastDeploymentStatusOk() (*string, bool)`

GetLastDeploymentStatusOk returns a tuple with the LastDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentStatus

`func (o *WafFirewallVersionResponseDataAttributes) SetLastDeploymentStatus(v string)`

SetLastDeploymentStatus sets LastDeploymentStatus field to given value.

### HasLastDeploymentStatus

`func (o *WafFirewallVersionResponseDataAttributes) HasLastDeploymentStatus() bool`

HasLastDeploymentStatus returns a boolean if a field has been set.

### SetLastDeploymentStatusNil

`func (o *WafFirewallVersionResponseDataAttributes) SetLastDeploymentStatusNil(b bool)`

 SetLastDeploymentStatusNil sets the value for LastDeploymentStatus to be an explicit nil

### UnsetLastDeploymentStatus
`func (o *WafFirewallVersionResponseDataAttributes) UnsetLastDeploymentStatus()`

UnsetLastDeploymentStatus ensures that no value is present for LastDeploymentStatus, not even an explicit nil
### GetDeployedAt

`func (o *WafFirewallVersionResponseDataAttributes) GetDeployedAt() string`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *WafFirewallVersionResponseDataAttributes) GetDeployedAtOk() (*string, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *WafFirewallVersionResponseDataAttributes) SetDeployedAt(v string)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *WafFirewallVersionResponseDataAttributes) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetError

`func (o *WafFirewallVersionResponseDataAttributes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WafFirewallVersionResponseDataAttributes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WafFirewallVersionResponseDataAttributes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WafFirewallVersionResponseDataAttributes) HasError() bool`

HasError returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
