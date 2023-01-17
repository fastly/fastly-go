# WafExclusionsResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafExclusionResponseData**](WafExclusionResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]IncludedWithWafExclusionItem**](IncludedWithWafExclusionItem.md) |  | [optional] 

## Methods

### NewWafExclusionsResponseAllOf

`func NewWafExclusionsResponseAllOf() *WafExclusionsResponseAllOf`

NewWafExclusionsResponseAllOf instantiates a new WafExclusionsResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafExclusionsResponseAllOfWithDefaults

`func NewWafExclusionsResponseAllOfWithDefaults() *WafExclusionsResponseAllOf`

NewWafExclusionsResponseAllOfWithDefaults instantiates a new WafExclusionsResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafExclusionsResponseAllOf) GetData() []WafExclusionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafExclusionsResponseAllOf) GetDataOk() (*[]WafExclusionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafExclusionsResponseAllOf) SetData(v []WafExclusionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafExclusionsResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafExclusionsResponseAllOf) GetIncluded() []IncludedWithWafExclusionItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafExclusionsResponseAllOf) GetIncludedOk() (*[]IncludedWithWafExclusionItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafExclusionsResponseAllOf) SetIncluded(v []IncludedWithWafExclusionItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafExclusionsResponseAllOf) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
