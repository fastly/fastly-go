# InlineResponse20017

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]KvStoreDetails**](KvStoreDetails.md) |  | [optional] 
**Meta** | Pointer to [**PaginationCursorMeta**](PaginationCursorMeta.md) |  | [optional] 

## Methods

### NewInlineResponse20017

`func NewInlineResponse20017() *InlineResponse20017`

NewInlineResponse20017 instantiates a new InlineResponse20017 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20017WithDefaults

`func NewInlineResponse20017WithDefaults() *InlineResponse20017`

NewInlineResponse20017WithDefaults instantiates a new InlineResponse20017 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineResponse20017) GetData() []KvStoreDetails`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20017) GetDataOk() (*[]KvStoreDetails, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20017) SetData(v []KvStoreDetails)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse20017) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse20017) GetMeta() PaginationCursorMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse20017) GetMetaOk() (*PaginationCursorMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse20017) SetMeta(v PaginationCursorMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse20017) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


