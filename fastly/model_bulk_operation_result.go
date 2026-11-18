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

// BulkOperationResult struct for BulkOperationResult
type BulkOperationResult struct {
	// The operation ID.
	ID *string `json:"id,omitempty"`
	// HTTP status code for this operation.
	StatusCode *int32 `json:"status_code,omitempty"`
	// Error reason if the operation failed.
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkOperationResult BulkOperationResult

// NewBulkOperationResult instantiates a new BulkOperationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkOperationResult() *BulkOperationResult {
	this := BulkOperationResult{}
	return &this
}

// NewBulkOperationResultWithDefaults instantiates a new BulkOperationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkOperationResultWithDefaults() *BulkOperationResult {
	this := BulkOperationResult{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *BulkOperationResult) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkOperationResult) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *BulkOperationResult) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *BulkOperationResult) SetID(v string) {
	o.ID = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *BulkOperationResult) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkOperationResult) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *BulkOperationResult) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *BulkOperationResult) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BulkOperationResult) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkOperationResult) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BulkOperationResult) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BulkOperationResult) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkOperationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BulkOperationResult) UnmarshalJSON(bytes []byte) (err error) {
	varBulkOperationResult := _BulkOperationResult{}

	if err = json.Unmarshal(bytes, &varBulkOperationResult); err == nil {
		*o = BulkOperationResult(varBulkOperationResult)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status_code")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkOperationResult is a helper abstraction for handling nullable bulkoperationresult types.
type NullableBulkOperationResult struct {
	value *BulkOperationResult
	isSet bool
}

// Get returns the value.
func (v NullableBulkOperationResult) Get() *BulkOperationResult {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkOperationResult) Set(val *BulkOperationResult) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkOperationResult) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkOperationResult) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkOperationResult returns a pointer to a new instance of NullableBulkOperationResult.
func NewNullableBulkOperationResult(val *BulkOperationResult) *NullableBulkOperationResult {
	return &NullableBulkOperationResult{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkOperationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBulkOperationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
