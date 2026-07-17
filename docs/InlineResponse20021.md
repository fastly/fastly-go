# InlineResponse20021

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SecretResponse**](SecretResponse.md) |  | [optional] 
**Meta** | Pointer to [**PaginationCursorMeta**](PaginationCursorMeta.md) |  | [optional] 

## Methods

### NewInlineResponse20021

`func NewInlineResponse20021() *InlineResponse20021`

NewInlineResponse20021 instantiates a new InlineResponse20021 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20021WithDefaults

`func NewInlineResponse20021WithDefaults() *InlineResponse20021`

NewInlineResponse20021WithDefaults instantiates a new InlineResponse20021 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineResponse20021) GetData() []SecretResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20021) GetDataOk() (*[]SecretResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20021) SetData(v []SecretResponse)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse20021) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse20021) GetMeta() PaginationCursorMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse20021) GetMetaOk() (*PaginationCursorMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse20021) SetMeta(v PaginationCursorMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse20021) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


