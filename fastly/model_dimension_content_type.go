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

// DimensionContentType struct for DimensionContentType
type DimensionContentType struct {
	// The content type of the response for this dimension.
	ContentType          *string `json:"content_type,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionContentType DimensionContentType

// NewDimensionContentType instantiates a new DimensionContentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionContentType() *DimensionContentType {
	this := DimensionContentType{}
	return &this
}

// NewDimensionContentTypeWithDefaults instantiates a new DimensionContentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionContentTypeWithDefaults() *DimensionContentType {
	this := DimensionContentType{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *DimensionContentType) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionContentType) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *DimensionContentType) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *DimensionContentType) SetContentType(v string) {
	o.ContentType = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionContentType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionContentType) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionContentType := _DimensionContentType{}

	if err = json.Unmarshal(bytes, &varDimensionContentType); err == nil {
		*o = DimensionContentType(varDimensionContentType)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionContentType is a helper abstraction for handling nullable dimensioncontenttype types.
type NullableDimensionContentType struct {
	value *DimensionContentType
	isSet bool
}

// Get returns the value.
func (v NullableDimensionContentType) Get() *DimensionContentType {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionContentType) Set(val *DimensionContentType) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionContentType) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionContentType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionContentType returns a pointer to a new instance of NullableDimensionContentType.
func NewNullableDimensionContentType(val *DimensionContentType) *NullableDimensionContentType {
	return &NullableDimensionContentType{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
