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

// DomainInspectorResponseBodyEnable struct for DomainInspectorResponseBodyEnable
type DomainInspectorResponseBodyEnable struct {
	Product              *DomainInspectorResponseProductProduct `json:"product,omitempty"`
	Service              *ApiDiscoveryResponseServiceService    `json:"service,omitempty"`
	Links                *DomainInspectorResponseLinksLinks     `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainInspectorResponseBodyEnable DomainInspectorResponseBodyEnable

// NewDomainInspectorResponseBodyEnable instantiates a new DomainInspectorResponseBodyEnable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainInspectorResponseBodyEnable() *DomainInspectorResponseBodyEnable {
	this := DomainInspectorResponseBodyEnable{}
	return &this
}

// NewDomainInspectorResponseBodyEnableWithDefaults instantiates a new DomainInspectorResponseBodyEnable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainInspectorResponseBodyEnableWithDefaults() *DomainInspectorResponseBodyEnable {
	this := DomainInspectorResponseBodyEnable{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *DomainInspectorResponseBodyEnable) GetProduct() DomainInspectorResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret DomainInspectorResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorResponseBodyEnable) GetProductOk() (*DomainInspectorResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *DomainInspectorResponseBodyEnable) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given DomainInspectorResponseProductProduct and assigns it to the Product field.
func (o *DomainInspectorResponseBodyEnable) SetProduct(v DomainInspectorResponseProductProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *DomainInspectorResponseBodyEnable) GetService() ApiDiscoveryResponseServiceService {
	if o == nil || o.Service == nil {
		var ret ApiDiscoveryResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorResponseBodyEnable) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *DomainInspectorResponseBodyEnable) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given ApiDiscoveryResponseServiceService and assigns it to the Service field.
func (o *DomainInspectorResponseBodyEnable) SetService(v ApiDiscoveryResponseServiceService) {
	o.Service = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DomainInspectorResponseBodyEnable) GetLinks() DomainInspectorResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret DomainInspectorResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorResponseBodyEnable) GetLinksOk() (*DomainInspectorResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DomainInspectorResponseBodyEnable) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DomainInspectorResponseLinksLinks and assigns it to the Links field.
func (o *DomainInspectorResponseBodyEnable) SetLinks(v DomainInspectorResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainInspectorResponseBodyEnable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
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
func (o *DomainInspectorResponseBodyEnable) UnmarshalJSON(bytes []byte) (err error) {
	varDomainInspectorResponseBodyEnable := _DomainInspectorResponseBodyEnable{}

	if err = json.Unmarshal(bytes, &varDomainInspectorResponseBodyEnable); err == nil {
		*o = DomainInspectorResponseBodyEnable(varDomainInspectorResponseBodyEnable)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "service")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDomainInspectorResponseBodyEnable is a helper abstraction for handling nullable domaininspectorresponsebodyenable types.
type NullableDomainInspectorResponseBodyEnable struct {
	value *DomainInspectorResponseBodyEnable
	isSet bool
}

// Get returns the value.
func (v NullableDomainInspectorResponseBodyEnable) Get() *DomainInspectorResponseBodyEnable {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainInspectorResponseBodyEnable) Set(val *DomainInspectorResponseBodyEnable) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainInspectorResponseBodyEnable) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainInspectorResponseBodyEnable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainInspectorResponseBodyEnable returns a pointer to a new instance of NullableDomainInspectorResponseBodyEnable.
func NewNullableDomainInspectorResponseBodyEnable(val *DomainInspectorResponseBodyEnable) *NullableDomainInspectorResponseBodyEnable {
	return &NullableDomainInspectorResponseBodyEnable{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainInspectorResponseBodyEnable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainInspectorResponseBodyEnable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
