# ArcVirtualKeysMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The `limit` value used when making the request. | [optional] 
**Sort** | Pointer to **string** | The `sort` value used when making the request. | [optional] 
**Total** | Pointer to **int32** | The count of results matching the request. | [optional] 
**NextCursor** | Pointer to **NullableString** | The cursor to use to request the next page of results. `null` when there are no more results. | [optional] 
**PrevCursor** | Pointer to **NullableString** | The cursor to use to request the previous page of results. `null` when there is no previous page. | [optional] 

## Methods

### NewArcVirtualKeysMeta

`func NewArcVirtualKeysMeta() *ArcVirtualKeysMeta`

NewArcVirtualKeysMeta instantiates a new ArcVirtualKeysMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeysMetaWithDefaults

`func NewArcVirtualKeysMetaWithDefaults() *ArcVirtualKeysMeta`

NewArcVirtualKeysMetaWithDefaults instantiates a new ArcVirtualKeysMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ArcVirtualKeysMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ArcVirtualKeysMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ArcVirtualKeysMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ArcVirtualKeysMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *ArcVirtualKeysMeta) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ArcVirtualKeysMeta) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ArcVirtualKeysMeta) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ArcVirtualKeysMeta) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotal

`func (o *ArcVirtualKeysMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ArcVirtualKeysMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ArcVirtualKeysMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ArcVirtualKeysMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNextCursor

`func (o *ArcVirtualKeysMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ArcVirtualKeysMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ArcVirtualKeysMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ArcVirtualKeysMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ArcVirtualKeysMeta) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ArcVirtualKeysMeta) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPrevCursor

`func (o *ArcVirtualKeysMeta) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *ArcVirtualKeysMeta) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *ArcVirtualKeysMeta) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.

### HasPrevCursor

`func (o *ArcVirtualKeysMeta) HasPrevCursor() bool`

HasPrevCursor returns a boolean if a field has been set.

### SetPrevCursorNil

`func (o *ArcVirtualKeysMeta) SetPrevCursorNil(b bool)`

 SetPrevCursorNil sets the value for PrevCursor to be an explicit nil

### UnsetPrevCursor
`func (o *ArcVirtualKeysMeta) UnsetPrevCursor()`

UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
