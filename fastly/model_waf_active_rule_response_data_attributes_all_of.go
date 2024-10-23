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

// WafActiveRuleResponseDataAttributesAllOf struct for WafActiveRuleResponseDataAttributesAllOf
type WafActiveRuleResponseDataAttributesAllOf struct {
	// The latest rule revision number that is available for the associated rule revision.
	LatestRevision *int32 `json:"latest_revision,omitempty"`
	// Indicates if the associated rule revision is up to date or not.
	Outdated             *bool `json:"outdated,omitempty"`
	AdditionalProperties map[string]any
}

type _WafActiveRuleResponseDataAttributesAllOf WafActiveRuleResponseDataAttributesAllOf

// NewWafActiveRuleResponseDataAttributesAllOf instantiates a new WafActiveRuleResponseDataAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafActiveRuleResponseDataAttributesAllOf() *WafActiveRuleResponseDataAttributesAllOf {
	this := WafActiveRuleResponseDataAttributesAllOf{}
	return &this
}

// NewWafActiveRuleResponseDataAttributesAllOfWithDefaults instantiates a new WafActiveRuleResponseDataAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafActiveRuleResponseDataAttributesAllOfWithDefaults() *WafActiveRuleResponseDataAttributesAllOf {
	this := WafActiveRuleResponseDataAttributesAllOf{}
	return &this
}

// GetLatestRevision returns the LatestRevision field value if set, zero value otherwise.
func (o *WafActiveRuleResponseDataAttributesAllOf) GetLatestRevision() int32 {
	if o == nil || o.LatestRevision == nil {
		var ret int32
		return ret
	}
	return *o.LatestRevision
}

// GetLatestRevisionOk returns a tuple with the LatestRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponseDataAttributesAllOf) GetLatestRevisionOk() (*int32, bool) {
	if o == nil || o.LatestRevision == nil {
		return nil, false
	}
	return o.LatestRevision, true
}

// HasLatestRevision returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAttributesAllOf) HasLatestRevision() bool {
	if o != nil && o.LatestRevision != nil {
		return true
	}

	return false
}

// SetLatestRevision gets a reference to the given int32 and assigns it to the LatestRevision field.
func (o *WafActiveRuleResponseDataAttributesAllOf) SetLatestRevision(v int32) {
	o.LatestRevision = &v
}

// GetOutdated returns the Outdated field value if set, zero value otherwise.
func (o *WafActiveRuleResponseDataAttributesAllOf) GetOutdated() bool {
	if o == nil || o.Outdated == nil {
		var ret bool
		return ret
	}
	return *o.Outdated
}

// GetOutdatedOk returns a tuple with the Outdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponseDataAttributesAllOf) GetOutdatedOk() (*bool, bool) {
	if o == nil || o.Outdated == nil {
		return nil, false
	}
	return o.Outdated, true
}

// HasOutdated returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAttributesAllOf) HasOutdated() bool {
	if o != nil && o.Outdated != nil {
		return true
	}

	return false
}

// SetOutdated gets a reference to the given bool and assigns it to the Outdated field.
func (o *WafActiveRuleResponseDataAttributesAllOf) SetOutdated(v bool) {
	o.Outdated = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRuleResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.LatestRevision != nil {
		toSerialize["latest_revision"] = o.LatestRevision
	}
	if o.Outdated != nil {
		toSerialize["outdated"] = o.Outdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafActiveRuleResponseDataAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafActiveRuleResponseDataAttributesAllOf := _WafActiveRuleResponseDataAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varWafActiveRuleResponseDataAttributesAllOf); err == nil {
		*o = WafActiveRuleResponseDataAttributesAllOf(varWafActiveRuleResponseDataAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "latest_revision")
		delete(additionalProperties, "outdated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafActiveRuleResponseDataAttributesAllOf is a helper abstraction for handling nullable wafactiveruleresponsedataattributesallof types.
type NullableWafActiveRuleResponseDataAttributesAllOf struct {
	value *WafActiveRuleResponseDataAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRuleResponseDataAttributesAllOf) Get() *WafActiveRuleResponseDataAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRuleResponseDataAttributesAllOf) Set(val *WafActiveRuleResponseDataAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRuleResponseDataAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRuleResponseDataAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRuleResponseDataAttributesAllOf returns a pointer to a new instance of NullableWafActiveRuleResponseDataAttributesAllOf.
func NewNullableWafActiveRuleResponseDataAttributesAllOf(val *WafActiveRuleResponseDataAttributesAllOf) *NullableWafActiveRuleResponseDataAttributesAllOf {
	return &NullableWafActiveRuleResponseDataAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRuleResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafActiveRuleResponseDataAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
