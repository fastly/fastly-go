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

// Serviceusagetype struct for Serviceusagetype
type Serviceusagetype struct {
	// The product identifier associated with the usage type. This corresponds to a Fastly product offering.
	ProductId *string `json:"product_id,omitempty"`
	// Full name of the product usage type as it might appear on a customer's invoice.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _Serviceusagetype Serviceusagetype

// NewServiceusagetype instantiates a new Serviceusagetype object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceusagetype() *Serviceusagetype {
	this := Serviceusagetype{}
	return &this
}

// NewServiceusagetypeWithDefaults instantiates a new Serviceusagetype object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceusagetypeWithDefaults() *Serviceusagetype {
	this := Serviceusagetype{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *Serviceusagetype) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceusagetype) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *Serviceusagetype) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *Serviceusagetype) SetProductId(v string) {
	o.ProductId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Serviceusagetype) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceusagetype) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Serviceusagetype) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Serviceusagetype) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Serviceusagetype) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ProductId != nil {
		toSerialize["product_id"] = o.ProductId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Serviceusagetype) UnmarshalJSON(bytes []byte) (err error) {
	varServiceusagetype := _Serviceusagetype{}

	if err = json.Unmarshal(bytes, &varServiceusagetype); err == nil {
		*o = Serviceusagetype(varServiceusagetype)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceusagetype is a helper abstraction for handling nullable serviceusagetype types.
type NullableServiceusagetype struct {
	value *Serviceusagetype
	isSet bool
}

// Get returns the value.
func (v NullableServiceusagetype) Get() *Serviceusagetype {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceusagetype) Set(val *Serviceusagetype) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceusagetype) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceusagetype) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceusagetype returns a pointer to a new instance of NullableServiceusagetype.
func NewNullableServiceusagetype(val *Serviceusagetype) *NullableServiceusagetype {
	return &NullableServiceusagetype{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceusagetype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceusagetype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
