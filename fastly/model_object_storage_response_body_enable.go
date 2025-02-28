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

// ObjectStorageResponseBodyEnable struct for ObjectStorageResponseBodyEnable
type ObjectStorageResponseBodyEnable struct {
	Product              *ObjectStorageResponseProductProduct   `json:"product,omitempty"`
	Customer             *AiAcceleratorResponseCustomerCustomer `json:"customer,omitempty"`
	Links                *ObjectStorageResponseLinksLinks       `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _ObjectStorageResponseBodyEnable ObjectStorageResponseBodyEnable

// NewObjectStorageResponseBodyEnable instantiates a new ObjectStorageResponseBodyEnable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageResponseBodyEnable() *ObjectStorageResponseBodyEnable {
	this := ObjectStorageResponseBodyEnable{}
	return &this
}

// NewObjectStorageResponseBodyEnableWithDefaults instantiates a new ObjectStorageResponseBodyEnable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageResponseBodyEnableWithDefaults() *ObjectStorageResponseBodyEnable {
	this := ObjectStorageResponseBodyEnable{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ObjectStorageResponseBodyEnable) GetProduct() ObjectStorageResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret ObjectStorageResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponseBodyEnable) GetProductOk() (*ObjectStorageResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ObjectStorageResponseBodyEnable) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ObjectStorageResponseProductProduct and assigns it to the Product field.
func (o *ObjectStorageResponseBodyEnable) SetProduct(v ObjectStorageResponseProductProduct) {
	o.Product = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ObjectStorageResponseBodyEnable) GetCustomer() AiAcceleratorResponseCustomerCustomer {
	if o == nil || o.Customer == nil {
		var ret AiAcceleratorResponseCustomerCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponseBodyEnable) GetCustomerOk() (*AiAcceleratorResponseCustomerCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ObjectStorageResponseBodyEnable) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given AiAcceleratorResponseCustomerCustomer and assigns it to the Customer field.
func (o *ObjectStorageResponseBodyEnable) SetCustomer(v AiAcceleratorResponseCustomerCustomer) {
	o.Customer = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ObjectStorageResponseBodyEnable) GetLinks() ObjectStorageResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret ObjectStorageResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponseBodyEnable) GetLinksOk() (*ObjectStorageResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ObjectStorageResponseBodyEnable) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ObjectStorageResponseLinksLinks and assigns it to the Links field.
func (o *ObjectStorageResponseBodyEnable) SetLinks(v ObjectStorageResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ObjectStorageResponseBodyEnable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ObjectStorageResponseBodyEnable) UnmarshalJSON(bytes []byte) (err error) {
	varObjectStorageResponseBodyEnable := _ObjectStorageResponseBodyEnable{}

	if err = json.Unmarshal(bytes, &varObjectStorageResponseBodyEnable); err == nil {
		*o = ObjectStorageResponseBodyEnable(varObjectStorageResponseBodyEnable)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "customer")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableObjectStorageResponseBodyEnable is a helper abstraction for handling nullable objectstorageresponsebodyenable types.
type NullableObjectStorageResponseBodyEnable struct {
	value *ObjectStorageResponseBodyEnable
	isSet bool
}

// Get returns the value.
func (v NullableObjectStorageResponseBodyEnable) Get() *ObjectStorageResponseBodyEnable {
	return v.value
}

// Set modifies the value.
func (v *NullableObjectStorageResponseBodyEnable) Set(val *ObjectStorageResponseBodyEnable) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableObjectStorageResponseBodyEnable) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableObjectStorageResponseBodyEnable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableObjectStorageResponseBodyEnable returns a pointer to a new instance of NullableObjectStorageResponseBodyEnable.
func NewNullableObjectStorageResponseBodyEnable(val *ObjectStorageResponseBodyEnable) *NullableObjectStorageResponseBodyEnable {
	return &NullableObjectStorageResponseBodyEnable{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableObjectStorageResponseBodyEnable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableObjectStorageResponseBodyEnable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
