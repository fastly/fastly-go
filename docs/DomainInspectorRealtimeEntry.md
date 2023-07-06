# DomainInspectorRealtimeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recorded** | Pointer to [**RecordedTimestamp**](RecordedTimestamp.md) |  | [optional] 
**Aggregated** | Pointer to [**map[string]DomainInspectorMeasurements**](domain_inspector_measurements.md) | Groups [measurements](#measurements-data-model) by backend name and then by IP address. | [optional] 
**Datacenter** | Pointer to [**map[string]map[string]DomainInspectorMeasurements**](map.md) | Groups [measurements](#measurements-data-model) by POP, then backend name, and then IP address. See the [POPs API](/reference/api/utils/pops/) for details about POP identifiers. | [optional] 

## Methods

### NewDomainInspectorRealtimeEntry

`func NewDomainInspectorRealtimeEntry() *DomainInspectorRealtimeEntry`

NewDomainInspectorRealtimeEntry instantiates a new DomainInspectorRealtimeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInspectorRealtimeEntryWithDefaults

`func NewDomainInspectorRealtimeEntryWithDefaults() *DomainInspectorRealtimeEntry`

NewDomainInspectorRealtimeEntryWithDefaults instantiates a new DomainInspectorRealtimeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecorded

`func (o *DomainInspectorRealtimeEntry) GetRecorded() RecordedTimestamp`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *DomainInspectorRealtimeEntry) GetRecordedOk() (*RecordedTimestamp, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *DomainInspectorRealtimeEntry) SetRecorded(v RecordedTimestamp)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *DomainInspectorRealtimeEntry) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetAggregated

`func (o *DomainInspectorRealtimeEntry) GetAggregated() map[string]DomainInspectorMeasurements`

GetAggregated returns the Aggregated field if non-nil, zero value otherwise.

### GetAggregatedOk

`func (o *DomainInspectorRealtimeEntry) GetAggregatedOk() (*map[string]DomainInspectorMeasurements, bool)`

GetAggregatedOk returns a tuple with the Aggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregated

`func (o *DomainInspectorRealtimeEntry) SetAggregated(v map[string]DomainInspectorMeasurements)`

SetAggregated sets Aggregated field to given value.

### HasAggregated

`func (o *DomainInspectorRealtimeEntry) HasAggregated() bool`

HasAggregated returns a boolean if a field has been set.

### GetDatacenter

`func (o *DomainInspectorRealtimeEntry) GetDatacenter() map[string]map[string]DomainInspectorMeasurements`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *DomainInspectorRealtimeEntry) GetDatacenterOk() (*map[string]map[string]DomainInspectorMeasurements, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *DomainInspectorRealtimeEntry) SetDatacenter(v map[string]map[string]DomainInspectorMeasurements)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *DomainInspectorRealtimeEntry) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
