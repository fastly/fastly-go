# WafRuleRevisionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]WafRuleRevisionResponseData**](WafRuleRevisionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]WafRule**](WafRule.md) |  | [optional] 

## Methods

### NewWafRuleRevisionsResponse

`func NewWafRuleRevisionsResponse() *WafRuleRevisionsResponse`

NewWafRuleRevisionsResponse instantiates a new WafRuleRevisionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleRevisionsResponseWithDefaults

`func NewWafRuleRevisionsResponseWithDefaults() *WafRuleRevisionsResponse`

NewWafRuleRevisionsResponseWithDefaults instantiates a new WafRuleRevisionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WafRuleRevisionsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WafRuleRevisionsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WafRuleRevisionsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WafRuleRevisionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *WafRuleRevisionsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WafRuleRevisionsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WafRuleRevisionsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WafRuleRevisionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *WafRuleRevisionsResponse) GetData() []WafRuleRevisionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafRuleRevisionsResponse) GetDataOk() (*[]WafRuleRevisionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafRuleRevisionsResponse) SetData(v []WafRuleRevisionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafRuleRevisionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafRuleRevisionsResponse) GetIncluded() []WafRule`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafRuleRevisionsResponse) GetIncludedOk() (*[]WafRule, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafRuleRevisionsResponse) SetIncluded(v []WafRule)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafRuleRevisionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
