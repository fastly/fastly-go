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

// ListAttackReportMeta Metadata about the request.
type ListAttackReportMeta struct {
	// The count of attack reports matching the filter.
	Total                *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _ListAttackReportMeta ListAttackReportMeta

// NewListAttackReportMeta instantiates a new ListAttackReportMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAttackReportMeta() *ListAttackReportMeta {
	this := ListAttackReportMeta{}
	return &this
}

// NewListAttackReportMetaWithDefaults instantiates a new ListAttackReportMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAttackReportMetaWithDefaults() *ListAttackReportMeta {
	this := ListAttackReportMeta{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ListAttackReportMeta) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAttackReportMeta) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ListAttackReportMeta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ListAttackReportMeta) SetTotal(v int32) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ListAttackReportMeta) MarshalJSON() ([]byte, error) {
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
func (o *ListAttackReportMeta) UnmarshalJSON(bytes []byte) (err error) {
	varListAttackReportMeta := _ListAttackReportMeta{}

	if err = json.Unmarshal(bytes, &varListAttackReportMeta); err == nil {
		*o = ListAttackReportMeta(varListAttackReportMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableListAttackReportMeta is a helper abstraction for handling nullable listattackreportmeta types.
type NullableListAttackReportMeta struct {
	value *ListAttackReportMeta
	isSet bool
}

// Get returns the value.
func (v NullableListAttackReportMeta) Get() *ListAttackReportMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableListAttackReportMeta) Set(val *ListAttackReportMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableListAttackReportMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableListAttackReportMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableListAttackReportMeta returns a pointer to a new instance of NullableListAttackReportMeta.
func NewNullableListAttackReportMeta(val *ListAttackReportMeta) *NullableListAttackReportMeta {
	return &NullableListAttackReportMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableListAttackReportMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableListAttackReportMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
