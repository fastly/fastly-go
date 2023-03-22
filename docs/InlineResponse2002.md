# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]StoreResponse**](StoreResponse.md) |  | [optional] 
**Meta** | Pointer to [**InlineResponse2002Meta**](InlineResponse2002Meta.md) |  | [optional] 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002() *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineResponse2002) GetData() []StoreResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse2002) GetDataOk() (*[]StoreResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse2002) SetData(v []StoreResponse)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse2002) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse2002) GetMeta() InlineResponse2002Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse2002) GetMetaOk() (*InlineResponse2002Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse2002) SetMeta(v InlineResponse2002Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse2002) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
