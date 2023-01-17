# WafTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]WafTagsResponseDataItem**](WafTagsResponseDataItem.md) |  | [optional] 
**Included** | Pointer to [**[]WafRule**](WafRule.md) |  | [optional] 

## Methods

### NewWafTagsResponse

`func NewWafTagsResponse() *WafTagsResponse`

NewWafTagsResponse instantiates a new WafTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTagsResponseWithDefaults

`func NewWafTagsResponseWithDefaults() *WafTagsResponse`

NewWafTagsResponseWithDefaults instantiates a new WafTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WafTagsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WafTagsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WafTagsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WafTagsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *WafTagsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WafTagsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WafTagsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WafTagsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *WafTagsResponse) GetData() []WafTagsResponseDataItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafTagsResponse) GetDataOk() (*[]WafTagsResponseDataItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafTagsResponse) SetData(v []WafTagsResponseDataItem)`

SetData sets Data field to given value.

### HasData

`func (o *WafTagsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafTagsResponse) GetIncluded() []WafRule`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafTagsResponse) GetIncludedOk() (*[]WafRule, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafTagsResponse) SetIncluded(v []WafRule)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafTagsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
