// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// EnabledProduct struct for EnabledProduct
type EnabledProduct struct {
	Product *EnabledProductProduct `json:"product,omitempty"`
	Service *EnabledProductProduct `json:"service,omitempty"`
	Links *EnabledProductLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _EnabledProduct EnabledProduct

// NewEnabledProduct instantiates a new EnabledProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnabledProduct() *EnabledProduct {
	this := EnabledProduct{}
	return &this
}

// NewEnabledProductWithDefaults instantiates a new EnabledProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnabledProductWithDefaults() *EnabledProduct {
	this := EnabledProduct{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *EnabledProduct) GetProduct() EnabledProductProduct {
	if o == nil || o.Product == nil {
		var ret EnabledProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProduct) GetProductOk() (*EnabledProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *EnabledProduct) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given EnabledProductProduct and assigns it to the Product field.
func (o *EnabledProduct) SetProduct(v EnabledProductProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *EnabledProduct) GetService() EnabledProductProduct {
	if o == nil || o.Service == nil {
		var ret EnabledProductProduct
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProduct) GetServiceOk() (*EnabledProductProduct, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *EnabledProduct) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given EnabledProductProduct and assigns it to the Service field.
func (o *EnabledProduct) SetService(v EnabledProductProduct) {
	o.Service = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnabledProduct) GetLinks() EnabledProductLinks {
	if o == nil || o.Links == nil {
		var ret EnabledProductLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProduct) GetLinksOk() (*EnabledProductLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnabledProduct) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EnabledProductLinks and assigns it to the Links field.
func (o *EnabledProduct) SetLinks(v EnabledProductLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o EnabledProduct) MarshalJSON() ([]byte, error) {
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
func (o *EnabledProduct) UnmarshalJSON(bytes []byte) (err error) {
	varEnabledProduct := _EnabledProduct{}

	if err = json.Unmarshal(bytes, &varEnabledProduct); err == nil {
		*o = EnabledProduct(varEnabledProduct)
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

// NullableEnabledProduct is a helper abstraction for handling nullable enabledproduct types. 
type NullableEnabledProduct struct {
	value *EnabledProduct
	isSet bool
}

// Get returns the value.
func (v NullableEnabledProduct) Get() *EnabledProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableEnabledProduct) Set(val *EnabledProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableEnabledProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableEnabledProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableEnabledProduct returns a pointer to a new instance of NullableEnabledProduct.
func NewNullableEnabledProduct(val *EnabledProduct) *NullableEnabledProduct {
	return &NullableEnabledProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableEnabledProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableEnabledProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
