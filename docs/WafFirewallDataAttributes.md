# WafFirewallDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | The status of the firewall. | [optional] [default to false]
**PrefetchCondition** | Pointer to **string** | Name of the corresponding condition object. | [optional] 
**Response** | Pointer to **string** | Name of the corresponding response object. | [optional] 
**ServiceVersionNumber** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWafFirewallDataAttributes

`func NewWafFirewallDataAttributes() *WafFirewallDataAttributes`

NewWafFirewallDataAttributes instantiates a new WafFirewallDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallDataAttributesWithDefaults

`func NewWafFirewallDataAttributesWithDefaults() *WafFirewallDataAttributes`

NewWafFirewallDataAttributesWithDefaults instantiates a new WafFirewallDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *WafFirewallDataAttributes) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *WafFirewallDataAttributes) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *WafFirewallDataAttributes) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *WafFirewallDataAttributes) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetPrefetchCondition

`func (o *WafFirewallDataAttributes) GetPrefetchCondition() string`

GetPrefetchCondition returns the PrefetchCondition field if non-nil, zero value otherwise.

### GetPrefetchConditionOk

`func (o *WafFirewallDataAttributes) GetPrefetchConditionOk() (*string, bool)`

GetPrefetchConditionOk returns a tuple with the PrefetchCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetchCondition

`func (o *WafFirewallDataAttributes) SetPrefetchCondition(v string)`

SetPrefetchCondition sets PrefetchCondition field to given value.

### HasPrefetchCondition

`func (o *WafFirewallDataAttributes) HasPrefetchCondition() bool`

HasPrefetchCondition returns a boolean if a field has been set.

### GetResponse

`func (o *WafFirewallDataAttributes) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WafFirewallDataAttributes) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WafFirewallDataAttributes) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WafFirewallDataAttributes) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetServiceVersionNumber

`func (o *WafFirewallDataAttributes) GetServiceVersionNumber() int32`

GetServiceVersionNumber returns the ServiceVersionNumber field if non-nil, zero value otherwise.

### GetServiceVersionNumberOk

`func (o *WafFirewallDataAttributes) GetServiceVersionNumberOk() (*int32, bool)`

GetServiceVersionNumberOk returns a tuple with the ServiceVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersionNumber

`func (o *WafFirewallDataAttributes) SetServiceVersionNumber(v int32)`

SetServiceVersionNumber sets ServiceVersionNumber field to given value.

### HasServiceVersionNumber

`func (o *WafFirewallDataAttributes) HasServiceVersionNumber() bool`

HasServiceVersionNumber returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
