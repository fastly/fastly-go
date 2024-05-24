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

// LoggingRequestCapsCommon struct for LoggingRequestCapsCommon
type LoggingRequestCapsCommon struct {
	// The maximum number of logs sent in one request. Defaults `0` for unbounded.
	RequestMaxEntries *int32 `json:"request_max_entries,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` for unbounded.
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingRequestCapsCommon LoggingRequestCapsCommon

// NewLoggingRequestCapsCommon instantiates a new LoggingRequestCapsCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingRequestCapsCommon() *LoggingRequestCapsCommon {
	this := LoggingRequestCapsCommon{}
	var requestMaxEntries int32 = 0
	this.RequestMaxEntries = &requestMaxEntries
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	return &this
}

// NewLoggingRequestCapsCommonWithDefaults instantiates a new LoggingRequestCapsCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingRequestCapsCommonWithDefaults() *LoggingRequestCapsCommon {
	this := LoggingRequestCapsCommon{}
	var requestMaxEntries int32 = 0
	this.RequestMaxEntries = &requestMaxEntries
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	return &this
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingRequestCapsCommon) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingRequestCapsCommon) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingRequestCapsCommon) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingRequestCapsCommon) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingRequestCapsCommon) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingRequestCapsCommon) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingRequestCapsCommon) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingRequestCapsCommon) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingRequestCapsCommon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.RequestMaxEntries != nil {
		toSerialize["request_max_entries"] = o.RequestMaxEntries
	}
	if o.RequestMaxBytes != nil {
		toSerialize["request_max_bytes"] = o.RequestMaxBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingRequestCapsCommon) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingRequestCapsCommon := _LoggingRequestCapsCommon{}

	if err = json.Unmarshal(bytes, &varLoggingRequestCapsCommon); err == nil {
		*o = LoggingRequestCapsCommon(varLoggingRequestCapsCommon)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_max_entries")
		delete(additionalProperties, "request_max_bytes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingRequestCapsCommon is a helper abstraction for handling nullable loggingrequestcapscommon types. 
type NullableLoggingRequestCapsCommon struct {
	value *LoggingRequestCapsCommon
	isSet bool
}

// Get returns the value.
func (v NullableLoggingRequestCapsCommon) Get() *LoggingRequestCapsCommon {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingRequestCapsCommon) Set(val *LoggingRequestCapsCommon) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingRequestCapsCommon) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingRequestCapsCommon) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingRequestCapsCommon returns a pointer to a new instance of NullableLoggingRequestCapsCommon.
func NewNullableLoggingRequestCapsCommon(val *LoggingRequestCapsCommon) *NullableLoggingRequestCapsCommon {
	return &NullableLoggingRequestCapsCommon{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingRequestCapsCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingRequestCapsCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
