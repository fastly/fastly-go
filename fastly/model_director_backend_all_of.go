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

// DirectorBackendAllOf struct for DirectorBackendAllOf
type DirectorBackendAllOf struct {
	// The name of the backend.
	BackendName *string `json:"backend_name,omitempty"`
	// Name for the Director.
	Director *string `json:"director,omitempty"`
	AdditionalProperties map[string]any
}

type _DirectorBackendAllOf DirectorBackendAllOf

// NewDirectorBackendAllOf instantiates a new DirectorBackendAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectorBackendAllOf() *DirectorBackendAllOf {
	this := DirectorBackendAllOf{}
	return &this
}

// NewDirectorBackendAllOfWithDefaults instantiates a new DirectorBackendAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectorBackendAllOfWithDefaults() *DirectorBackendAllOf {
	this := DirectorBackendAllOf{}
	return &this
}

// GetBackendName returns the BackendName field value if set, zero value otherwise.
func (o *DirectorBackendAllOf) GetBackendName() string {
	if o == nil || o.BackendName == nil {
		var ret string
		return ret
	}
	return *o.BackendName
}

// GetBackendNameOk returns a tuple with the BackendName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorBackendAllOf) GetBackendNameOk() (*string, bool) {
	if o == nil || o.BackendName == nil {
		return nil, false
	}
	return o.BackendName, true
}

// HasBackendName returns a boolean if a field has been set.
func (o *DirectorBackendAllOf) HasBackendName() bool {
	if o != nil && o.BackendName != nil {
		return true
	}

	return false
}

// SetBackendName gets a reference to the given string and assigns it to the BackendName field.
func (o *DirectorBackendAllOf) SetBackendName(v string) {
	o.BackendName = &v
}

// GetDirector returns the Director field value if set, zero value otherwise.
func (o *DirectorBackendAllOf) GetDirector() string {
	if o == nil || o.Director == nil {
		var ret string
		return ret
	}
	return *o.Director
}

// GetDirectorOk returns a tuple with the Director field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorBackendAllOf) GetDirectorOk() (*string, bool) {
	if o == nil || o.Director == nil {
		return nil, false
	}
	return o.Director, true
}

// HasDirector returns a boolean if a field has been set.
func (o *DirectorBackendAllOf) HasDirector() bool {
	if o != nil && o.Director != nil {
		return true
	}

	return false
}

// SetDirector gets a reference to the given string and assigns it to the Director field.
func (o *DirectorBackendAllOf) SetDirector(v string) {
	o.Director = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DirectorBackendAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BackendName != nil {
		toSerialize["backend_name"] = o.BackendName
	}
	if o.Director != nil {
		toSerialize["director"] = o.Director
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *DirectorBackendAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDirectorBackendAllOf := _DirectorBackendAllOf{}

	if err = json.Unmarshal(bytes, &varDirectorBackendAllOf); err == nil {
		*o = DirectorBackendAllOf(varDirectorBackendAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "backend_name")
		delete(additionalProperties, "director")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDirectorBackendAllOf is a helper abstraction for handling nullable directorbackendallof types. 
type NullableDirectorBackendAllOf struct {
	value *DirectorBackendAllOf
	isSet bool
}

// Get returns the value.
func (v NullableDirectorBackendAllOf) Get() *DirectorBackendAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableDirectorBackendAllOf) Set(val *DirectorBackendAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDirectorBackendAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDirectorBackendAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDirectorBackendAllOf returns a pointer to a new instance of NullableDirectorBackendAllOf.
func NewNullableDirectorBackendAllOf(val *DirectorBackendAllOf) *NullableDirectorBackendAllOf {
	return &NullableDirectorBackendAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDirectorBackendAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDirectorBackendAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
