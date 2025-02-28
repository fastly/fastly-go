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

// NgwafResponseProduct struct for NgwafResponseProduct
type NgwafResponseProduct struct {
	Product              *NgwafResponseProductProduct `json:"product,omitempty"`
	AdditionalProperties map[string]any
}

type _NgwafResponseProduct NgwafResponseProduct

// NewNgwafResponseProduct instantiates a new NgwafResponseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgwafResponseProduct() *NgwafResponseProduct {
	this := NgwafResponseProduct{}
	return &this
}

// NewNgwafResponseProductWithDefaults instantiates a new NgwafResponseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgwafResponseProductWithDefaults() *NgwafResponseProduct {
	this := NgwafResponseProduct{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *NgwafResponseProduct) GetProduct() NgwafResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret NgwafResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafResponseProduct) GetProductOk() (*NgwafResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *NgwafResponseProduct) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NgwafResponseProductProduct and assigns it to the Product field.
func (o *NgwafResponseProduct) SetProduct(v NgwafResponseProductProduct) {
	o.Product = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NgwafResponseProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NgwafResponseProduct) UnmarshalJSON(bytes []byte) (err error) {
	varNgwafResponseProduct := _NgwafResponseProduct{}

	if err = json.Unmarshal(bytes, &varNgwafResponseProduct); err == nil {
		*o = NgwafResponseProduct(varNgwafResponseProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNgwafResponseProduct is a helper abstraction for handling nullable ngwafresponseproduct types.
type NullableNgwafResponseProduct struct {
	value *NgwafResponseProduct
	isSet bool
}

// Get returns the value.
func (v NullableNgwafResponseProduct) Get() *NgwafResponseProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableNgwafResponseProduct) Set(val *NgwafResponseProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNgwafResponseProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNgwafResponseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNgwafResponseProduct returns a pointer to a new instance of NullableNgwafResponseProduct.
func NewNullableNgwafResponseProduct(val *NgwafResponseProduct) *NullableNgwafResponseProduct {
	return &NullableNgwafResponseProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNgwafResponseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNgwafResponseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
