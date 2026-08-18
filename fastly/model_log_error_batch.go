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

// LogErrorBatch struct for LogErrorBatch
type LogErrorBatch struct {
	// Unique identifier for this batch of logs.
	BatchId              *string    `json:"batch_id,omitempty"`
	Logs                 []LogError `json:"logs,omitempty"`
	AdditionalProperties map[string]any
}

type _LogErrorBatch LogErrorBatch

// NewLogErrorBatch instantiates a new LogErrorBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogErrorBatch() *LogErrorBatch {
	this := LogErrorBatch{}
	return &this
}

// NewLogErrorBatchWithDefaults instantiates a new LogErrorBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogErrorBatchWithDefaults() *LogErrorBatch {
	this := LogErrorBatch{}
	return &this
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *LogErrorBatch) GetBatchId() string {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogErrorBatch) GetBatchIdOk() (*string, bool) {
	if o == nil || o.BatchId == nil {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *LogErrorBatch) HasBatchId() bool {
	if o != nil && o.BatchId != nil {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *LogErrorBatch) SetBatchId(v string) {
	o.BatchId = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *LogErrorBatch) GetLogs() []LogError {
	if o == nil || o.Logs == nil {
		var ret []LogError
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogErrorBatch) GetLogsOk() ([]LogError, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *LogErrorBatch) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []LogError and assigns it to the Logs field.
func (o *LogErrorBatch) SetLogs(v []LogError) {
	o.Logs = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogErrorBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BatchId != nil {
		toSerialize["batch_id"] = o.BatchId
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogErrorBatch) UnmarshalJSON(bytes []byte) (err error) {
	varLogErrorBatch := _LogErrorBatch{}

	if err = json.Unmarshal(bytes, &varLogErrorBatch); err == nil {
		*o = LogErrorBatch(varLogErrorBatch)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "batch_id")
		delete(additionalProperties, "logs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogErrorBatch is a helper abstraction for handling nullable logerrorbatch types.
type NullableLogErrorBatch struct {
	value *LogErrorBatch
	isSet bool
}

// Get returns the value.
func (v NullableLogErrorBatch) Get() *LogErrorBatch {
	return v.value
}

// Set modifies the value.
func (v *NullableLogErrorBatch) Set(val *LogErrorBatch) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogErrorBatch) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogErrorBatch) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogErrorBatch returns a pointer to a new instance of NullableLogErrorBatch.
func NewNullableLogErrorBatch(val *LogErrorBatch) *NullableLogErrorBatch {
	return &NullableLogErrorBatch{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogErrorBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogErrorBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
