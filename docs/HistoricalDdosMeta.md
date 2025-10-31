# HistoricalDdosMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | Start time that was used to perform the query as an ISO-8601-formatted date and time. | [optional] 
**End** | Pointer to **string** | End time that was used to perform the query as an ISO-8601-formatted date and time. | [optional] 
**Downsample** | Pointer to **string** | Downsample that was used to perform the query. One of `hour` or `day`. | [optional] 
**Metric** | Pointer to **string** | A comma-separated list of the metrics that were requested. | [optional] 

## Methods

### NewHistoricalDdosMeta

`func NewHistoricalDdosMeta() *HistoricalDdosMeta`

NewHistoricalDdosMeta instantiates a new HistoricalDdosMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalDdosMetaWithDefaults

`func NewHistoricalDdosMetaWithDefaults() *HistoricalDdosMeta`

NewHistoricalDdosMetaWithDefaults instantiates a new HistoricalDdosMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *HistoricalDdosMeta) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HistoricalDdosMeta) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HistoricalDdosMeta) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *HistoricalDdosMeta) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *HistoricalDdosMeta) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HistoricalDdosMeta) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *HistoricalDdosMeta) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *HistoricalDdosMeta) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetDownsample

`func (o *HistoricalDdosMeta) GetDownsample() string`

GetDownsample returns the Downsample field if non-nil, zero value otherwise.

### GetDownsampleOk

`func (o *HistoricalDdosMeta) GetDownsampleOk() (*string, bool)`

GetDownsampleOk returns a tuple with the Downsample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsample

`func (o *HistoricalDdosMeta) SetDownsample(v string)`

SetDownsample sets Downsample field to given value.

### HasDownsample

`func (o *HistoricalDdosMeta) HasDownsample() bool`

HasDownsample returns a boolean if a field has been set.

### GetMetric

`func (o *HistoricalDdosMeta) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *HistoricalDdosMeta) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *HistoricalDdosMeta) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *HistoricalDdosMeta) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


