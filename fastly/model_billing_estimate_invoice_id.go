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

// BillingEstimateInvoiceID struct for BillingEstimateInvoiceID
type BillingEstimateInvoiceID struct {
	InvoiceID            *string `json:"invoice_id,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingEstimateInvoiceID BillingEstimateInvoiceID

// NewBillingEstimateInvoiceID instantiates a new BillingEstimateInvoiceID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEstimateInvoiceID() *BillingEstimateInvoiceID {
	this := BillingEstimateInvoiceID{}
	return &this
}

// NewBillingEstimateInvoiceIDWithDefaults instantiates a new BillingEstimateInvoiceID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEstimateInvoiceIDWithDefaults() *BillingEstimateInvoiceID {
	this := BillingEstimateInvoiceID{}
	return &this
}

// GetInvoiceID returns the InvoiceID field value if set, zero value otherwise.
func (o *BillingEstimateInvoiceID) GetInvoiceID() string {
	if o == nil || o.InvoiceID == nil {
		var ret string
		return ret
	}
	return *o.InvoiceID
}

// GetInvoiceIDOk returns a tuple with the InvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateInvoiceID) GetInvoiceIDOk() (*string, bool) {
	if o == nil || o.InvoiceID == nil {
		return nil, false
	}
	return o.InvoiceID, true
}

// HasInvoiceID returns a boolean if a field has been set.
func (o *BillingEstimateInvoiceID) HasInvoiceID() bool {
	if o != nil && o.InvoiceID != nil {
		return true
	}

	return false
}

// SetInvoiceID gets a reference to the given string and assigns it to the InvoiceID field.
func (o *BillingEstimateInvoiceID) SetInvoiceID(v string) {
	o.InvoiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingEstimateInvoiceID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.InvoiceID != nil {
		toSerialize["invoice_id"] = o.InvoiceID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BillingEstimateInvoiceID) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEstimateInvoiceID := _BillingEstimateInvoiceID{}

	if err = json.Unmarshal(bytes, &varBillingEstimateInvoiceID); err == nil {
		*o = BillingEstimateInvoiceID(varBillingEstimateInvoiceID)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "invoice_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingEstimateInvoiceID is a helper abstraction for handling nullable billingestimateinvoiceid types.
type NullableBillingEstimateInvoiceID struct {
	value *BillingEstimateInvoiceID
	isSet bool
}

// Get returns the value.
func (v NullableBillingEstimateInvoiceID) Get() *BillingEstimateInvoiceID {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingEstimateInvoiceID) Set(val *BillingEstimateInvoiceID) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingEstimateInvoiceID) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingEstimateInvoiceID) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingEstimateInvoiceID returns a pointer to a new instance of NullableBillingEstimateInvoiceID.
func NewNullableBillingEstimateInvoiceID(val *BillingEstimateInvoiceID) *NullableBillingEstimateInvoiceID {
	return &NullableBillingEstimateInvoiceID{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingEstimateInvoiceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBillingEstimateInvoiceID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
