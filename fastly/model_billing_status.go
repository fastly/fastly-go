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
	"time"
)

// BillingStatus struct for BillingStatus
type BillingStatus struct {
	// What the current status of this invoice can be.
	Status *string `json:"status,omitempty"`
	// Date and time in ISO 8601 format.
	// Deprecated
	SentAt               NullableTime `json:"sent_at,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingStatus BillingStatus

// NewBillingStatus instantiates a new BillingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingStatus() *BillingStatus {
	this := BillingStatus{}
	return &this
}

// NewBillingStatusWithDefaults instantiates a new BillingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingStatusWithDefaults() *BillingStatus {
	this := BillingStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BillingStatus) SetStatus(v string) {
	o.Status = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *BillingStatus) GetSentAt() time.Time {
	if o == nil || o.SentAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *BillingStatus) GetSentAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// HasSentAt returns a boolean if a field has been set.
func (o *BillingStatus) HasSentAt() bool {
	if o != nil && o.SentAt.IsSet() {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given NullableTime and assigns it to the SentAt field.
// Deprecated
func (o *BillingStatus) SetSentAt(v time.Time) {
	o.SentAt.Set(&v)
}

// SetSentAtNil sets the value for SentAt to be an explicit nil
func (o *BillingStatus) SetSentAtNil() {
	o.SentAt.Set(nil)
}

// UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
func (o *BillingStatus) UnsetSentAt() {
	o.SentAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SentAt.IsSet() {
		toSerialize["sent_at"] = o.SentAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BillingStatus) UnmarshalJSON(bytes []byte) (err error) {
	varBillingStatus := _BillingStatus{}

	if err = json.Unmarshal(bytes, &varBillingStatus); err == nil {
		*o = BillingStatus(varBillingStatus)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "sent_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingStatus is a helper abstraction for handling nullable billingstatus types.
type NullableBillingStatus struct {
	value *BillingStatus
	isSet bool
}

// Get returns the value.
func (v NullableBillingStatus) Get() *BillingStatus {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingStatus) Set(val *BillingStatus) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingStatus) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingStatus returns a pointer to a new instance of NullableBillingStatus.
func NewNullableBillingStatus(val *BillingStatus) *NullableBillingStatus {
	return &NullableBillingStatus{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBillingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
