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

// NetworkProtectConflictError An optional conflict error detail from the API.
type NetworkProtectConflictError struct {
	// Request field associated with the conflict.
	Field *string `json:"field,omitempty"`
	// Explanation of the conflict for the field.
	Error                *string `json:"error,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectConflictError NetworkProtectConflictError

// NewNetworkProtectConflictError instantiates a new NetworkProtectConflictError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectConflictError() *NetworkProtectConflictError {
	this := NetworkProtectConflictError{}
	return &this
}

// NewNetworkProtectConflictErrorWithDefaults instantiates a new NetworkProtectConflictError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectConflictErrorWithDefaults() *NetworkProtectConflictError {
	this := NetworkProtectConflictError{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *NetworkProtectConflictError) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectConflictError) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *NetworkProtectConflictError) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *NetworkProtectConflictError) SetField(v string) {
	o.Field = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NetworkProtectConflictError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectConflictError) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NetworkProtectConflictError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *NetworkProtectConflictError) SetError(v string) {
	o.Error = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectConflictError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectConflictError) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectConflictError := _NetworkProtectConflictError{}

	if err = json.Unmarshal(bytes, &varNetworkProtectConflictError); err == nil {
		*o = NetworkProtectConflictError(varNetworkProtectConflictError)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectConflictError is a helper abstraction for handling nullable networkprotectconflicterror types.
type NullableNetworkProtectConflictError struct {
	value *NetworkProtectConflictError
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectConflictError) Get() *NetworkProtectConflictError {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectConflictError) Set(val *NetworkProtectConflictError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectConflictError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectConflictError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectConflictError returns a pointer to a new instance of NullableNetworkProtectConflictError.
func NewNullableNetworkProtectConflictError(val *NetworkProtectConflictError) *NullableNetworkProtectConflictError {
	return &NullableNetworkProtectConflictError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectConflictError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
