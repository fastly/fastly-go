# InlineResponse20013

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DdosProtectionRuleWithStats**](DdosProtectionRuleWithStats.md) |  | 
**Meta** | [**PaginationCursorMeta**](PaginationCursorMeta.md) |  | 

## Methods

### NewInlineResponse20013

`func NewInlineResponse20013(data []DdosProtectionRuleWithStats, meta PaginationCursorMeta, ) *InlineResponse20013`

NewInlineResponse20013 instantiates a new InlineResponse20013 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013WithDefaults

`func NewInlineResponse20013WithDefaults() *InlineResponse20013`

NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineResponse20013) GetData() []DdosProtectionRuleWithStats`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20013) GetDataOk() (*[]DdosProtectionRuleWithStats, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20013) SetData(v []DdosProtectionRuleWithStats)`

SetData sets Data field to given value.


### GetMeta

`func (o *InlineResponse20013) GetMeta() PaginationCursorMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse20013) GetMetaOk() (*PaginationCursorMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse20013) SetMeta(v PaginationCursorMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


