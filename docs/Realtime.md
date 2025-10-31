# Realtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int32** | Value to use for subsequent requests. | [optional] 
**AggregateDelay** | Pointer to **int64** | How long the system will wait before aggregating messages for each second. The most recent data returned will have happened at the moment of the request, minus the aggregation delay. | [optional] 
**Data** | Pointer to [**[]RealtimeEntry**](RealtimeEntry.md) | A list of [records](#record-data-model), each representing one second of time. | [optional] 

## Methods

### NewRealtime

`func NewRealtime() *Realtime`

NewRealtime instantiates a new Realtime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeWithDefaults

`func NewRealtimeWithDefaults() *Realtime`

NewRealtimeWithDefaults instantiates a new Realtime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *Realtime) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Realtime) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Realtime) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Realtime) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAggregateDelay

`func (o *Realtime) GetAggregateDelay() int64`

GetAggregateDelay returns the AggregateDelay field if non-nil, zero value otherwise.

### GetAggregateDelayOk

`func (o *Realtime) GetAggregateDelayOk() (*int64, bool)`

GetAggregateDelayOk returns a tuple with the AggregateDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateDelay

`func (o *Realtime) SetAggregateDelay(v int64)`

SetAggregateDelay sets AggregateDelay field to given value.

### HasAggregateDelay

`func (o *Realtime) HasAggregateDelay() bool`

HasAggregateDelay returns a boolean if a field has been set.

### GetData

`func (o *Realtime) GetData() []RealtimeEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Realtime) GetDataOk() (*[]RealtimeEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Realtime) SetData(v []RealtimeEntry)`

SetData sets Data field to given value.

### HasData

`func (o *Realtime) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


