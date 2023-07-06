# LegacyWafRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message metadata for the rule. | [optional] 
**RuleID** | Pointer to **string** | Corresponding ModSecurity rule ID. | [optional] 
**Severity** | Pointer to **int32** | Severity metadata for the rule. | [optional] 
**Source** | Pointer to **string** | The ModSecurity rule logic. | [optional] 
**Vcl** | Pointer to **string** | The VCL representation of the rule logic. | [optional] 

## Methods

### NewLegacyWafRule

`func NewLegacyWafRule() *LegacyWafRule`

NewLegacyWafRule instantiates a new LegacyWafRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWafRuleWithDefaults

`func NewLegacyWafRuleWithDefaults() *LegacyWafRule`

NewLegacyWafRuleWithDefaults instantiates a new LegacyWafRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LegacyWafRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LegacyWafRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LegacyWafRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LegacyWafRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRuleID

`func (o *LegacyWafRule) GetRuleID() string`

GetRuleID returns the RuleID field if non-nil, zero value otherwise.

### GetRuleIDOk

`func (o *LegacyWafRule) GetRuleIDOk() (*string, bool)`

GetRuleIDOk returns a tuple with the RuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleID

`func (o *LegacyWafRule) SetRuleID(v string)`

SetRuleID sets RuleID field to given value.

### HasRuleID

`func (o *LegacyWafRule) HasRuleID() bool`

HasRuleID returns a boolean if a field has been set.

### GetSeverity

`func (o *LegacyWafRule) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *LegacyWafRule) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *LegacyWafRule) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *LegacyWafRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *LegacyWafRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LegacyWafRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LegacyWafRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LegacyWafRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVcl

`func (o *LegacyWafRule) GetVcl() string`

GetVcl returns the Vcl field if non-nil, zero value otherwise.

### GetVclOk

`func (o *LegacyWafRule) GetVclOk() (*string, bool)`

GetVclOk returns a tuple with the Vcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcl

`func (o *LegacyWafRule) SetVcl(v string)`

SetVcl sets Vcl field to given value.

### HasVcl

`func (o *LegacyWafRule) HasVcl() bool`

HasVcl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
