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

// ObjectStorageResponseProduct struct for ObjectStorageResponseProduct
type ObjectStorageResponseProduct struct {
	Product              *ObjectStorageResponseProductProduct `json:"product,omitempty"`
	AdditionalProperties map[string]any
}

type _ObjectStorageResponseProduct ObjectStorageResponseProduct

// NewObjectStorageResponseProduct instantiates a new ObjectStorageResponseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageResponseProduct() *ObjectStorageResponseProduct {
	this := ObjectStorageResponseProduct{}
	return &this
}

// NewObjectStorageResponseProductWithDefaults instantiates a new ObjectStorageResponseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageResponseProductWithDefaults() *ObjectStorageResponseProduct {
	this := ObjectStorageResponseProduct{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ObjectStorageResponseProduct) GetProduct() ObjectStorageResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret ObjectStorageResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponseProduct) GetProductOk() (*ObjectStorageResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ObjectStorageResponseProduct) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ObjectStorageResponseProductProduct and assigns it to the Product field.
func (o *ObjectStorageResponseProduct) SetProduct(v ObjectStorageResponseProductProduct) {
	o.Product = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ObjectStorageResponseProduct) MarshalJSON() ([]byte, error) {
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
func (o *ObjectStorageResponseProduct) UnmarshalJSON(bytes []byte) (err error) {
	varObjectStorageResponseProduct := _ObjectStorageResponseProduct{}

	if err = json.Unmarshal(bytes, &varObjectStorageResponseProduct); err == nil {
		*o = ObjectStorageResponseProduct(varObjectStorageResponseProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableObjectStorageResponseProduct is a helper abstraction for handling nullable objectstorageresponseproduct types.
type NullableObjectStorageResponseProduct struct {
	value *ObjectStorageResponseProduct
	isSet bool
}

// Get returns the value.
func (v NullableObjectStorageResponseProduct) Get() *ObjectStorageResponseProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableObjectStorageResponseProduct) Set(val *ObjectStorageResponseProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableObjectStorageResponseProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableObjectStorageResponseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableObjectStorageResponseProduct returns a pointer to a new instance of NullableObjectStorageResponseProduct.
func NewNullableObjectStorageResponseProduct(val *ObjectStorageResponseProduct) *NullableObjectStorageResponseProduct {
	return &NullableObjectStorageResponseProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableObjectStorageResponseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableObjectStorageResponseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
