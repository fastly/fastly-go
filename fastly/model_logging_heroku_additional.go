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

// LoggingHerokuAdditional struct for LoggingHerokuAdditional
type LoggingHerokuAdditional struct {
	// The token to use for authentication ([https://devcenter.heroku.com/articles/add-on-partner-log-integration](https://devcenter.heroku.com/articles/add-on-partner-log-integration)).
	Token *string `json:"token,omitempty"`
	// The URL to stream logs to.
	URL *string `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingHerokuAdditional LoggingHerokuAdditional

// NewLoggingHerokuAdditional instantiates a new LoggingHerokuAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingHerokuAdditional() *LoggingHerokuAdditional {
	this := LoggingHerokuAdditional{}
	return &this
}

// NewLoggingHerokuAdditionalWithDefaults instantiates a new LoggingHerokuAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingHerokuAdditionalWithDefaults() *LoggingHerokuAdditional {
	this := LoggingHerokuAdditional{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingHerokuAdditional) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHerokuAdditional) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingHerokuAdditional) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingHerokuAdditional) SetToken(v string) {
	o.Token = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingHerokuAdditional) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHerokuAdditional) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingHerokuAdditional) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingHerokuAdditional) SetURL(v string) {
	o.URL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingHerokuAdditional) MarshalJSON() ([]byte, error) {
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
func (o *LoggingHerokuAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingHerokuAdditional := _LoggingHerokuAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingHerokuAdditional); err == nil {
		*o = LoggingHerokuAdditional(varLoggingHerokuAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingHerokuAdditional is a helper abstraction for handling nullable loggingherokuadditional types. 
type NullableLoggingHerokuAdditional struct {
	value *LoggingHerokuAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingHerokuAdditional) Get() *LoggingHerokuAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingHerokuAdditional) Set(val *LoggingHerokuAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingHerokuAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingHerokuAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingHerokuAdditional returns a pointer to a new instance of NullableLoggingHerokuAdditional.
func NewNullableLoggingHerokuAdditional(val *LoggingHerokuAdditional) *NullableLoggingHerokuAdditional {
	return &NullableLoggingHerokuAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingHerokuAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingHerokuAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
