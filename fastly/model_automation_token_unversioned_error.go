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

// AutomationTokenUnversionedError struct for AutomationTokenUnversionedError
type AutomationTokenUnversionedError struct {
	Title                *string `json:"title,omitempty"`
	Detail               *string `json:"detail,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenUnversionedError AutomationTokenUnversionedError

// NewAutomationTokenUnversionedError instantiates a new AutomationTokenUnversionedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenUnversionedError() *AutomationTokenUnversionedError {
	this := AutomationTokenUnversionedError{}
	return &this
}

// NewAutomationTokenUnversionedErrorWithDefaults instantiates a new AutomationTokenUnversionedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenUnversionedErrorWithDefaults() *AutomationTokenUnversionedError {
	this := AutomationTokenUnversionedError{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AutomationTokenUnversionedError) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenUnversionedError) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AutomationTokenUnversionedError) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AutomationTokenUnversionedError) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AutomationTokenUnversionedError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenUnversionedError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *AutomationTokenUnversionedError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *AutomationTokenUnversionedError) SetDetail(v string) {
	o.Detail = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenUnversionedError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenUnversionedError) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenUnversionedError := _AutomationTokenUnversionedError{}

	if err = json.Unmarshal(bytes, &varAutomationTokenUnversionedError); err == nil {
		*o = AutomationTokenUnversionedError(varAutomationTokenUnversionedError)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenUnversionedError is a helper abstraction for handling nullable automationtokenunversionederror types.
type NullableAutomationTokenUnversionedError struct {
	value *AutomationTokenUnversionedError
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenUnversionedError) Get() *AutomationTokenUnversionedError {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenUnversionedError) Set(val *AutomationTokenUnversionedError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenUnversionedError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenUnversionedError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenUnversionedError returns a pointer to a new instance of NullableAutomationTokenUnversionedError.
func NewNullableAutomationTokenUnversionedError(val *AutomationTokenUnversionedError) *NullableAutomationTokenUnversionedError {
	return &NullableAutomationTokenUnversionedError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenUnversionedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenUnversionedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
