# InlineResponse20011

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]HeaderEvent**](HeaderEvent.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewInlineResponse20011

`func NewInlineResponse20011() *InlineResponse20011`

NewInlineResponse20011 instantiates a new InlineResponse20011 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011WithDefaults

`func NewInlineResponse20011WithDefaults() *InlineResponse20011`

NewInlineResponse20011WithDefaults instantiates a new InlineResponse20011 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineResponse20011) GetData() []HeaderEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20011) GetDataOk() (*[]HeaderEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20011) SetData(v []HeaderEvent)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse20011) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse20011) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse20011) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse20011) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse20011) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


