# WafRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]WafRuleResponseData**](WafRuleResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafRuleItem**](IncludedWithWafRuleItem.md) |  | [optional] 

## Methods

### NewWafRulesResponse

`func NewWafRulesResponse() *WafRulesResponse`

NewWafRulesResponse instantiates a new WafRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRulesResponseWithDefaults

`func NewWafRulesResponseWithDefaults() *WafRulesResponse`

NewWafRulesResponseWithDefaults instantiates a new WafRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WafRulesResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WafRulesResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WafRulesResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WafRulesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *WafRulesResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WafRulesResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WafRulesResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WafRulesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *WafRulesResponse) GetData() []WafRuleResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafRulesResponse) GetDataOk() (*[]WafRuleResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafRulesResponse) SetData(v []WafRuleResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafRulesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafRulesResponse) GetIncluded() []IncludedWithWafRuleItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafRulesResponse) GetIncludedOk() (*[]IncludedWithWafRuleItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafRulesResponse) SetIncluded(v []IncludedWithWafRuleItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafRulesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
