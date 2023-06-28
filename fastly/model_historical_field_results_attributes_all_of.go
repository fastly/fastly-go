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

// HistoricalFieldResultsAttributesAllOf struct for HistoricalFieldResultsAttributesAllOf
type HistoricalFieldResultsAttributesAllOf struct {
	ServiceID *ReadOnlyIDService `json:"service_id,omitempty"`
	StartTime *int32 `json:"start_time,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalFieldResultsAttributesAllOf HistoricalFieldResultsAttributesAllOf

// NewHistoricalFieldResultsAttributesAllOf instantiates a new HistoricalFieldResultsAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalFieldResultsAttributesAllOf() *HistoricalFieldResultsAttributesAllOf {
	this := HistoricalFieldResultsAttributesAllOf{}
	return &this
}

// NewHistoricalFieldResultsAttributesAllOfWithDefaults instantiates a new HistoricalFieldResultsAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalFieldResultsAttributesAllOfWithDefaults() *HistoricalFieldResultsAttributesAllOf {
	this := HistoricalFieldResultsAttributesAllOf{}
	return &this
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *HistoricalFieldResultsAttributesAllOf) GetServiceID() ReadOnlyIDService {
	if o == nil || o.ServiceID == nil {
		var ret ReadOnlyIDService
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalFieldResultsAttributesAllOf) GetServiceIDOk() (*ReadOnlyIDService, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *HistoricalFieldResultsAttributesAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given ReadOnlyIDService and assigns it to the ServiceID field.
func (o *HistoricalFieldResultsAttributesAllOf) SetServiceID(v ReadOnlyIDService) {
	o.ServiceID = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *HistoricalFieldResultsAttributesAllOf) GetStartTime() int32 {
	if o == nil || o.StartTime == nil {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalFieldResultsAttributesAllOf) GetStartTimeOk() (*int32, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *HistoricalFieldResultsAttributesAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *HistoricalFieldResultsAttributesAllOf) SetStartTime(v int32) {
	o.StartTime = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalFieldResultsAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *HistoricalFieldResultsAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalFieldResultsAttributesAllOf := _HistoricalFieldResultsAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varHistoricalFieldResultsAttributesAllOf); err == nil {
		*o = HistoricalFieldResultsAttributesAllOf(varHistoricalFieldResultsAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "start_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalFieldResultsAttributesAllOf is a helper abstraction for handling nullable historicalfieldresultsattributesallof types. 
type NullableHistoricalFieldResultsAttributesAllOf struct {
	value *HistoricalFieldResultsAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalFieldResultsAttributesAllOf) Get() *HistoricalFieldResultsAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalFieldResultsAttributesAllOf) Set(val *HistoricalFieldResultsAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalFieldResultsAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalFieldResultsAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalFieldResultsAttributesAllOf returns a pointer to a new instance of NullableHistoricalFieldResultsAttributesAllOf.
func NewNullableHistoricalFieldResultsAttributesAllOf(val *HistoricalFieldResultsAttributesAllOf) *NullableHistoricalFieldResultsAttributesAllOf {
	return &NullableHistoricalFieldResultsAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalFieldResultsAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalFieldResultsAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
