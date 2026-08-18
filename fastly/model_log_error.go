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

// LogError struct for LogError
type LogError struct {
	// Sequence number for ordering messages.
	SequenceNumber *int32 `json:"sequence_number,omitempty"`
	// Timestamp of the error in microseconds.
	ErrorTimeUs *int64 `json:"error_time_us,omitempty"`
	// The stream type, always 'logging_error' for logging endpoint errors.
	Stream *string `json:"stream,omitempty"`
	// User-friendly error message.
	Message *string `json:"message,omitempty"`
	// Name of the logging endpoint that generated the error.
	Endpoint *string `json:"endpoint,omitempty"`
	// Additional error details as a JSON string.
	Details              *string `json:"details,omitempty"`
	AdditionalProperties map[string]any
}

type _LogError LogError

// NewLogError instantiates a new LogError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogError() *LogError {
	this := LogError{}
	return &this
}

// NewLogErrorWithDefaults instantiates a new LogError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogErrorWithDefaults() *LogError {
	this := LogError{}
	return &this
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *LogError) GetSequenceNumber() int32 {
	if o == nil || o.SequenceNumber == nil {
		var ret int32
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogError) GetSequenceNumberOk() (*int32, bool) {
	if o == nil || o.SequenceNumber == nil {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *LogError) HasSequenceNumber() bool {
	if o != nil && o.SequenceNumber != nil {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int32 and assigns it to the SequenceNumber field.
func (o *LogError) SetSequenceNumber(v int32) {
	o.SequenceNumber = &v
}

// GetErrorTimeUs returns the ErrorTimeUs field value if set, zero value otherwise.
func (o *LogError) GetErrorTimeUs() int64 {
	if o == nil || o.ErrorTimeUs == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTimeUs
}

// GetErrorTimeUsOk returns a tuple with the ErrorTimeUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogError) GetErrorTimeUsOk() (*int64, bool) {
	if o == nil || o.ErrorTimeUs == nil {
		return nil, false
	}
	return o.ErrorTimeUs, true
}

// HasErrorTimeUs returns a boolean if a field has been set.
func (o *LogError) HasErrorTimeUs() bool {
	if o != nil && o.ErrorTimeUs != nil {
		return true
	}

	return false
}

// SetErrorTimeUs gets a reference to the given int64 and assigns it to the ErrorTimeUs field.
func (o *LogError) SetErrorTimeUs(v int64) {
	o.ErrorTimeUs = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *LogError) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogError) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *LogError) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *LogError) SetStream(v string) {
	o.Stream = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LogError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LogError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LogError) SetMessage(v string) {
	o.Message = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *LogError) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogError) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *LogError) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *LogError) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *LogError) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogError) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *LogError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *LogError) SetDetails(v string) {
	o.Details = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.SequenceNumber != nil {
		toSerialize["sequence_number"] = o.SequenceNumber
	}
	if o.ErrorTimeUs != nil {
		toSerialize["error_time_us"] = o.ErrorTimeUs
	}
	if o.Stream != nil {
		toSerialize["stream"] = o.Stream
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogError) UnmarshalJSON(bytes []byte) (err error) {
	varLogError := _LogError{}

	if err = json.Unmarshal(bytes, &varLogError); err == nil {
		*o = LogError(varLogError)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sequence_number")
		delete(additionalProperties, "error_time_us")
		delete(additionalProperties, "stream")
		delete(additionalProperties, "message")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogError is a helper abstraction for handling nullable logerror types.
type NullableLogError struct {
	value *LogError
	isSet bool
}

// Get returns the value.
func (v NullableLogError) Get() *LogError {
	return v.value
}

// Set modifies the value.
func (v *NullableLogError) Set(val *LogError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogError returns a pointer to a new instance of NullableLogError.
func NewNullableLogError(val *LogError) *NullableLogError {
	return &NullableLogError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
