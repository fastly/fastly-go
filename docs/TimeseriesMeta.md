# TimeseriesMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Start time for the query as supplied in the request. | [optional] 
**To** | Pointer to **string** | End time for the query as supplied in the request. | [optional] 
**Granularity** | Pointer to **string** | The granularity of the time buckets in the response. | [optional] 
**Limit** | Pointer to **string** | Maximum number of results returned in the request. | [optional] 

## Methods

### NewTimeseriesMeta

`func NewTimeseriesMeta() *TimeseriesMeta`

NewTimeseriesMeta instantiates a new TimeseriesMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesMetaWithDefaults

`func NewTimeseriesMetaWithDefaults() *TimeseriesMeta`

NewTimeseriesMetaWithDefaults instantiates a new TimeseriesMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TimeseriesMeta) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TimeseriesMeta) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TimeseriesMeta) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TimeseriesMeta) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *TimeseriesMeta) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TimeseriesMeta) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TimeseriesMeta) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TimeseriesMeta) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGranularity

`func (o *TimeseriesMeta) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TimeseriesMeta) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TimeseriesMeta) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *TimeseriesMeta) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetLimit

`func (o *TimeseriesMeta) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TimeseriesMeta) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TimeseriesMeta) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TimeseriesMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
