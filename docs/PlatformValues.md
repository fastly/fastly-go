# PlatformValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Timestamp of the metrics data point. | [optional] 
**TtfbOriginP25Us** | Pointer to **float32** | 25th percentile of time to first byte from origin, in microseconds. | [optional] 
**TtfbOriginP50Us** | Pointer to **float32** | 50th percentile of time to first byte from origin, in microseconds. | [optional] 
**TtfbOriginP75Us** | Pointer to **float32** | 75th percentile of time to first byte from origin, in microseconds. | [optional] 
**TtfbOriginP95Us** | Pointer to **float32** | 95th percentile of time to first byte from origin, in microseconds. | [optional] 
**TtfbOriginP99Us** | Pointer to **float32** | 99th percentile of time to first byte from origin, in microseconds. | [optional] 
**TtfbShieldP25Us** | Pointer to **float32** | 25th percentile of time to first byte from shield, in microseconds. | [optional] 
**TtfbShieldP50Us** | Pointer to **float32** | 50th percentile of time to first byte from shield, in microseconds. | [optional] 
**TtfbShieldP75Us** | Pointer to **float32** | 75th percentile of time to first byte from shield, in microseconds. | [optional] 
**TtfbShieldP95Us** | Pointer to **float32** | 95th percentile of time to first byte from shield, in microseconds. | [optional] 
**TtfbShieldP99Us** | Pointer to **float32** | 99th percentile of time to first byte from shield, in microseconds. | [optional] 
**TtfbEdgeP25Us** | Pointer to **float32** | 25th percentile of time to first byte from edge, in microseconds. | [optional] 
**TtfbEdgeP50Us** | Pointer to **float32** | 50th percentile of time to first byte from edge, in microseconds. | [optional] 
**TtfbEdgeP75Us** | Pointer to **float32** | 75th percentile of time to first byte from edge, in microseconds. | [optional] 
**TtfbEdgeP95Us** | Pointer to **float32** | 95th percentile of time to first byte from edge, in microseconds. | [optional] 
**TtfbEdgeP99Us** | Pointer to **float32** | 99th percentile of time to first byte from edge, in microseconds. | [optional] 

## Methods

### NewPlatformValues

`func NewPlatformValues() *PlatformValues`

NewPlatformValues instantiates a new PlatformValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformValuesWithDefaults

`func NewPlatformValuesWithDefaults() *PlatformValues`

NewPlatformValuesWithDefaults instantiates a new PlatformValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *PlatformValues) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PlatformValues) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PlatformValues) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PlatformValues) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTtfbOriginP25Us

`func (o *PlatformValues) GetTtfbOriginP25Us() float32`

GetTtfbOriginP25Us returns the TtfbOriginP25Us field if non-nil, zero value otherwise.

### GetTtfbOriginP25UsOk

`func (o *PlatformValues) GetTtfbOriginP25UsOk() (*float32, bool)`

GetTtfbOriginP25UsOk returns a tuple with the TtfbOriginP25Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbOriginP25Us

`func (o *PlatformValues) SetTtfbOriginP25Us(v float32)`

SetTtfbOriginP25Us sets TtfbOriginP25Us field to given value.

### HasTtfbOriginP25Us

`func (o *PlatformValues) HasTtfbOriginP25Us() bool`

HasTtfbOriginP25Us returns a boolean if a field has been set.

### GetTtfbOriginP50Us

`func (o *PlatformValues) GetTtfbOriginP50Us() float32`

GetTtfbOriginP50Us returns the TtfbOriginP50Us field if non-nil, zero value otherwise.

### GetTtfbOriginP50UsOk

`func (o *PlatformValues) GetTtfbOriginP50UsOk() (*float32, bool)`

GetTtfbOriginP50UsOk returns a tuple with the TtfbOriginP50Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbOriginP50Us

`func (o *PlatformValues) SetTtfbOriginP50Us(v float32)`

SetTtfbOriginP50Us sets TtfbOriginP50Us field to given value.

### HasTtfbOriginP50Us

`func (o *PlatformValues) HasTtfbOriginP50Us() bool`

HasTtfbOriginP50Us returns a boolean if a field has been set.

### GetTtfbOriginP75Us

`func (o *PlatformValues) GetTtfbOriginP75Us() float32`

GetTtfbOriginP75Us returns the TtfbOriginP75Us field if non-nil, zero value otherwise.

### GetTtfbOriginP75UsOk

`func (o *PlatformValues) GetTtfbOriginP75UsOk() (*float32, bool)`

GetTtfbOriginP75UsOk returns a tuple with the TtfbOriginP75Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbOriginP75Us

`func (o *PlatformValues) SetTtfbOriginP75Us(v float32)`

SetTtfbOriginP75Us sets TtfbOriginP75Us field to given value.

### HasTtfbOriginP75Us

`func (o *PlatformValues) HasTtfbOriginP75Us() bool`

HasTtfbOriginP75Us returns a boolean if a field has been set.

