# WafRulesResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafRuleResponseData**](WafRuleResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafRuleItem**](IncludedWithWafRuleItem.md) |  | [optional] 

## Methods

### NewWafRulesResponseAllOf

`func NewWafRulesResponseAllOf() *WafRulesResponseAllOf`

NewWafRulesResponseAllOf instantiates a new WafRulesResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRulesResponseAllOfWithDefaults

`func NewWafRulesResponseAllOfWithDefaults() *WafRulesResponseAllOf`

NewWafRulesResponseAllOfWithDefaults instantiates a new WafRulesResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafRulesResponseAllOf) GetData() []WafRuleResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafRulesResponseAllOf) GetDataOk() (*[]WafRuleResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafRulesResponseAllOf) SetData(v []WafRuleResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafRulesResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafRulesResponseAllOf) GetIncluded() []IncludedWithWafRuleItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafRulesResponseAllOf) GetIncludedOk() (*[]IncludedWithWafRuleItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafRulesResponseAllOf) SetIncluded(v []IncludedWithWafRuleItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafRulesResponseAllOf) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
