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

// ApiDiscoveryResponseCustomer struct for ApiDiscoveryResponseCustomer
type ApiDiscoveryResponseCustomer struct {
	Customer             *ApiDiscoveryResponseCustomerCustomer `json:"customer,omitempty"`
	AdditionalProperties map[string]any
}

type _ApiDiscoveryResponseCustomer ApiDiscoveryResponseCustomer

// NewApiDiscoveryResponseCustomer instantiates a new ApiDiscoveryResponseCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDiscoveryResponseCustomer() *ApiDiscoveryResponseCustomer {
	this := ApiDiscoveryResponseCustomer{}
	return &this
}

// NewApiDiscoveryResponseCustomerWithDefaults instantiates a new ApiDiscoveryResponseCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDiscoveryResponseCustomerWithDefaults() *ApiDiscoveryResponseCustomer {
	this := ApiDiscoveryResponseCustomer{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ApiDiscoveryResponseCustomer) GetCustomer() ApiDiscoveryResponseCustomerCustomer {
	if o == nil || o.Customer == nil {
		var ret ApiDiscoveryResponseCustomerCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDiscoveryResponseCustomer) GetCustomerOk() (*ApiDiscoveryResponseCustomerCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ApiDiscoveryResponseCustomer) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given ApiDiscoveryResponseCustomerCustomer and assigns it to the Customer field.
func (o *ApiDiscoveryResponseCustomer) SetCustomer(v ApiDiscoveryResponseCustomerCustomer) {
	o.Customer = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ApiDiscoveryResponseCustomer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ApiDiscoveryResponseCustomer) UnmarshalJSON(bytes []byte) (err error) {
	varApiDiscoveryResponseCustomer := _ApiDiscoveryResponseCustomer{}

	if err = json.Unmarshal(bytes, &varApiDiscoveryResponseCustomer); err == nil {
		*o = ApiDiscoveryResponseCustomer(varApiDiscoveryResponseCustomer)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableApiDiscoveryResponseCustomer is a helper abstraction for handling nullable apidiscoveryresponsecustomer types.
type NullableApiDiscoveryResponseCustomer struct {
	value *ApiDiscoveryResponseCustomer
	isSet bool
}

// Get returns the value.
func (v NullableApiDiscoveryResponseCustomer) Get() *ApiDiscoveryResponseCustomer {
	return v.value
}

// Set modifies the value.
func (v *NullableApiDiscoveryResponseCustomer) Set(val *ApiDiscoveryResponseCustomer) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableApiDiscoveryResponseCustomer) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableApiDiscoveryResponseCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableApiDiscoveryResponseCustomer returns a pointer to a new instance of NullableApiDiscoveryResponseCustomer.
func NewNullableApiDiscoveryResponseCustomer(val *ApiDiscoveryResponseCustomer) *NullableApiDiscoveryResponseCustomer {
	return &NullableApiDiscoveryResponseCustomer{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableApiDiscoveryResponseCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableApiDiscoveryResponseCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
