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
	"time"
)

// SecurityHeader struct for SecurityHeader
type SecurityHeader struct {
	Name                 *string    `json:"name,omitempty"`
	Value                *string    `json:"value,omitempty"`
	ObservedAt           *time.Time `json:"observed_at,omitempty"`
	AdditionalProperties map[string]any
}

type _SecurityHeader SecurityHeader

// NewSecurityHeader instantiates a new SecurityHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityHeader() *SecurityHeader {
	this := SecurityHeader{}
	return &this
}

// NewSecurityHeaderWithDefaults instantiates a new SecurityHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityHeaderWithDefaults() *SecurityHeader {
	this := SecurityHeader{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityHeader) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityHeader) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityHeader) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityHeader) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SecurityHeader) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityHeader) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SecurityHeader) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SecurityHeader) SetValue(v string) {
	o.Value = &v
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *SecurityHeader) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityHeader) GetObservedAtOk() (*time.Time, bool) {
	if o == nil || o.ObservedAt == nil {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *SecurityHeader) HasObservedAt() bool {
	if o != nil && o.ObservedAt != nil {
		return true
	}

	return false
}

// SetObservedAt gets a reference to the given time.Time and assigns it to the ObservedAt field.
func (o *SecurityHeader) SetObservedAt(v time.Time) {
	o.ObservedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SecurityHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ObservedAt != nil {
		toSerialize["observed_at"] = o.ObservedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *SecurityHeader) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityHeader := _SecurityHeader{}

	if err = json.Unmarshal(bytes, &varSecurityHeader); err == nil {
		*o = SecurityHeader(varSecurityHeader)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		delete(additionalProperties, "observed_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSecurityHeader is a helper abstraction for handling nullable securityheader types.
type NullableSecurityHeader struct {
	value *SecurityHeader
	isSet bool
}

// Get returns the value.
func (v NullableSecurityHeader) Get() *SecurityHeader {
	return v.value
}

// Set modifies the value.
func (v *NullableSecurityHeader) Set(val *SecurityHeader) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSecurityHeader) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSecurityHeader) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSecurityHeader returns a pointer to a new instance of NullableSecurityHeader.
func NewNullableSecurityHeader(val *SecurityHeader) *NullableSecurityHeader {
	return &NullableSecurityHeader{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSecurityHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSecurityHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
