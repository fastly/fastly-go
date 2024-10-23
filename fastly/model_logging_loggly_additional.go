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

// LoggingLogglyAdditional struct for LoggingLogglyAdditional
type LoggingLogglyAdditional struct {
	// The token to use for authentication ([https://www.loggly.com/docs/customer-token-authentication-token/](https://www.loggly.com/docs/customer-token-authentication-token/)).
	Token                *string `json:"token,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingLogglyAdditional LoggingLogglyAdditional

// NewLoggingLogglyAdditional instantiates a new LoggingLogglyAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingLogglyAdditional() *LoggingLogglyAdditional {
	this := LoggingLogglyAdditional{}
	return &this
}

// NewLoggingLogglyAdditionalWithDefaults instantiates a new LoggingLogglyAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingLogglyAdditionalWithDefaults() *LoggingLogglyAdditional {
	this := LoggingLogglyAdditional{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingLogglyAdditional) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingLogglyAdditional) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingLogglyAdditional) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingLogglyAdditional) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingLogglyAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingLogglyAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingLogglyAdditional := _LoggingLogglyAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingLogglyAdditional); err == nil {
		*o = LoggingLogglyAdditional(varLoggingLogglyAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingLogglyAdditional is a helper abstraction for handling nullable logginglogglyadditional types.
type NullableLoggingLogglyAdditional struct {
	value *LoggingLogglyAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingLogglyAdditional) Get() *LoggingLogglyAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingLogglyAdditional) Set(val *LoggingLogglyAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingLogglyAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingLogglyAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingLogglyAdditional returns a pointer to a new instance of NullableLoggingLogglyAdditional.
func NewNullableLoggingLogglyAdditional(val *LoggingLogglyAdditional) *NullableLoggingLogglyAdditional {
	return &NullableLoggingLogglyAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingLogglyAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingLogglyAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
