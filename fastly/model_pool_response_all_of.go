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

// PoolResponseAllOf struct for PoolResponseAllOf
type PoolResponseAllOf struct {
	// Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up.
	Quorum *string `json:"quorum,omitempty"`
	AdditionalProperties map[string]any
}

type _PoolResponseAllOf PoolResponseAllOf

// NewPoolResponseAllOf instantiates a new PoolResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolResponseAllOf() *PoolResponseAllOf {
	this := PoolResponseAllOf{}
	var quorum string = "75"
	this.Quorum = &quorum
	return &this
}

// NewPoolResponseAllOfWithDefaults instantiates a new PoolResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolResponseAllOfWithDefaults() *PoolResponseAllOf {
	this := PoolResponseAllOf{}
	var quorum string = "75"
	this.Quorum = &quorum
	return &this
}

// GetQuorum returns the Quorum field value if set, zero value otherwise.
func (o *PoolResponseAllOf) GetQuorum() string {
	if o == nil || o.Quorum == nil {
		var ret string
		return ret
	}
	return *o.Quorum
}

// GetQuorumOk returns a tuple with the Quorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponseAllOf) GetQuorumOk() (*string, bool) {
	if o == nil || o.Quorum == nil {
		return nil, false
	}
	return o.Quorum, true
}

// HasQuorum returns a boolean if a field has been set.
func (o *PoolResponseAllOf) HasQuorum() bool {
	if o != nil && o.Quorum != nil {
		return true
	}

	return false
}

// SetQuorum gets a reference to the given string and assigns it to the Quorum field.
func (o *PoolResponseAllOf) SetQuorum(v string) {
	o.Quorum = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PoolResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Quorum != nil {
		toSerialize["quorum"] = o.Quorum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PoolResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPoolResponseAllOf := _PoolResponseAllOf{}

	if err = json.Unmarshal(bytes, &varPoolResponseAllOf); err == nil {
		*o = PoolResponseAllOf(varPoolResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "quorum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePoolResponseAllOf is a helper abstraction for handling nullable poolresponseallof types. 
type NullablePoolResponseAllOf struct {
	value *PoolResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullablePoolResponseAllOf) Get() *PoolResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullablePoolResponseAllOf) Set(val *PoolResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePoolResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePoolResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePoolResponseAllOf returns a pointer to a new instance of NullablePoolResponseAllOf.
func NewNullablePoolResponseAllOf(val *PoolResponseAllOf) *NullablePoolResponseAllOf {
	return &NullablePoolResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePoolResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePoolResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
