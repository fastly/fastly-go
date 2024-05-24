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

// LoggingFtpAdditional struct for LoggingFtpAdditional
type LoggingFtpAdditional struct {
	// An hostname or IPv4 address.
	Address *string `json:"address,omitempty"`
	// Hostname used.
	Hostname *string `json:"hostname,omitempty"`
	// IPv4 address of the host.
	Ipv4 *string `json:"ipv4,omitempty"`
	// The password for the server. For anonymous use an email address.
	Password *string `json:"password,omitempty"`
	// The path to upload log files to. If the path ends in `/` then it is treated as a directory.
	Path *string `json:"path,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// The username for the server. Can be anonymous.
	User *string `json:"user,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingFtpAdditional LoggingFtpAdditional

// NewLoggingFtpAdditional instantiates a new LoggingFtpAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingFtpAdditional() *LoggingFtpAdditional {
	this := LoggingFtpAdditional{}
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// NewLoggingFtpAdditionalWithDefaults instantiates a new LoggingFtpAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingFtpAdditionalWithDefaults() *LoggingFtpAdditional {
	this := LoggingFtpAdditional{}
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LoggingFtpAdditional) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFtpAdditional) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LoggingFtpAdditional) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *LoggingFtpAdditional) SetAddress(v string) {
	o.Address = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *LoggingFtpAdditional) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFtpAdditional) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *LoggingFtpAdditional) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *LoggingFtpAdditional) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *LoggingFtpAdditional) GetIpv4() string {
	if o == nil || o.Ipv4 == nil {
		var ret string
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFtpAdditional) GetIpv4Ok() (*string, bool) {
	if o == nil || o.Ipv4 == nil {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *LoggingFtpAdditional) HasIpv4() bool {
	if o != nil && o.Ipv4 != nil {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given string and assigns it to the Ipv4 field.
func (o *LoggingFtpAdditional) SetIpv4(v string) {
	o.Ipv4 = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoggingFtpAdditional) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFtpAdditional) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoggingFtpAdditional) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LoggingFtpAdditional) SetPassword(v string) {
	o.Password = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *LoggingFtpAdditional) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFtpAdditional) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingFtpAdditional) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *LoggingFtpAdditional) SetPath(v string) {
	o.Path = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingFtpAdditional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingFtpAdditional) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingFtpAdditional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingFtpAdditional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingFtpAdditional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingFtpAdditional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingFtpAdditional) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFtpAdditional) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingFtpAdditional) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingFtpAdditional) SetUser(v string) {
	o.User = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingFtpAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Ipv4 != nil {
		toSerialize["ipv4"] = o.Ipv4
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
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
func (o *LoggingFtpAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingFtpAdditional := _LoggingFtpAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingFtpAdditional); err == nil {
		*o = LoggingFtpAdditional(varLoggingFtpAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "password")
		delete(additionalProperties, "path")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingFtpAdditional is a helper abstraction for handling nullable loggingftpadditional types. 
type NullableLoggingFtpAdditional struct {
	value *LoggingFtpAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingFtpAdditional) Get() *LoggingFtpAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingFtpAdditional) Set(val *LoggingFtpAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingFtpAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingFtpAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingFtpAdditional returns a pointer to a new instance of NullableLoggingFtpAdditional.
func NewNullableLoggingFtpAdditional(val *LoggingFtpAdditional) *NullableLoggingFtpAdditional {
	return &NullableLoggingFtpAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingFtpAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingFtpAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
