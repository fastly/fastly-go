# WafFirewallData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafFirewall**](TypeWafFirewall.md) |  | [optional] [default to TYPEWAFFIREWALL_WAF_FIREWALL]
**Attributes** | Pointer to [**WafFirewallDataAttributes**](WafFirewallDataAttributes.md) |  | [optional] 

## Methods

### NewWafFirewallData

`func NewWafFirewallData() *WafFirewallData`

NewWafFirewallData instantiates a new WafFirewallData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallDataWithDefaults

`func NewWafFirewallDataWithDefaults() *WafFirewallData`

NewWafFirewallDataWithDefaults instantiates a new WafFirewallData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafFirewallData) GetType() TypeWafFirewall`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafFirewallData) GetTypeOk() (*TypeWafFirewall, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafFirewallData) SetType(v TypeWafFirewall)`

SetType sets Type field to given value.

### HasType

`func (o *WafFirewallData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WafFirewallData) GetAttributes() WafFirewallDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafFirewallData) GetAttributesOk() (*WafFirewallDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafFirewallData) SetAttributes(v WafFirewallDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafFirewallData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
