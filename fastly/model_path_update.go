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

// PathUpdate All attributes for updating a path.
type PathUpdate struct {
	// The URL path pattern, beginning with `/`. Maximum 2048 characters.
	Path                 *string `json:"path,omitempty"`
	AdditionalProperties map[string]any
}

type _PathUpdate PathUpdate

// NewPathUpdate instantiates a new PathUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathUpdate() *PathUpdate {
	this := PathUpdate{}
	return &this
}

// NewPathUpdateWithDefaults instantiates a new PathUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathUpdateWithDefaults() *PathUpdate {
	this := PathUpdate{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PathUpdate) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathUpdate) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PathUpdate) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PathUpdate) SetPath(v string) {
	o.Path = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PathUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PathUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varPathUpdate := _PathUpdate{}

	if err = json.Unmarshal(bytes, &varPathUpdate); err == nil {
		*o = PathUpdate(varPathUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePathUpdate is a helper abstraction for handling nullable pathupdate types.
type NullablePathUpdate struct {
	value *PathUpdate
	isSet bool
}

// Get returns the value.
func (v NullablePathUpdate) Get() *PathUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullablePathUpdate) Set(val *PathUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePathUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePathUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePathUpdate returns a pointer to a new instance of NullablePathUpdate.
func NewNullablePathUpdate(val *PathUpdate) *NullablePathUpdate {
	return &NullablePathUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePathUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePathUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
