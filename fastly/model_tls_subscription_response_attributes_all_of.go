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

// TlsSubscriptionResponseAttributesAllOf struct for TlsSubscriptionResponseAttributesAllOf
type TlsSubscriptionResponseAttributesAllOf struct {
	// The current state of your subscription.
	State *string `json:"state,omitempty"`
	// Subscription has an active order
	HasActiveOrder       *bool `json:"has_active_order,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsSubscriptionResponseAttributesAllOf TlsSubscriptionResponseAttributesAllOf

// NewTlsSubscriptionResponseAttributesAllOf instantiates a new TlsSubscriptionResponseAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsSubscriptionResponseAttributesAllOf() *TlsSubscriptionResponseAttributesAllOf {
	this := TlsSubscriptionResponseAttributesAllOf{}
	return &this
}

// NewTlsSubscriptionResponseAttributesAllOfWithDefaults instantiates a new TlsSubscriptionResponseAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsSubscriptionResponseAttributesAllOfWithDefaults() *TlsSubscriptionResponseAttributesAllOf {
	this := TlsSubscriptionResponseAttributesAllOf{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TlsSubscriptionResponseAttributesAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionResponseAttributesAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TlsSubscriptionResponseAttributesAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TlsSubscriptionResponseAttributesAllOf) SetState(v string) {
	o.State = &v
}

// GetHasActiveOrder returns the HasActiveOrder field value if set, zero value otherwise.
func (o *TlsSubscriptionResponseAttributesAllOf) GetHasActiveOrder() bool {
	if o == nil || o.HasActiveOrder == nil {
		var ret bool
		return ret
	}
	return *o.HasActiveOrder
}

// GetHasActiveOrderOk returns a tuple with the HasActiveOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionResponseAttributesAllOf) GetHasActiveOrderOk() (*bool, bool) {
	if o == nil || o.HasActiveOrder == nil {
		return nil, false
	}
	return o.HasActiveOrder, true
}

// HasHasActiveOrder returns a boolean if a field has been set.
func (o *TlsSubscriptionResponseAttributesAllOf) HasHasActiveOrder() bool {
	if o != nil && o.HasActiveOrder != nil {
		return true
	}

	return false
}

// SetHasActiveOrder gets a reference to the given bool and assigns it to the HasActiveOrder field.
func (o *TlsSubscriptionResponseAttributesAllOf) SetHasActiveOrder(v bool) {
	o.HasActiveOrder = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsSubscriptionResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.HasActiveOrder != nil {
		toSerialize["has_active_order"] = o.HasActiveOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsSubscriptionResponseAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTlsSubscriptionResponseAttributesAllOf := _TlsSubscriptionResponseAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varTlsSubscriptionResponseAttributesAllOf); err == nil {
		*o = TlsSubscriptionResponseAttributesAllOf(varTlsSubscriptionResponseAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "has_active_order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsSubscriptionResponseAttributesAllOf is a helper abstraction for handling nullable tlssubscriptionresponseattributesallof types.
type NullableTlsSubscriptionResponseAttributesAllOf struct {
	value *TlsSubscriptionResponseAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTlsSubscriptionResponseAttributesAllOf) Get() *TlsSubscriptionResponseAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsSubscriptionResponseAttributesAllOf) Set(val *TlsSubscriptionResponseAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsSubscriptionResponseAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsSubscriptionResponseAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsSubscriptionResponseAttributesAllOf returns a pointer to a new instance of NullableTlsSubscriptionResponseAttributesAllOf.
func NewNullableTlsSubscriptionResponseAttributesAllOf(val *TlsSubscriptionResponseAttributesAllOf) *NullableTlsSubscriptionResponseAttributesAllOf {
	return &NullableTlsSubscriptionResponseAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsSubscriptionResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsSubscriptionResponseAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
