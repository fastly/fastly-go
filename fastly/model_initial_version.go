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

// InitialVersion Optional initial version payload to seed the new routing config with paths and rules in a single request.
type InitialVersion struct {
	// Whether to activate the initial version on creation.
	Activate *bool `json:"activate,omitempty"`
	// A freeform comment for the initial version.
	Comment *string `json:"comment,omitempty"`
	// The paths to create on the initial version.
	Paths                []InitialVersionPath `json:"paths,omitempty"`
	AdditionalProperties map[string]any
}

type _InitialVersion InitialVersion

// NewInitialVersion instantiates a new InitialVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialVersion() *InitialVersion {
	this := InitialVersion{}
	var activate bool = false
	this.Activate = &activate
	return &this
}

// NewInitialVersionWithDefaults instantiates a new InitialVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialVersionWithDefaults() *InitialVersion {
	this := InitialVersion{}
	var activate bool = false
	this.Activate = &activate
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *InitialVersion) GetActivate() bool {
	if o == nil || o.Activate == nil {
		var ret bool
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialVersion) GetActivateOk() (*bool, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *InitialVersion) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given bool and assigns it to the Activate field.
func (o *InitialVersion) SetActivate(v bool) {
	o.Activate = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InitialVersion) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialVersion) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InitialVersion) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InitialVersion) SetComment(v string) {
	o.Comment = &v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *InitialVersion) GetPaths() []InitialVersionPath {
	if o == nil || o.Paths == nil {
		var ret []InitialVersionPath
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialVersion) GetPathsOk() ([]InitialVersionPath, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *InitialVersion) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []InitialVersionPath and assigns it to the Paths field.
func (o *InitialVersion) SetPaths(v []InitialVersionPath) {
	o.Paths = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InitialVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *InitialVersion) UnmarshalJSON(bytes []byte) (err error) {
	varInitialVersion := _InitialVersion{}

	if err = json.Unmarshal(bytes, &varInitialVersion); err == nil {
		*o = InitialVersion(varInitialVersion)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "paths")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInitialVersion is a helper abstraction for handling nullable initialversion types.
type NullableInitialVersion struct {
	value *InitialVersion
	isSet bool
}

// Get returns the value.
func (v NullableInitialVersion) Get() *InitialVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableInitialVersion) Set(val *InitialVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInitialVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInitialVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInitialVersion returns a pointer to a new instance of NullableInitialVersion.
func NewNullableInitialVersion(val *InitialVersion) *NullableInitialVersion {
	return &NullableInitialVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInitialVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInitialVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
