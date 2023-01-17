# WafFirewallVersionsResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafFirewallVersionResponseData**](WafFirewallVersionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafFirewallVersionItem**](IncludedWithWafFirewallVersionItem.md) |  | [optional] 

## Methods

### NewWafFirewallVersionsResponseAllOf

`func NewWafFirewallVersionsResponseAllOf() *WafFirewallVersionsResponseAllOf`

NewWafFirewallVersionsResponseAllOf instantiates a new WafFirewallVersionsResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionsResponseAllOfWithDefaults

`func NewWafFirewallVersionsResponseAllOfWithDefaults() *WafFirewallVersionsResponseAllOf`

NewWafFirewallVersionsResponseAllOfWithDefaults instantiates a new WafFirewallVersionsResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafFirewallVersionsResponseAllOf) GetData() []WafFirewallVersionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafFirewallVersionsResponseAllOf) GetDataOk() (*[]WafFirewallVersionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafFirewallVersionsResponseAllOf) SetData(v []WafFirewallVersionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafFirewallVersionsResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafFirewallVersionsResponseAllOf) GetIncluded() []IncludedWithWafFirewallVersionItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafFirewallVersionsResponseAllOf) GetIncludedOk() (*[]IncludedWithWafFirewallVersionItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafFirewallVersionsResponseAllOf) SetIncluded(v []IncludedWithWafFirewallVersionItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafFirewallVersionsResponseAllOf) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
