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

// ComputeACLLookup Defines the structure of an ACL Lookup response.
type ComputeACLLookup struct {
	// A valid IPv4 or IPv6 address.
	Prefix *string `json:"prefix,omitempty"`
	// The length of address in the IP addressing space.
	Length *int32 `json:"length,omitempty"`
	// One of \"ALLOW\" or \"BLOCK\".
	Action               *string `json:"action,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeACLLookup ComputeACLLookup

// NewComputeACLLookup instantiates a new ComputeACLLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeACLLookup() *ComputeACLLookup {
	this := ComputeACLLookup{}
	return &this
}

// NewComputeACLLookupWithDefaults instantiates a new ComputeACLLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeACLLookupWithDefaults() *ComputeACLLookup {
	this := ComputeACLLookup{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ComputeACLLookup) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLLookup) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ComputeACLLookup) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ComputeACLLookup) SetPrefix(v string) {
	o.Prefix = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *ComputeACLLookup) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLLookup) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *ComputeACLLookup) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *ComputeACLLookup) SetLength(v int32) {
	o.Length = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ComputeACLLookup) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLLookup) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ComputeACLLookup) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ComputeACLLookup) SetAction(v string) {
	o.Action = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeACLLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeACLLookup) UnmarshalJSON(bytes []byte) (err error) {
	varComputeACLLookup := _ComputeACLLookup{}

	if err = json.Unmarshal(bytes, &varComputeACLLookup); err == nil {
		*o = ComputeACLLookup(varComputeACLLookup)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "length")
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeACLLookup is a helper abstraction for handling nullable computeacllookup types.
type NullableComputeACLLookup struct {
	value *ComputeACLLookup
	isSet bool
}

// Get returns the value.
func (v NullableComputeACLLookup) Get() *ComputeACLLookup {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeACLLookup) Set(val *ComputeACLLookup) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeACLLookup) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeACLLookup) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeACLLookup returns a pointer to a new instance of NullableComputeACLLookup.
func NewNullableComputeACLLookup(val *ComputeACLLookup) *NullableComputeACLLookup {
	return &NullableComputeACLLookup{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeACLLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeACLLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
