# WafRuleRevisionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message metadata for the rule. | [optional] [readonly] 
**ModsecRuleID** | Pointer to **int32** | Corresponding ModSecurity rule ID. | [optional] [readonly] 
**ParanoiaLevel** | Pointer to **int32** | Paranoia level for the rule. | [optional] [readonly] 
**Revision** | Pointer to **int32** | Revision number. | [optional] 
**Severity** | Pointer to **int32** | Severity metadata for the rule. | [optional] [readonly] 
**Source** | Pointer to **string** | The ModSecurity rule logic. | [optional] [readonly] 
**State** | Pointer to **string** | The state, indicating if the revision is the most recent version of the rule. | [optional] [readonly] 
**Vcl** | Pointer to **string** | The VCL representation of the rule logic. | [optional] [readonly] 

## Methods

### NewWafRuleRevisionAttributes

`func NewWafRuleRevisionAttributes() *WafRuleRevisionAttributes`

NewWafRuleRevisionAttributes instantiates a new WafRuleRevisionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleRevisionAttributesWithDefaults

`func NewWafRuleRevisionAttributesWithDefaults() *WafRuleRevisionAttributes`

NewWafRuleRevisionAttributesWithDefaults instantiates a new WafRuleRevisionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *WafRuleRevisionAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WafRuleRevisionAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WafRuleRevisionAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WafRuleRevisionAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModsecRuleID

`func (o *WafRuleRevisionAttributes) GetModsecRuleID() int32`

GetModsecRuleID returns the ModsecRuleID field if non-nil, zero value otherwise.

### GetModsecRuleIDOk

`func (o *WafRuleRevisionAttributes) GetModsecRuleIDOk() (*int32, bool)`

GetModsecRuleIDOk returns a tuple with the ModsecRuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModsecRuleID

`func (o *WafRuleRevisionAttributes) SetModsecRuleID(v int32)`

SetModsecRuleID sets ModsecRuleID field to given value.

### HasModsecRuleID

`func (o *WafRuleRevisionAttributes) HasModsecRuleID() bool`

HasModsecRuleID returns a boolean if a field has been set.

### GetParanoiaLevel

`func (o *WafRuleRevisionAttributes) GetParanoiaLevel() int32`

GetParanoiaLevel returns the ParanoiaLevel field if non-nil, zero value otherwise.

### GetParanoiaLevelOk

`func (o *WafRuleRevisionAttributes) GetParanoiaLevelOk() (*int32, bool)`

GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParanoiaLevel

`func (o *WafRuleRevisionAttributes) SetParanoiaLevel(v int32)`

SetParanoiaLevel sets ParanoiaLevel field to given value.

### HasParanoiaLevel

`func (o *WafRuleRevisionAttributes) HasParanoiaLevel() bool`

HasParanoiaLevel returns a boolean if a field has been set.

### GetRevision

`func (o *WafRuleRevisionAttributes) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *WafRuleRevisionAttributes) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *WafRuleRevisionAttributes) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *WafRuleRevisionAttributes) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSeverity

`func (o *WafRuleRevisionAttributes) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *WafRuleRevisionAttributes) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *WafRuleRevisionAttributes) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *WafRuleRevisionAttributes) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *WafRuleRevisionAttributes) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WafRuleRevisionAttributes) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WafRuleRevisionAttributes) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *WafRuleRevisionAttributes) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *WafRuleRevisionAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WafRuleRevisionAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WafRuleRevisionAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WafRuleRevisionAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVcl

`func (o *WafRuleRevisionAttributes) GetVcl() string`

GetVcl returns the Vcl field if non-nil, zero value otherwise.

### GetVclOk

`func (o *WafRuleRevisionAttributes) GetVclOk() (*string, bool)`

GetVclOk returns a tuple with the Vcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcl

`func (o *WafRuleRevisionAttributes) SetVcl(v string)`

SetVcl sets Vcl field to given value.

### HasVcl

`func (o *WafRuleRevisionAttributes) HasVcl() bool`

HasVcl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
