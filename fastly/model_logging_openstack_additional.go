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

// LoggingOpenstackAdditional struct for LoggingOpenstackAdditional
type LoggingOpenstackAdditional struct {
	// Your OpenStack account access key.
	AccessKey *string `json:"access_key,omitempty"`
	// The name of your OpenStack container.
	BucketName *string `json:"bucket_name,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// Your OpenStack auth url.
	URL *string `json:"url,omitempty"`
	// The username for your OpenStack account.
	User *string `json:"user,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingOpenstackAdditional LoggingOpenstackAdditional

// NewLoggingOpenstackAdditional instantiates a new LoggingOpenstackAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingOpenstackAdditional() *LoggingOpenstackAdditional {
	this := LoggingOpenstackAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// NewLoggingOpenstackAdditionalWithDefaults instantiates a new LoggingOpenstackAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingOpenstackAdditionalWithDefaults() *LoggingOpenstackAdditional {
	this := LoggingOpenstackAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *LoggingOpenstackAdditional) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingOpenstackAdditional) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoggingOpenstackAdditional) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *LoggingOpenstackAdditional) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *LoggingOpenstackAdditional) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingOpenstackAdditional) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LoggingOpenstackAdditional) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LoggingOpenstackAdditional) SetBucketName(v string) {
	o.BucketName = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingOpenstackAdditional) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingOpenstackAdditional) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingOpenstackAdditional) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingOpenstackAdditional) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingOpenstackAdditional) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingOpenstackAdditional) UnsetPath() {
	o.Path.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingOpenstackAdditional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingOpenstackAdditional) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingOpenstackAdditional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingOpenstackAdditional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingOpenstackAdditional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingOpenstackAdditional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingOpenstackAdditional) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingOpenstackAdditional) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingOpenstackAdditional) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingOpenstackAdditional) SetURL(v string) {
	o.URL = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingOpenstackAdditional) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingOpenstackAdditional) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingOpenstackAdditional) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingOpenstackAdditional) SetUser(v string) {
	o.User = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingOpenstackAdditional) MarshalJSON() ([]byte, error) {
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
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.URL != nil {
		toSerialize["url"] = o.URL
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
func (o *LoggingOpenstackAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingOpenstackAdditional := _LoggingOpenstackAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingOpenstackAdditional); err == nil {
		*o = LoggingOpenstackAdditional(varLoggingOpenstackAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "bucket_name")
		delete(additionalProperties, "path")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "url")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingOpenstackAdditional is a helper abstraction for handling nullable loggingopenstackadditional types. 
type NullableLoggingOpenstackAdditional struct {
	value *LoggingOpenstackAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingOpenstackAdditional) Get() *LoggingOpenstackAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingOpenstackAdditional) Set(val *LoggingOpenstackAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingOpenstackAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingOpenstackAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingOpenstackAdditional returns a pointer to a new instance of NullableLoggingOpenstackAdditional.
func NewNullableLoggingOpenstackAdditional(val *LoggingOpenstackAdditional) *NullableLoggingOpenstackAdditional {
	return &NullableLoggingOpenstackAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingOpenstackAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingOpenstackAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
