# HistoricalDomainsMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | Start time that was used to perform the query as an ISO-8601-formatted date and time. | [optional] 
**End** | Pointer to **string** | End time that was used to perform the query as an ISO-8601-formatted date and time. | [optional] 
**Downsample** | Pointer to **string** | Downsample that was used to perform the query. One of `minute`, `hour`, or `day`. | [optional] 
**Metrics** | Pointer to **string** | A comma-separated list of the metrics that were requested. | [optional] 
**Limit** | Pointer to **float32** | The maximum number of results shown per page. | [optional] 
**NextCursor** | Pointer to **string** | A string that can be used to request the next page of results, if any. | [optional] 
**Sort** | Pointer to **string** | A comma-separated list of keys the results are sorted by. | [optional] 
**GroupBy** | Pointer to **string** | A comma-separated list of dimensions being summed over in the query. | [optional] 
**Filters** | Pointer to [**HistoricalDomainsMetaFilters**](HistoricalDomainsMetaFilters.md) |  | [optional] 

## Methods

### NewHistoricalDomainsMeta

`func NewHistoricalDomainsMeta() *HistoricalDomainsMeta`

NewHistoricalDomainsMeta instantiates a new HistoricalDomainsMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalDomainsMetaWithDefaults

`func NewHistoricalDomainsMetaWithDefaults() *HistoricalDomainsMeta`

NewHistoricalDomainsMetaWithDefaults instantiates a new HistoricalDomainsMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *HistoricalDomainsMeta) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HistoricalDomainsMeta) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HistoricalDomainsMeta) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *HistoricalDomainsMeta) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *HistoricalDomainsMeta) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HistoricalDomainsMeta) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *HistoricalDomainsMeta) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *HistoricalDomainsMeta) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetDownsample

`func (o *HistoricalDomainsMeta) GetDownsample() string`

GetDownsample returns the Downsample field if non-nil, zero value otherwise.

### GetDownsampleOk

`func (o *HistoricalDomainsMeta) GetDownsampleOk() (*string, bool)`

GetDownsampleOk returns a tuple with the Downsample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsample

`func (o *HistoricalDomainsMeta) SetDownsample(v string)`

SetDownsample sets Downsample field to given value.

### HasDownsample

`func (o *HistoricalDomainsMeta) HasDownsample() bool`

HasDownsample returns a boolean if a field has been set.

### GetMetrics

`func (o *HistoricalDomainsMeta) GetMetrics() string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *HistoricalDomainsMeta) GetMetricsOk() (*string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *HistoricalDomainsMeta) SetMetrics(v string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *HistoricalDomainsMeta) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetLimit

`func (o *HistoricalDomainsMeta) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *HistoricalDomainsMeta) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *HistoricalDomainsMeta) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *HistoricalDomainsMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *HistoricalDomainsMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *HistoricalDomainsMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *HistoricalDomainsMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *HistoricalDomainsMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetSort

`func (o *HistoricalDomainsMeta) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *HistoricalDomainsMeta) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *HistoricalDomainsMeta) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *HistoricalDomainsMeta) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetGroupBy

`func (o *HistoricalDomainsMeta) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *HistoricalDomainsMeta) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *HistoricalDomainsMeta) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *HistoricalDomainsMeta) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetFilters

`func (o *HistoricalDomainsMeta) GetFilters() HistoricalDomainsMetaFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HistoricalDomainsMeta) GetFiltersOk() (*HistoricalDomainsMetaFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HistoricalDomainsMeta) SetFilters(v HistoricalDomainsMetaFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HistoricalDomainsMeta) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


