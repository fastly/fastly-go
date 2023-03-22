// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// MutualAuthenticationResponseAttributesAllOf struct for MutualAuthenticationResponseAttributesAllOf
type MutualAuthenticationResponseAttributesAllOf struct {
	// Determines whether Mutual TLS will fail closed (enforced) or fail open.
	Enforced *bool `json:"enforced,omitempty"`
	AdditionalProperties map[string]any
}

type _MutualAuthenticationResponseAttributesAllOf MutualAuthenticationResponseAttributesAllOf

// NewMutualAuthenticationResponseAttributesAllOf instantiates a new MutualAuthenticationResponseAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualAuthenticationResponseAttributesAllOf() *MutualAuthenticationResponseAttributesAllOf {
	this := MutualAuthenticationResponseAttributesAllOf{}
	return &this
}

// NewMutualAuthenticationResponseAttributesAllOfWithDefaults instantiates a new MutualAuthenticationResponseAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualAuthenticationResponseAttributesAllOfWithDefaults() *MutualAuthenticationResponseAttributesAllOf {
	this := MutualAuthenticationResponseAttributesAllOf{}
	return &this
}

// GetEnforced returns the Enforced field value if set, zero value otherwise.
func (o *MutualAuthenticationResponseAttributesAllOf) GetEnforced() bool {
	if o == nil || o.Enforced == nil {
		var ret bool
		return ret
	}
	return *o.Enforced
}

// GetEnforcedOk returns a tuple with the Enforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationResponseAttributesAllOf) GetEnforcedOk() (*bool, bool) {
	if o == nil || o.Enforced == nil {
		return nil, false
	}
	return o.Enforced, true
}

// HasEnforced returns a boolean if a field has been set.
func (o *MutualAuthenticationResponseAttributesAllOf) HasEnforced() bool {
	if o != nil && o.Enforced != nil {
		return true
	}

	return false
}

// SetEnforced gets a reference to the given bool and assigns it to the Enforced field.
func (o *MutualAuthenticationResponseAttributesAllOf) SetEnforced(v bool) {
	o.Enforced = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o MutualAuthenticationResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Enforced != nil {
		toSerialize["enforced"] = o.Enforced
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *MutualAuthenticationResponseAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMutualAuthenticationResponseAttributesAllOf := _MutualAuthenticationResponseAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varMutualAuthenticationResponseAttributesAllOf); err == nil {
		*o = MutualAuthenticationResponseAttributesAllOf(varMutualAuthenticationResponseAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enforced")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableMutualAuthenticationResponseAttributesAllOf is a helper abstraction for handling nullable mutualauthenticationresponseattributesallof types. 
type NullableMutualAuthenticationResponseAttributesAllOf struct {
	value *MutualAuthenticationResponseAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableMutualAuthenticationResponseAttributesAllOf) Get() *MutualAuthenticationResponseAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableMutualAuthenticationResponseAttributesAllOf) Set(val *MutualAuthenticationResponseAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMutualAuthenticationResponseAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMutualAuthenticationResponseAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMutualAuthenticationResponseAttributesAllOf returns a pointer to a new instance of NullableMutualAuthenticationResponseAttributesAllOf.
func NewNullableMutualAuthenticationResponseAttributesAllOf(val *MutualAuthenticationResponseAttributesAllOf) *NullableMutualAuthenticationResponseAttributesAllOf {
	return &NullableMutualAuthenticationResponseAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMutualAuthenticationResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableMutualAuthenticationResponseAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
