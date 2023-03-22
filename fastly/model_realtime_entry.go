// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// RealtimeEntry A list of records, each representing one second of time. The `Data` property provides access to [measurement data](#measurements-data-model) for that time period, grouped in various ways.
type RealtimeEntry struct {
	// The Unix timestamp at which this record's data was generated.
	Recorded *int32 `json:"recorded,omitempty"`
	// Aggregates [measurements](#measurements-data-model) across all Fastly POPs.
	Aggregated *RealtimeMeasurements `json:"aggregated,omitempty"`
	// Groups [measurements](#measurements-data-model) by POP. See the [POPs API](/reference/api/utils/pops/) for details of POP identifiers.
	Datacenter *map[string]RealtimeMeasurements `json:"datacenter,omitempty"`
	AdditionalProperties map[string]any
}

type _RealtimeEntry RealtimeEntry

// NewRealtimeEntry instantiates a new RealtimeEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeEntry() *RealtimeEntry {
	this := RealtimeEntry{}
	return &this
}

// NewRealtimeEntryWithDefaults instantiates a new RealtimeEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeEntryWithDefaults() *RealtimeEntry {
	this := RealtimeEntry{}
	return &this
}

// GetRecorded returns the Recorded field value if set, zero value otherwise.
func (o *RealtimeEntry) GetRecorded() int32 {
	if o == nil || o.Recorded == nil {
		var ret int32
		return ret
	}
	return *o.Recorded
}

// GetRecordedOk returns a tuple with the Recorded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntry) GetRecordedOk() (*int32, bool) {
	if o == nil || o.Recorded == nil {
		return nil, false
	}
	return o.Recorded, true
}

// HasRecorded returns a boolean if a field has been set.
func (o *RealtimeEntry) HasRecorded() bool {
	if o != nil && o.Recorded != nil {
		return true
	}

	return false
}

// SetRecorded gets a reference to the given int32 and assigns it to the Recorded field.
func (o *RealtimeEntry) SetRecorded(v int32) {
	o.Recorded = &v
}

// GetAggregated returns the Aggregated field value if set, zero value otherwise.
func (o *RealtimeEntry) GetAggregated() RealtimeMeasurements {
	if o == nil || o.Aggregated == nil {
		var ret RealtimeMeasurements
		return ret
	}
	return *o.Aggregated
}

// GetAggregatedOk returns a tuple with the Aggregated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntry) GetAggregatedOk() (*RealtimeMeasurements, bool) {
	if o == nil || o.Aggregated == nil {
		return nil, false
	}
	return o.Aggregated, true
}

// HasAggregated returns a boolean if a field has been set.
func (o *RealtimeEntry) HasAggregated() bool {
	if o != nil && o.Aggregated != nil {
		return true
	}

	return false
}

// SetAggregated gets a reference to the given RealtimeMeasurements and assigns it to the Aggregated field.
func (o *RealtimeEntry) SetAggregated(v RealtimeMeasurements) {
	o.Aggregated = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *RealtimeEntry) GetDatacenter() map[string]RealtimeMeasurements {
	if o == nil || o.Datacenter == nil {
		var ret map[string]RealtimeMeasurements
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntry) GetDatacenterOk() (*map[string]RealtimeMeasurements, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *RealtimeEntry) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given map[string]RealtimeMeasurements and assigns it to the Datacenter field.
func (o *RealtimeEntry) SetDatacenter(v map[string]RealtimeMeasurements) {
	o.Datacenter = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RealtimeEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Recorded != nil {
		toSerialize["recorded"] = o.Recorded
	}
	if o.Aggregated != nil {
		toSerialize["aggregated"] = o.Aggregated
	}
	if o.Datacenter != nil {
		toSerialize["datacenter"] = o.Datacenter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RealtimeEntry) UnmarshalJSON(bytes []byte) (err error) {
	varRealtimeEntry := _RealtimeEntry{}

	if err = json.Unmarshal(bytes, &varRealtimeEntry); err == nil {
		*o = RealtimeEntry(varRealtimeEntry)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recorded")
		delete(additionalProperties, "aggregated")
		delete(additionalProperties, "datacenter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRealtimeEntry is a helper abstraction for handling nullable realtimeentry types. 
type NullableRealtimeEntry struct {
	value *RealtimeEntry
	isSet bool
}

// Get returns the value.
func (v NullableRealtimeEntry) Get() *RealtimeEntry {
	return v.value
}

// Set modifies the value.
func (v *NullableRealtimeEntry) Set(val *RealtimeEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRealtimeEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRealtimeEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRealtimeEntry returns a pointer to a new instance of NullableRealtimeEntry.
func NewNullableRealtimeEntry(val *RealtimeEntry) *NullableRealtimeEntry {
	return &NullableRealtimeEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRealtimeEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRealtimeEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
