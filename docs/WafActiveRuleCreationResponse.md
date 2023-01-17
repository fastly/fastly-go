# WafActiveRuleCreationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafActiveRuleResponseData**](WafActiveRuleResponseData.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafActiveRuleItem**](IncludedWithWafActiveRuleItem.md) |  | [optional] 

## Methods

### NewWafActiveRuleCreationResponse

`func NewWafActiveRuleCreationResponse() *WafActiveRuleCreationResponse`

NewWafActiveRuleCreationResponse instantiates a new WafActiveRuleCreationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRuleCreationResponseWithDefaults

`func NewWafActiveRuleCreationResponseWithDefaults() *WafActiveRuleCreationResponse`

NewWafActiveRuleCreationResponseWithDefaults instantiates a new WafActiveRuleCreationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafActiveRuleCreationResponse) GetData() []WafActiveRuleResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafActiveRuleCreationResponse) GetDataOk() (*[]WafActiveRuleResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafActiveRuleCreationResponse) SetData(v []WafActiveRuleResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafActiveRuleCreationResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *WafActiveRuleCreationResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WafActiveRuleCreationResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WafActiveRuleCreationResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WafActiveRuleCreationResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *WafActiveRuleCreationResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WafActiveRuleCreationResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WafActiveRuleCreationResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WafActiveRuleCreationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetIncluded

`func (o *WafActiveRuleCreationResponse) GetIncluded() []IncludedWithWafActiveRuleItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafActiveRuleCreationResponse) GetIncludedOk() (*[]IncludedWithWafActiveRuleItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafActiveRuleCreationResponse) SetIncluded(v []IncludedWithWafActiveRuleItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafActiveRuleCreationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
