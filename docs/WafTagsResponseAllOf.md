# WafTagsResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WafTagsResponseDataItem**](WafTagsResponseDataItem.md) |  | [optional] 
**Included** | Pointer to [**[]WafRule**](WafRule.md) |  | [optional] 

## Methods

### NewWafTagsResponseAllOf

`func NewWafTagsResponseAllOf() *WafTagsResponseAllOf`

NewWafTagsResponseAllOf instantiates a new WafTagsResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTagsResponseAllOfWithDefaults

`func NewWafTagsResponseAllOfWithDefaults() *WafTagsResponseAllOf`

NewWafTagsResponseAllOfWithDefaults instantiates a new WafTagsResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafTagsResponseAllOf) GetData() []WafTagsResponseDataItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafTagsResponseAllOf) GetDataOk() (*[]WafTagsResponseDataItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafTagsResponseAllOf) SetData(v []WafTagsResponseDataItem)`

SetData sets Data field to given value.

### HasData

`func (o *WafTagsResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafTagsResponseAllOf) GetIncluded() []WafRule`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafTagsResponseAllOf) GetIncludedOk() (*[]WafRule, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafTagsResponseAllOf) SetIncluded(v []WafRule)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafTagsResponseAllOf) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
