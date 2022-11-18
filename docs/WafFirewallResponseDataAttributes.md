# WafFirewallResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
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

### NewWafFirewallResponseDataAttributes

`func NewWafFirewallResponseDataAttributes() *WafFirewallResponseDataAttributes`

NewWafFirewallResponseDataAttributes instantiates a new WafFirewallResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallResponseDataAttributesWithDefaults

`func NewWafFirewallResponseDataAttributesWithDefaults() *WafFirewallResponseDataAttributes`

NewWafFirewallResponseDataAttributesWithDefaults instantiates a new WafFirewallResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WafFirewallResponseDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WafFirewallResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WafFirewallResponseDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WafFirewallResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *WafFirewallResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WafFirewallResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *WafFirewallResponseDataAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WafFirewallResponseDataAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WafFirewallResponseDataAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WafFirewallResponseDataAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *WafFirewallResponseDataAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *WafFirewallResponseDataAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *WafFirewallResponseDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WafFirewallResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WafFirewallResponseDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WafFirewallResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *WafFirewallResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WafFirewallResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *WafFirewallResponseDataAttributes) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *WafFirewallResponseDataAttributes) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *WafFirewallResponseDataAttributes) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *WafFirewallResponseDataAttributes) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetActiveRulesFastlyBlockCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesFastlyBlockCount() int32`

GetActiveRulesFastlyBlockCount returns the ActiveRulesFastlyBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyBlockCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesFastlyBlockCountOk() (*int32, bool)`

GetActiveRulesFastlyBlockCountOk returns a tuple with the ActiveRulesFastlyBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyBlockCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesFastlyBlockCount(v int32)`

SetActiveRulesFastlyBlockCount sets ActiveRulesFastlyBlockCount field to given value.

### HasActiveRulesFastlyBlockCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesFastlyBlockCount() bool`

HasActiveRulesFastlyBlockCount returns a boolean if a field has been set.

### GetActiveRulesFastlyLogCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesFastlyLogCount() int32`

GetActiveRulesFastlyLogCount returns the ActiveRulesFastlyLogCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyLogCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesFastlyLogCountOk() (*int32, bool)`

GetActiveRulesFastlyLogCountOk returns a tuple with the ActiveRulesFastlyLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyLogCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesFastlyLogCount(v int32)`

SetActiveRulesFastlyLogCount sets ActiveRulesFastlyLogCount field to given value.

### HasActiveRulesFastlyLogCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesFastlyLogCount() bool`

HasActiveRulesFastlyLogCount returns a boolean if a field has been set.

### GetActiveRulesFastlyScoreCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesFastlyScoreCount() int32`

GetActiveRulesFastlyScoreCount returns the ActiveRulesFastlyScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesFastlyScoreCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesFastlyScoreCountOk() (*int32, bool)`

GetActiveRulesFastlyScoreCountOk returns a tuple with the ActiveRulesFastlyScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesFastlyScoreCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesFastlyScoreCount(v int32)`

SetActiveRulesFastlyScoreCount sets ActiveRulesFastlyScoreCount field to given value.

### HasActiveRulesFastlyScoreCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesFastlyScoreCount() bool`

HasActiveRulesFastlyScoreCount returns a boolean if a field has been set.

### GetActiveRulesOwaspBlockCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesOwaspBlockCount() int32`

GetActiveRulesOwaspBlockCount returns the ActiveRulesOwaspBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspBlockCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesOwaspBlockCountOk() (*int32, bool)`

GetActiveRulesOwaspBlockCountOk returns a tuple with the ActiveRulesOwaspBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspBlockCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesOwaspBlockCount(v int32)`

SetActiveRulesOwaspBlockCount sets ActiveRulesOwaspBlockCount field to given value.

### HasActiveRulesOwaspBlockCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesOwaspBlockCount() bool`

HasActiveRulesOwaspBlockCount returns a boolean if a field has been set.

### GetActiveRulesOwaspLogCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesOwaspLogCount() int32`

GetActiveRulesOwaspLogCount returns the ActiveRulesOwaspLogCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspLogCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesOwaspLogCountOk() (*int32, bool)`

GetActiveRulesOwaspLogCountOk returns a tuple with the ActiveRulesOwaspLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspLogCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesOwaspLogCount(v int32)`

SetActiveRulesOwaspLogCount sets ActiveRulesOwaspLogCount field to given value.

### HasActiveRulesOwaspLogCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesOwaspLogCount() bool`

HasActiveRulesOwaspLogCount returns a boolean if a field has been set.

### GetActiveRulesOwaspScoreCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesOwaspScoreCount() int32`

GetActiveRulesOwaspScoreCount returns the ActiveRulesOwaspScoreCount field if non-nil, zero value otherwise.

### GetActiveRulesOwaspScoreCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesOwaspScoreCountOk() (*int32, bool)`

GetActiveRulesOwaspScoreCountOk returns a tuple with the ActiveRulesOwaspScoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesOwaspScoreCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesOwaspScoreCount(v int32)`

SetActiveRulesOwaspScoreCount sets ActiveRulesOwaspScoreCount field to given value.

### HasActiveRulesOwaspScoreCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesOwaspScoreCount() bool`

HasActiveRulesOwaspScoreCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesTrustwaveBlockCount() int32`

GetActiveRulesTrustwaveBlockCount returns the ActiveRulesTrustwaveBlockCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveBlockCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesTrustwaveBlockCountOk() (*int32, bool)`

GetActiveRulesTrustwaveBlockCountOk returns a tuple with the ActiveRulesTrustwaveBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveBlockCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesTrustwaveBlockCount(v int32)`

SetActiveRulesTrustwaveBlockCount sets ActiveRulesTrustwaveBlockCount field to given value.

### HasActiveRulesTrustwaveBlockCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesTrustwaveBlockCount() bool`

HasActiveRulesTrustwaveBlockCount returns a boolean if a field has been set.

### GetActiveRulesTrustwaveLogCount

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesTrustwaveLogCount() int32`

GetActiveRulesTrustwaveLogCount returns the ActiveRulesTrustwaveLogCount field if non-nil, zero value otherwise.

### GetActiveRulesTrustwaveLogCountOk

`func (o *WafFirewallResponseDataAttributes) GetActiveRulesTrustwaveLogCountOk() (*int32, bool)`

GetActiveRulesTrustwaveLogCountOk returns a tuple with the ActiveRulesTrustwaveLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesTrustwaveLogCount

`func (o *WafFirewallResponseDataAttributes) SetActiveRulesTrustwaveLogCount(v int32)`

SetActiveRulesTrustwaveLogCount sets ActiveRulesTrustwaveLogCount field to given value.

### HasActiveRulesTrustwaveLogCount

`func (o *WafFirewallResponseDataAttributes) HasActiveRulesTrustwaveLogCount() bool`

HasActiveRulesTrustwaveLogCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
