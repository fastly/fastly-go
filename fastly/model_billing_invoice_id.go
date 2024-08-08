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

// BillingInvoiceID struct for BillingInvoiceID
type BillingInvoiceID struct {
	InvoiceID *int32 `json:"invoice_id,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingInvoiceID BillingInvoiceID

// NewBillingInvoiceID instantiates a new BillingInvoiceID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInvoiceID() *BillingInvoiceID {
	this := BillingInvoiceID{}
	return &this
}

// NewBillingInvoiceIDWithDefaults instantiates a new BillingInvoiceID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInvoiceIDWithDefaults() *BillingInvoiceID {
	this := BillingInvoiceID{}
	return &this
}

// GetInvoiceID returns the InvoiceID field value if set, zero value otherwise.
func (o *BillingInvoiceID) GetInvoiceID() int32 {
	if o == nil || o.InvoiceID == nil {
		var ret int32
		return ret
	}
	return *o.InvoiceID
}

// GetInvoiceIDOk returns a tuple with the InvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceID) GetInvoiceIDOk() (*int32, bool) {
	if o == nil || o.InvoiceID == nil {
		return nil, false
	}
	return o.InvoiceID, true
}

// HasInvoiceID returns a boolean if a field has been set.
func (o *BillingInvoiceID) HasInvoiceID() bool {
	if o != nil && o.InvoiceID != nil {
		return true
	}

	return false
}

// SetInvoiceID gets a reference to the given int32 and assigns it to the InvoiceID field.
func (o *BillingInvoiceID) SetInvoiceID(v int32) {
	o.InvoiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingInvoiceID) MarshalJSON() ([]byte, error) {
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
func (o *BillingInvoiceID) UnmarshalJSON(bytes []byte) (err error) {
	varBillingInvoiceID := _BillingInvoiceID{}

	if err = json.Unmarshal(bytes, &varBillingInvoiceID); err == nil {
		*o = BillingInvoiceID(varBillingInvoiceID)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "invoice_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingInvoiceID is a helper abstraction for handling nullable billinginvoiceid types. 
type NullableBillingInvoiceID struct {
	value *BillingInvoiceID
	isSet bool
}

// Get returns the value.
func (v NullableBillingInvoiceID) Get() *BillingInvoiceID {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingInvoiceID) Set(val *BillingInvoiceID) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingInvoiceID) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingInvoiceID) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingInvoiceID returns a pointer to a new instance of NullableBillingInvoiceID.
func NewNullableBillingInvoiceID(val *BillingInvoiceID) *NullableBillingInvoiceID {
	return &NullableBillingInvoiceID{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingInvoiceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingInvoiceID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
