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

// LoggingLogshuttleAllOf struct for LoggingLogshuttleAllOf
type LoggingLogshuttleAllOf struct {
	// The data authentication token associated with this endpoint.
	Token NullableString `json:"token,omitempty"`
	// The URL to stream logs to.
	URL *string `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingLogshuttleAllOf LoggingLogshuttleAllOf

// NewLoggingLogshuttleAllOf instantiates a new LoggingLogshuttleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingLogshuttleAllOf() *LoggingLogshuttleAllOf {
	this := LoggingLogshuttleAllOf{}
	return &this
}

// NewLoggingLogshuttleAllOfWithDefaults instantiates a new LoggingLogshuttleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingLogshuttleAllOfWithDefaults() *LoggingLogshuttleAllOf {
	this := LoggingLogshuttleAllOf{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingLogshuttleAllOf) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingLogshuttleAllOf) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingLogshuttleAllOf) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *LoggingLogshuttleAllOf) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *LoggingLogshuttleAllOf) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *LoggingLogshuttleAllOf) UnsetToken() {
	o.Token.Unset()
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingLogshuttleAllOf) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingLogshuttleAllOf) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingLogshuttleAllOf) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingLogshuttleAllOf) SetURL(v string) {
	o.URL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingLogshuttleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.URL != nil {
		toSerialize["url"] = o.URL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingLogshuttleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingLogshuttleAllOf := _LoggingLogshuttleAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingLogshuttleAllOf); err == nil {
		*o = LoggingLogshuttleAllOf(varLoggingLogshuttleAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingLogshuttleAllOf is a helper abstraction for handling nullable logginglogshuttleallof types. 
type NullableLoggingLogshuttleAllOf struct {
	value *LoggingLogshuttleAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingLogshuttleAllOf) Get() *LoggingLogshuttleAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingLogshuttleAllOf) Set(val *LoggingLogshuttleAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingLogshuttleAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingLogshuttleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingLogshuttleAllOf returns a pointer to a new instance of NullableLoggingLogshuttleAllOf.
func NewNullableLoggingLogshuttleAllOf(val *LoggingLogshuttleAllOf) *NullableLoggingLogshuttleAllOf {
	return &NullableLoggingLogshuttleAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingLogshuttleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingLogshuttleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
