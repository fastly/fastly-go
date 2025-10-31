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

// DomainInspectorResponseProductProduct struct for DomainInspectorResponseProductProduct
type DomainInspectorResponseProductProduct struct {
	// Product identifier
	Id *string `json:"id,omitempty"`
	// Name of the object
	Object               *string `json:"object,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainInspectorResponseProductProduct DomainInspectorResponseProductProduct

// NewDomainInspectorResponseProductProduct instantiates a new DomainInspectorResponseProductProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainInspectorResponseProductProduct() *DomainInspectorResponseProductProduct {
	this := DomainInspectorResponseProductProduct{}
	return &this
}

// NewDomainInspectorResponseProductProductWithDefaults instantiates a new DomainInspectorResponseProductProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainInspectorResponseProductProductWithDefaults() *DomainInspectorResponseProductProduct {
	this := DomainInspectorResponseProductProduct{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DomainInspectorResponseProductProduct) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorResponseProductProduct) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DomainInspectorResponseProductProduct) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DomainInspectorResponseProductProduct) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DomainInspectorResponseProductProduct) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorResponseProductProduct) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DomainInspectorResponseProductProduct) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *DomainInspectorResponseProductProduct) SetObject(v string) {
	o.Object = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainInspectorResponseProductProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DomainInspectorResponseProductProduct) UnmarshalJSON(bytes []byte) (err error) {
	varDomainInspectorResponseProductProduct := _DomainInspectorResponseProductProduct{}

	if err = json.Unmarshal(bytes, &varDomainInspectorResponseProductProduct); err == nil {
		*o = DomainInspectorResponseProductProduct(varDomainInspectorResponseProductProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDomainInspectorResponseProductProduct is a helper abstraction for handling nullable domaininspectorresponseproductproduct types.
type NullableDomainInspectorResponseProductProduct struct {
	value *DomainInspectorResponseProductProduct
	isSet bool
}

// Get returns the value.
func (v NullableDomainInspectorResponseProductProduct) Get() *DomainInspectorResponseProductProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainInspectorResponseProductProduct) Set(val *DomainInspectorResponseProductProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainInspectorResponseProductProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainInspectorResponseProductProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainInspectorResponseProductProduct returns a pointer to a new instance of NullableDomainInspectorResponseProductProduct.
func NewNullableDomainInspectorResponseProductProduct(val *DomainInspectorResponseProductProduct) *NullableDomainInspectorResponseProductProduct {
	return &NullableDomainInspectorResponseProductProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainInspectorResponseProductProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainInspectorResponseProductProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
