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

// AccessKey struct for AccessKey
type AccessKey struct {
	// A description of the access key.
	Description string `json:"description"`
	// The permissions granted to an access key.
	Permission           string   `json:"permission"`
	Buckets              []string `json:"buckets,omitempty"`
	AdditionalProperties map[string]any
}

type _AccessKey AccessKey

// NewAccessKey instantiates a new AccessKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKey(description string, permission string) *AccessKey {
	this := AccessKey{}
	this.Description = description
	this.Permission = permission
	return &this
}

// NewAccessKeyWithDefaults instantiates a new AccessKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyWithDefaults() *AccessKey {
	this := AccessKey{}
	return &this
}

// GetDescription returns the Description field value
func (o *AccessKey) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AccessKey) SetDescription(v string) {
	o.Description = v
}

// GetPermission returns the Permission field value
func (o *AccessKey) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *AccessKey) SetPermission(v string) {
	o.Permission = v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *AccessKey) GetBuckets() []string {
	if o == nil || o.Buckets == nil {
		var ret []string
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetBucketsOk() ([]string, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *AccessKey) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []string and assigns it to the Buckets field.
func (o *AccessKey) SetBuckets(v []string) {
	o.Buckets = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AccessKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AccessKey) UnmarshalJSON(bytes []byte) (err error) {
	varAccessKey := _AccessKey{}

	if err = json.Unmarshal(bytes, &varAccessKey); err == nil {
		*o = AccessKey(varAccessKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "permission")
		delete(additionalProperties, "buckets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAccessKey is a helper abstraction for handling nullable accesskey types.
type NullableAccessKey struct {
	value *AccessKey
	isSet bool
}

// Get returns the value.
func (v NullableAccessKey) Get() *AccessKey {
	return v.value
}

// Set modifies the value.
func (v *NullableAccessKey) Set(val *AccessKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAccessKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAccessKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAccessKey returns a pointer to a new instance of NullableAccessKey.
func NewNullableAccessKey(val *AccessKey) *NullableAccessKey {
	return &NullableAccessKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAccessKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
