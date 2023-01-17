# BulkWafActiveRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafActiveRuleData**](WafActiveRuleData.md) |  | [optional] 

## Methods

### NewBulkWafActiveRules

`func NewBulkWafActiveRules() *BulkWafActiveRules`

NewBulkWafActiveRules instantiates a new BulkWafActiveRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWafActiveRulesWithDefaults

`func NewBulkWafActiveRulesWithDefaults() *BulkWafActiveRules`

NewBulkWafActiveRulesWithDefaults instantiates a new BulkWafActiveRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BulkWafActiveRules) GetData() []WafActiveRuleData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkWafActiveRules) GetDataOk() (*[]WafActiveRuleData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkWafActiveRules) SetData(v []WafActiveRuleData)`

SetData sets Data field to given value.

### HasData

`func (o *BulkWafActiveRules) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
