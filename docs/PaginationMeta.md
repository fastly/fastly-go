# PaginationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The number of records returned per page. | [optional] 
**NextCursor** | Pointer to **string** | Cursor value used to retrieve the next page of results. Empty if there are no more results. | [optional] 
**PreviousCursor** | Pointer to **string** | Cursor value used to retrieve the previous page of results. Empty if there is no previous page. | [optional] 
**Sort** | Pointer to **string** | The sort order applied to the results. | [optional] 

## Methods

### NewPaginationMeta

`func NewPaginationMeta() *PaginationMeta`

NewPaginationMeta instantiates a new PaginationMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationMetaWithDefaults

`func NewPaginationMetaWithDefaults() *PaginationMeta`

NewPaginationMetaWithDefaults instantiates a new PaginationMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PaginationMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *PaginationMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginationMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginationMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PaginationMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetPreviousCursor

`func (o *PaginationMeta) GetPreviousCursor() string`

GetPreviousCursor returns the PreviousCursor field if non-nil, zero value otherwise.

### GetPreviousCursorOk

`func (o *PaginationMeta) GetPreviousCursorOk() (*string, bool)`

GetPreviousCursorOk returns a tuple with the PreviousCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCursor

`func (o *PaginationMeta) SetPreviousCursor(v string)`

SetPreviousCursor sets PreviousCursor field to given value.

### HasPreviousCursor

`func (o *PaginationMeta) HasPreviousCursor() bool`

HasPreviousCursor returns a boolean if a field has been set.

### GetSort

`func (o *PaginationMeta) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PaginationMeta) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PaginationMeta) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PaginationMeta) HasSort() bool`

HasSort returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


