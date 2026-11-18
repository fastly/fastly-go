# ArcUsageMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ArcUsageMetric**](ArcUsageMetric.md) | The usage metrics returned by the request. | 
**Meta** | [**ArcUsageMetricsMeta**](ArcUsageMetricsMeta.md) |  | 

## Methods

### NewArcUsageMetricsResponse

`func NewArcUsageMetricsResponse(data []ArcUsageMetric, meta ArcUsageMetricsMeta, ) *ArcUsageMetricsResponse`

NewArcUsageMetricsResponse instantiates a new ArcUsageMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcUsageMetricsResponseWithDefaults

`func NewArcUsageMetricsResponseWithDefaults() *ArcUsageMetricsResponse`

NewArcUsageMetricsResponseWithDefaults instantiates a new ArcUsageMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ArcUsageMetricsResponse) GetData() []ArcUsageMetric`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArcUsageMetricsResponse) GetDataOk() (*[]ArcUsageMetric, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArcUsageMetricsResponse) SetData(v []ArcUsageMetric)`

SetData sets Data field to given value.


### GetMeta

`func (o *ArcUsageMetricsResponse) GetMeta() ArcUsageMetricsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ArcUsageMetricsResponse) GetMetaOk() (*ArcUsageMetricsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ArcUsageMetricsResponse) SetMeta(v ArcUsageMetricsMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
