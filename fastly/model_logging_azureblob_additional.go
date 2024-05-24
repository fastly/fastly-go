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

// LoggingAzureblobAdditional struct for LoggingAzureblobAdditional
type LoggingAzureblobAdditional struct {
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// The unique Azure Blob Storage namespace in which your data objects are stored. Required.
	AccountName *string `json:"account_name,omitempty"`
	// The name of the Azure Blob Storage container in which to store logs. Required.
	Container *string `json:"container,omitempty"`
	// The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required.
	SasToken *string `json:"sas_token,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
	FileMaxBytes *int32 `json:"file_max_bytes,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingAzureblobAdditional LoggingAzureblobAdditional

// NewLoggingAzureblobAdditional instantiates a new LoggingAzureblobAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingAzureblobAdditional() *LoggingAzureblobAdditional {
	this := LoggingAzureblobAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// NewLoggingAzureblobAdditionalWithDefaults instantiates a new LoggingAzureblobAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingAzureblobAdditionalWithDefaults() *LoggingAzureblobAdditional {
	this := LoggingAzureblobAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingAzureblobAdditional) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingAzureblobAdditional) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingAzureblobAdditional) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingAzureblobAdditional) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingAzureblobAdditional) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingAzureblobAdditional) UnsetPath() {
	o.Path.Unset()
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *LoggingAzureblobAdditional) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblobAdditional) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *LoggingAzureblobAdditional) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *LoggingAzureblobAdditional) SetAccountName(v string) {
	o.AccountName = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *LoggingAzureblobAdditional) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblobAdditional) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *LoggingAzureblobAdditional) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *LoggingAzureblobAdditional) SetContainer(v string) {
	o.Container = &v
}

// GetSasToken returns the SasToken field value if set, zero value otherwise.
func (o *LoggingAzureblobAdditional) GetSasToken() string {
	if o == nil || o.SasToken == nil {
		var ret string
		return ret
	}
	return *o.SasToken
}

// GetSasTokenOk returns a tuple with the SasToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblobAdditional) GetSasTokenOk() (*string, bool) {
	if o == nil || o.SasToken == nil {
		return nil, false
	}
	return o.SasToken, true
}

// HasSasToken returns a boolean if a field has been set.
func (o *LoggingAzureblobAdditional) HasSasToken() bool {
	if o != nil && o.SasToken != nil {
		return true
	}

	return false
}

// SetSasToken gets a reference to the given string and assigns it to the SasToken field.
func (o *LoggingAzureblobAdditional) SetSasToken(v string) {
	o.SasToken = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingAzureblobAdditional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingAzureblobAdditional) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingAzureblobAdditional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingAzureblobAdditional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingAzureblobAdditional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingAzureblobAdditional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetFileMaxBytes returns the FileMaxBytes field value if set, zero value otherwise.
func (o *LoggingAzureblobAdditional) GetFileMaxBytes() int32 {
	if o == nil || o.FileMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.FileMaxBytes
}

// GetFileMaxBytesOk returns a tuple with the FileMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblobAdditional) GetFileMaxBytesOk() (*int32, bool) {
	if o == nil || o.FileMaxBytes == nil {
		return nil, false
	}
	return o.FileMaxBytes, true
}

// HasFileMaxBytes returns a boolean if a field has been set.
func (o *LoggingAzureblobAdditional) HasFileMaxBytes() bool {
	if o != nil && o.FileMaxBytes != nil {
		return true
	}

	return false
}

// SetFileMaxBytes gets a reference to the given int32 and assigns it to the FileMaxBytes field.
func (o *LoggingAzureblobAdditional) SetFileMaxBytes(v int32) {
	o.FileMaxBytes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingAzureblobAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.SasToken != nil {
		toSerialize["sas_token"] = o.SasToken
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.FileMaxBytes != nil {
		toSerialize["file_max_bytes"] = o.FileMaxBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingAzureblobAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingAzureblobAdditional := _LoggingAzureblobAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingAzureblobAdditional); err == nil {
		*o = LoggingAzureblobAdditional(varLoggingAzureblobAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "container")
		delete(additionalProperties, "sas_token")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "file_max_bytes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingAzureblobAdditional is a helper abstraction for handling nullable loggingazureblobadditional types. 
type NullableLoggingAzureblobAdditional struct {
	value *LoggingAzureblobAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingAzureblobAdditional) Get() *LoggingAzureblobAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingAzureblobAdditional) Set(val *LoggingAzureblobAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingAzureblobAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingAzureblobAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingAzureblobAdditional returns a pointer to a new instance of NullableLoggingAzureblobAdditional.
func NewNullableLoggingAzureblobAdditional(val *LoggingAzureblobAdditional) *NullableLoggingAzureblobAdditional {
	return &NullableLoggingAzureblobAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingAzureblobAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingAzureblobAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
