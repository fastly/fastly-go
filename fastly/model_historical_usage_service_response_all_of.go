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

// HistoricalUsageServiceResponseAllOf struct for HistoricalUsageServiceResponseAllOf
type HistoricalUsageServiceResponseAllOf struct {
	Data *HistoricalUsageResults `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalUsageServiceResponseAllOf HistoricalUsageServiceResponseAllOf

// NewHistoricalUsageServiceResponseAllOf instantiates a new HistoricalUsageServiceResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalUsageServiceResponseAllOf() *HistoricalUsageServiceResponseAllOf {
	this := HistoricalUsageServiceResponseAllOf{}
	return &this
}

// NewHistoricalUsageServiceResponseAllOfWithDefaults instantiates a new HistoricalUsageServiceResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalUsageServiceResponseAllOfWithDefaults() *HistoricalUsageServiceResponseAllOf {
	this := HistoricalUsageServiceResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HistoricalUsageServiceResponseAllOf) GetData() HistoricalUsageResults {
	if o == nil || o.Data == nil {
		var ret HistoricalUsageResults
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageServiceResponseAllOf) GetDataOk() (*HistoricalUsageResults, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoricalUsageServiceResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given HistoricalUsageResults and assigns it to the Data field.
func (o *HistoricalUsageServiceResponseAllOf) SetData(v HistoricalUsageResults) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalUsageServiceResponseAllOf) MarshalJSON() ([]byte, error) {
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
func (o *HistoricalUsageServiceResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalUsageServiceResponseAllOf := _HistoricalUsageServiceResponseAllOf{}

	if err = json.Unmarshal(bytes, &varHistoricalUsageServiceResponseAllOf); err == nil {
		*o = HistoricalUsageServiceResponseAllOf(varHistoricalUsageServiceResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalUsageServiceResponseAllOf is a helper abstraction for handling nullable historicalusageserviceresponseallof types. 
type NullableHistoricalUsageServiceResponseAllOf struct {
	value *HistoricalUsageServiceResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalUsageServiceResponseAllOf) Get() *HistoricalUsageServiceResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalUsageServiceResponseAllOf) Set(val *HistoricalUsageServiceResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalUsageServiceResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalUsageServiceResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalUsageServiceResponseAllOf returns a pointer to a new instance of NullableHistoricalUsageServiceResponseAllOf.
func NewNullableHistoricalUsageServiceResponseAllOf(val *HistoricalUsageServiceResponseAllOf) *NullableHistoricalUsageServiceResponseAllOf {
	return &NullableHistoricalUsageServiceResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalUsageServiceResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalUsageServiceResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
