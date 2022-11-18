# WafRuleRevisionsResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafRuleRevisionResponseData**](WafRuleRevisionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]WafRule**](WafRule.md) |  | [optional] 

## Methods

### NewWafRuleRevisionsResponseAllOf

`func NewWafRuleRevisionsResponseAllOf() *WafRuleRevisionsResponseAllOf`

NewWafRuleRevisionsResponseAllOf instantiates a new WafRuleRevisionsResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleRevisionsResponseAllOfWithDefaults

`func NewWafRuleRevisionsResponseAllOfWithDefaults() *WafRuleRevisionsResponseAllOf`

NewWafRuleRevisionsResponseAllOfWithDefaults instantiates a new WafRuleRevisionsResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafRuleRevisionsResponseAllOf) GetData() []WafRuleRevisionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafRuleRevisionsResponseAllOf) GetDataOk() (*[]WafRuleRevisionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafRuleRevisionsResponseAllOf) SetData(v []WafRuleRevisionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafRuleRevisionsResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafRuleRevisionsResponseAllOf) GetIncluded() []WafRule`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafRuleRevisionsResponseAllOf) GetIncludedOk() (*[]WafRule, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafRuleRevisionsResponseAllOf) SetIncluded(v []WafRule)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafRuleRevisionsResponseAllOf) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
