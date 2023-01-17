# WafFirewallVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]WafFirewallVersionResponseData**](WafFirewallVersionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafFirewallVersionItem**](IncludedWithWafFirewallVersionItem.md) |  | [optional] 

## Methods

### NewWafFirewallVersionsResponse

`func NewWafFirewallVersionsResponse() *WafFirewallVersionsResponse`

NewWafFirewallVersionsResponse instantiates a new WafFirewallVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionsResponseWithDefaults

`func NewWafFirewallVersionsResponseWithDefaults() *WafFirewallVersionsResponse`

NewWafFirewallVersionsResponseWithDefaults instantiates a new WafFirewallVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WafFirewallVersionsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WafFirewallVersionsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WafFirewallVersionsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WafFirewallVersionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *WafFirewallVersionsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WafFirewallVersionsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WafFirewallVersionsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WafFirewallVersionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *WafFirewallVersionsResponse) GetData() []WafFirewallVersionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafFirewallVersionsResponse) GetDataOk() (*[]WafFirewallVersionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafFirewallVersionsResponse) SetData(v []WafFirewallVersionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafFirewallVersionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafFirewallVersionsResponse) GetIncluded() []IncludedWithWafFirewallVersionItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafFirewallVersionsResponse) GetIncludedOk() (*[]IncludedWithWafFirewallVersionItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafFirewallVersionsResponse) SetIncluded(v []IncludedWithWafFirewallVersionItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafFirewallVersionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
