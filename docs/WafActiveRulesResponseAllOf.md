# WafActiveRulesResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafActiveRuleResponseData**](WafActiveRuleResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafActiveRuleItem**](IncludedWithWafActiveRuleItem.md) |  | [optional] 

## Methods

### NewWafActiveRulesResponseAllOf

`func NewWafActiveRulesResponseAllOf() *WafActiveRulesResponseAllOf`

NewWafActiveRulesResponseAllOf instantiates a new WafActiveRulesResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRulesResponseAllOfWithDefaults

`func NewWafActiveRulesResponseAllOfWithDefaults() *WafActiveRulesResponseAllOf`

NewWafActiveRulesResponseAllOfWithDefaults instantiates a new WafActiveRulesResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafActiveRulesResponseAllOf) GetData() []WafActiveRuleResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafActiveRulesResponseAllOf) GetDataOk() (*[]WafActiveRuleResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafActiveRulesResponseAllOf) SetData(v []WafActiveRuleResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafActiveRulesResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafActiveRulesResponseAllOf) GetIncluded() []IncludedWithWafActiveRuleItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafActiveRulesResponseAllOf) GetIncludedOk() (*[]IncludedWithWafActiveRuleItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafActiveRulesResponseAllOf) SetIncluded(v []IncludedWithWafActiveRuleItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafActiveRulesResponseAllOf) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
