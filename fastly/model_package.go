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

// ModelPackage struct for ModelPackage
type ModelPackage struct {
	Metadata             *PackageMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]any
}

type _ModelPackage ModelPackage

// NewModelPackage instantiates a new ModelPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPackage() *ModelPackage {
	this := ModelPackage{}
	return &this
}

// NewModelPackageWithDefaults instantiates a new ModelPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPackageWithDefaults() *ModelPackage {
	this := ModelPackage{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ModelPackage) GetMetadata() PackageMetadata {
	if o == nil || o.Metadata == nil {
		var ret PackageMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetMetadataOk() (*PackageMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModelPackage) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given PackageMetadata and assigns it to the Metadata field.
func (o *ModelPackage) SetMetadata(v PackageMetadata) {
	o.Metadata = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ModelPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ModelPackage) UnmarshalJSON(bytes []byte) (err error) {
	varModelPackage := _ModelPackage{}

	if err = json.Unmarshal(bytes, &varModelPackage); err == nil {
		*o = ModelPackage(varModelPackage)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableModelPackage is a helper abstraction for handling nullable modelpackage types.
type NullableModelPackage struct {
	value *ModelPackage
	isSet bool
}

// Get returns the value.
func (v NullableModelPackage) Get() *ModelPackage {
	return v.value
}

// Set modifies the value.
func (v *NullableModelPackage) Set(val *ModelPackage) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableModelPackage) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableModelPackage) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableModelPackage returns a pointer to a new instance of NullableModelPackage.
func NewNullableModelPackage(val *ModelPackage) *NullableModelPackage {
	return &NullableModelPackage{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableModelPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableModelPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
