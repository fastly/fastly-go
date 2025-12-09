# PlatformMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PlatformMetadata**](PlatformMetadata.md) |  | [optional] 
**Data** | Pointer to [**[]PlatformValues**](PlatformValues.md) | An array of values representing the metric values at each point in time. Note that this dataset is sparse: only the keys with non-zero values will be included in the record.  | [optional] 

## Methods

### NewPlatformMetricsResponse

`func NewPlatformMetricsResponse() *PlatformMetricsResponse`

NewPlatformMetricsResponse instantiates a new PlatformMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformMetricsResponseWithDefaults

`func NewPlatformMetricsResponseWithDefaults() *PlatformMetricsResponse`

NewPlatformMetricsResponseWithDefaults instantiates a new PlatformMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PlatformMetricsResponse) GetMeta() PlatformMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PlatformMetricsResponse) GetMetaOk() (*PlatformMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PlatformMetricsResponse) SetMeta(v PlatformMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PlatformMetricsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *PlatformMetricsResponse) GetData() []PlatformValues`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlatformMetricsResponse) GetDataOk() (*[]PlatformValues, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlatformMetricsResponse) SetData(v []PlatformValues)`

SetData sets Data field to given value.

### HasData

`func (o *PlatformMetricsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


