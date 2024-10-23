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

// ServiceDetailAllOf struct for ServiceDetailAllOf
type ServiceDetailAllOf struct {
	ActiveVersion        NullableServiceVersionDetailOrNull `json:"active_version,omitempty"`
	Version              *ServiceVersionDetail              `json:"version,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceDetailAllOf ServiceDetailAllOf

// NewServiceDetailAllOf instantiates a new ServiceDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDetailAllOf() *ServiceDetailAllOf {
	this := ServiceDetailAllOf{}
	return &this
}

// NewServiceDetailAllOfWithDefaults instantiates a new ServiceDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDetailAllOfWithDefaults() *ServiceDetailAllOf {
	this := ServiceDetailAllOf{}
	return &this
}

// GetActiveVersion returns the ActiveVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDetailAllOf) GetActiveVersion() ServiceVersionDetailOrNull {
	if o == nil || o.ActiveVersion.Get() == nil {
		var ret ServiceVersionDetailOrNull
		return ret
	}
	return *o.ActiveVersion.Get()
}

// GetActiveVersionOk returns a tuple with the ActiveVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDetailAllOf) GetActiveVersionOk() (*ServiceVersionDetailOrNull, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveVersion.Get(), o.ActiveVersion.IsSet()
}

// HasActiveVersion returns a boolean if a field has been set.
func (o *ServiceDetailAllOf) HasActiveVersion() bool {
	if o != nil && o.ActiveVersion.IsSet() {
		return true
	}

	return false
}

// SetActiveVersion gets a reference to the given NullableServiceVersionDetailOrNull and assigns it to the ActiveVersion field.
func (o *ServiceDetailAllOf) SetActiveVersion(v ServiceVersionDetailOrNull) {
	o.ActiveVersion.Set(&v)
}

// SetActiveVersionNil sets the value for ActiveVersion to be an explicit nil
func (o *ServiceDetailAllOf) SetActiveVersionNil() {
	o.ActiveVersion.Set(nil)
}

// UnsetActiveVersion ensures that no value is present for ActiveVersion, not even an explicit nil
func (o *ServiceDetailAllOf) UnsetActiveVersion() {
	o.ActiveVersion.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceDetailAllOf) GetVersion() ServiceVersionDetail {
	if o == nil || o.Version == nil {
		var ret ServiceVersionDetail
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetailAllOf) GetVersionOk() (*ServiceVersionDetail, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceDetailAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given ServiceVersionDetail and assigns it to the Version field.
func (o *ServiceDetailAllOf) SetVersion(v ServiceVersionDetail) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ActiveVersion.IsSet() {
		toSerialize["active_version"] = o.ActiveVersion.Get()
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ServiceDetailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServiceDetailAllOf := _ServiceDetailAllOf{}

	if err = json.Unmarshal(bytes, &varServiceDetailAllOf); err == nil {
		*o = ServiceDetailAllOf(varServiceDetailAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active_version")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceDetailAllOf is a helper abstraction for handling nullable servicedetailallof types.
type NullableServiceDetailAllOf struct {
	value *ServiceDetailAllOf
	isSet bool
}

// Get returns the value.
func (v NullableServiceDetailAllOf) Get() *ServiceDetailAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceDetailAllOf) Set(val *ServiceDetailAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceDetailAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceDetailAllOf returns a pointer to a new instance of NullableServiceDetailAllOf.
func NewNullableServiceDetailAllOf(val *ServiceDetailAllOf) *NullableServiceDetailAllOf {
	return &NullableServiceDetailAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
