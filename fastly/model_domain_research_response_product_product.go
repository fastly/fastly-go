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

// DomainResearchResponseProductProduct struct for DomainResearchResponseProductProduct
type DomainResearchResponseProductProduct struct {
	// Product identifier
	ID *string `json:"id,omitempty"`
	// Name of the object
	Object               *string `json:"object,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainResearchResponseProductProduct DomainResearchResponseProductProduct

// NewDomainResearchResponseProductProduct instantiates a new DomainResearchResponseProductProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainResearchResponseProductProduct() *DomainResearchResponseProductProduct {
	this := DomainResearchResponseProductProduct{}
	return &this
}

// NewDomainResearchResponseProductProductWithDefaults instantiates a new DomainResearchResponseProductProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainResearchResponseProductProductWithDefaults() *DomainResearchResponseProductProduct {
	this := DomainResearchResponseProductProduct{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *DomainResearchResponseProductProduct) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResearchResponseProductProduct) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *DomainResearchResponseProductProduct) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *DomainResearchResponseProductProduct) SetID(v string) {
	o.ID = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DomainResearchResponseProductProduct) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResearchResponseProductProduct) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DomainResearchResponseProductProduct) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *DomainResearchResponseProductProduct) SetObject(v string) {
	o.Object = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainResearchResponseProductProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *DomainResearchResponseProductProduct) UnmarshalJSON(bytes []byte) (err error) {
	varDomainResearchResponseProductProduct := _DomainResearchResponseProductProduct{}

	if err = json.Unmarshal(bytes, &varDomainResearchResponseProductProduct); err == nil {
		*o = DomainResearchResponseProductProduct(varDomainResearchResponseProductProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDomainResearchResponseProductProduct is a helper abstraction for handling nullable domainresearchresponseproductproduct types.
type NullableDomainResearchResponseProductProduct struct {
	value *DomainResearchResponseProductProduct
	isSet bool
}

// Get returns the value.
func (v NullableDomainResearchResponseProductProduct) Get() *DomainResearchResponseProductProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainResearchResponseProductProduct) Set(val *DomainResearchResponseProductProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainResearchResponseProductProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainResearchResponseProductProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainResearchResponseProductProduct returns a pointer to a new instance of NullableDomainResearchResponseProductProduct.
func NewNullableDomainResearchResponseProductProduct(val *DomainResearchResponseProductProduct) *NullableDomainResearchResponseProductProduct {
	return &NullableDomainResearchResponseProductProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainResearchResponseProductProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainResearchResponseProductProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
