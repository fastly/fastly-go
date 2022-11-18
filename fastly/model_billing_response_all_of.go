// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// BillingResponseAllOf struct for BillingResponseAllOf
type BillingResponseAllOf struct {
	LineItems []BillingResponseLineItem `json:"line_items,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingResponseAllOf BillingResponseAllOf

// NewBillingResponseAllOf instantiates a new BillingResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingResponseAllOf() *BillingResponseAllOf {
	this := BillingResponseAllOf{}
	return &this
}

// NewBillingResponseAllOfWithDefaults instantiates a new BillingResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingResponseAllOfWithDefaults() *BillingResponseAllOf {
	this := BillingResponseAllOf{}
	return &this
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *BillingResponseAllOf) GetLineItems() []BillingResponseLineItem {
	if o == nil || o.LineItems == nil {
		var ret []BillingResponseLineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseAllOf) GetLineItemsOk() ([]BillingResponseLineItem, bool) {
	if o == nil || o.LineItems == nil {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *BillingResponseAllOf) HasLineItems() bool {
	if o != nil && o.LineItems != nil {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []BillingResponseLineItem and assigns it to the LineItems field.
func (o *BillingResponseAllOf) SetLineItems(v []BillingResponseLineItem) {
	o.LineItems = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.LineItems != nil {
		toSerialize["line_items"] = o.LineItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBillingResponseAllOf := _BillingResponseAllOf{}

	if err = json.Unmarshal(bytes, &varBillingResponseAllOf); err == nil {
		*o = BillingResponseAllOf(varBillingResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "line_items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingResponseAllOf is a helper abstraction for handling nullable billingresponseallof types. 
type NullableBillingResponseAllOf struct {
	value *BillingResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBillingResponseAllOf) Get() *BillingResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingResponseAllOf) Set(val *BillingResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingResponseAllOf returns a pointer to a new instance of NullableBillingResponseAllOf.
func NewNullableBillingResponseAllOf(val *BillingResponseAllOf) *NullableBillingResponseAllOf {
	return &NullableBillingResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
