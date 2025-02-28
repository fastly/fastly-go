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

// BrotliCompressionResponseProduct struct for BrotliCompressionResponseProduct
type BrotliCompressionResponseProduct struct {
	Product              *BrotliCompressionResponseProductProduct `json:"product,omitempty"`
	AdditionalProperties map[string]any
}

type _BrotliCompressionResponseProduct BrotliCompressionResponseProduct

// NewBrotliCompressionResponseProduct instantiates a new BrotliCompressionResponseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrotliCompressionResponseProduct() *BrotliCompressionResponseProduct {
	this := BrotliCompressionResponseProduct{}
	return &this
}

// NewBrotliCompressionResponseProductWithDefaults instantiates a new BrotliCompressionResponseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrotliCompressionResponseProductWithDefaults() *BrotliCompressionResponseProduct {
	this := BrotliCompressionResponseProduct{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *BrotliCompressionResponseProduct) GetProduct() BrotliCompressionResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret BrotliCompressionResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrotliCompressionResponseProduct) GetProductOk() (*BrotliCompressionResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *BrotliCompressionResponseProduct) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given BrotliCompressionResponseProductProduct and assigns it to the Product field.
func (o *BrotliCompressionResponseProduct) SetProduct(v BrotliCompressionResponseProductProduct) {
	o.Product = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BrotliCompressionResponseProduct) MarshalJSON() ([]byte, error) {
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
func (o *BrotliCompressionResponseProduct) UnmarshalJSON(bytes []byte) (err error) {
	varBrotliCompressionResponseProduct := _BrotliCompressionResponseProduct{}

	if err = json.Unmarshal(bytes, &varBrotliCompressionResponseProduct); err == nil {
		*o = BrotliCompressionResponseProduct(varBrotliCompressionResponseProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBrotliCompressionResponseProduct is a helper abstraction for handling nullable brotlicompressionresponseproduct types.
type NullableBrotliCompressionResponseProduct struct {
	value *BrotliCompressionResponseProduct
	isSet bool
}

// Get returns the value.
func (v NullableBrotliCompressionResponseProduct) Get() *BrotliCompressionResponseProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableBrotliCompressionResponseProduct) Set(val *BrotliCompressionResponseProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBrotliCompressionResponseProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBrotliCompressionResponseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBrotliCompressionResponseProduct returns a pointer to a new instance of NullableBrotliCompressionResponseProduct.
func NewNullableBrotliCompressionResponseProduct(val *BrotliCompressionResponseProduct) *NullableBrotliCompressionResponseProduct {
	return &NullableBrotliCompressionResponseProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBrotliCompressionResponseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBrotliCompressionResponseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
