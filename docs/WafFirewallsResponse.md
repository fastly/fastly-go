# WafFirewallsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]WafFirewallResponseData**](WafFirewallResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]SchemasWafFirewallVersion**](SchemasWafFirewallVersion.md) |  | [optional] 

## Methods

### NewWafFirewallsResponse

`func NewWafFirewallsResponse() *WafFirewallsResponse`

NewWafFirewallsResponse instantiates a new WafFirewallsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallsResponseWithDefaults

`func NewWafFirewallsResponseWithDefaults() *WafFirewallsResponse`

NewWafFirewallsResponseWithDefaults instantiates a new WafFirewallsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WafFirewallsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WafFirewallsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WafFirewallsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WafFirewallsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *WafFirewallsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WafFirewallsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WafFirewallsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WafFirewallsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *WafFirewallsResponse) GetData() []WafFirewallResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafFirewallsResponse) GetDataOk() (*[]WafFirewallResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafFirewallsResponse) SetData(v []WafFirewallResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafFirewallsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafFirewallsResponse) GetIncluded() []SchemasWafFirewallVersion`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafFirewallsResponse) GetIncludedOk() (*[]SchemasWafFirewallVersion, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafFirewallsResponse) SetIncluded(v []SchemasWafFirewallVersion)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafFirewallsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
