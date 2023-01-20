# GetStoresResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **string** | Cursor for the next page. | [optional] 
**Limit** | Pointer to **int32** | Entries returned. | [optional] 

## Methods

### NewGetStoresResponseMeta

`func NewGetStoresResponseMeta() *GetStoresResponseMeta`

NewGetStoresResponseMeta instantiates a new GetStoresResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoresResponseMetaWithDefaults

`func NewGetStoresResponseMetaWithDefaults() *GetStoresResponseMeta`

NewGetStoresResponseMetaWithDefaults instantiates a new GetStoresResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *GetStoresResponseMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetStoresResponseMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetStoresResponseMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *GetStoresResponseMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetLimit

`func (o *GetStoresResponseMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetStoresResponseMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetStoresResponseMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetStoresResponseMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
