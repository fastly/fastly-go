// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// OriginInspectorRealtimeEntry Each reporting period is represented by an entry in the `Data` property of the top level response and provides access to [measurement data](#measurements-data-model) for that time period, grouped by origin name and optionally by POP. The `datacenter` property organizes the measurements by Fastly POP, while the `aggregated` property combines the measurements of all POPs. 
type OriginInspectorRealtimeEntry struct {
	Recorded *OriginInspectorRealtimeEntryRecorded `json:"recorded,omitempty"`
	// Groups [measurements](#measurements-data-model) by backend name.
	Aggregated *map[string]OriginInspectorMeasurements `json:"aggregated,omitempty"`
	// Groups [measurements](#measurements-data-model) by POP, then backend name. See the [POPs API](https://www.fastly.com/documentation/reference/api/utils/pops/) for details about POP identifiers.
	Datacenter *map[string]map[string]OriginInspectorMeasurements `json:"datacenter,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorRealtimeEntry OriginInspectorRealtimeEntry

// NewOriginInspectorRealtimeEntry instantiates a new OriginInspectorRealtimeEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorRealtimeEntry() *OriginInspectorRealtimeEntry {
	this := OriginInspectorRealtimeEntry{}
	return &this
}

// NewOriginInspectorRealtimeEntryWithDefaults instantiates a new OriginInspectorRealtimeEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorRealtimeEntryWithDefaults() *OriginInspectorRealtimeEntry {
	this := OriginInspectorRealtimeEntry{}
	return &this
}

// GetRecorded returns the Recorded field value if set, zero value otherwise.
func (o *OriginInspectorRealtimeEntry) GetRecorded() OriginInspectorRealtimeEntryRecorded {
	if o == nil || o.Recorded == nil {
		var ret OriginInspectorRealtimeEntryRecorded
		return ret
	}
	return *o.Recorded
}

// GetRecordedOk returns a tuple with the Recorded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorRealtimeEntry) GetRecordedOk() (*OriginInspectorRealtimeEntryRecorded, bool) {
	if o == nil || o.Recorded == nil {
		return nil, false
	}
	return o.Recorded, true
}

// HasRecorded returns a boolean if a field has been set.
func (o *OriginInspectorRealtimeEntry) HasRecorded() bool {
	if o != nil && o.Recorded != nil {
		return true
	}

	return false
}

// SetRecorded gets a reference to the given OriginInspectorRealtimeEntryRecorded and assigns it to the Recorded field.
func (o *OriginInspectorRealtimeEntry) SetRecorded(v OriginInspectorRealtimeEntryRecorded) {
	o.Recorded = &v
}

// GetAggregated returns the Aggregated field value if set, zero value otherwise.
func (o *OriginInspectorRealtimeEntry) GetAggregated() map[string]OriginInspectorMeasurements {
	if o == nil || o.Aggregated == nil {
		var ret map[string]OriginInspectorMeasurements
		return ret
	}
	return *o.Aggregated
}

// GetAggregatedOk returns a tuple with the Aggregated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorRealtimeEntry) GetAggregatedOk() (*map[string]OriginInspectorMeasurements, bool) {
	if o == nil || o.Aggregated == nil {
		return nil, false
	}
	return o.Aggregated, true
}

// HasAggregated returns a boolean if a field has been set.
func (o *OriginInspectorRealtimeEntry) HasAggregated() bool {
	if o != nil && o.Aggregated != nil {
		return true
	}

	return false
}

// SetAggregated gets a reference to the given map[string]OriginInspectorMeasurements and assigns it to the Aggregated field.
func (o *OriginInspectorRealtimeEntry) SetAggregated(v map[string]OriginInspectorMeasurements) {
	o.Aggregated = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *OriginInspectorRealtimeEntry) GetDatacenter() map[string]map[string]OriginInspectorMeasurements {
	if o == nil || o.Datacenter == nil {
		var ret map[string]map[string]OriginInspectorMeasurements
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorRealtimeEntry) GetDatacenterOk() (*map[string]map[string]OriginInspectorMeasurements, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *OriginInspectorRealtimeEntry) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given map[string]map[string]OriginInspectorMeasurements and assigns it to the Datacenter field.
func (o *OriginInspectorRealtimeEntry) SetDatacenter(v map[string]map[string]OriginInspectorMeasurements) {
	o.Datacenter = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorRealtimeEntry) MarshalJSON() ([]byte, error) {
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
func (o *OriginInspectorRealtimeEntry) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorRealtimeEntry := _OriginInspectorRealtimeEntry{}

	if err = json.Unmarshal(bytes, &varOriginInspectorRealtimeEntry); err == nil {
		*o = OriginInspectorRealtimeEntry(varOriginInspectorRealtimeEntry)
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

// NullableOriginInspectorRealtimeEntry is a helper abstraction for handling nullable origininspectorrealtimeentry types. 
type NullableOriginInspectorRealtimeEntry struct {
	value *OriginInspectorRealtimeEntry
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorRealtimeEntry) Get() *OriginInspectorRealtimeEntry {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorRealtimeEntry) Set(val *OriginInspectorRealtimeEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorRealtimeEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorRealtimeEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorRealtimeEntry returns a pointer to a new instance of NullableOriginInspectorRealtimeEntry.
func NewNullableOriginInspectorRealtimeEntry(val *OriginInspectorRealtimeEntry) *NullableOriginInspectorRealtimeEntry {
	return &NullableOriginInspectorRealtimeEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorRealtimeEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableOriginInspectorRealtimeEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
