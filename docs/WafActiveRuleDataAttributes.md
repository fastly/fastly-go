# WafActiveRuleDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModsecRuleID** | Pointer to **int32** | The ModSecurity rule ID of the associated rule revision. | [optional] 
**Revision** | Pointer to [**WafRuleRevisionOrLatest**](WafRuleRevisionOrLatest.md) |  | [optional] 
**Status** | Pointer to **string** | Describes the behavior for the particular rule revision within this firewall version. | [optional] 

## Methods

### NewWafActiveRuleDataAttributes

`func NewWafActiveRuleDataAttributes() *WafActiveRuleDataAttributes`

NewWafActiveRuleDataAttributes instantiates a new WafActiveRuleDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRuleDataAttributesWithDefaults

`func NewWafActiveRuleDataAttributesWithDefaults() *WafActiveRuleDataAttributes`

NewWafActiveRuleDataAttributesWithDefaults instantiates a new WafActiveRuleDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModsecRuleID

`func (o *WafActiveRuleDataAttributes) GetModsecRuleID() int32`

GetModsecRuleID returns the ModsecRuleID field if non-nil, zero value otherwise.

### GetModsecRuleIDOk

`func (o *WafActiveRuleDataAttributes) GetModsecRuleIDOk() (*int32, bool)`

GetModsecRuleIDOk returns a tuple with the ModsecRuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModsecRuleID

`func (o *WafActiveRuleDataAttributes) SetModsecRuleID(v int32)`

SetModsecRuleID sets ModsecRuleID field to given value.

### HasModsecRuleID

`func (o *WafActiveRuleDataAttributes) HasModsecRuleID() bool`

HasModsecRuleID returns a boolean if a field has been set.

### GetRevision

`func (o *WafActiveRuleDataAttributes) GetRevision() WafRuleRevisionOrLatest`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *WafActiveRuleDataAttributes) GetRevisionOk() (*WafRuleRevisionOrLatest, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *WafActiveRuleDataAttributes) SetRevision(v WafRuleRevisionOrLatest)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *WafActiveRuleDataAttributes) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStatus

`func (o *WafActiveRuleDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WafActiveRuleDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WafActiveRuleDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WafActiveRuleDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
