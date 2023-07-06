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

// HistoricalDomainsResponseAllOf struct for HistoricalDomainsResponseAllOf
type HistoricalDomainsResponseAllOf struct {
	// A list of timeseries. Each individual timeseries represents a unique combination of dimensions, such as domain, region or POP.
	Data []DomainInspectorEntry `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalDomainsResponseAllOf HistoricalDomainsResponseAllOf

// NewHistoricalDomainsResponseAllOf instantiates a new HistoricalDomainsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalDomainsResponseAllOf() *HistoricalDomainsResponseAllOf {
	this := HistoricalDomainsResponseAllOf{}
	return &this
}

// NewHistoricalDomainsResponseAllOfWithDefaults instantiates a new HistoricalDomainsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalDomainsResponseAllOfWithDefaults() *HistoricalDomainsResponseAllOf {
	this := HistoricalDomainsResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HistoricalDomainsResponseAllOf) GetData() []DomainInspectorEntry {
	if o == nil || o.Data == nil {
		var ret []DomainInspectorEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsResponseAllOf) GetDataOk() ([]DomainInspectorEntry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoricalDomainsResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DomainInspectorEntry and assigns it to the Data field.
func (o *HistoricalDomainsResponseAllOf) SetData(v []DomainInspectorEntry) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalDomainsResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *HistoricalDomainsResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalDomainsResponseAllOf := _HistoricalDomainsResponseAllOf{}

	if err = json.Unmarshal(bytes, &varHistoricalDomainsResponseAllOf); err == nil {
		*o = HistoricalDomainsResponseAllOf(varHistoricalDomainsResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalDomainsResponseAllOf is a helper abstraction for handling nullable historicaldomainsresponseallof types. 
type NullableHistoricalDomainsResponseAllOf struct {
	value *HistoricalDomainsResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalDomainsResponseAllOf) Get() *HistoricalDomainsResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalDomainsResponseAllOf) Set(val *HistoricalDomainsResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalDomainsResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalDomainsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalDomainsResponseAllOf returns a pointer to a new instance of NullableHistoricalDomainsResponseAllOf.
func NewNullableHistoricalDomainsResponseAllOf(val *HistoricalDomainsResponseAllOf) *NullableHistoricalDomainsResponseAllOf {
	return &NullableHistoricalDomainsResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalDomainsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalDomainsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