### GetTtfbOriginP95Us

`func (o *PlatformValues) GetTtfbOriginP95Us() float32`

GetTtfbOriginP95Us returns the TtfbOriginP95Us field if non-nil, zero value otherwise.

### GetTtfbOriginP95UsOk

`func (o *PlatformValues) GetTtfbOriginP95UsOk() (*float32, bool)`

GetTtfbOriginP95UsOk returns a tuple with the TtfbOriginP95Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbOriginP95Us

`func (o *PlatformValues) SetTtfbOriginP95Us(v float32)`

SetTtfbOriginP95Us sets TtfbOriginP95Us field to given value.

### HasTtfbOriginP95Us

`func (o *PlatformValues) HasTtfbOriginP95Us() bool`

HasTtfbOriginP95Us returns a boolean if a field has been set.

### GetTtfbOriginP99Us

`func (o *PlatformValues) GetTtfbOriginP99Us() float32`

GetTtfbOriginP99Us returns the TtfbOriginP99Us field if non-nil, zero value otherwise.

### GetTtfbOriginP99UsOk

`func (o *PlatformValues) GetTtfbOriginP99UsOk() (*float32, bool)`

GetTtfbOriginP99UsOk returns a tuple with the TtfbOriginP99Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbOriginP99Us

`func (o *PlatformValues) SetTtfbOriginP99Us(v float32)`

SetTtfbOriginP99Us sets TtfbOriginP99Us field to given value.

### HasTtfbOriginP99Us

`func (o *PlatformValues) HasTtfbOriginP99Us() bool`

HasTtfbOriginP99Us returns a boolean if a field has been set.

### GetTtfbShieldP25Us

`func (o *PlatformValues) GetTtfbShieldP25Us() float32`

GetTtfbShieldP25Us returns the TtfbShieldP25Us field if non-nil, zero value otherwise.

### GetTtfbShieldP25UsOk

`func (o *PlatformValues) GetTtfbShieldP25UsOk() (*float32, bool)`

GetTtfbShieldP25UsOk returns a tuple with the TtfbShieldP25Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbShieldP25Us

`func (o *PlatformValues) SetTtfbShieldP25Us(v float32)`

SetTtfbShieldP25Us sets TtfbShieldP25Us field to given value.

### HasTtfbShieldP25Us

`func (o *PlatformValues) HasTtfbShieldP25Us() bool`

HasTtfbShieldP25Us returns a boolean if a field has been set.

### GetTtfbShieldP50Us

`func (o *PlatformValues) GetTtfbShieldP50Us() float32`

GetTtfbShieldP50Us returns the TtfbShieldP50Us field if non-nil, zero value otherwise.

### GetTtfbShieldP50UsOk

`func (o *PlatformValues) GetTtfbShieldP50UsOk() (*float32, bool)`

GetTtfbShieldP50UsOk returns a tuple with the TtfbShieldP50Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbShieldP50Us

`func (o *PlatformValues) SetTtfbShieldP50Us(v float32)`

SetTtfbShieldP50Us sets TtfbShieldP50Us field to given value.

### HasTtfbShieldP50Us

`func (o *PlatformValues) HasTtfbShieldP50Us() bool`

HasTtfbShieldP50Us returns a boolean if a field has been set.

### GetTtfbShieldP75Us

`func (o *PlatformValues) GetTtfbShieldP75Us() float32`

GetTtfbShieldP75Us returns the TtfbShieldP75Us field if non-nil, zero value otherwise.

### GetTtfbShieldP75UsOk

`func (o *PlatformValues) GetTtfbShieldP75UsOk() (*float32, bool)`

GetTtfbShieldP75UsOk returns a tuple with the TtfbShieldP75Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbShieldP75Us

`func (o *PlatformValues) SetTtfbShieldP75Us(v float32)`

SetTtfbShieldP75Us sets TtfbShieldP75Us field to given value.

### HasTtfbShieldP75Us

`func (o *PlatformValues) HasTtfbShieldP75Us() bool`

HasTtfbShieldP75Us returns a boolean if a field has been set.

### GetTtfbShieldP95Us

`func (o *PlatformValues) GetTtfbShieldP95Us() float32`

GetTtfbShieldP95Us returns the TtfbShieldP95Us field if non-nil, zero value otherwise.

### GetTtfbShieldP95UsOk

`func (o *PlatformValues) GetTtfbShieldP95UsOk() (*float32, bool)`

GetTtfbShieldP95UsOk returns a tuple with the TtfbShieldP95Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbShieldP95Us

`func (o *PlatformValues) SetTtfbShieldP95Us(v float32)`

SetTtfbShieldP95Us sets TtfbShieldP95Us field to given value.

### HasTtfbShieldP95Us

`func (o *PlatformValues) HasTtfbShieldP95Us() bool`

HasTtfbShieldP95Us returns a boolean if a field has been set.

### GetTtfbShieldP99Us

`func (o *PlatformValues) GetTtfbShieldP99Us() float32`

