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

// LoggingLogglyAllOf struct for LoggingLogglyAllOf
type LoggingLogglyAllOf struct {
	// The token to use for authentication ([https://www.loggly.com/docs/customer-token-authentication-token/](https://www.loggly.com/docs/customer-token-authentication-token/)).
	Token *string `json:"token,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingLogglyAllOf LoggingLogglyAllOf

// NewLoggingLogglyAllOf instantiates a new LoggingLogglyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingLogglyAllOf() *LoggingLogglyAllOf {
	this := LoggingLogglyAllOf{}
	return &this
}

// NewLoggingLogglyAllOfWithDefaults instantiates a new LoggingLogglyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingLogglyAllOfWithDefaults() *LoggingLogglyAllOf {
	this := LoggingLogglyAllOf{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingLogglyAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingLogglyAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingLogglyAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingLogglyAllOf) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingLogglyAllOf) MarshalJSON() ([]byte, error) {
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
func (o *LoggingLogglyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingLogglyAllOf := _LoggingLogglyAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingLogglyAllOf); err == nil {
		*o = LoggingLogglyAllOf(varLoggingLogglyAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingLogglyAllOf is a helper abstraction for handling nullable logginglogglyallof types. 
type NullableLoggingLogglyAllOf struct {
	value *LoggingLogglyAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingLogglyAllOf) Get() *LoggingLogglyAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingLogglyAllOf) Set(val *LoggingLogglyAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingLogglyAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingLogglyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingLogglyAllOf returns a pointer to a new instance of NullableLoggingLogglyAllOf.
func NewNullableLoggingLogglyAllOf(val *LoggingLogglyAllOf) *NullableLoggingLogglyAllOf {
	return &NullableLoggingLogglyAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingLogglyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingLogglyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
