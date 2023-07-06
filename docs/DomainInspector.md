# DomainInspector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to [**SubsequentRequestTimestamp**](SubsequentRequestTimestamp.md) |  | [optional] 
**AggregateDelay** | Pointer to **int32** | Offset of entry timestamps from the current time due to processing time. | [optional] 
**Data** | Pointer to [**[]DomainInspectorRealtimeEntry**](DomainInspectorRealtimeEntry.md) | A list of report [entries](#entry-data-model), each representing one second of time. | [optional] 

## Methods

### NewDomainInspector

`func NewDomainInspector() *DomainInspector`

NewDomainInspector instantiates a new DomainInspector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInspectorWithDefaults

`func NewDomainInspectorWithDefaults() *DomainInspector`

NewDomainInspectorWithDefaults instantiates a new DomainInspector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *DomainInspector) GetTimestamp() SubsequentRequestTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DomainInspector) GetTimestampOk() (*SubsequentRequestTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DomainInspector) SetTimestamp(v SubsequentRequestTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DomainInspector) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAggregateDelay

`func (o *DomainInspector) GetAggregateDelay() int32`

GetAggregateDelay returns the AggregateDelay field if non-nil, zero value otherwise.

### GetAggregateDelayOk

`func (o *DomainInspector) GetAggregateDelayOk() (*int32, bool)`

GetAggregateDelayOk returns a tuple with the AggregateDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateDelay

`func (o *DomainInspector) SetAggregateDelay(v int32)`

SetAggregateDelay sets AggregateDelay field to given value.

### HasAggregateDelay

`func (o *DomainInspector) HasAggregateDelay() bool`

HasAggregateDelay returns a boolean if a field has been set.

### GetData

`func (o *DomainInspector) GetData() []DomainInspectorRealtimeEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainInspector) GetDataOk() (*[]DomainInspectorRealtimeEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainInspector) SetData(v []DomainInspectorRealtimeEntry)`

SetData sets Data field to given value.

### HasData

`func (o *DomainInspector) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
