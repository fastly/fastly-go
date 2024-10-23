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

// LoggingSyslogAdditional struct for LoggingSyslogAdditional
type LoggingSyslogAdditional struct {
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// The hostname used for the syslog endpoint.
	Hostname *string `json:"hostname,omitempty"`
	// The IPv4 address used for the syslog endpoint.
	Ipv4 NullableString `json:"ipv4,omitempty"`
	// Whether to prepend each message with a specific token.
	Token                NullableString       `json:"token,omitempty"`
	UseTLS               *LoggingUseTLSString `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSyslogAdditional LoggingSyslogAdditional

// NewLoggingSyslogAdditional instantiates a new LoggingSyslogAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSyslogAdditional() *LoggingSyslogAdditional {
	this := LoggingSyslogAdditional{}
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var token string = "null"
	this.Token = *NewNullableString(&token)
	var useTLS LoggingUseTLSString = LOGGINGUSETLSSTRING_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingSyslogAdditionalWithDefaults instantiates a new LoggingSyslogAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSyslogAdditionalWithDefaults() *LoggingSyslogAdditional {
	this := LoggingSyslogAdditional{}
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var token string = "null"
	this.Token = *NewNullableString(&token)
	var useTLS LoggingUseTLSString = LOGGINGUSETLSSTRING_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingSyslogAdditional) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslogAdditional) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingSyslogAdditional) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingSyslogAdditional) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *LoggingSyslogAdditional) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslogAdditional) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *LoggingSyslogAdditional) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *LoggingSyslogAdditional) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslogAdditional) GetIpv4() string {
	if o == nil || o.Ipv4.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ipv4.Get()
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslogAdditional) GetIpv4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv4.Get(), o.Ipv4.IsSet()
}

// HasIpv4 returns a boolean if a field has been set.
func (o *LoggingSyslogAdditional) HasIpv4() bool {
	if o != nil && o.Ipv4.IsSet() {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NullableString and assigns it to the Ipv4 field.
func (o *LoggingSyslogAdditional) SetIpv4(v string) {
	o.Ipv4.Set(&v)
}

// SetIpv4Nil sets the value for Ipv4 to be an explicit nil
func (o *LoggingSyslogAdditional) SetIpv4Nil() {
	o.Ipv4.Set(nil)
}

// UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
func (o *LoggingSyslogAdditional) UnsetIpv4() {
	o.Ipv4.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslogAdditional) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslogAdditional) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingSyslogAdditional) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *LoggingSyslogAdditional) SetToken(v string) {
	o.Token.Set(&v)
}

// SetTokenNil sets the value for Token to be an explicit nil
func (o *LoggingSyslogAdditional) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *LoggingSyslogAdditional) UnsetToken() {
	o.Token.Unset()
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingSyslogAdditional) GetUseTLS() LoggingUseTLSString {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLSString
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslogAdditional) GetUseTLSOk() (*LoggingUseTLSString, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingSyslogAdditional) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLSString and assigns it to the UseTLS field.
func (o *LoggingSyslogAdditional) SetUseTLS(v LoggingUseTLSString) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSyslogAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.MessageType != nil {
		toSerialize["message_type"] = o.MessageType
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Ipv4.IsSet() {
		toSerialize["ipv4"] = o.Ipv4.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
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
func (o *LoggingSyslogAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSyslogAdditional := _LoggingSyslogAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingSyslogAdditional); err == nil {
		*o = LoggingSyslogAdditional(varLoggingSyslogAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "token")
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSyslogAdditional is a helper abstraction for handling nullable loggingsyslogadditional types.
type NullableLoggingSyslogAdditional struct {
	value *LoggingSyslogAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSyslogAdditional) Get() *LoggingSyslogAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSyslogAdditional) Set(val *LoggingSyslogAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSyslogAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSyslogAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSyslogAdditional returns a pointer to a new instance of NullableLoggingSyslogAdditional.
func NewNullableLoggingSyslogAdditional(val *LoggingSyslogAdditional) *NullableLoggingSyslogAdditional {
	return &NullableLoggingSyslogAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSyslogAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingSyslogAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
