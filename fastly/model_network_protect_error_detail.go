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

// NetworkProtectErrorDetail A field-level validation error.
type NetworkProtectErrorDetail struct {
	// Request field associated with the error.
	Field string `json:"field"`
	// Explanation of why the field is invalid.
	Error                string `json:"error"`
	AdditionalProperties map[string]any
}

type _NetworkProtectErrorDetail NetworkProtectErrorDetail

// NewNetworkProtectErrorDetail instantiates a new NetworkProtectErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectErrorDetail(field string, error_ string) *NetworkProtectErrorDetail {
	this := NetworkProtectErrorDetail{}
	this.Field = field
	this.Error = error_
	return &this
}

// NewNetworkProtectErrorDetailWithDefaults instantiates a new NetworkProtectErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectErrorDetailWithDefaults() *NetworkProtectErrorDetail {
	this := NetworkProtectErrorDetail{}
	return &this
}

// GetField returns the Field field value
func (o *NetworkProtectErrorDetail) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectErrorDetail) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *NetworkProtectErrorDetail) SetField(v string) {
	o.Field = v
}

// GetError returns the Error field value
func (o *NetworkProtectErrorDetail) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectErrorDetail) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *NetworkProtectErrorDetail) SetError(v string) {
	o.Error = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectErrorDetail) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectErrorDetail := _NetworkProtectErrorDetail{}

	if err = json.Unmarshal(bytes, &varNetworkProtectErrorDetail); err == nil {
		*o = NetworkProtectErrorDetail(varNetworkProtectErrorDetail)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectErrorDetail is a helper abstraction for handling nullable networkprotecterrordetail types.
type NullableNetworkProtectErrorDetail struct {
	value *NetworkProtectErrorDetail
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectErrorDetail) Get() *NetworkProtectErrorDetail {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectErrorDetail) Set(val *NetworkProtectErrorDetail) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectErrorDetail) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectErrorDetail returns a pointer to a new instance of NullableNetworkProtectErrorDetail.
func NewNullableNetworkProtectErrorDetail(val *NetworkProtectErrorDetail) *NullableNetworkProtectErrorDetail {
	return &NullableNetworkProtectErrorDetail{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
