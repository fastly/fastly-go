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

// LoggingGcsAdditional struct for LoggingGcsAdditional
type LoggingGcsAdditional struct {
	// The name of the GCS bucket.
	BucketName *string `json:"bucket_name,omitempty"`
	Path       *string `json:"path,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// Your Google Cloud Platform project ID. Required
	ProjectId            *string `json:"project_id,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGcsAdditional LoggingGcsAdditional

// NewLoggingGcsAdditional instantiates a new LoggingGcsAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGcsAdditional() *LoggingGcsAdditional {
	this := LoggingGcsAdditional{}
	var path string = "/"
	this.Path = &path
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// NewLoggingGcsAdditionalWithDefaults instantiates a new LoggingGcsAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGcsAdditionalWithDefaults() *LoggingGcsAdditional {
	this := LoggingGcsAdditional{}
	var path string = "/"
	this.Path = &path
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *LoggingGcsAdditional) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGcsAdditional) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LoggingGcsAdditional) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LoggingGcsAdditional) SetBucketName(v string) {
	o.BucketName = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *LoggingGcsAdditional) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGcsAdditional) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingGcsAdditional) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *LoggingGcsAdditional) SetPath(v string) {
	o.Path = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingGcsAdditional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingGcsAdditional) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingGcsAdditional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingGcsAdditional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}

// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingGcsAdditional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingGcsAdditional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *LoggingGcsAdditional) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGcsAdditional) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *LoggingGcsAdditional) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *LoggingGcsAdditional) SetProjectId(v string) {
	o.ProjectId = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGcsAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingGcsAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGcsAdditional := _LoggingGcsAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingGcsAdditional); err == nil {
		*o = LoggingGcsAdditional(varLoggingGcsAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bucket_name")
		delete(additionalProperties, "path")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingGcsAdditional is a helper abstraction for handling nullable logginggcsadditional types.
type NullableLoggingGcsAdditional struct {
	value *LoggingGcsAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGcsAdditional) Get() *LoggingGcsAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGcsAdditional) Set(val *LoggingGcsAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGcsAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGcsAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGcsAdditional returns a pointer to a new instance of NullableLoggingGcsAdditional.
func NewNullableLoggingGcsAdditional(val *LoggingGcsAdditional) *NullableLoggingGcsAdditional {
	return &NullableLoggingGcsAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGcsAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingGcsAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
