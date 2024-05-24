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

// LoggingSftpAdditional struct for LoggingSftpAdditional
type LoggingSftpAdditional struct {
	// The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference.
	Password *string `json:"password,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference.
	SecretKey NullableString `json:"secret_key,omitempty"`
	// A list of host keys for all hosts we can connect to over SFTP.
	SSHKnownHosts *string `json:"ssh_known_hosts,omitempty"`
	// The username for the server.
	User *string `json:"user,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSftpAdditional LoggingSftpAdditional

// NewLoggingSftpAdditional instantiates a new LoggingSftpAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSftpAdditional() *LoggingSftpAdditional {
	this := LoggingSftpAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	var secretKey string = "null"
	this.SecretKey = *NewNullableString(&secretKey)
	return &this
}

// NewLoggingSftpAdditionalWithDefaults instantiates a new LoggingSftpAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSftpAdditionalWithDefaults() *LoggingSftpAdditional {
	this := LoggingSftpAdditional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	var secretKey string = "null"
	this.SecretKey = *NewNullableString(&secretKey)
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoggingSftpAdditional) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpAdditional) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoggingSftpAdditional) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LoggingSftpAdditional) SetPassword(v string) {
	o.Password = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSftpAdditional) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSftpAdditional) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingSftpAdditional) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingSftpAdditional) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingSftpAdditional) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingSftpAdditional) UnsetPath() {
	o.Path.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSftpAdditional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSftpAdditional) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingSftpAdditional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingSftpAdditional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingSftpAdditional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingSftpAdditional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSftpAdditional) GetSecretKey() string {
	if o == nil || o.SecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSftpAdditional) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingSftpAdditional) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *LoggingSftpAdditional) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}
// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *LoggingSftpAdditional) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *LoggingSftpAdditional) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetSSHKnownHosts returns the SSHKnownHosts field value if set, zero value otherwise.
func (o *LoggingSftpAdditional) GetSSHKnownHosts() string {
	if o == nil || o.SSHKnownHosts == nil {
		var ret string
		return ret
	}
	return *o.SSHKnownHosts
}

// GetSSHKnownHostsOk returns a tuple with the SSHKnownHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpAdditional) GetSSHKnownHostsOk() (*string, bool) {
	if o == nil || o.SSHKnownHosts == nil {
		return nil, false
	}
	return o.SSHKnownHosts, true
}

// HasSSHKnownHosts returns a boolean if a field has been set.
func (o *LoggingSftpAdditional) HasSSHKnownHosts() bool {
	if o != nil && o.SSHKnownHosts != nil {
		return true
	}

	return false
}

// SetSSHKnownHosts gets a reference to the given string and assigns it to the SSHKnownHosts field.
func (o *LoggingSftpAdditional) SetSSHKnownHosts(v string) {
	o.SSHKnownHosts = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingSftpAdditional) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpAdditional) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingSftpAdditional) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingSftpAdditional) SetUser(v string) {
	o.User = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSftpAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.SecretKey.IsSet() {
		toSerialize["secret_key"] = o.SecretKey.Get()
	}
	if o.SSHKnownHosts != nil {
		toSerialize["ssh_known_hosts"] = o.SSHKnownHosts
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
func (o *LoggingSftpAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSftpAdditional := _LoggingSftpAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingSftpAdditional); err == nil {
		*o = LoggingSftpAdditional(varLoggingSftpAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "path")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "ssh_known_hosts")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSftpAdditional is a helper abstraction for handling nullable loggingsftpadditional types. 
type NullableLoggingSftpAdditional struct {
	value *LoggingSftpAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSftpAdditional) Get() *LoggingSftpAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSftpAdditional) Set(val *LoggingSftpAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSftpAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSftpAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSftpAdditional returns a pointer to a new instance of NullableLoggingSftpAdditional.
func NewNullableLoggingSftpAdditional(val *LoggingSftpAdditional) *NullableLoggingSftpAdditional {
	return &NullableLoggingSftpAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSftpAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSftpAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
