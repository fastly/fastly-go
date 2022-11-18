// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// HTTP3AllOf struct for HTTP3AllOf
type HTTP3AllOf struct {
	// Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision.
	FeatureRevision *int32 `json:"feature_revision,omitempty"`
	AdditionalProperties map[string]any
}

type _HTTP3AllOf HTTP3AllOf

// NewHTTP3AllOf instantiates a new HTTP3AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTP3AllOf() *HTTP3AllOf {
	this := HTTP3AllOf{}
	return &this
}

// NewHTTP3AllOfWithDefaults instantiates a new HTTP3AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTP3AllOfWithDefaults() *HTTP3AllOf {
	this := HTTP3AllOf{}
	return &this
}

// GetFeatureRevision returns the FeatureRevision field value if set, zero value otherwise.
func (o *HTTP3AllOf) GetFeatureRevision() int32 {
	if o == nil || o.FeatureRevision == nil {
		var ret int32
		return ret
	}
	return *o.FeatureRevision
}

// GetFeatureRevisionOk returns a tuple with the FeatureRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP3AllOf) GetFeatureRevisionOk() (*int32, bool) {
	if o == nil || o.FeatureRevision == nil {
		return nil, false
	}
	return o.FeatureRevision, true
}

// HasFeatureRevision returns a boolean if a field has been set.
func (o *HTTP3AllOf) HasFeatureRevision() bool {
	if o != nil && o.FeatureRevision != nil {
		return true
	}

	return false
}

// SetFeatureRevision gets a reference to the given int32 and assigns it to the FeatureRevision field.
func (o *HTTP3AllOf) SetFeatureRevision(v int32) {
	o.FeatureRevision = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HTTP3AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.FeatureRevision != nil {
		toSerialize["feature_revision"] = o.FeatureRevision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *HTTP3AllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHTTP3AllOf := _HTTP3AllOf{}

	if err = json.Unmarshal(bytes, &varHTTP3AllOf); err == nil {
		*o = HTTP3AllOf(varHTTP3AllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "feature_revision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHTTP3AllOf is a helper abstraction for handling nullable http3allof types. 
type NullableHTTP3AllOf struct {
	value *HTTP3AllOf
	isSet bool
}

// Get returns the value.
func (v NullableHTTP3AllOf) Get() *HTTP3AllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableHTTP3AllOf) Set(val *HTTP3AllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHTTP3AllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHTTP3AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHTTP3AllOf returns a pointer to a new instance of NullableHTTP3AllOf.
func NewNullableHTTP3AllOf(val *HTTP3AllOf) *NullableHTTP3AllOf {
	return &NullableHTTP3AllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHTTP3AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHTTP3AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
