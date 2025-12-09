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

// DomainResearchResponseBodyEnable struct for DomainResearchResponseBodyEnable
type DomainResearchResponseBodyEnable struct {
	Product              *DomainResearchResponseProductProduct  `json:"product,omitempty"`
	Customer             *AiAcceleratorResponseCustomerCustomer `json:"customer,omitempty"`
	Links                *DomainResearchResponseLinksLinks      `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainResearchResponseBodyEnable DomainResearchResponseBodyEnable

// NewDomainResearchResponseBodyEnable instantiates a new DomainResearchResponseBodyEnable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainResearchResponseBodyEnable() *DomainResearchResponseBodyEnable {
	this := DomainResearchResponseBodyEnable{}
	return &this
}

// NewDomainResearchResponseBodyEnableWithDefaults instantiates a new DomainResearchResponseBodyEnable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainResearchResponseBodyEnableWithDefaults() *DomainResearchResponseBodyEnable {
	this := DomainResearchResponseBodyEnable{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *DomainResearchResponseBodyEnable) GetProduct() DomainResearchResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret DomainResearchResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResearchResponseBodyEnable) GetProductOk() (*DomainResearchResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *DomainResearchResponseBodyEnable) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given DomainResearchResponseProductProduct and assigns it to the Product field.
func (o *DomainResearchResponseBodyEnable) SetProduct(v DomainResearchResponseProductProduct) {
	o.Product = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *DomainResearchResponseBodyEnable) GetCustomer() AiAcceleratorResponseCustomerCustomer {
	if o == nil || o.Customer == nil {
		var ret AiAcceleratorResponseCustomerCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResearchResponseBodyEnable) GetCustomerOk() (*AiAcceleratorResponseCustomerCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *DomainResearchResponseBodyEnable) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given AiAcceleratorResponseCustomerCustomer and assigns it to the Customer field.
func (o *DomainResearchResponseBodyEnable) SetCustomer(v AiAcceleratorResponseCustomerCustomer) {
	o.Customer = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DomainResearchResponseBodyEnable) GetLinks() DomainResearchResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret DomainResearchResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResearchResponseBodyEnable) GetLinksOk() (*DomainResearchResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DomainResearchResponseBodyEnable) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DomainResearchResponseLinksLinks and assigns it to the Links field.
func (o *DomainResearchResponseBodyEnable) SetLinks(v DomainResearchResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainResearchResponseBodyEnable) MarshalJSON() ([]byte, error) {
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
func (o *DomainResearchResponseBodyEnable) UnmarshalJSON(bytes []byte) (err error) {
	varDomainResearchResponseBodyEnable := _DomainResearchResponseBodyEnable{}

	if err = json.Unmarshal(bytes, &varDomainResearchResponseBodyEnable); err == nil {
		*o = DomainResearchResponseBodyEnable(varDomainResearchResponseBodyEnable)
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

// NullableDomainResearchResponseBodyEnable is a helper abstraction for handling nullable domainresearchresponsebodyenable types.
type NullableDomainResearchResponseBodyEnable struct {
	value *DomainResearchResponseBodyEnable
	isSet bool
}

// Get returns the value.
func (v NullableDomainResearchResponseBodyEnable) Get() *DomainResearchResponseBodyEnable {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainResearchResponseBodyEnable) Set(val *DomainResearchResponseBodyEnable) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainResearchResponseBodyEnable) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainResearchResponseBodyEnable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainResearchResponseBodyEnable returns a pointer to a new instance of NullableDomainResearchResponseBodyEnable.
func NewNullableDomainResearchResponseBodyEnable(val *DomainResearchResponseBodyEnable) *NullableDomainResearchResponseBodyEnable {
	return &NullableDomainResearchResponseBodyEnable{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainResearchResponseBodyEnable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainResearchResponseBodyEnable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
