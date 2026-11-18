# ArcSessionsMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The `limit` value used when making the request. | [optional] 
**Sort** | Pointer to **string** | The `sort` value used when making the request. | [optional] 
**Total** | Pointer to **int32** | The count of results matching the request. | [optional] 
**NextCursor** | Pointer to **NullableString** | The cursor to use to request the next page of results. `null` when there are no more results. | [optional] 
**PrevCursor** | Pointer to **NullableString** | The cursor to use to request the previous page of results. `null` when there is no previous page. | [optional] 

## Methods

### NewArcSessionsMeta

`func NewArcSessionsMeta() *ArcSessionsMeta`

NewArcSessionsMeta instantiates a new ArcSessionsMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcSessionsMetaWithDefaults

`func NewArcSessionsMetaWithDefaults() *ArcSessionsMeta`

NewArcSessionsMetaWithDefaults instantiates a new ArcSessionsMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ArcSessionsMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ArcSessionsMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ArcSessionsMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ArcSessionsMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *ArcSessionsMeta) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ArcSessionsMeta) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ArcSessionsMeta) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ArcSessionsMeta) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotal

`func (o *ArcSessionsMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ArcSessionsMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ArcSessionsMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ArcSessionsMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNextCursor

`func (o *ArcSessionsMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ArcSessionsMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ArcSessionsMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ArcSessionsMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ArcSessionsMeta) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ArcSessionsMeta) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPrevCursor

`func (o *ArcSessionsMeta) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *ArcSessionsMeta) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *ArcSessionsMeta) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.

### HasPrevCursor

`func (o *ArcSessionsMeta) HasPrevCursor() bool`

HasPrevCursor returns a boolean if a field has been set.

### SetPrevCursorNil

`func (o *ArcSessionsMeta) SetPrevCursorNil(b bool)`

 SetPrevCursorNil sets the value for PrevCursor to be an explicit nil

### UnsetPrevCursor
`func (o *ArcSessionsMeta) UnsetPrevCursor()`

UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
