# OriginInspectorRealtimeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recorded** | Pointer to **int32** | The Unix timestamp at which this record&#39;s data was generated. | [optional] 
**Aggregated** | Pointer to [**map[string]OriginInspectorMeasurements**](OriginInspectorMeasurements.md) | Groups [measurements](#measurements-data-model) by backend name. | [optional] 
**Datacenter** | Pointer to [**map[string]map[string]OriginInspectorMeasurements**](map.md) | Groups [measurements](#measurements-data-model) by POP, then backend name. See the [POPs API](https://www.fastly.com/documentation/reference/api/utils/pops/) for details about POP identifiers. | [optional] 

## Methods

### NewOriginInspectorRealtimeEntry

`func NewOriginInspectorRealtimeEntry() *OriginInspectorRealtimeEntry`

NewOriginInspectorRealtimeEntry instantiates a new OriginInspectorRealtimeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorRealtimeEntryWithDefaults

`func NewOriginInspectorRealtimeEntryWithDefaults() *OriginInspectorRealtimeEntry`

NewOriginInspectorRealtimeEntryWithDefaults instantiates a new OriginInspectorRealtimeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecorded

`func (o *OriginInspectorRealtimeEntry) GetRecorded() int32`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *OriginInspectorRealtimeEntry) GetRecordedOk() (*int32, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *OriginInspectorRealtimeEntry) SetRecorded(v int32)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *OriginInspectorRealtimeEntry) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetAggregated

`func (o *OriginInspectorRealtimeEntry) GetAggregated() map[string]OriginInspectorMeasurements`

GetAggregated returns the Aggregated field if non-nil, zero value otherwise.

### GetAggregatedOk

`func (o *OriginInspectorRealtimeEntry) GetAggregatedOk() (*map[string]OriginInspectorMeasurements, bool)`

GetAggregatedOk returns a tuple with the Aggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregated

`func (o *OriginInspectorRealtimeEntry) SetAggregated(v map[string]OriginInspectorMeasurements)`

SetAggregated sets Aggregated field to given value.

### HasAggregated

`func (o *OriginInspectorRealtimeEntry) HasAggregated() bool`

HasAggregated returns a boolean if a field has been set.

### GetDatacenter

`func (o *OriginInspectorRealtimeEntry) GetDatacenter() map[string]map[string]OriginInspectorMeasurements`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *OriginInspectorRealtimeEntry) GetDatacenterOk() (*map[string]map[string]OriginInspectorMeasurements, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *OriginInspectorRealtimeEntry) SetDatacenter(v map[string]map[string]OriginInspectorMeasurements)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *OriginInspectorRealtimeEntry) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


