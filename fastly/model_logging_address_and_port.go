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

// LoggingAddressAndPort struct for LoggingAddressAndPort
type LoggingAddressAndPort struct {
	// A hostname or IPv4 address.
	Address *string `json:"address,omitempty"`
	// The port number.
	Port *int32 `json:"port,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingAddressAndPort LoggingAddressAndPort

// NewLoggingAddressAndPort instantiates a new LoggingAddressAndPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingAddressAndPort() *LoggingAddressAndPort {
	this := LoggingAddressAndPort{}
	var port int32 = 514
	this.Port = &port
	return &this
}

// NewLoggingAddressAndPortWithDefaults instantiates a new LoggingAddressAndPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingAddressAndPortWithDefaults() *LoggingAddressAndPort {
	this := LoggingAddressAndPort{}
	var port int32 = 514
	this.Port = &port
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LoggingAddressAndPort) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAddressAndPort) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LoggingAddressAndPort) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *LoggingAddressAndPort) SetAddress(v string) {
	o.Address = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoggingAddressAndPort) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAddressAndPort) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoggingAddressAndPort) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LoggingAddressAndPort) SetPort(v int32) {
	o.Port = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingAddressAndPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingAddressAndPort) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingAddressAndPort := _LoggingAddressAndPort{}

	if err = json.Unmarshal(bytes, &varLoggingAddressAndPort); err == nil {
		*o = LoggingAddressAndPort(varLoggingAddressAndPort)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingAddressAndPort is a helper abstraction for handling nullable loggingaddressandport types. 
type NullableLoggingAddressAndPort struct {
	value *LoggingAddressAndPort
	isSet bool
}

// Get returns the value.
func (v NullableLoggingAddressAndPort) Get() *LoggingAddressAndPort {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingAddressAndPort) Set(val *LoggingAddressAndPort) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingAddressAndPort) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingAddressAndPort) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingAddressAndPort returns a pointer to a new instance of NullableLoggingAddressAndPort.
func NewNullableLoggingAddressAndPort(val *LoggingAddressAndPort) *NullableLoggingAddressAndPort {
	return &NullableLoggingAddressAndPort{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingAddressAndPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingAddressAndPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
