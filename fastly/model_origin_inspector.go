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

// OriginInspector struct for OriginInspector
type OriginInspector struct {
	// Value to use for subsequent requests.
	Timestamp *int32 `json:"Timestamp,omitempty"`
	// Offset of entry timestamps from the current time due to processing time.
	AggregateDelay *int32 `json:"AggregateDelay,omitempty"`
	// A list of report [entries](#entry-data-model), each representing one second of time.
	Data                 []OriginInspectorRealtimeEntry `json:"Data,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspector OriginInspector

// NewOriginInspector instantiates a new OriginInspector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspector() *OriginInspector {
	this := OriginInspector{}
	return &this
}

// NewOriginInspectorWithDefaults instantiates a new OriginInspector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorWithDefaults() *OriginInspector {
	this := OriginInspector{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OriginInspector) GetTimestamp() int32 {
	if o == nil || o.Timestamp == nil {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspector) GetTimestampOk() (*int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OriginInspector) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *OriginInspector) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetAggregateDelay returns the AggregateDelay field value if set, zero value otherwise.
func (o *OriginInspector) GetAggregateDelay() int32 {
	if o == nil || o.AggregateDelay == nil {
		var ret int32
		return ret
	}
	return *o.AggregateDelay
}

// GetAggregateDelayOk returns a tuple with the AggregateDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspector) GetAggregateDelayOk() (*int32, bool) {
	if o == nil || o.AggregateDelay == nil {
		return nil, false
	}
	return o.AggregateDelay, true
}

// HasAggregateDelay returns a boolean if a field has been set.
func (o *OriginInspector) HasAggregateDelay() bool {
	if o != nil && o.AggregateDelay != nil {
		return true
	}

	return false
}

// SetAggregateDelay gets a reference to the given int32 and assigns it to the AggregateDelay field.
func (o *OriginInspector) SetAggregateDelay(v int32) {
	o.AggregateDelay = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OriginInspector) GetData() []OriginInspectorRealtimeEntry {
	if o == nil || o.Data == nil {
		var ret []OriginInspectorRealtimeEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspector) GetDataOk() ([]OriginInspectorRealtimeEntry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OriginInspector) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []OriginInspectorRealtimeEntry and assigns it to the Data field.
func (o *OriginInspector) SetData(v []OriginInspectorRealtimeEntry) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.AggregateDelay != nil {
		toSerialize["AggregateDelay"] = o.AggregateDelay
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OriginInspector) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspector := _OriginInspector{}

	if err = json.Unmarshal(bytes, &varOriginInspector); err == nil {
		*o = OriginInspector(varOriginInspector)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "AggregateDelay")
		delete(additionalProperties, "Data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspector is a helper abstraction for handling nullable origininspector types.
type NullableOriginInspector struct {
	value *OriginInspector
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspector) Get() *OriginInspector {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspector) Set(val *OriginInspector) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspector) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspector) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspector returns a pointer to a new instance of NullableOriginInspector.
func NewNullableOriginInspector(val *OriginInspector) *NullableOriginInspector {
	return &NullableOriginInspector{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
