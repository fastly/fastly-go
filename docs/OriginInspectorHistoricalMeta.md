# OriginInspectorHistoricalMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | Start time that was used to perform the query as an ISO-8601-formatted date and time. | [optional] 
**End** | Pointer to **string** | End time that was used to perform the query as an ISO-8601-formatted date and time. | [optional] 
**Downsample** | Pointer to **string** | Downsample that was used to perform the query. One of `minute`, `hour` or `day`. | [optional] 
**Metrics** | Pointer to **string** | A comma-separated list of the metrics that were requested. | [optional] 
**Limit** | Pointer to **float32** | The maximum number of results shown per page. | [optional] 
**NextCursor** | Pointer to **string** | A string that can be used to request the next page of results, if any. | [optional] 
**Sort** | Pointer to **string** | A comma-separated list of keys the results are sorted by. | [optional] 
**GroupBy** | Pointer to **string** | A comma-separated list of dimensions being summed over in the query. | [optional] 
**Filters** | Pointer to [**OriginInspectorHistoricalMetaFilters**](OriginInspectorHistoricalMetaFilters.md) |  | [optional] 

## Methods

### NewOriginInspectorHistoricalMeta

`func NewOriginInspectorHistoricalMeta() *OriginInspectorHistoricalMeta`

NewOriginInspectorHistoricalMeta instantiates a new OriginInspectorHistoricalMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorHistoricalMetaWithDefaults

`func NewOriginInspectorHistoricalMetaWithDefaults() *OriginInspectorHistoricalMeta`

NewOriginInspectorHistoricalMetaWithDefaults instantiates a new OriginInspectorHistoricalMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *OriginInspectorHistoricalMeta) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OriginInspectorHistoricalMeta) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OriginInspectorHistoricalMeta) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *OriginInspectorHistoricalMeta) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *OriginInspectorHistoricalMeta) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OriginInspectorHistoricalMeta) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OriginInspectorHistoricalMeta) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *OriginInspectorHistoricalMeta) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetDownsample

`func (o *OriginInspectorHistoricalMeta) GetDownsample() string`

GetDownsample returns the Downsample field if non-nil, zero value otherwise.

### GetDownsampleOk

`func (o *OriginInspectorHistoricalMeta) GetDownsampleOk() (*string, bool)`

GetDownsampleOk returns a tuple with the Downsample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsample

`func (o *OriginInspectorHistoricalMeta) SetDownsample(v string)`

SetDownsample sets Downsample field to given value.

### HasDownsample

`func (o *OriginInspectorHistoricalMeta) HasDownsample() bool`

HasDownsample returns a boolean if a field has been set.

### GetMetrics

`func (o *OriginInspectorHistoricalMeta) GetMetrics() string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *OriginInspectorHistoricalMeta) GetMetricsOk() (*string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *OriginInspectorHistoricalMeta) SetMetrics(v string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *OriginInspectorHistoricalMeta) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetLimit

`func (o *OriginInspectorHistoricalMeta) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OriginInspectorHistoricalMeta) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OriginInspectorHistoricalMeta) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *OriginInspectorHistoricalMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *OriginInspectorHistoricalMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OriginInspectorHistoricalMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OriginInspectorHistoricalMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *OriginInspectorHistoricalMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetSort

`func (o *OriginInspectorHistoricalMeta) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *OriginInspectorHistoricalMeta) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *OriginInspectorHistoricalMeta) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *OriginInspectorHistoricalMeta) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetGroupBy

`func (o *OriginInspectorHistoricalMeta) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *OriginInspectorHistoricalMeta) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *OriginInspectorHistoricalMeta) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *OriginInspectorHistoricalMeta) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetFilters

`func (o *OriginInspectorHistoricalMeta) GetFilters() OriginInspectorHistoricalMetaFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OriginInspectorHistoricalMeta) GetFiltersOk() (*OriginInspectorHistoricalMetaFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OriginInspectorHistoricalMeta) SetFilters(v OriginInspectorHistoricalMetaFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OriginInspectorHistoricalMeta) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
