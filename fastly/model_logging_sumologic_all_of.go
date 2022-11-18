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

// LoggingSumologicAllOf struct for LoggingSumologicAllOf
type LoggingSumologicAllOf struct {
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// The URL to post logs to.
	URL *string `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSumologicAllOf LoggingSumologicAllOf

// NewLoggingSumologicAllOf instantiates a new LoggingSumologicAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSumologicAllOf() *LoggingSumologicAllOf {
	this := LoggingSumologicAllOf{}
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	return &this
}

// NewLoggingSumologicAllOfWithDefaults instantiates a new LoggingSumologicAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSumologicAllOfWithDefaults() *LoggingSumologicAllOf {
	this := LoggingSumologicAllOf{}
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingSumologicAllOf) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSumologicAllOf) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingSumologicAllOf) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingSumologicAllOf) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingSumologicAllOf) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSumologicAllOf) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingSumologicAllOf) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingSumologicAllOf) SetURL(v string) {
	o.URL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSumologicAllOf) MarshalJSON() ([]byte, error) {
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
func (o *LoggingSumologicAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSumologicAllOf := _LoggingSumologicAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingSumologicAllOf); err == nil {
		*o = LoggingSumologicAllOf(varLoggingSumologicAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSumologicAllOf is a helper abstraction for handling nullable loggingsumologicallof types. 
type NullableLoggingSumologicAllOf struct {
	value *LoggingSumologicAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSumologicAllOf) Get() *LoggingSumologicAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSumologicAllOf) Set(val *LoggingSumologicAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSumologicAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSumologicAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSumologicAllOf returns a pointer to a new instance of NullableLoggingSumologicAllOf.
func NewNullableLoggingSumologicAllOf(val *LoggingSumologicAllOf) *NullableLoggingSumologicAllOf {
	return &NullableLoggingSumologicAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSumologicAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSumologicAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
