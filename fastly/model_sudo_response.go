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
	"time"
)

// SudoResponse struct for SudoResponse
type SudoResponse struct {
	// A UTC time-stamp of when sudo access will expire. If blank, sudo access expires five minutes after the request.
	ExpiryTime           *time.Time `json:"expiry_time,omitempty"`
	AdditionalProperties map[string]any
}

type _SudoResponse SudoResponse

// NewSudoResponse instantiates a new SudoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSudoResponse() *SudoResponse {
	this := SudoResponse{}
	return &this
}

// NewSudoResponseWithDefaults instantiates a new SudoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSudoResponseWithDefaults() *SudoResponse {
	this := SudoResponse{}
	return &this
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *SudoResponse) GetExpiryTime() time.Time {
	if o == nil || o.ExpiryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoResponse) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *SudoResponse) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given time.Time and assigns it to the ExpiryTime field.
func (o *SudoResponse) SetExpiryTime(v time.Time) {
	o.ExpiryTime = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SudoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ExpiryTime != nil {
		toSerialize["expiry_time"] = o.ExpiryTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *SudoResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSudoResponse := _SudoResponse{}

	if err = json.Unmarshal(bytes, &varSudoResponse); err == nil {
		*o = SudoResponse(varSudoResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expiry_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSudoResponse is a helper abstraction for handling nullable sudoresponse types.
type NullableSudoResponse struct {
	value *SudoResponse
	isSet bool
}

// Get returns the value.
func (v NullableSudoResponse) Get() *SudoResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableSudoResponse) Set(val *SudoResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSudoResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSudoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSudoResponse returns a pointer to a new instance of NullableSudoResponse.
func NewNullableSudoResponse(val *SudoResponse) *NullableSudoResponse {
	return &NullableSudoResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSudoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSudoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
