# TimeseriesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TimeseriesResult**](TimeseriesResult.md) |  | [optional] 
**Meta** | Pointer to [**TimeseriesMeta**](TimeseriesMeta.md) |  | [optional] 

## Methods

### NewTimeseriesGetResponse

`func NewTimeseriesGetResponse() *TimeseriesGetResponse`

NewTimeseriesGetResponse instantiates a new TimeseriesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesGetResponseWithDefaults

`func NewTimeseriesGetResponseWithDefaults() *TimeseriesGetResponse`

NewTimeseriesGetResponseWithDefaults instantiates a new TimeseriesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TimeseriesGetResponse) GetData() []TimeseriesResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TimeseriesGetResponse) GetDataOk() (*[]TimeseriesResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TimeseriesGetResponse) SetData(v []TimeseriesResult)`

SetData sets Data field to given value.

### HasData

`func (o *TimeseriesGetResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *TimeseriesGetResponse) GetMeta() TimeseriesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TimeseriesGetResponse) GetMetaOk() (*TimeseriesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TimeseriesGetResponse) SetMeta(v TimeseriesMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TimeseriesGetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
