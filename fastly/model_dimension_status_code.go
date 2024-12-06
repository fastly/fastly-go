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

// DimensionStatusCode struct for DimensionStatusCode
type DimensionStatusCode struct {
	// The HTTP response code for this dimension.
	StatusCode           *string `json:"status-code,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionStatusCode DimensionStatusCode

// NewDimensionStatusCode instantiates a new DimensionStatusCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionStatusCode() *DimensionStatusCode {
	this := DimensionStatusCode{}
	return &this
}

// NewDimensionStatusCodeWithDefaults instantiates a new DimensionStatusCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionStatusCodeWithDefaults() *DimensionStatusCode {
	this := DimensionStatusCode{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *DimensionStatusCode) GetStatusCode() string {
	if o == nil || o.StatusCode == nil {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionStatusCode) GetStatusCodeOk() (*string, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *DimensionStatusCode) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *DimensionStatusCode) SetStatusCode(v string) {
	o.StatusCode = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionStatusCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.StatusCode != nil {
		toSerialize["status-code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionStatusCode) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionStatusCode := _DimensionStatusCode{}

	if err = json.Unmarshal(bytes, &varDimensionStatusCode); err == nil {
		*o = DimensionStatusCode(varDimensionStatusCode)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status-code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionStatusCode is a helper abstraction for handling nullable dimensionstatuscode types.
type NullableDimensionStatusCode struct {
	value *DimensionStatusCode
	isSet bool
}

// Get returns the value.
func (v NullableDimensionStatusCode) Get() *DimensionStatusCode {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionStatusCode) Set(val *DimensionStatusCode) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionStatusCode) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionStatusCode) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionStatusCode returns a pointer to a new instance of NullableDimensionStatusCode.
func NewNullableDimensionStatusCode(val *DimensionStatusCode) *NullableDimensionStatusCode {
	return &NullableDimensionStatusCode{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionStatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionStatusCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
