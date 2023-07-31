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

// LoggingDigitaloceanAdditional struct for LoggingDigitaloceanAdditional
type LoggingDigitaloceanAdditional struct {
	// The name of the DigitalOcean Space.
	BucketName *string `json:"bucket_name,omitempty"`
	// Your DigitalOcean Spaces account access key.
	AccessKey *string `json:"access_key,omitempty"`
	// Your DigitalOcean Spaces account secret key.
	SecretKey *string `json:"secret_key,omitempty"`
	// The domain of the DigitalOcean Spaces endpoint.
	Domain *string `json:"domain,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingDigitaloceanAdditional LoggingDigitaloceanAdditional

// NewLoggingDigitaloceanAdditional instantiates a new LoggingDigitaloceanAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingDigitaloceanAdditional() *LoggingDigitaloceanAdditional {
	this := LoggingDigitaloceanAdditional{}
	var domain string = "nyc3.digitaloceanspaces.com"
	this.Domain = &domain
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// NewLoggingDigitaloceanAdditionalWithDefaults instantiates a new LoggingDigitaloceanAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingDigitaloceanAdditionalWithDefaults() *LoggingDigitaloceanAdditional {
	this := LoggingDigitaloceanAdditional{}
	var domain string = "nyc3.digitaloceanspaces.com"
	this.Domain = &domain
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *LoggingDigitaloceanAdditional) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingDigitaloceanAdditional) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LoggingDigitaloceanAdditional) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LoggingDigitaloceanAdditional) SetBucketName(v string) {
	o.BucketName = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *LoggingDigitaloceanAdditional) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingDigitaloceanAdditional) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoggingDigitaloceanAdditional) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *LoggingDigitaloceanAdditional) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *LoggingDigitaloceanAdditional) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingDigitaloceanAdditional) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingDigitaloceanAdditional) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *LoggingDigitaloceanAdditional) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LoggingDigitaloceanAdditional) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingDigitaloceanAdditional) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LoggingDigitaloceanAdditional) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *LoggingDigitaloceanAdditional) SetDomain(v string) {
	o.Domain = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingDigitaloceanAdditional) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingDigitaloceanAdditional) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingDigitaloceanAdditional) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingDigitaloceanAdditional) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingDigitaloceanAdditional) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingDigitaloceanAdditional) UnsetPath() {
	o.Path.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingDigitaloceanAdditional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingDigitaloceanAdditional) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingDigitaloceanAdditional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingDigitaloceanAdditional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingDigitaloceanAdditional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingDigitaloceanAdditional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingDigitaloceanAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.AccessKey != nil {
		toSerialize["access_key"] = o.AccessKey
	}
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingDigitaloceanAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingDigitaloceanAdditional := _LoggingDigitaloceanAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingDigitaloceanAdditional); err == nil {
		*o = LoggingDigitaloceanAdditional(varLoggingDigitaloceanAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bucket_name")
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "path")
		delete(additionalProperties, "public_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingDigitaloceanAdditional is a helper abstraction for handling nullable loggingdigitaloceanadditional types. 
type NullableLoggingDigitaloceanAdditional struct {
	value *LoggingDigitaloceanAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingDigitaloceanAdditional) Get() *LoggingDigitaloceanAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingDigitaloceanAdditional) Set(val *LoggingDigitaloceanAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingDigitaloceanAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingDigitaloceanAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingDigitaloceanAdditional returns a pointer to a new instance of NullableLoggingDigitaloceanAdditional.
func NewNullableLoggingDigitaloceanAdditional(val *LoggingDigitaloceanAdditional) *NullableLoggingDigitaloceanAdditional {
	return &NullableLoggingDigitaloceanAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingDigitaloceanAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingDigitaloceanAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
