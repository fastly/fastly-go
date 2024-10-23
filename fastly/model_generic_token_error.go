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

// GenericTokenError struct for GenericTokenError
type GenericTokenError struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]any
}

type _GenericTokenError GenericTokenError

// NewGenericTokenError instantiates a new GenericTokenError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericTokenError() *GenericTokenError {
	this := GenericTokenError{}
	return &this
}

// NewGenericTokenErrorWithDefaults instantiates a new GenericTokenError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericTokenErrorWithDefaults() *GenericTokenError {
	this := GenericTokenError{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GenericTokenError) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericTokenError) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GenericTokenError) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GenericTokenError) SetMsg(v string) {
	o.Msg = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o GenericTokenError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *GenericTokenError) UnmarshalJSON(bytes []byte) (err error) {
	varGenericTokenError := _GenericTokenError{}

	if err = json.Unmarshal(bytes, &varGenericTokenError); err == nil {
		*o = GenericTokenError(varGenericTokenError)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableGenericTokenError is a helper abstraction for handling nullable generictokenerror types.
type NullableGenericTokenError struct {
	value *GenericTokenError
	isSet bool
}

// Get returns the value.
func (v NullableGenericTokenError) Get() *GenericTokenError {
	return v.value
}

// Set modifies the value.
func (v *NullableGenericTokenError) Set(val *GenericTokenError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableGenericTokenError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableGenericTokenError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGenericTokenError returns a pointer to a new instance of NullableGenericTokenError.
func NewNullableGenericTokenError(val *GenericTokenError) *NullableGenericTokenError {
	return &NullableGenericTokenError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableGenericTokenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableGenericTokenError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
