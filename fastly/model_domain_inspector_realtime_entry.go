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

// DomainInspectorRealtimeEntry Each reporting period is represented by an entry in the `Data` property of the top level response and provides access to [measurement data](#measurements-data-model) for that time period, grouped in various ways: by domain name, domain IP address, and optionally by POP. The `datacenter` property organizes the measurements by Fastly POP, while the `aggregated` property combines the measurements of all POPs (but still splits by backend name and IP).
type DomainInspectorRealtimeEntry struct {
	// The Unix timestamp at which this record's data was generated.
	Recorded *int32 `json:"recorded,omitempty"`
	// Groups [measurements](#measurements-data-model) by backend name and then by IP address.
	Aggregated map[string]DomainInspectorMeasurements `json:"aggregated,omitempty"`
	// Groups [measurements](#measurements-data-model) by POP, then backend name, and then IP address. See the [POPs API](https://www.fastly.com/documentation/reference/api/utils/pops/) for details about POP identifiers.
	Datacenter           *map[string]map[string]DomainInspectorMeasurements `json:"datacenter,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainInspectorRealtimeEntry DomainInspectorRealtimeEntry

// NewDomainInspectorRealtimeEntry instantiates a new DomainInspectorRealtimeEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainInspectorRealtimeEntry() *DomainInspectorRealtimeEntry {
	this := DomainInspectorRealtimeEntry{}
	return &this
}

// NewDomainInspectorRealtimeEntryWithDefaults instantiates a new DomainInspectorRealtimeEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainInspectorRealtimeEntryWithDefaults() *DomainInspectorRealtimeEntry {
	this := DomainInspectorRealtimeEntry{}
	return &this
}

// GetRecorded returns the Recorded field value if set, zero value otherwise.
func (o *DomainInspectorRealtimeEntry) GetRecorded() int32 {
	if o == nil || o.Recorded == nil {
		var ret int32
		return ret
	}
	return *o.Recorded
}

// GetRecordedOk returns a tuple with the Recorded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorRealtimeEntry) GetRecordedOk() (*int32, bool) {
	if o == nil || o.Recorded == nil {
		return nil, false
	}
	return o.Recorded, true
}

// HasRecorded returns a boolean if a field has been set.
func (o *DomainInspectorRealtimeEntry) HasRecorded() bool {
	if o != nil && o.Recorded != nil {
		return true
	}

	return false
}

// SetRecorded gets a reference to the given int32 and assigns it to the Recorded field.
func (o *DomainInspectorRealtimeEntry) SetRecorded(v int32) {
	o.Recorded = &v
}

// GetAggregated returns the Aggregated field value if set, zero value otherwise.
func (o *DomainInspectorRealtimeEntry) GetAggregated() map[string]DomainInspectorMeasurements {
	if o == nil || o.Aggregated == nil {
		var ret map[string]DomainInspectorMeasurements
		return ret
	}
	return o.Aggregated
}

// GetAggregatedOk returns a tuple with the Aggregated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorRealtimeEntry) GetAggregatedOk() (map[string]DomainInspectorMeasurements, bool) {
	if o == nil || o.Aggregated == nil {
		return nil, false
	}
	return o.Aggregated, true
}

// HasAggregated returns a boolean if a field has been set.
func (o *DomainInspectorRealtimeEntry) HasAggregated() bool {
	if o != nil && o.Aggregated != nil {
		return true
	}

	return false
}

// SetAggregated gets a reference to the given map[string]DomainInspectorMeasurements and assigns it to the Aggregated field.
func (o *DomainInspectorRealtimeEntry) SetAggregated(v map[string]DomainInspectorMeasurements) {
	o.Aggregated = v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *DomainInspectorRealtimeEntry) GetDatacenter() map[string]map[string]DomainInspectorMeasurements {
	if o == nil || o.Datacenter == nil {
		var ret map[string]map[string]DomainInspectorMeasurements
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorRealtimeEntry) GetDatacenterOk() (*map[string]map[string]DomainInspectorMeasurements, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *DomainInspectorRealtimeEntry) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given map[string]map[string]DomainInspectorMeasurements and assigns it to the Datacenter field.
func (o *DomainInspectorRealtimeEntry) SetDatacenter(v map[string]map[string]DomainInspectorMeasurements) {
	o.Datacenter = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainInspectorRealtimeEntry) MarshalJSON() ([]byte, error) {
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
func (o *DomainInspectorRealtimeEntry) UnmarshalJSON(bytes []byte) (err error) {
	varDomainInspectorRealtimeEntry := _DomainInspectorRealtimeEntry{}

	if err = json.Unmarshal(bytes, &varDomainInspectorRealtimeEntry); err == nil {
		*o = DomainInspectorRealtimeEntry(varDomainInspectorRealtimeEntry)
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

// NullableDomainInspectorRealtimeEntry is a helper abstraction for handling nullable domaininspectorrealtimeentry types.
type NullableDomainInspectorRealtimeEntry struct {
	value *DomainInspectorRealtimeEntry
	isSet bool
}

// Get returns the value.
func (v NullableDomainInspectorRealtimeEntry) Get() *DomainInspectorRealtimeEntry {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainInspectorRealtimeEntry) Set(val *DomainInspectorRealtimeEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainInspectorRealtimeEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainInspectorRealtimeEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainInspectorRealtimeEntry returns a pointer to a new instance of NullableDomainInspectorRealtimeEntry.
func NewNullableDomainInspectorRealtimeEntry(val *DomainInspectorRealtimeEntry) *NullableDomainInspectorRealtimeEntry {
	return &NullableDomainInspectorRealtimeEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainInspectorRealtimeEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainInspectorRealtimeEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
