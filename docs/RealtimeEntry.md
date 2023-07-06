# RealtimeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recorded** | Pointer to [**RealtimeEntryRecorded**](RealtimeEntryRecorded.md) |  | [optional] 
**Aggregated** | Pointer to [**RealtimeEntryAggregated**](RealtimeEntryAggregated.md) |  | [optional] 
**Datacenter** | Pointer to [**map[string]RealtimeMeasurements**](RealtimeMeasurements.md) | Groups [measurements](#measurements-data-model) by POP. See the [POPs API](/reference/api/utils/pops/) for details of POP identifiers. | [optional] 

## Methods

### NewRealtimeEntry

`func NewRealtimeEntry() *RealtimeEntry`

NewRealtimeEntry instantiates a new RealtimeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeEntryWithDefaults

`func NewRealtimeEntryWithDefaults() *RealtimeEntry`

NewRealtimeEntryWithDefaults instantiates a new RealtimeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecorded

`func (o *RealtimeEntry) GetRecorded() RealtimeEntryRecorded`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *RealtimeEntry) GetRecordedOk() (*RealtimeEntryRecorded, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *RealtimeEntry) SetRecorded(v RealtimeEntryRecorded)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *RealtimeEntry) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetAggregated

`func (o *RealtimeEntry) GetAggregated() RealtimeEntryAggregated`

GetAggregated returns the Aggregated field if non-nil, zero value otherwise.

### GetAggregatedOk

`func (o *RealtimeEntry) GetAggregatedOk() (*RealtimeEntryAggregated, bool)`

GetAggregatedOk returns a tuple with the Aggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregated

`func (o *RealtimeEntry) SetAggregated(v RealtimeEntryAggregated)`

SetAggregated sets Aggregated field to given value.

### HasAggregated

`func (o *RealtimeEntry) HasAggregated() bool`

HasAggregated returns a boolean if a field has been set.

### GetDatacenter

`func (o *RealtimeEntry) GetDatacenter() map[string]RealtimeMeasurements`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *RealtimeEntry) GetDatacenterOk() (*map[string]RealtimeMeasurements, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *RealtimeEntry) SetDatacenter(v map[string]RealtimeMeasurements)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *RealtimeEntry) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
