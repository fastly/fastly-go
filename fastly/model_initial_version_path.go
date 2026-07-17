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

// InitialVersionPath A path on the initial version, with its rules.
type InitialVersionPath struct {
	// The URL path pattern, beginning with `/`. Maximum 2048 characters.
	Path string `json:"path"`
	// The rules to create on this path.
	Rules                []RuleCreate `json:"rules,omitempty"`
	AdditionalProperties map[string]any
}

type _InitialVersionPath InitialVersionPath

// NewInitialVersionPath instantiates a new InitialVersionPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialVersionPath(path string) *InitialVersionPath {
	this := InitialVersionPath{}
	this.Path = path
	return &this
}

// NewInitialVersionPathWithDefaults instantiates a new InitialVersionPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialVersionPathWithDefaults() *InitialVersionPath {
	this := InitialVersionPath{}
	return &this
}

// GetPath returns the Path field value
func (o *InitialVersionPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *InitialVersionPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *InitialVersionPath) SetPath(v string) {
	o.Path = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InitialVersionPath) GetRules() []RuleCreate {
	if o == nil || o.Rules == nil {
		var ret []RuleCreate
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialVersionPath) GetRulesOk() ([]RuleCreate, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InitialVersionPath) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []RuleCreate and assigns it to the Rules field.
func (o *InitialVersionPath) SetRules(v []RuleCreate) {
	o.Rules = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InitialVersionPath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *InitialVersionPath) UnmarshalJSON(bytes []byte) (err error) {
	varInitialVersionPath := _InitialVersionPath{}

	if err = json.Unmarshal(bytes, &varInitialVersionPath); err == nil {
		*o = InitialVersionPath(varInitialVersionPath)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInitialVersionPath is a helper abstraction for handling nullable initialversionpath types.
type NullableInitialVersionPath struct {
	value *InitialVersionPath
	isSet bool
}

// Get returns the value.
func (v NullableInitialVersionPath) Get() *InitialVersionPath {
	return v.value
}

// Set modifies the value.
func (v *NullableInitialVersionPath) Set(val *InitialVersionPath) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInitialVersionPath) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInitialVersionPath) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInitialVersionPath returns a pointer to a new instance of NullableInitialVersionPath.
func NewNullableInitialVersionPath(val *InitialVersionPath) *NullableInitialVersionPath {
	return &NullableInitialVersionPath{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInitialVersionPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInitialVersionPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
