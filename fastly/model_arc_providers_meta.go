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

// ArcProvidersMeta struct for ArcProvidersMeta
type ArcProvidersMeta struct {
	// The count of providers returned.
	Total                *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcProvidersMeta ArcProvidersMeta

// NewArcProvidersMeta instantiates a new ArcProvidersMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProvidersMeta() *ArcProvidersMeta {
	this := ArcProvidersMeta{}
	return &this
}

// NewArcProvidersMetaWithDefaults instantiates a new ArcProvidersMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProvidersMetaWithDefaults() *ArcProvidersMeta {
	this := ArcProvidersMeta{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ArcProvidersMeta) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcProvidersMeta) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ArcProvidersMeta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ArcProvidersMeta) SetTotal(v int32) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProvidersMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProvidersMeta) UnmarshalJSON(bytes []byte) (err error) {
	varArcProvidersMeta := _ArcProvidersMeta{}

	if err = json.Unmarshal(bytes, &varArcProvidersMeta); err == nil {
		*o = ArcProvidersMeta(varArcProvidersMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProvidersMeta is a helper abstraction for handling nullable arcprovidersmeta types.
type NullableArcProvidersMeta struct {
	value *ArcProvidersMeta
	isSet bool
}

// Get returns the value.
func (v NullableArcProvidersMeta) Get() *ArcProvidersMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProvidersMeta) Set(val *ArcProvidersMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProvidersMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProvidersMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProvidersMeta returns a pointer to a new instance of NullableArcProvidersMeta.
func NewNullableArcProvidersMeta(val *ArcProvidersMeta) *NullableArcProvidersMeta {
	return &NullableArcProvidersMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProvidersMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProvidersMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
