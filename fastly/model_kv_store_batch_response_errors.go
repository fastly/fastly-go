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

// KvStoreBatchResponseErrors struct for KvStoreBatchResponseErrors
type KvStoreBatchResponseErrors struct {
	// The key that the error corresponds to. This field will be empty if the object or one of its fields was not parseable.
	Key *string `json:"key,omitempty"`
	// The line number of the batch request body on which the error occurred (starting from 0 for the first line).
	Index *int32 `json:"index,omitempty"`
	// The HTTP response code for the item, or a 400 if the item's operation was not completed.
	Code *string `json:"code,omitempty"`
	// A descriptor of this particular item's error.
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]any
}

type _KvStoreBatchResponseErrors KvStoreBatchResponseErrors

// NewKvStoreBatchResponseErrors instantiates a new KvStoreBatchResponseErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvStoreBatchResponseErrors() *KvStoreBatchResponseErrors {
	this := KvStoreBatchResponseErrors{}
	return &this
}

// NewKvStoreBatchResponseErrorsWithDefaults instantiates a new KvStoreBatchResponseErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvStoreBatchResponseErrorsWithDefaults() *KvStoreBatchResponseErrors {
	this := KvStoreBatchResponseErrors{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KvStoreBatchResponseErrors) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreBatchResponseErrors) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KvStoreBatchResponseErrors) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KvStoreBatchResponseErrors) SetKey(v string) {
	o.Key = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *KvStoreBatchResponseErrors) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreBatchResponseErrors) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *KvStoreBatchResponseErrors) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *KvStoreBatchResponseErrors) SetIndex(v int32) {
	o.Index = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *KvStoreBatchResponseErrors) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreBatchResponseErrors) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *KvStoreBatchResponseErrors) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *KvStoreBatchResponseErrors) SetCode(v string) {
	o.Code = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *KvStoreBatchResponseErrors) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreBatchResponseErrors) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *KvStoreBatchResponseErrors) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *KvStoreBatchResponseErrors) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o KvStoreBatchResponseErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
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
func (o *KvStoreBatchResponseErrors) UnmarshalJSON(bytes []byte) (err error) {
	varKvStoreBatchResponseErrors := _KvStoreBatchResponseErrors{}

	if err = json.Unmarshal(bytes, &varKvStoreBatchResponseErrors); err == nil {
		*o = KvStoreBatchResponseErrors(varKvStoreBatchResponseErrors)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "index")
		delete(additionalProperties, "code")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableKvStoreBatchResponseErrors is a helper abstraction for handling nullable kvstorebatchresponseerrors types.
type NullableKvStoreBatchResponseErrors struct {
	value *KvStoreBatchResponseErrors
	isSet bool
}

// Get returns the value.
func (v NullableKvStoreBatchResponseErrors) Get() *KvStoreBatchResponseErrors {
	return v.value
}

// Set modifies the value.
func (v *NullableKvStoreBatchResponseErrors) Set(val *KvStoreBatchResponseErrors) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableKvStoreBatchResponseErrors) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableKvStoreBatchResponseErrors) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableKvStoreBatchResponseErrors returns a pointer to a new instance of NullableKvStoreBatchResponseErrors.
func NewNullableKvStoreBatchResponseErrors(val *KvStoreBatchResponseErrors) *NullableKvStoreBatchResponseErrors {
	return &NullableKvStoreBatchResponseErrors{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableKvStoreBatchResponseErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableKvStoreBatchResponseErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
