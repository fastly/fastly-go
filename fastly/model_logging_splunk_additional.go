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

// LoggingSplunkAdditional struct for LoggingSplunkAdditional
type LoggingSplunkAdditional struct {
	// The URL to post logs to.
	URL *string `json:"url,omitempty"`
	// A Splunk token for use in posting logs over HTTP to your collector.
	Token *string `json:"token,omitempty"`
	UseTLS *LoggingUseTLSString `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSplunkAdditional LoggingSplunkAdditional

// NewLoggingSplunkAdditional instantiates a new LoggingSplunkAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSplunkAdditional() *LoggingSplunkAdditional {
	this := LoggingSplunkAdditional{}
	var useTLS LoggingUseTLSString = LOGGINGUSETLSSTRING_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingSplunkAdditionalWithDefaults instantiates a new LoggingSplunkAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSplunkAdditionalWithDefaults() *LoggingSplunkAdditional {
	this := LoggingSplunkAdditional{}
	var useTLS LoggingUseTLSString = LOGGINGUSETLSSTRING_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingSplunkAdditional) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkAdditional) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingSplunkAdditional) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingSplunkAdditional) SetURL(v string) {
	o.URL = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingSplunkAdditional) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkAdditional) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingSplunkAdditional) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingSplunkAdditional) SetToken(v string) {
	o.Token = &v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingSplunkAdditional) GetUseTLS() LoggingUseTLSString {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLSString
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkAdditional) GetUseTLSOk() (*LoggingUseTLSString, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingSplunkAdditional) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLSString and assigns it to the UseTLS field.
func (o *LoggingSplunkAdditional) SetUseTLS(v LoggingUseTLSString) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSplunkAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.URL != nil {
		toSerialize["url"] = o.URL
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UseTLS != nil {
		toSerialize["use_tls"] = o.UseTLS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingSplunkAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSplunkAdditional := _LoggingSplunkAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingSplunkAdditional); err == nil {
		*o = LoggingSplunkAdditional(varLoggingSplunkAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "token")
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSplunkAdditional is a helper abstraction for handling nullable loggingsplunkadditional types. 
type NullableLoggingSplunkAdditional struct {
	value *LoggingSplunkAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSplunkAdditional) Get() *LoggingSplunkAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSplunkAdditional) Set(val *LoggingSplunkAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSplunkAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSplunkAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSplunkAdditional returns a pointer to a new instance of NullableLoggingSplunkAdditional.
func NewNullableLoggingSplunkAdditional(val *LoggingSplunkAdditional) *NullableLoggingSplunkAdditional {
	return &NullableLoggingSplunkAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSplunkAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSplunkAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
