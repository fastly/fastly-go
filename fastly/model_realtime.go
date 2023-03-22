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

// Realtime struct for Realtime
type Realtime struct {
	// Value to use for subsequent requests.
	Timestamp *int32 `json:"Timestamp,omitempty"`
	// How long the system will wait before aggregating messages for each second. The most recent data returned will have happened at the moment of the request, minus the aggregation delay.
	AggregateDelay *int32 `json:"AggregateDelay,omitempty"`
	// A list of [records](#record-data-model), each representing one second of time.
	Data []RealtimeEntry `json:"Data,omitempty"`
	AdditionalProperties map[string]any
}

type _Realtime Realtime

// NewRealtime instantiates a new Realtime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtime() *Realtime {
	this := Realtime{}
	return &this
}

// NewRealtimeWithDefaults instantiates a new Realtime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeWithDefaults() *Realtime {
	this := Realtime{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Realtime) GetTimestamp() int32 {
	if o == nil || o.Timestamp == nil {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Realtime) GetTimestampOk() (*int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Realtime) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *Realtime) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetAggregateDelay returns the AggregateDelay field value if set, zero value otherwise.
func (o *Realtime) GetAggregateDelay() int32 {
	if o == nil || o.AggregateDelay == nil {
		var ret int32
		return ret
	}
	return *o.AggregateDelay
}

// GetAggregateDelayOk returns a tuple with the AggregateDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Realtime) GetAggregateDelayOk() (*int32, bool) {
	if o == nil || o.AggregateDelay == nil {
		return nil, false
	}
	return o.AggregateDelay, true
}

// HasAggregateDelay returns a boolean if a field has been set.
func (o *Realtime) HasAggregateDelay() bool {
	if o != nil && o.AggregateDelay != nil {
		return true
	}

	return false
}

// SetAggregateDelay gets a reference to the given int32 and assigns it to the AggregateDelay field.
func (o *Realtime) SetAggregateDelay(v int32) {
	o.AggregateDelay = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Realtime) GetData() []RealtimeEntry {
	if o == nil || o.Data == nil {
		var ret []RealtimeEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Realtime) GetDataOk() ([]RealtimeEntry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Realtime) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RealtimeEntry and assigns it to the Data field.
func (o *Realtime) SetData(v []RealtimeEntry) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Realtime) MarshalJSON() ([]byte, error) {
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
func (o *Realtime) UnmarshalJSON(bytes []byte) (err error) {
	varRealtime := _Realtime{}

	if err = json.Unmarshal(bytes, &varRealtime); err == nil {
		*o = Realtime(varRealtime)
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

// NullableRealtime is a helper abstraction for handling nullable realtime types. 
type NullableRealtime struct {
	value *Realtime
	isSet bool
}

// Get returns the value.
func (v NullableRealtime) Get() *Realtime {
	return v.value
}

// Set modifies the value.
func (v *NullableRealtime) Set(val *Realtime) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRealtime) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRealtime) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRealtime returns a pointer to a new instance of NullableRealtime.
func NewNullableRealtime(val *Realtime) *NullableRealtime {
	return &NullableRealtime{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRealtime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRealtime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
