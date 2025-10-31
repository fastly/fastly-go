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

// DdosProtectionRuleWithStatsAllOf struct for DdosProtectionRuleWithStatsAllOf
type DdosProtectionRuleWithStatsAllOf struct {
	TrafficStats         *DdosProtectionTrafficStats `json:"traffic_stats,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionRuleWithStatsAllOf DdosProtectionRuleWithStatsAllOf

// NewDdosProtectionRuleWithStatsAllOf instantiates a new DdosProtectionRuleWithStatsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionRuleWithStatsAllOf() *DdosProtectionRuleWithStatsAllOf {
	this := DdosProtectionRuleWithStatsAllOf{}
	return &this
}

// NewDdosProtectionRuleWithStatsAllOfWithDefaults instantiates a new DdosProtectionRuleWithStatsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionRuleWithStatsAllOfWithDefaults() *DdosProtectionRuleWithStatsAllOf {
	this := DdosProtectionRuleWithStatsAllOf{}
	return &this
}

// GetTrafficStats returns the TrafficStats field value if set, zero value otherwise.
func (o *DdosProtectionRuleWithStatsAllOf) GetTrafficStats() DdosProtectionTrafficStats {
	if o == nil || o.TrafficStats == nil {
		var ret DdosProtectionTrafficStats
		return ret
	}
	return *o.TrafficStats
}

// GetTrafficStatsOk returns a tuple with the TrafficStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRuleWithStatsAllOf) GetTrafficStatsOk() (*DdosProtectionTrafficStats, bool) {
	if o == nil || o.TrafficStats == nil {
		return nil, false
	}
	return o.TrafficStats, true
}

// HasTrafficStats returns a boolean if a field has been set.
func (o *DdosProtectionRuleWithStatsAllOf) HasTrafficStats() bool {
	if o != nil && o.TrafficStats != nil {
		return true
	}

	return false
}

// SetTrafficStats gets a reference to the given DdosProtectionTrafficStats and assigns it to the TrafficStats field.
func (o *DdosProtectionRuleWithStatsAllOf) SetTrafficStats(v DdosProtectionTrafficStats) {
	o.TrafficStats = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionRuleWithStatsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TrafficStats != nil {
		toSerialize["traffic_stats"] = o.TrafficStats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionRuleWithStatsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionRuleWithStatsAllOf := _DdosProtectionRuleWithStatsAllOf{}

	if err = json.Unmarshal(bytes, &varDdosProtectionRuleWithStatsAllOf); err == nil {
		*o = DdosProtectionRuleWithStatsAllOf(varDdosProtectionRuleWithStatsAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "traffic_stats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionRuleWithStatsAllOf is a helper abstraction for handling nullable ddosprotectionrulewithstatsallof types.
type NullableDdosProtectionRuleWithStatsAllOf struct {
	value *DdosProtectionRuleWithStatsAllOf
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionRuleWithStatsAllOf) Get() *DdosProtectionRuleWithStatsAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionRuleWithStatsAllOf) Set(val *DdosProtectionRuleWithStatsAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionRuleWithStatsAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionRuleWithStatsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionRuleWithStatsAllOf returns a pointer to a new instance of NullableDdosProtectionRuleWithStatsAllOf.
func NewNullableDdosProtectionRuleWithStatsAllOf(val *DdosProtectionRuleWithStatsAllOf) *NullableDdosProtectionRuleWithStatsAllOf {
	return &NullableDdosProtectionRuleWithStatsAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionRuleWithStatsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionRuleWithStatsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
