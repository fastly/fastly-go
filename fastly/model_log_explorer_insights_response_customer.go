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

// LogExplorerInsightsResponseCustomer struct for LogExplorerInsightsResponseCustomer
type LogExplorerInsightsResponseCustomer struct {
	Customer             *BotManagementResponseCustomerCustomer `json:"customer,omitempty"`
	AdditionalProperties map[string]any
}

type _LogExplorerInsightsResponseCustomer LogExplorerInsightsResponseCustomer

// NewLogExplorerInsightsResponseCustomer instantiates a new LogExplorerInsightsResponseCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExplorerInsightsResponseCustomer() *LogExplorerInsightsResponseCustomer {
	this := LogExplorerInsightsResponseCustomer{}
	return &this
}

// NewLogExplorerInsightsResponseCustomerWithDefaults instantiates a new LogExplorerInsightsResponseCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExplorerInsightsResponseCustomerWithDefaults() *LogExplorerInsightsResponseCustomer {
	this := LogExplorerInsightsResponseCustomer{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseCustomer) GetCustomer() BotManagementResponseCustomerCustomer {
	if o == nil || o.Customer == nil {
		var ret BotManagementResponseCustomerCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseCustomer) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseCustomer) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given BotManagementResponseCustomerCustomer and assigns it to the Customer field.
func (o *LogExplorerInsightsResponseCustomer) SetCustomer(v BotManagementResponseCustomerCustomer) {
	o.Customer = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogExplorerInsightsResponseCustomer) MarshalJSON() ([]byte, error) {
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
func (o *LogExplorerInsightsResponseCustomer) UnmarshalJSON(bytes []byte) (err error) {
	varLogExplorerInsightsResponseCustomer := _LogExplorerInsightsResponseCustomer{}

	if err = json.Unmarshal(bytes, &varLogExplorerInsightsResponseCustomer); err == nil {
		*o = LogExplorerInsightsResponseCustomer(varLogExplorerInsightsResponseCustomer)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogExplorerInsightsResponseCustomer is a helper abstraction for handling nullable logexplorerinsightsresponsecustomer types.
type NullableLogExplorerInsightsResponseCustomer struct {
	value *LogExplorerInsightsResponseCustomer
	isSet bool
}

// Get returns the value.
func (v NullableLogExplorerInsightsResponseCustomer) Get() *LogExplorerInsightsResponseCustomer {
	return v.value
}

// Set modifies the value.
func (v *NullableLogExplorerInsightsResponseCustomer) Set(val *LogExplorerInsightsResponseCustomer) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogExplorerInsightsResponseCustomer) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogExplorerInsightsResponseCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogExplorerInsightsResponseCustomer returns a pointer to a new instance of NullableLogExplorerInsightsResponseCustomer.
func NewNullableLogExplorerInsightsResponseCustomer(val *LogExplorerInsightsResponseCustomer) *NullableLogExplorerInsightsResponseCustomer {
	return &NullableLogExplorerInsightsResponseCustomer{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogExplorerInsightsResponseCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogExplorerInsightsResponseCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
