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

// VersionResponseAllOf struct for VersionResponseAllOf
type VersionResponseAllOf struct {
	ServiceID *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _VersionResponseAllOf VersionResponseAllOf

// NewVersionResponseAllOf instantiates a new VersionResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionResponseAllOf() *VersionResponseAllOf {
	this := VersionResponseAllOf{}
	return &this
}

// NewVersionResponseAllOfWithDefaults instantiates a new VersionResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionResponseAllOfWithDefaults() *VersionResponseAllOf {
	this := VersionResponseAllOf{}
	return &this
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *VersionResponseAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionResponseAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *VersionResponseAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *VersionResponseAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o VersionResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *VersionResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVersionResponseAllOf := _VersionResponseAllOf{}

	if err = json.Unmarshal(bytes, &varVersionResponseAllOf); err == nil {
		*o = VersionResponseAllOf(varVersionResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVersionResponseAllOf is a helper abstraction for handling nullable versionresponseallof types. 
type NullableVersionResponseAllOf struct {
	value *VersionResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableVersionResponseAllOf) Get() *VersionResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableVersionResponseAllOf) Set(val *VersionResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVersionResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVersionResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVersionResponseAllOf returns a pointer to a new instance of NullableVersionResponseAllOf.
func NewNullableVersionResponseAllOf(val *VersionResponseAllOf) *NullableVersionResponseAllOf {
	return &NullableVersionResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVersionResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableVersionResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