GetTtfbShieldP99Us returns the TtfbShieldP99Us field if non-nil, zero value otherwise.

### GetTtfbShieldP99UsOk

`func (o *PlatformValues) GetTtfbShieldP99UsOk() (*float32, bool)`

GetTtfbShieldP99UsOk returns a tuple with the TtfbShieldP99Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbShieldP99Us

`func (o *PlatformValues) SetTtfbShieldP99Us(v float32)`

SetTtfbShieldP99Us sets TtfbShieldP99Us field to given value.

### HasTtfbShieldP99Us

`func (o *PlatformValues) HasTtfbShieldP99Us() bool`

HasTtfbShieldP99Us returns a boolean if a field has been set.

### GetTtfbEdgeP25Us

`func (o *PlatformValues) GetTtfbEdgeP25Us() float32`

GetTtfbEdgeP25Us returns the TtfbEdgeP25Us field if non-nil, zero value otherwise.

### GetTtfbEdgeP25UsOk

`func (o *PlatformValues) GetTtfbEdgeP25UsOk() (*float32, bool)`

GetTtfbEdgeP25UsOk returns a tuple with the TtfbEdgeP25Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbEdgeP25Us

`func (o *PlatformValues) SetTtfbEdgeP25Us(v float32)`

SetTtfbEdgeP25Us sets TtfbEdgeP25Us field to given value.

### HasTtfbEdgeP25Us

`func (o *PlatformValues) HasTtfbEdgeP25Us() bool`

HasTtfbEdgeP25Us returns a boolean if a field has been set.

### GetTtfbEdgeP50Us

`func (o *PlatformValues) GetTtfbEdgeP50Us() float32`

GetTtfbEdgeP50Us returns the TtfbEdgeP50Us field if non-nil, zero value otherwise.

### GetTtfbEdgeP50UsOk

`func (o *PlatformValues) GetTtfbEdgeP50UsOk() (*float32, bool)`

GetTtfbEdgeP50UsOk returns a tuple with the TtfbEdgeP50Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbEdgeP50Us

`func (o *PlatformValues) SetTtfbEdgeP50Us(v float32)`

SetTtfbEdgeP50Us sets TtfbEdgeP50Us field to given value.

### HasTtfbEdgeP50Us

`func (o *PlatformValues) HasTtfbEdgeP50Us() bool`

HasTtfbEdgeP50Us returns a boolean if a field has been set.

### GetTtfbEdgeP75Us

`func (o *PlatformValues) GetTtfbEdgeP75Us() float32`

GetTtfbEdgeP75Us returns the TtfbEdgeP75Us field if non-nil, zero value otherwise.

### GetTtfbEdgeP75UsOk

`func (o *PlatformValues) GetTtfbEdgeP75UsOk() (*float32, bool)`

GetTtfbEdgeP75UsOk returns a tuple with the TtfbEdgeP75Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbEdgeP75Us

`func (o *PlatformValues) SetTtfbEdgeP75Us(v float32)`

SetTtfbEdgeP75Us sets TtfbEdgeP75Us field to given value.

### HasTtfbEdgeP75Us

`func (o *PlatformValues) HasTtfbEdgeP75Us() bool`

HasTtfbEdgeP75Us returns a boolean if a field has been set.

### GetTtfbEdgeP95Us

`func (o *PlatformValues) GetTtfbEdgeP95Us() float32`

GetTtfbEdgeP95Us returns the TtfbEdgeP95Us field if non-nil, zero value otherwise.

### GetTtfbEdgeP95UsOk

`func (o *PlatformValues) GetTtfbEdgeP95UsOk() (*float32, bool)`

GetTtfbEdgeP95UsOk returns a tuple with the TtfbEdgeP95Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbEdgeP95Us

`func (o *PlatformValues) SetTtfbEdgeP95Us(v float32)`

SetTtfbEdgeP95Us sets TtfbEdgeP95Us field to given value.

### HasTtfbEdgeP95Us

`func (o *PlatformValues) HasTtfbEdgeP95Us() bool`

HasTtfbEdgeP95Us returns a boolean if a field has been set.

### GetTtfbEdgeP99Us

`func (o *PlatformValues) GetTtfbEdgeP99Us() float32`

GetTtfbEdgeP99Us returns the TtfbEdgeP99Us field if non-nil, zero value otherwise.

### GetTtfbEdgeP99UsOk

`func (o *PlatformValues) GetTtfbEdgeP99UsOk() (*float32, bool)`

GetTtfbEdgeP99UsOk returns a tuple with the TtfbEdgeP99Us field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtfbEdgeP99Us

`func (o *PlatformValues) SetTtfbEdgeP99Us(v float32)`

SetTtfbEdgeP99Us sets TtfbEdgeP99Us field to given value.

### HasTtfbEdgeP99Us

`func (o *PlatformValues) HasTtfbEdgeP99Us() bool`

HasTtfbEdgeP99Us returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


