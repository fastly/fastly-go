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

// DdosProtectionErrorErrors struct for DdosProtectionErrorErrors
type DdosProtectionErrorErrors struct {
	Field                *string `json:"field,omitempty"`
	Error                *string `json:"error,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionErrorErrors DdosProtectionErrorErrors

// NewDdosProtectionErrorErrors instantiates a new DdosProtectionErrorErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionErrorErrors() *DdosProtectionErrorErrors {
	this := DdosProtectionErrorErrors{}
	return &this
}

// NewDdosProtectionErrorErrorsWithDefaults instantiates a new DdosProtectionErrorErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionErrorErrorsWithDefaults() *DdosProtectionErrorErrors {
	this := DdosProtectionErrorErrors{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *DdosProtectionErrorErrors) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionErrorErrors) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *DdosProtectionErrorErrors) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *DdosProtectionErrorErrors) SetField(v string) {
	o.Field = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DdosProtectionErrorErrors) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionErrorErrors) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DdosProtectionErrorErrors) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *DdosProtectionErrorErrors) SetError(v string) {
	o.Error = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionErrorErrors) MarshalJSON() ([]byte, error) {
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
func (o *DdosProtectionErrorErrors) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionErrorErrors := _DdosProtectionErrorErrors{}

	if err = json.Unmarshal(bytes, &varDdosProtectionErrorErrors); err == nil {
		*o = DdosProtectionErrorErrors(varDdosProtectionErrorErrors)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionErrorErrors is a helper abstraction for handling nullable ddosprotectionerrorerrors types.
type NullableDdosProtectionErrorErrors struct {
	value *DdosProtectionErrorErrors
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionErrorErrors) Get() *DdosProtectionErrorErrors {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionErrorErrors) Set(val *DdosProtectionErrorErrors) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionErrorErrors) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionErrorErrors) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionErrorErrors returns a pointer to a new instance of NullableDdosProtectionErrorErrors.
func NewNullableDdosProtectionErrorErrors(val *DdosProtectionErrorErrors) *NullableDdosProtectionErrorErrors {
	return &NullableDdosProtectionErrorErrors{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionErrorErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionErrorErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
