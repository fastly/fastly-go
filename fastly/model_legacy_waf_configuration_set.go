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

// LegacyWafConfigurationSet struct for LegacyWafConfigurationSet
type LegacyWafConfigurationSet struct {
	// The active configuration set is the default configuration set when creating a new WAF. When Fastly adds configuration sets, the new versions become the default (active).
	Active *bool `json:"active,omitempty"`
	// The name of the configuration set.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _LegacyWafConfigurationSet LegacyWafConfigurationSet

// NewLegacyWafConfigurationSet instantiates a new LegacyWafConfigurationSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyWafConfigurationSet() *LegacyWafConfigurationSet {
	this := LegacyWafConfigurationSet{}
	return &this
}

// NewLegacyWafConfigurationSetWithDefaults instantiates a new LegacyWafConfigurationSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyWafConfigurationSetWithDefaults() *LegacyWafConfigurationSet {
	this := LegacyWafConfigurationSet{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *LegacyWafConfigurationSet) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafConfigurationSet) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *LegacyWafConfigurationSet) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *LegacyWafConfigurationSet) SetActive(v bool) {
	o.Active = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LegacyWafConfigurationSet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafConfigurationSet) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LegacyWafConfigurationSet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LegacyWafConfigurationSet) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LegacyWafConfigurationSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LegacyWafConfigurationSet) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyWafConfigurationSet := _LegacyWafConfigurationSet{}

	if err = json.Unmarshal(bytes, &varLegacyWafConfigurationSet); err == nil {
		*o = LegacyWafConfigurationSet(varLegacyWafConfigurationSet)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLegacyWafConfigurationSet is a helper abstraction for handling nullable legacywafconfigurationset types.
type NullableLegacyWafConfigurationSet struct {
	value *LegacyWafConfigurationSet
	isSet bool
}

// Get returns the value.
func (v NullableLegacyWafConfigurationSet) Get() *LegacyWafConfigurationSet {
	return v.value
}

// Set modifies the value.
func (v *NullableLegacyWafConfigurationSet) Set(val *LegacyWafConfigurationSet) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLegacyWafConfigurationSet) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLegacyWafConfigurationSet) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLegacyWafConfigurationSet returns a pointer to a new instance of NullableLegacyWafConfigurationSet.
func NewNullableLegacyWafConfigurationSet(val *LegacyWafConfigurationSet) *NullableLegacyWafConfigurationSet {
	return &NullableLegacyWafConfigurationSet{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLegacyWafConfigurationSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLegacyWafConfigurationSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
