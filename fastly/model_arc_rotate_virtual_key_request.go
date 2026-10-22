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

// ArcRotateVirtualKeyRequest Options for rotating a virtual key.
type ArcRotateVirtualKeyRequest struct {
	// The date and time the new access token expires, in RFC 3339 format.
	ExpiresAt            *time.Time `json:"expires_at,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcRotateVirtualKeyRequest ArcRotateVirtualKeyRequest

// NewArcRotateVirtualKeyRequest instantiates a new ArcRotateVirtualKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcRotateVirtualKeyRequest() *ArcRotateVirtualKeyRequest {
	this := ArcRotateVirtualKeyRequest{}
	return &this
}

// NewArcRotateVirtualKeyRequestWithDefaults instantiates a new ArcRotateVirtualKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcRotateVirtualKeyRequestWithDefaults() *ArcRotateVirtualKeyRequest {
	this := ArcRotateVirtualKeyRequest{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ArcRotateVirtualKeyRequest) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ArcRotateVirtualKeyRequest) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ArcRotateVirtualKeyRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcRotateVirtualKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcRotateVirtualKeyRequest) UnmarshalJSON(bytes []byte) (err error) {
	varArcRotateVirtualKeyRequest := _ArcRotateVirtualKeyRequest{}

	if err = json.Unmarshal(bytes, &varArcRotateVirtualKeyRequest); err == nil {
		*o = ArcRotateVirtualKeyRequest(varArcRotateVirtualKeyRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expires_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcRotateVirtualKeyRequest is a helper abstraction for handling nullable arcrotatevirtualkeyrequest types.
type NullableArcRotateVirtualKeyRequest struct {
	value *ArcRotateVirtualKeyRequest
	isSet bool
}

// Get returns the value.
func (v NullableArcRotateVirtualKeyRequest) Get() *ArcRotateVirtualKeyRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableArcRotateVirtualKeyRequest) Set(val *ArcRotateVirtualKeyRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcRotateVirtualKeyRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcRotateVirtualKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcRotateVirtualKeyRequest returns a pointer to a new instance of NullableArcRotateVirtualKeyRequest.
func NewNullableArcRotateVirtualKeyRequest(val *ArcRotateVirtualKeyRequest) *NullableArcRotateVirtualKeyRequest {
	return &NullableArcRotateVirtualKeyRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcRotateVirtualKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcRotateVirtualKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
