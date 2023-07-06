# LegacyWafRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPush** | Pointer to **string** | Date and time WAF ruleset VCL was last deployed. | [optional] 
**Vcl** | Pointer to **string** | The WAF ruleset VCL currently deployed. | [optional] 

## Methods

### NewLegacyWafRuleset

`func NewLegacyWafRuleset() *LegacyWafRuleset`

NewLegacyWafRuleset instantiates a new LegacyWafRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWafRulesetWithDefaults

`func NewLegacyWafRulesetWithDefaults() *LegacyWafRuleset`

NewLegacyWafRulesetWithDefaults instantiates a new LegacyWafRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPush

`func (o *LegacyWafRuleset) GetLastPush() string`

GetLastPush returns the LastPush field if non-nil, zero value otherwise.

### GetLastPushOk

`func (o *LegacyWafRuleset) GetLastPushOk() (*string, bool)`

GetLastPushOk returns a tuple with the LastPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPush

`func (o *LegacyWafRuleset) SetLastPush(v string)`

SetLastPush sets LastPush field to given value.

### HasLastPush

`func (o *LegacyWafRuleset) HasLastPush() bool`

HasLastPush returns a boolean if a field has been set.

### GetVcl

`func (o *LegacyWafRuleset) GetVcl() string`

GetVcl returns the Vcl field if non-nil, zero value otherwise.

### GetVclOk

`func (o *LegacyWafRuleset) GetVclOk() (*string, bool)`

GetVclOk returns a tuple with the Vcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcl

`func (o *LegacyWafRuleset) SetVcl(v string)`

SetVcl sets Vcl field to given value.

### HasVcl

`func (o *LegacyWafRuleset) HasVcl() bool`

HasVcl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
