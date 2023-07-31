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

// LoggingSftpResponseAllOf struct for LoggingSftpResponseAllOf
type LoggingSftpResponseAllOf struct {
	// A hostname or IPv4 address.
	Address *string `json:"address,omitempty"`
	// The port number.
	Port *string `json:"port,omitempty"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *string `json:"period,omitempty"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int32 `json:"gzip_level,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSftpResponseAllOf LoggingSftpResponseAllOf

// NewLoggingSftpResponseAllOf instantiates a new LoggingSftpResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSftpResponseAllOf() *LoggingSftpResponseAllOf {
	this := LoggingSftpResponseAllOf{}
	var port string = "22"
	this.Port = &port
	var period string = "3600"
	this.Period = &period
	var gzipLevel int32 = 0
	this.GzipLevel = &gzipLevel
	return &this
}

// NewLoggingSftpResponseAllOfWithDefaults instantiates a new LoggingSftpResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSftpResponseAllOfWithDefaults() *LoggingSftpResponseAllOf {
	this := LoggingSftpResponseAllOf{}
	var port string = "22"
	this.Port = &port
	var period string = "3600"
	this.Period = &period
	var gzipLevel int32 = 0
	this.GzipLevel = &gzipLevel
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LoggingSftpResponseAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpResponseAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LoggingSftpResponseAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *LoggingSftpResponseAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoggingSftpResponseAllOf) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpResponseAllOf) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoggingSftpResponseAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *LoggingSftpResponseAllOf) SetPort(v string) {
	o.Port = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LoggingSftpResponseAllOf) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpResponseAllOf) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LoggingSftpResponseAllOf) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *LoggingSftpResponseAllOf) SetPeriod(v string) {
	o.Period = &v
}

// GetGzipLevel returns the GzipLevel field value if set, zero value otherwise.
func (o *LoggingSftpResponseAllOf) GetGzipLevel() int32 {
	if o == nil || o.GzipLevel == nil {
		var ret int32
		return ret
	}
	return *o.GzipLevel
}

// GetGzipLevelOk returns a tuple with the GzipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSftpResponseAllOf) GetGzipLevelOk() (*int32, bool) {
	if o == nil || o.GzipLevel == nil {
		return nil, false
	}
	return o.GzipLevel, true
}

// HasGzipLevel returns a boolean if a field has been set.
func (o *LoggingSftpResponseAllOf) HasGzipLevel() bool {
	if o != nil && o.GzipLevel != nil {
		return true
	}

	return false
}

// SetGzipLevel gets a reference to the given int32 and assigns it to the GzipLevel field.
func (o *LoggingSftpResponseAllOf) SetGzipLevel(v int32) {
	o.GzipLevel = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSftpResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.GzipLevel != nil {
		toSerialize["gzip_level"] = o.GzipLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingSftpResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSftpResponseAllOf := _LoggingSftpResponseAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingSftpResponseAllOf); err == nil {
		*o = LoggingSftpResponseAllOf(varLoggingSftpResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "port")
		delete(additionalProperties, "period")
		delete(additionalProperties, "gzip_level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSftpResponseAllOf is a helper abstraction for handling nullable loggingsftpresponseallof types. 
type NullableLoggingSftpResponseAllOf struct {
	value *LoggingSftpResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSftpResponseAllOf) Get() *LoggingSftpResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSftpResponseAllOf) Set(val *LoggingSftpResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSftpResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSftpResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSftpResponseAllOf returns a pointer to a new instance of NullableLoggingSftpResponseAllOf.
func NewNullableLoggingSftpResponseAllOf(val *LoggingSftpResponseAllOf) *NullableLoggingSftpResponseAllOf {
	return &NullableLoggingSftpResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSftpResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSftpResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
