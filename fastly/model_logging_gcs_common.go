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

// LoggingGcsCommon struct for LoggingGcsCommon
type LoggingGcsCommon struct {
	// Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified.
	User *string `json:"user,omitempty"`
	// Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified.
	SecretKey *string `json:"secret_key,omitempty"`
	// The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided.
	AccountName          *string `json:"account_name,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGcsCommon LoggingGcsCommon

// NewLoggingGcsCommon instantiates a new LoggingGcsCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGcsCommon() *LoggingGcsCommon {
	this := LoggingGcsCommon{}
	return &this
}

// NewLoggingGcsCommonWithDefaults instantiates a new LoggingGcsCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGcsCommonWithDefaults() *LoggingGcsCommon {
	this := LoggingGcsCommon{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingGcsCommon) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGcsCommon) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingGcsCommon) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingGcsCommon) SetUser(v string) {
	o.User = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *LoggingGcsCommon) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGcsCommon) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingGcsCommon) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *LoggingGcsCommon) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *LoggingGcsCommon) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGcsCommon) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *LoggingGcsCommon) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *LoggingGcsCommon) SetAccountName(v string) {
	o.AccountName = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGcsCommon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingGcsCommon) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGcsCommon := _LoggingGcsCommon{}

	if err = json.Unmarshal(bytes, &varLoggingGcsCommon); err == nil {
		*o = LoggingGcsCommon(varLoggingGcsCommon)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "account_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingGcsCommon is a helper abstraction for handling nullable logginggcscommon types.
type NullableLoggingGcsCommon struct {
	value *LoggingGcsCommon
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGcsCommon) Get() *LoggingGcsCommon {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGcsCommon) Set(val *LoggingGcsCommon) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGcsCommon) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGcsCommon) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGcsCommon returns a pointer to a new instance of NullableLoggingGcsCommon.
func NewNullableLoggingGcsCommon(val *LoggingGcsCommon) *NullableLoggingGcsCommon {
	return &NullableLoggingGcsCommon{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGcsCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingGcsCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
