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

// OriginInspectorResponseProduct struct for OriginInspectorResponseProduct
type OriginInspectorResponseProduct struct {
	Product              *OriginInspectorResponseProductProduct `json:"product,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorResponseProduct OriginInspectorResponseProduct

// NewOriginInspectorResponseProduct instantiates a new OriginInspectorResponseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorResponseProduct() *OriginInspectorResponseProduct {
	this := OriginInspectorResponseProduct{}
	return &this
}

// NewOriginInspectorResponseProductWithDefaults instantiates a new OriginInspectorResponseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorResponseProductWithDefaults() *OriginInspectorResponseProduct {
	this := OriginInspectorResponseProduct{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *OriginInspectorResponseProduct) GetProduct() OriginInspectorResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret OriginInspectorResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorResponseProduct) GetProductOk() (*OriginInspectorResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *OriginInspectorResponseProduct) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given OriginInspectorResponseProductProduct and assigns it to the Product field.
func (o *OriginInspectorResponseProduct) SetProduct(v OriginInspectorResponseProductProduct) {
	o.Product = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorResponseProduct) MarshalJSON() ([]byte, error) {
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
func (o *OriginInspectorResponseProduct) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorResponseProduct := _OriginInspectorResponseProduct{}

	if err = json.Unmarshal(bytes, &varOriginInspectorResponseProduct); err == nil {
		*o = OriginInspectorResponseProduct(varOriginInspectorResponseProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorResponseProduct is a helper abstraction for handling nullable origininspectorresponseproduct types.
type NullableOriginInspectorResponseProduct struct {
	value *OriginInspectorResponseProduct
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorResponseProduct) Get() *OriginInspectorResponseProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorResponseProduct) Set(val *OriginInspectorResponseProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorResponseProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorResponseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorResponseProduct returns a pointer to a new instance of NullableOriginInspectorResponseProduct.
func NewNullableOriginInspectorResponseProduct(val *OriginInspectorResponseProduct) *NullableOriginInspectorResponseProduct {
	return &NullableOriginInspectorResponseProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorResponseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorResponseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
