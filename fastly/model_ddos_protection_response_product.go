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

// DdosProtectionResponseProduct struct for DdosProtectionResponseProduct
type DdosProtectionResponseProduct struct {
	Product              *DdosProtectionResponseProductProduct `json:"product,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionResponseProduct DdosProtectionResponseProduct

// NewDdosProtectionResponseProduct instantiates a new DdosProtectionResponseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionResponseProduct() *DdosProtectionResponseProduct {
	this := DdosProtectionResponseProduct{}
	return &this
}

// NewDdosProtectionResponseProductWithDefaults instantiates a new DdosProtectionResponseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionResponseProductWithDefaults() *DdosProtectionResponseProduct {
	this := DdosProtectionResponseProduct{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *DdosProtectionResponseProduct) GetProduct() DdosProtectionResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret DdosProtectionResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseProduct) GetProductOk() (*DdosProtectionResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *DdosProtectionResponseProduct) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given DdosProtectionResponseProductProduct and assigns it to the Product field.
func (o *DdosProtectionResponseProduct) SetProduct(v DdosProtectionResponseProductProduct) {
	o.Product = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionResponseProduct) MarshalJSON() ([]byte, error) {
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
func (o *DdosProtectionResponseProduct) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionResponseProduct := _DdosProtectionResponseProduct{}

	if err = json.Unmarshal(bytes, &varDdosProtectionResponseProduct); err == nil {
		*o = DdosProtectionResponseProduct(varDdosProtectionResponseProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionResponseProduct is a helper abstraction for handling nullable ddosprotectionresponseproduct types.
type NullableDdosProtectionResponseProduct struct {
	value *DdosProtectionResponseProduct
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionResponseProduct) Get() *DdosProtectionResponseProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionResponseProduct) Set(val *DdosProtectionResponseProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionResponseProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionResponseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionResponseProduct returns a pointer to a new instance of NullableDdosProtectionResponseProduct.
func NewNullableDdosProtectionResponseProduct(val *DdosProtectionResponseProduct) *NullableDdosProtectionResponseProduct {
	return &NullableDdosProtectionResponseProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionResponseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionResponseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
