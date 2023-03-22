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

// LoggingSftpAllOf struct for LoggingSftpAllOf
type LoggingSftpAllOf struct {
	// The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference.
	Password *string `json:"password,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// The port number.
	Port *int32 `json:"port,omitempty"`
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

type _LoggingSftpAllOf LoggingSftpAllOf

// NewLoggingSftpAllOf instantiates a new LoggingSftpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSftpAllOf() *LoggingSftpAllOf {
	this := LoggingSftpAllOf{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var port int32 = 22
	this.Port = &port
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	var secretKey string = "null"
	this.SecretKey = *NewNullableString(&secretKey)
	return &this
}

// NewLoggingSftpAllOfWithDefaults instantiates a new LoggingSftpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSftpAllOfWithDefaults() *LoggingSftpAllOf {
	this := LoggingSftpAllOf{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var port int32 = 22
	this.Port = &port
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	var secretKey string = "null"
	this.SecretKey = *NewNullableString(&secretKey)
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoggingSftpAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoggingSftpAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LoggingSftpAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSftpAllOf) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSftpAllOf) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingSftpAllOf) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingSftpAllOf) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingSftpAllOf) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingSftpAllOf) UnsetPath() {
	o.Path.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoggingSftpAllOf) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpAllOf) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoggingSftpAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LoggingSftpAllOf) SetPort(v int32) {
	o.Port = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSftpAllOf) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSftpAllOf) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingSftpAllOf) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingSftpAllOf) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingSftpAllOf) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingSftpAllOf) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSftpAllOf) GetSecretKey() string {
	if o == nil || o.SecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSftpAllOf) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingSftpAllOf) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *LoggingSftpAllOf) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}
// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *LoggingSftpAllOf) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *LoggingSftpAllOf) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetSSHKnownHosts returns the SSHKnownHosts field value if set, zero value otherwise.
func (o *LoggingSftpAllOf) GetSSHKnownHosts() string {
	if o == nil || o.SSHKnownHosts == nil {
		var ret string
		return ret
	}
	return *o.SSHKnownHosts
}

// GetSSHKnownHostsOk returns a tuple with the SSHKnownHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpAllOf) GetSSHKnownHostsOk() (*string, bool) {
	if o == nil || o.SSHKnownHosts == nil {
		return nil, false
	}
	return o.SSHKnownHosts, true
}

// HasSSHKnownHosts returns a boolean if a field has been set.
func (o *LoggingSftpAllOf) HasSSHKnownHosts() bool {
	if o != nil && o.SSHKnownHosts != nil {
		return true
	}

	return false
}

// SetSSHKnownHosts gets a reference to the given string and assigns it to the SSHKnownHosts field.
func (o *LoggingSftpAllOf) SetSSHKnownHosts(v string) {
	o.SSHKnownHosts = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingSftpAllOf) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpAllOf) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingSftpAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingSftpAllOf) SetUser(v string) {
	o.User = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSftpAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
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
func (o *LoggingSftpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSftpAllOf := _LoggingSftpAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingSftpAllOf); err == nil {
		*o = LoggingSftpAllOf(varLoggingSftpAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "path")
		delete(additionalProperties, "port")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "ssh_known_hosts")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSftpAllOf is a helper abstraction for handling nullable loggingsftpallof types. 
type NullableLoggingSftpAllOf struct {
	value *LoggingSftpAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSftpAllOf) Get() *LoggingSftpAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSftpAllOf) Set(val *LoggingSftpAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSftpAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSftpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSftpAllOf returns a pointer to a new instance of NullableLoggingSftpAllOf.
func NewNullableLoggingSftpAllOf(val *LoggingSftpAllOf) *NullableLoggingSftpAllOf {
	return &NullableLoggingSftpAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSftpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSftpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
