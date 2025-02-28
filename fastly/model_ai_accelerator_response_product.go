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

// AiAcceleratorResponseProduct struct for AiAcceleratorResponseProduct
type AiAcceleratorResponseProduct struct {
	Product              *AiAcceleratorResponseProductProduct `json:"product,omitempty"`
	AdditionalProperties map[string]any
}

type _AiAcceleratorResponseProduct AiAcceleratorResponseProduct

// NewAiAcceleratorResponseProduct instantiates a new AiAcceleratorResponseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAcceleratorResponseProduct() *AiAcceleratorResponseProduct {
	this := AiAcceleratorResponseProduct{}
	return &this
}

// NewAiAcceleratorResponseProductWithDefaults instantiates a new AiAcceleratorResponseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAcceleratorResponseProductWithDefaults() *AiAcceleratorResponseProduct {
	this := AiAcceleratorResponseProduct{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AiAcceleratorResponseProduct) GetProduct() AiAcceleratorResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret AiAcceleratorResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAcceleratorResponseProduct) GetProductOk() (*AiAcceleratorResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AiAcceleratorResponseProduct) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AiAcceleratorResponseProductProduct and assigns it to the Product field.
func (o *AiAcceleratorResponseProduct) SetProduct(v AiAcceleratorResponseProductProduct) {
	o.Product = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AiAcceleratorResponseProduct) MarshalJSON() ([]byte, error) {
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
func (o *AiAcceleratorResponseProduct) UnmarshalJSON(bytes []byte) (err error) {
	varAiAcceleratorResponseProduct := _AiAcceleratorResponseProduct{}

	if err = json.Unmarshal(bytes, &varAiAcceleratorResponseProduct); err == nil {
		*o = AiAcceleratorResponseProduct(varAiAcceleratorResponseProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAiAcceleratorResponseProduct is a helper abstraction for handling nullable aiacceleratorresponseproduct types.
type NullableAiAcceleratorResponseProduct struct {
	value *AiAcceleratorResponseProduct
	isSet bool
}

// Get returns the value.
func (v NullableAiAcceleratorResponseProduct) Get() *AiAcceleratorResponseProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableAiAcceleratorResponseProduct) Set(val *AiAcceleratorResponseProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAiAcceleratorResponseProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAiAcceleratorResponseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiAcceleratorResponseProduct returns a pointer to a new instance of NullableAiAcceleratorResponseProduct.
func NewNullableAiAcceleratorResponseProduct(val *AiAcceleratorResponseProduct) *NullableAiAcceleratorResponseProduct {
	return &NullableAiAcceleratorResponseProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAiAcceleratorResponseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAiAcceleratorResponseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
