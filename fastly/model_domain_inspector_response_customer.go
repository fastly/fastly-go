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

// DomainInspectorResponseCustomer struct for DomainInspectorResponseCustomer
type DomainInspectorResponseCustomer struct {
	Customer             *BotManagementResponseCustomerCustomer `json:"customer,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainInspectorResponseCustomer DomainInspectorResponseCustomer

// NewDomainInspectorResponseCustomer instantiates a new DomainInspectorResponseCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainInspectorResponseCustomer() *DomainInspectorResponseCustomer {
	this := DomainInspectorResponseCustomer{}
	return &this
}

// NewDomainInspectorResponseCustomerWithDefaults instantiates a new DomainInspectorResponseCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainInspectorResponseCustomerWithDefaults() *DomainInspectorResponseCustomer {
	this := DomainInspectorResponseCustomer{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *DomainInspectorResponseCustomer) GetCustomer() BotManagementResponseCustomerCustomer {
	if o == nil || o.Customer == nil {
		var ret BotManagementResponseCustomerCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorResponseCustomer) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *DomainInspectorResponseCustomer) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given BotManagementResponseCustomerCustomer and assigns it to the Customer field.
func (o *DomainInspectorResponseCustomer) SetCustomer(v BotManagementResponseCustomerCustomer) {
	o.Customer = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainInspectorResponseCustomer) MarshalJSON() ([]byte, error) {
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
func (o *DomainInspectorResponseCustomer) UnmarshalJSON(bytes []byte) (err error) {
	varDomainInspectorResponseCustomer := _DomainInspectorResponseCustomer{}

	if err = json.Unmarshal(bytes, &varDomainInspectorResponseCustomer); err == nil {
		*o = DomainInspectorResponseCustomer(varDomainInspectorResponseCustomer)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDomainInspectorResponseCustomer is a helper abstraction for handling nullable domaininspectorresponsecustomer types.
type NullableDomainInspectorResponseCustomer struct {
	value *DomainInspectorResponseCustomer
	isSet bool
}

// Get returns the value.
func (v NullableDomainInspectorResponseCustomer) Get() *DomainInspectorResponseCustomer {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainInspectorResponseCustomer) Set(val *DomainInspectorResponseCustomer) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainInspectorResponseCustomer) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainInspectorResponseCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainInspectorResponseCustomer returns a pointer to a new instance of NullableDomainInspectorResponseCustomer.
func NewNullableDomainInspectorResponseCustomer(val *DomainInspectorResponseCustomer) *NullableDomainInspectorResponseCustomer {
	return &NullableDomainInspectorResponseCustomer{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainInspectorResponseCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainInspectorResponseCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
