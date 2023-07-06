# LegacyWafFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPush** | Pointer to **string** | Date and time that VCL was last pushed to cache nodes. | [optional] 
**PrefetchCondition** | Pointer to **string** | Name of the corresponding condition object. | [optional] 
**Response** | Pointer to **string** | Name of the corresponding response object. | [optional] 
**Version** | Pointer to [**ReadOnlyVersion**](ReadOnlyVersion.md) |  | [optional] 
**ServiceID** | Pointer to [**ReadOnlyServiceID**](ReadOnlyServiceID.md) |  | [optional] 
**Disabled** | Pointer to **bool** | The status of the firewall. | [optional] [default to false]
**RuleStatusesLogCount** | Pointer to **int32** | The number of rule statuses set to log. | [optional] 
**RuleStatusesBlockCount** | Pointer to **int32** | The number of rule statuses set to block. | [optional] 
**RuleStatusesDisabledCount** | Pointer to **int32** | The number of rule statuses set to disabled. | [optional] 

## Methods

### NewLegacyWafFirewall

`func NewLegacyWafFirewall() *LegacyWafFirewall`

NewLegacyWafFirewall instantiates a new LegacyWafFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWafFirewallWithDefaults

`func NewLegacyWafFirewallWithDefaults() *LegacyWafFirewall`

NewLegacyWafFirewallWithDefaults instantiates a new LegacyWafFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPush

`func (o *LegacyWafFirewall) GetLastPush() string`

GetLastPush returns the LastPush field if non-nil, zero value otherwise.

### GetLastPushOk

`func (o *LegacyWafFirewall) GetLastPushOk() (*string, bool)`

GetLastPushOk returns a tuple with the LastPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPush

`func (o *LegacyWafFirewall) SetLastPush(v string)`

SetLastPush sets LastPush field to given value.

### HasLastPush

`func (o *LegacyWafFirewall) HasLastPush() bool`

HasLastPush returns a boolean if a field has been set.

### GetPrefetchCondition

`func (o *LegacyWafFirewall) GetPrefetchCondition() string`

GetPrefetchCondition returns the PrefetchCondition field if non-nil, zero value otherwise.

### GetPrefetchConditionOk

`func (o *LegacyWafFirewall) GetPrefetchConditionOk() (*string, bool)`

GetPrefetchConditionOk returns a tuple with the PrefetchCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetchCondition

`func (o *LegacyWafFirewall) SetPrefetchCondition(v string)`

SetPrefetchCondition sets PrefetchCondition field to given value.

### HasPrefetchCondition

`func (o *LegacyWafFirewall) HasPrefetchCondition() bool`

HasPrefetchCondition returns a boolean if a field has been set.

### GetResponse

`func (o *LegacyWafFirewall) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *LegacyWafFirewall) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *LegacyWafFirewall) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *LegacyWafFirewall) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetVersion

`func (o *LegacyWafFirewall) GetVersion() ReadOnlyVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LegacyWafFirewall) GetVersionOk() (*ReadOnlyVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LegacyWafFirewall) SetVersion(v ReadOnlyVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LegacyWafFirewall) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetServiceID

`func (o *LegacyWafFirewall) GetServiceID() ReadOnlyServiceID`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LegacyWafFirewall) GetServiceIDOk() (*ReadOnlyServiceID, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LegacyWafFirewall) SetServiceID(v ReadOnlyServiceID)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LegacyWafFirewall) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetDisabled

`func (o *LegacyWafFirewall) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *LegacyWafFirewall) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *LegacyWafFirewall) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *LegacyWafFirewall) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetRuleStatusesLogCount

`func (o *LegacyWafFirewall) GetRuleStatusesLogCount() int32`

GetRuleStatusesLogCount returns the RuleStatusesLogCount field if non-nil, zero value otherwise.

### GetRuleStatusesLogCountOk

`func (o *LegacyWafFirewall) GetRuleStatusesLogCountOk() (*int32, bool)`

GetRuleStatusesLogCountOk returns a tuple with the RuleStatusesLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleStatusesLogCount

`func (o *LegacyWafFirewall) SetRuleStatusesLogCount(v int32)`

SetRuleStatusesLogCount sets RuleStatusesLogCount field to given value.

### HasRuleStatusesLogCount

`func (o *LegacyWafFirewall) HasRuleStatusesLogCount() bool`

HasRuleStatusesLogCount returns a boolean if a field has been set.

### GetRuleStatusesBlockCount

`func (o *LegacyWafFirewall) GetRuleStatusesBlockCount() int32`

GetRuleStatusesBlockCount returns the RuleStatusesBlockCount field if non-nil, zero value otherwise.

### GetRuleStatusesBlockCountOk

`func (o *LegacyWafFirewall) GetRuleStatusesBlockCountOk() (*int32, bool)`

GetRuleStatusesBlockCountOk returns a tuple with the RuleStatusesBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleStatusesBlockCount

`func (o *LegacyWafFirewall) SetRuleStatusesBlockCount(v int32)`

SetRuleStatusesBlockCount sets RuleStatusesBlockCount field to given value.

### HasRuleStatusesBlockCount

`func (o *LegacyWafFirewall) HasRuleStatusesBlockCount() bool`

HasRuleStatusesBlockCount returns a boolean if a field has been set.

### GetRuleStatusesDisabledCount

`func (o *LegacyWafFirewall) GetRuleStatusesDisabledCount() int32`

GetRuleStatusesDisabledCount returns the RuleStatusesDisabledCount field if non-nil, zero value otherwise.

### GetRuleStatusesDisabledCountOk

`func (o *LegacyWafFirewall) GetRuleStatusesDisabledCountOk() (*int32, bool)`

GetRuleStatusesDisabledCountOk returns a tuple with the RuleStatusesDisabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleStatusesDisabledCount

`func (o *LegacyWafFirewall) SetRuleStatusesDisabledCount(v int32)`

SetRuleStatusesDisabledCount sets RuleStatusesDisabledCount field to given value.

### HasRuleStatusesDisabledCount

`func (o *LegacyWafFirewall) HasRuleStatusesDisabledCount() bool`

HasRuleStatusesDisabledCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
