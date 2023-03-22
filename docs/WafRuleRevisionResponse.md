# WafRuleRevisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WafRuleRevisionResponseData**](WafRuleRevisionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]WafRule**](WafRule.md) |  | [optional] 

## Methods

### NewWafRuleRevisionResponse

`func NewWafRuleRevisionResponse() *WafRuleRevisionResponse`

NewWafRuleRevisionResponse instantiates a new WafRuleRevisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleRevisionResponseWithDefaults

`func NewWafRuleRevisionResponseWithDefaults() *WafRuleRevisionResponse`

NewWafRuleRevisionResponseWithDefaults instantiates a new WafRuleRevisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafRuleRevisionResponse) GetData() WafRuleRevisionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafRuleRevisionResponse) GetDataOk() (*WafRuleRevisionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafRuleRevisionResponse) SetData(v WafRuleRevisionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafRuleRevisionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafRuleRevisionResponse) GetIncluded() []WafRule`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafRuleRevisionResponse) GetIncludedOk() (*[]WafRule, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafRuleRevisionResponse) SetIncluded(v []WafRule)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafRuleRevisionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
