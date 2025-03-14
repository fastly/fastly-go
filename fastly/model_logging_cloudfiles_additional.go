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

// LoggingCloudfilesAdditional struct for LoggingCloudfilesAdditional
type LoggingCloudfilesAdditional struct {
	// Your Cloud Files account access key.
	AccessKey *string `json:"access_key,omitempty"`
	// The name of your Cloud Files container.
	BucketName *string `json:"bucket_name,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// The region to stream logs to.
	Region NullableString `json:"region,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// The username for your Cloud Files account.
	User                 *string `json:"user,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingCloudfilesAdditional LoggingCloudfilesAdditional

// NewLoggingCloudfilesAdditional instantiates a new LoggingCloudfilesAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingCloudfilesAdditional() *LoggingCloudfilesAdditional {
	this := LoggingCloudfilesAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// NewLoggingCloudfilesAdditionalWithDefaults instantiates a new LoggingCloudfilesAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingCloudfilesAdditionalWithDefaults() *LoggingCloudfilesAdditional {
	this := LoggingCloudfilesAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *LoggingCloudfilesAdditional) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingCloudfilesAdditional) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoggingCloudfilesAdditional) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *LoggingCloudfilesAdditional) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *LoggingCloudfilesAdditional) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingCloudfilesAdditional) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LoggingCloudfilesAdditional) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LoggingCloudfilesAdditional) SetBucketName(v string) {
	o.BucketName = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingCloudfilesAdditional) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingCloudfilesAdditional) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingCloudfilesAdditional) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingCloudfilesAdditional) SetPath(v string) {
	o.Path.Set(&v)
}

// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingCloudfilesAdditional) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingCloudfilesAdditional) UnsetPath() {
	o.Path.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingCloudfilesAdditional) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingCloudfilesAdditional) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *LoggingCloudfilesAdditional) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *LoggingCloudfilesAdditional) SetRegion(v string) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *LoggingCloudfilesAdditional) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *LoggingCloudfilesAdditional) UnsetRegion() {
	o.Region.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingCloudfilesAdditional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingCloudfilesAdditional) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingCloudfilesAdditional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingCloudfilesAdditional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}

// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingCloudfilesAdditional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingCloudfilesAdditional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingCloudfilesAdditional) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingCloudfilesAdditional) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingCloudfilesAdditional) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingCloudfilesAdditional) SetUser(v string) {
	o.User = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingCloudfilesAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AccessKey != nil {
		toSerialize["access_key"] = o.AccessKey
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingCloudfilesAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingCloudfilesAdditional := _LoggingCloudfilesAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingCloudfilesAdditional); err == nil {
		*o = LoggingCloudfilesAdditional(varLoggingCloudfilesAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "bucket_name")
		delete(additionalProperties, "path")
		delete(additionalProperties, "region")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingCloudfilesAdditional is a helper abstraction for handling nullable loggingcloudfilesadditional types.
type NullableLoggingCloudfilesAdditional struct {
	value *LoggingCloudfilesAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingCloudfilesAdditional) Get() *LoggingCloudfilesAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingCloudfilesAdditional) Set(val *LoggingCloudfilesAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingCloudfilesAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingCloudfilesAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingCloudfilesAdditional returns a pointer to a new instance of NullableLoggingCloudfilesAdditional.
func NewNullableLoggingCloudfilesAdditional(val *LoggingCloudfilesAdditional) *NullableLoggingCloudfilesAdditional {
	return &NullableLoggingCloudfilesAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingCloudfilesAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingCloudfilesAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
