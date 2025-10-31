# OriginInspector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int32** | Value to use for subsequent requests. | [optional] 
**AggregateDelay** | Pointer to **int32** | Offset of entry timestamps from the current time due to processing time. | [optional] 
**Data** | Pointer to [**[]OriginInspectorRealtimeEntry**](OriginInspectorRealtimeEntry.md) | A list of report [entries](#entry-data-model), each representing one second of time. | [optional] 

## Methods

### NewOriginInspector

`func NewOriginInspector() *OriginInspector`

NewOriginInspector instantiates a new OriginInspector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorWithDefaults

`func NewOriginInspectorWithDefaults() *OriginInspector`

NewOriginInspectorWithDefaults instantiates a new OriginInspector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *OriginInspector) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OriginInspector) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OriginInspector) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OriginInspector) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAggregateDelay

`func (o *OriginInspector) GetAggregateDelay() int32`

GetAggregateDelay returns the AggregateDelay field if non-nil, zero value otherwise.

### GetAggregateDelayOk

`func (o *OriginInspector) GetAggregateDelayOk() (*int32, bool)`

GetAggregateDelayOk returns a tuple with the AggregateDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateDelay

`func (o *OriginInspector) SetAggregateDelay(v int32)`

SetAggregateDelay sets AggregateDelay field to given value.

### HasAggregateDelay

`func (o *OriginInspector) HasAggregateDelay() bool`

HasAggregateDelay returns a boolean if a field has been set.

### GetData

`func (o *OriginInspector) GetData() []OriginInspectorRealtimeEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OriginInspector) GetDataOk() (*[]OriginInspectorRealtimeEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OriginInspector) SetData(v []OriginInspectorRealtimeEntry)`

SetData sets Data field to given value.

### HasData

`func (o *OriginInspector) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


