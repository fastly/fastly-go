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

// LoggingSumologicAdditional struct for LoggingSumologicAdditional
type LoggingSumologicAdditional struct {
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// The URL to post logs to.
	URL *string `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSumologicAdditional LoggingSumologicAdditional

// NewLoggingSumologicAdditional instantiates a new LoggingSumologicAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSumologicAdditional() *LoggingSumologicAdditional {
	this := LoggingSumologicAdditional{}
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	return &this
}

// NewLoggingSumologicAdditionalWithDefaults instantiates a new LoggingSumologicAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSumologicAdditionalWithDefaults() *LoggingSumologicAdditional {
	this := LoggingSumologicAdditional{}
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingSumologicAdditional) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSumologicAdditional) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingSumologicAdditional) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingSumologicAdditional) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingSumologicAdditional) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSumologicAdditional) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingSumologicAdditional) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingSumologicAdditional) SetURL(v string) {
	o.URL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSumologicAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.MessageType != nil {
		toSerialize["message_type"] = o.MessageType
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
func (o *LoggingSumologicAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSumologicAdditional := _LoggingSumologicAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingSumologicAdditional); err == nil {
		*o = LoggingSumologicAdditional(varLoggingSumologicAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSumologicAdditional is a helper abstraction for handling nullable loggingsumologicadditional types. 
type NullableLoggingSumologicAdditional struct {
	value *LoggingSumologicAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSumologicAdditional) Get() *LoggingSumologicAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSumologicAdditional) Set(val *LoggingSumologicAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSumologicAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSumologicAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSumologicAdditional returns a pointer to a new instance of NullableLoggingSumologicAdditional.
func NewNullableLoggingSumologicAdditional(val *LoggingSumologicAdditional) *NullableLoggingSumologicAdditional {
	return &NullableLoggingSumologicAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSumologicAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSumologicAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
