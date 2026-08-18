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

// ErrorResponseErrors struct for ErrorResponseErrors
type ErrorResponseErrors struct {
	Object               *string `json:"object,omitempty"`
	Property             *string `json:"property,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]any
}

type _ErrorResponseErrors ErrorResponseErrors

// NewErrorResponseErrors instantiates a new ErrorResponseErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseErrors() *ErrorResponseErrors {
	this := ErrorResponseErrors{}
	return &this
}

// NewErrorResponseErrorsWithDefaults instantiates a new ErrorResponseErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseErrorsWithDefaults() *ErrorResponseErrors {
	this := ErrorResponseErrors{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ErrorResponseErrors) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrors) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ErrorResponseErrors) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ErrorResponseErrors) SetObject(v string) {
	o.Object = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ErrorResponseErrors) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrors) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ErrorResponseErrors) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *ErrorResponseErrors) SetProperty(v string) {
	o.Property = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrorResponseErrors) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrors) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorResponseErrors) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrorResponseErrors) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ErrorResponseErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ErrorResponseErrors) UnmarshalJSON(bytes []byte) (err error) {
	varErrorResponseErrors := _ErrorResponseErrors{}

	if err = json.Unmarshal(bytes, &varErrorResponseErrors); err == nil {
		*o = ErrorResponseErrors(varErrorResponseErrors)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "property")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableErrorResponseErrors is a helper abstraction for handling nullable errorresponseerrors types.
type NullableErrorResponseErrors struct {
	value *ErrorResponseErrors
	isSet bool
}

// Get returns the value.
func (v NullableErrorResponseErrors) Get() *ErrorResponseErrors {
	return v.value
}

// Set modifies the value.
func (v *NullableErrorResponseErrors) Set(val *ErrorResponseErrors) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableErrorResponseErrors) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableErrorResponseErrors) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableErrorResponseErrors returns a pointer to a new instance of NullableErrorResponseErrors.
func NewNullableErrorResponseErrors(val *ErrorResponseErrors) *NullableErrorResponseErrors {
	return &NullableErrorResponseErrors{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableErrorResponseErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableErrorResponseErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
