# WafFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WafFirewallData**](WafFirewallData.md) |  | [optional] 

## Methods

### NewWafFirewall

`func NewWafFirewall() *WafFirewall`

NewWafFirewall instantiates a new WafFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallWithDefaults

`func NewWafFirewallWithDefaults() *WafFirewall`

NewWafFirewallWithDefaults instantiates a new WafFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafFirewall) GetData() WafFirewallData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafFirewall) GetDataOk() (*WafFirewallData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafFirewall) SetData(v WafFirewallData)`

SetData sets Data field to given value.

### HasData

`func (o *WafFirewall) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
