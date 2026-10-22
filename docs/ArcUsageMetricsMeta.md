# ArcUsageMetricsMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sort** | Pointer to **string** | The `sort` value used when making the request. | [optional] 
**Total** | Pointer to **int32** | The count of results matching the request. | [optional] 

## Methods

### NewArcUsageMetricsMeta

`func NewArcUsageMetricsMeta() *ArcUsageMetricsMeta`

NewArcUsageMetricsMeta instantiates a new ArcUsageMetricsMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcUsageMetricsMetaWithDefaults

`func NewArcUsageMetricsMetaWithDefaults() *ArcUsageMetricsMeta`

NewArcUsageMetricsMetaWithDefaults instantiates a new ArcUsageMetricsMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSort

`func (o *ArcUsageMetricsMeta) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ArcUsageMetricsMeta) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ArcUsageMetricsMeta) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ArcUsageMetricsMeta) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotal

`func (o *ArcUsageMetricsMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ArcUsageMetricsMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ArcUsageMetricsMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ArcUsageMetricsMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


