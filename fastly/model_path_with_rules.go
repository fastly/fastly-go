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

// PathWithRules A path together with all of its rules.
type PathWithRules struct {
	Path                 *PathResponse  `json:"path,omitempty"`
	Rules                []RuleResponse `json:"rules,omitempty"`
	AdditionalProperties map[string]any
}

type _PathWithRules PathWithRules

// NewPathWithRules instantiates a new PathWithRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathWithRules() *PathWithRules {
	this := PathWithRules{}
	return &this
}

// NewPathWithRulesWithDefaults instantiates a new PathWithRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathWithRulesWithDefaults() *PathWithRules {
	this := PathWithRules{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PathWithRules) GetPath() PathResponse {
	if o == nil || o.Path == nil {
		var ret PathResponse
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathWithRules) GetPathOk() (*PathResponse, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PathWithRules) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given PathResponse and assigns it to the Path field.
func (o *PathWithRules) SetPath(v PathResponse) {
	o.Path = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *PathWithRules) GetRules() []RuleResponse {
	if o == nil || o.Rules == nil {
		var ret []RuleResponse
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathWithRules) GetRulesOk() ([]RuleResponse, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *PathWithRules) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []RuleResponse and assigns it to the Rules field.
func (o *PathWithRules) SetRules(v []RuleResponse) {
	o.Rules = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PathWithRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Path != nil {
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
func (o *PathWithRules) UnmarshalJSON(bytes []byte) (err error) {
	varPathWithRules := _PathWithRules{}

	if err = json.Unmarshal(bytes, &varPathWithRules); err == nil {
		*o = PathWithRules(varPathWithRules)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePathWithRules is a helper abstraction for handling nullable pathwithrules types.
type NullablePathWithRules struct {
	value *PathWithRules
	isSet bool
}

// Get returns the value.
func (v NullablePathWithRules) Get() *PathWithRules {
	return v.value
}

// Set modifies the value.
func (v *NullablePathWithRules) Set(val *PathWithRules) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePathWithRules) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePathWithRules) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePathWithRules returns a pointer to a new instance of NullablePathWithRules.
func NewNullablePathWithRules(val *PathWithRules) *NullablePathWithRules {
	return &NullablePathWithRules{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePathWithRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePathWithRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
