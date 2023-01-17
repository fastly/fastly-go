# WafExclusionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]WafExclusionResponseData**](WafExclusionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafExclusionItem**](IncludedWithWafExclusionItem.md) |  | [optional] 

## Methods

### NewWafExclusionsResponse

`func NewWafExclusionsResponse() *WafExclusionsResponse`

NewWafExclusionsResponse instantiates a new WafExclusionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafExclusionsResponseWithDefaults

`func NewWafExclusionsResponseWithDefaults() *WafExclusionsResponse`

NewWafExclusionsResponseWithDefaults instantiates a new WafExclusionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WafExclusionsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WafExclusionsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WafExclusionsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WafExclusionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *WafExclusionsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WafExclusionsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WafExclusionsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WafExclusionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *WafExclusionsResponse) GetData() []WafExclusionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafExclusionsResponse) GetDataOk() (*[]WafExclusionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafExclusionsResponse) SetData(v []WafExclusionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafExclusionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafExclusionsResponse) GetIncluded() []IncludedWithWafExclusionItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafExclusionsResponse) GetIncludedOk() (*[]IncludedWithWafExclusionItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafExclusionsResponse) SetIncluded(v []IncludedWithWafExclusionItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafExclusionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
