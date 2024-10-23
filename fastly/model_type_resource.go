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
	"fmt"
)

// TypeResource Resource type
type TypeResource string

// List of resourceTyperesource
const (
	TYPERESOURCE_KV_STORE     TypeResource = "kv-store"
	TYPERESOURCE_SECRET_STORE TypeResource = "secret-store"
	TYPERESOURCE_CONFIG       TypeResource = "config"
)

// AllowedTypeResourceEnumValues All allowed values of TypeResource enum
var AllowedTypeResourceEnumValues = []TypeResource{
	"kv-store",
	"secret-store",
	"config",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeResource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeResource(value)
	for _, existing := range AllowedTypeResourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeResource", value)
}

// NewTypeResourceFromValue returns a pointer to a valid TypeResource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeResourceFromValue(v string) (*TypeResource, error) {
	ev := TypeResource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeResource: valid values are %v", v, AllowedTypeResourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeResource) IsValid() bool {
	for _, existing := range AllowedTypeResourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTyperesource value
func (v TypeResource) Ptr() *TypeResource {
	return &v
}

// NullableTypeResource is a helper abstraction for handling nullable typeresource types.
type NullableTypeResource struct {
	value *TypeResource
	isSet bool
}

// Get returns the value.
func (v NullableTypeResource) Get() *TypeResource {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeResource) Set(val *TypeResource) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeResource) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeResource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeResource returns a pointer to a new instance of NullableTypeResource.
func NewNullableTypeResource(val *TypeResource) *NullableTypeResource {
	return &NullableTypeResource{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
