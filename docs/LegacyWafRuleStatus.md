# LegacyWafRuleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Describes the behavior for the particular rule within this firewall. Permitted values: `log`, `block`, and `disabled`.  | [optional] 
**ModsecRuleID** | Pointer to **string** | The ModSecurity rule ID. | [optional] 
**UniqueRuleID** | Pointer to **string** | The Rule ID. | [optional] 

## Methods

### NewLegacyWafRuleStatus

`func NewLegacyWafRuleStatus() *LegacyWafRuleStatus`

NewLegacyWafRuleStatus instantiates a new LegacyWafRuleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWafRuleStatusWithDefaults

`func NewLegacyWafRuleStatusWithDefaults() *LegacyWafRuleStatus`

NewLegacyWafRuleStatusWithDefaults instantiates a new LegacyWafRuleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LegacyWafRuleStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegacyWafRuleStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegacyWafRuleStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegacyWafRuleStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetModsecRuleID

`func (o *LegacyWafRuleStatus) GetModsecRuleID() string`

GetModsecRuleID returns the ModsecRuleID field if non-nil, zero value otherwise.

### GetModsecRuleIDOk

`func (o *LegacyWafRuleStatus) GetModsecRuleIDOk() (*string, bool)`

GetModsecRuleIDOk returns a tuple with the ModsecRuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModsecRuleID

`func (o *LegacyWafRuleStatus) SetModsecRuleID(v string)`

SetModsecRuleID sets ModsecRuleID field to given value.

### HasModsecRuleID

`func (o *LegacyWafRuleStatus) HasModsecRuleID() bool`

HasModsecRuleID returns a boolean if a field has been set.

### GetUniqueRuleID

`func (o *LegacyWafRuleStatus) GetUniqueRuleID() string`

GetUniqueRuleID returns the UniqueRuleID field if non-nil, zero value otherwise.

### GetUniqueRuleIDOk

`func (o *LegacyWafRuleStatus) GetUniqueRuleIDOk() (*string, bool)`

GetUniqueRuleIDOk returns a tuple with the UniqueRuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueRuleID

`func (o *LegacyWafRuleStatus) SetUniqueRuleID(v string)`

SetUniqueRuleID sets UniqueRuleID field to given value.

### HasUniqueRuleID

`func (o *LegacyWafRuleStatus) HasUniqueRuleID() bool`

HasUniqueRuleID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
