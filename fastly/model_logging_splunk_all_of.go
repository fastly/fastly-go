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

// LoggingSplunkAllOf struct for LoggingSplunkAllOf
type LoggingSplunkAllOf struct {
	// The URL to post logs to.
	URL *string `json:"url,omitempty"`
	// A Splunk token for use in posting logs over HTTP to your collector.
	Token *string `json:"token,omitempty"`
	UseTLS *LoggingUseTLS `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSplunkAllOf LoggingSplunkAllOf

// NewLoggingSplunkAllOf instantiates a new LoggingSplunkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSplunkAllOf() *LoggingSplunkAllOf {
	this := LoggingSplunkAllOf{}
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingSplunkAllOfWithDefaults instantiates a new LoggingSplunkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSplunkAllOfWithDefaults() *LoggingSplunkAllOf {
	this := LoggingSplunkAllOf{}
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingSplunkAllOf) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkAllOf) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingSplunkAllOf) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingSplunkAllOf) SetURL(v string) {
	o.URL = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingSplunkAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingSplunkAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingSplunkAllOf) SetToken(v string) {
	o.Token = &v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingSplunkAllOf) GetUseTLS() LoggingUseTLS {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLS
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkAllOf) GetUseTLSOk() (*LoggingUseTLS, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingSplunkAllOf) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLS and assigns it to the UseTLS field.
func (o *LoggingSplunkAllOf) SetUseTLS(v LoggingUseTLS) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSplunkAllOf) MarshalJSON() ([]byte, error) {
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
func (o *LoggingSplunkAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSplunkAllOf := _LoggingSplunkAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingSplunkAllOf); err == nil {
		*o = LoggingSplunkAllOf(varLoggingSplunkAllOf)
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

// NullableLoggingSplunkAllOf is a helper abstraction for handling nullable loggingsplunkallof types. 
type NullableLoggingSplunkAllOf struct {
	value *LoggingSplunkAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSplunkAllOf) Get() *LoggingSplunkAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSplunkAllOf) Set(val *LoggingSplunkAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSplunkAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSplunkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSplunkAllOf returns a pointer to a new instance of NullableLoggingSplunkAllOf.
func NewNullableLoggingSplunkAllOf(val *LoggingSplunkAllOf) *NullableLoggingSplunkAllOf {
	return &NullableLoggingSplunkAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSplunkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSplunkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
