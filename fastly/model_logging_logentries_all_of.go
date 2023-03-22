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

// LoggingLogentriesAllOf struct for LoggingLogentriesAllOf
type LoggingLogentriesAllOf struct {
	// The port number.
	Port *int32 `json:"port,omitempty"`
	// Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)).
	Token *string `json:"token,omitempty"`
	UseTLS *LoggingUseTLS `json:"use_tls,omitempty"`
	// The region to which to stream logs.
	Region *string `json:"region,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingLogentriesAllOf LoggingLogentriesAllOf

// NewLoggingLogentriesAllOf instantiates a new LoggingLogentriesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingLogentriesAllOf() *LoggingLogentriesAllOf {
	this := LoggingLogentriesAllOf{}
	var port int32 = 20000
	this.Port = &port
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingLogentriesAllOfWithDefaults instantiates a new LoggingLogentriesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingLogentriesAllOfWithDefaults() *LoggingLogentriesAllOf {
	this := LoggingLogentriesAllOf{}
	var port int32 = 20000
	this.Port = &port
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoggingLogentriesAllOf) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingLogentriesAllOf) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoggingLogentriesAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LoggingLogentriesAllOf) SetPort(v int32) {
	o.Port = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingLogentriesAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingLogentriesAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingLogentriesAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingLogentriesAllOf) SetToken(v string) {
	o.Token = &v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingLogentriesAllOf) GetUseTLS() LoggingUseTLS {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLS
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingLogentriesAllOf) GetUseTLSOk() (*LoggingUseTLS, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingLogentriesAllOf) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLS and assigns it to the UseTLS field.
func (o *LoggingLogentriesAllOf) SetUseTLS(v LoggingUseTLS) {
	o.UseTLS = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LoggingLogentriesAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingLogentriesAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LoggingLogentriesAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *LoggingLogentriesAllOf) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingLogentriesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UseTLS != nil {
		toSerialize["use_tls"] = o.UseTLS
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingLogentriesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingLogentriesAllOf := _LoggingLogentriesAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingLogentriesAllOf); err == nil {
		*o = LoggingLogentriesAllOf(varLoggingLogentriesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "port")
		delete(additionalProperties, "token")
		delete(additionalProperties, "use_tls")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingLogentriesAllOf is a helper abstraction for handling nullable logginglogentriesallof types. 
type NullableLoggingLogentriesAllOf struct {
	value *LoggingLogentriesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingLogentriesAllOf) Get() *LoggingLogentriesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingLogentriesAllOf) Set(val *LoggingLogentriesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingLogentriesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingLogentriesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingLogentriesAllOf returns a pointer to a new instance of NullableLoggingLogentriesAllOf.
func NewNullableLoggingLogentriesAllOf(val *LoggingLogentriesAllOf) *NullableLoggingLogentriesAllOf {
	return &NullableLoggingLogentriesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingLogentriesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingLogentriesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
