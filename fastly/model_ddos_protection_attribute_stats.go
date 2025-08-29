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

// DdosProtectionAttributeStats struct for DdosProtectionAttributeStats
type DdosProtectionAttributeStats struct {
	// Name of an attribute type used in traffic stats. Currently, supported values are source_ip, country_code, host, asn, source_ip_prefix, user_agent, method_path.
	Name *string `json:"name,omitempty"`
	// Values for traffic attribute.
	Values               []DdosProtectionAttributeValue `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionAttributeStats DdosProtectionAttributeStats

// NewDdosProtectionAttributeStats instantiates a new DdosProtectionAttributeStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionAttributeStats() *DdosProtectionAttributeStats {
	this := DdosProtectionAttributeStats{}
	return &this
}

// NewDdosProtectionAttributeStatsWithDefaults instantiates a new DdosProtectionAttributeStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionAttributeStatsWithDefaults() *DdosProtectionAttributeStats {
	this := DdosProtectionAttributeStats{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DdosProtectionAttributeStats) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionAttributeStats) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DdosProtectionAttributeStats) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DdosProtectionAttributeStats) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *DdosProtectionAttributeStats) GetValues() []DdosProtectionAttributeValue {
	if o == nil || o.Values == nil {
		var ret []DdosProtectionAttributeValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionAttributeStats) GetValuesOk() ([]DdosProtectionAttributeValue, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *DdosProtectionAttributeStats) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []DdosProtectionAttributeValue and assigns it to the Values field.
func (o *DdosProtectionAttributeStats) SetValues(v []DdosProtectionAttributeValue) {
	o.Values = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionAttributeStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionAttributeStats) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionAttributeStats := _DdosProtectionAttributeStats{}

	if err = json.Unmarshal(bytes, &varDdosProtectionAttributeStats); err == nil {
		*o = DdosProtectionAttributeStats(varDdosProtectionAttributeStats)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionAttributeStats is a helper abstraction for handling nullable ddosprotectionattributestats types.
type NullableDdosProtectionAttributeStats struct {
	value *DdosProtectionAttributeStats
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionAttributeStats) Get() *DdosProtectionAttributeStats {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionAttributeStats) Set(val *DdosProtectionAttributeStats) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionAttributeStats) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionAttributeStats) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionAttributeStats returns a pointer to a new instance of NullableDdosProtectionAttributeStats.
func NewNullableDdosProtectionAttributeStats(val *DdosProtectionAttributeStats) *NullableDdosProtectionAttributeStats {
	return &NullableDdosProtectionAttributeStats{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionAttributeStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionAttributeStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
