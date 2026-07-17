# InlineResponse20012

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DdosProtectionEvent**](DdosProtectionEvent.md) |  | [optional] 
**Meta** | Pointer to [**PaginationCursorMeta**](PaginationCursorMeta.md) |  | [optional] 

## Methods

### NewInlineResponse20012

`func NewInlineResponse20012() *InlineResponse20012`

NewInlineResponse20012 instantiates a new InlineResponse20012 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20012WithDefaults

`func NewInlineResponse20012WithDefaults() *InlineResponse20012`

NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineResponse20012) GetData() []DdosProtectionEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20012) GetDataOk() (*[]DdosProtectionEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20012) SetData(v []DdosProtectionEvent)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse20012) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse20012) GetMeta() PaginationCursorMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse20012) GetMetaOk() (*PaginationCursorMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse20012) SetMeta(v PaginationCursorMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse20012) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


