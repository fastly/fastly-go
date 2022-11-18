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

// LoggingHerokuAllOf struct for LoggingHerokuAllOf
type LoggingHerokuAllOf struct {
	// The token to use for authentication ([https://devcenter.heroku.com/articles/add-on-partner-log-integration](https://devcenter.heroku.com/articles/add-on-partner-log-integration)).
	Token *string `json:"token,omitempty"`
	// The URL to stream logs to.
	URL *string `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingHerokuAllOf LoggingHerokuAllOf

// NewLoggingHerokuAllOf instantiates a new LoggingHerokuAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingHerokuAllOf() *LoggingHerokuAllOf {
	this := LoggingHerokuAllOf{}
	return &this
}

// NewLoggingHerokuAllOfWithDefaults instantiates a new LoggingHerokuAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingHerokuAllOfWithDefaults() *LoggingHerokuAllOf {
	this := LoggingHerokuAllOf{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingHerokuAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHerokuAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingHerokuAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingHerokuAllOf) SetToken(v string) {
	o.Token = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingHerokuAllOf) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHerokuAllOf) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingHerokuAllOf) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingHerokuAllOf) SetURL(v string) {
	o.URL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingHerokuAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
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
func (o *LoggingHerokuAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingHerokuAllOf := _LoggingHerokuAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingHerokuAllOf); err == nil {
		*o = LoggingHerokuAllOf(varLoggingHerokuAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingHerokuAllOf is a helper abstraction for handling nullable loggingherokuallof types. 
type NullableLoggingHerokuAllOf struct {
	value *LoggingHerokuAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingHerokuAllOf) Get() *LoggingHerokuAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingHerokuAllOf) Set(val *LoggingHerokuAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingHerokuAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingHerokuAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingHerokuAllOf returns a pointer to a new instance of NullableLoggingHerokuAllOf.
func NewNullableLoggingHerokuAllOf(val *LoggingHerokuAllOf) *NullableLoggingHerokuAllOf {
	return &NullableLoggingHerokuAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingHerokuAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingHerokuAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
