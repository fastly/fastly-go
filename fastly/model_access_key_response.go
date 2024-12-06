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

// AccessKeyResponse struct for AccessKeyResponse
type AccessKeyResponse struct {
	// Generated access key.
	AccessKey *string `json:"access_key,omitempty"`
	// Generated secret key.
	SecretKey *string `json:"secret_key,omitempty"`
	// Description for the access key.
	Description *string `json:"description,omitempty"`
	// Permissions granted to an access key.
	Permission *string  `json:"permission,omitempty"`
	Buckets    []string `json:"buckets,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt            NullableTime `json:"created_at,omitempty"`
	AdditionalProperties map[string]any
}

type _AccessKeyResponse AccessKeyResponse

// NewAccessKeyResponse instantiates a new AccessKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeyResponse() *AccessKeyResponse {
	this := AccessKeyResponse{}
	return &this
}

// NewAccessKeyResponseWithDefaults instantiates a new AccessKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyResponseWithDefaults() *AccessKeyResponse {
	this := AccessKeyResponse{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *AccessKeyResponse) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyResponse) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *AccessKeyResponse) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *AccessKeyResponse) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AccessKeyResponse) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyResponse) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AccessKeyResponse) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AccessKeyResponse) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessKeyResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessKeyResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessKeyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *AccessKeyResponse) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyResponse) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *AccessKeyResponse) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *AccessKeyResponse) SetPermission(v string) {
	o.Permission = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *AccessKeyResponse) GetBuckets() []string {
	if o == nil || o.Buckets == nil {
		var ret []string
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyResponse) GetBucketsOk() ([]string, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *AccessKeyResponse) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []string and assigns it to the Buckets field.
func (o *AccessKeyResponse) SetBuckets(v []string) {
	o.Buckets = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessKeyResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessKeyResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccessKeyResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *AccessKeyResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *AccessKeyResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *AccessKeyResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AccessKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AccessKey != nil {
		toSerialize["access_key"] = o.AccessKey
	}
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AccessKeyResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccessKeyResponse := _AccessKeyResponse{}

	if err = json.Unmarshal(bytes, &varAccessKeyResponse); err == nil {
		*o = AccessKeyResponse(varAccessKeyResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "description")
		delete(additionalProperties, "permission")
		delete(additionalProperties, "buckets")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAccessKeyResponse is a helper abstraction for handling nullable accesskeyresponse types.
type NullableAccessKeyResponse struct {
	value *AccessKeyResponse
	isSet bool
}

// Get returns the value.
func (v NullableAccessKeyResponse) Get() *AccessKeyResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableAccessKeyResponse) Set(val *AccessKeyResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAccessKeyResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAccessKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAccessKeyResponse returns a pointer to a new instance of NullableAccessKeyResponse.
func NewNullableAccessKeyResponse(val *AccessKeyResponse) *NullableAccessKeyResponse {
	return &NullableAccessKeyResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAccessKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAccessKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
