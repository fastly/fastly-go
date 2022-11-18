# WafFirewallVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WafFirewallVersionResponseData**](WafFirewallVersionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafFirewallVersionItem**](IncludedWithWafFirewallVersionItem.md) |  | [optional] 

## Methods

### NewWafFirewallVersionResponse

`func NewWafFirewallVersionResponse() *WafFirewallVersionResponse`

NewWafFirewallVersionResponse instantiates a new WafFirewallVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionResponseWithDefaults

`func NewWafFirewallVersionResponseWithDefaults() *WafFirewallVersionResponse`

NewWafFirewallVersionResponseWithDefaults instantiates a new WafFirewallVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafFirewallVersionResponse) GetData() WafFirewallVersionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafFirewallVersionResponse) GetDataOk() (*WafFirewallVersionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafFirewallVersionResponse) SetData(v WafFirewallVersionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafFirewallVersionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafFirewallVersionResponse) GetIncluded() []IncludedWithWafFirewallVersionItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafFirewallVersionResponse) GetIncludedOk() (*[]IncludedWithWafFirewallVersionItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafFirewallVersionResponse) SetIncluded(v []IncludedWithWafFirewallVersionItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafFirewallVersionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
