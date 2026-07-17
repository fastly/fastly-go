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

// InlineResponse20015 struct for InlineResponse20015
type InlineResponse20015 struct {
	// Time-stamp (GMT) when the domain_ownership validation will expire.
	ExpiresAt            *string `json:"expires_at,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse20015 InlineResponse20015

// NewInlineResponse20015 instantiates a new InlineResponse20015 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20015() *InlineResponse20015 {
	this := InlineResponse20015{}
	return &this
}

// NewInlineResponse20015WithDefaults instantiates a new InlineResponse20015 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20015WithDefaults() *InlineResponse20015 {
	this := InlineResponse20015{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *InlineResponse20015) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *InlineResponse20015) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *InlineResponse20015) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse20015) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse20015) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse20015 := _InlineResponse20015{}

	if err = json.Unmarshal(bytes, &varInlineResponse20015); err == nil {
		*o = InlineResponse20015(varInlineResponse20015)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expires_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse20015 is a helper abstraction for handling nullable inlineresponse20015 types.
type NullableInlineResponse20015 struct {
	value *InlineResponse20015
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse20015) Get() *InlineResponse20015 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse20015) Set(val *InlineResponse20015) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse20015) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse20015) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse20015 returns a pointer to a new instance of NullableInlineResponse20015.
func NewNullableInlineResponse20015(val *InlineResponse20015) *NullableInlineResponse20015 {
	return &NullableInlineResponse20015{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse20015) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse20015) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
