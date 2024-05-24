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

// EnabledProductResponse struct for EnabledProductResponse
type EnabledProductResponse struct {
	Product *EnabledProductResponseProduct `json:"product,omitempty"`
	Service *EnabledProductResponseService `json:"service,omitempty"`
	Links *EnabledProductResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _EnabledProductResponse EnabledProductResponse

// NewEnabledProductResponse instantiates a new EnabledProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnabledProductResponse() *EnabledProductResponse {
	this := EnabledProductResponse{}
	return &this
}

// NewEnabledProductResponseWithDefaults instantiates a new EnabledProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnabledProductResponseWithDefaults() *EnabledProductResponse {
	this := EnabledProductResponse{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *EnabledProductResponse) GetProduct() EnabledProductResponseProduct {
	if o == nil || o.Product == nil {
		var ret EnabledProductResponseProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProductResponse) GetProductOk() (*EnabledProductResponseProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *EnabledProductResponse) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given EnabledProductResponseProduct and assigns it to the Product field.
func (o *EnabledProductResponse) SetProduct(v EnabledProductResponseProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *EnabledProductResponse) GetService() EnabledProductResponseService {
	if o == nil || o.Service == nil {
		var ret EnabledProductResponseService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProductResponse) GetServiceOk() (*EnabledProductResponseService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *EnabledProductResponse) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given EnabledProductResponseService and assigns it to the Service field.
func (o *EnabledProductResponse) SetService(v EnabledProductResponseService) {
	o.Service = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnabledProductResponse) GetLinks() EnabledProductResponseLinks {
	if o == nil || o.Links == nil {
		var ret EnabledProductResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProductResponse) GetLinksOk() (*EnabledProductResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnabledProductResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EnabledProductResponseLinks and assigns it to the Links field.
func (o *EnabledProductResponse) SetLinks(v EnabledProductResponseLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o EnabledProductResponse) MarshalJSON() ([]byte, error) {
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
func (o *EnabledProductResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEnabledProductResponse := _EnabledProductResponse{}

	if err = json.Unmarshal(bytes, &varEnabledProductResponse); err == nil {
		*o = EnabledProductResponse(varEnabledProductResponse)
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

// NullableEnabledProductResponse is a helper abstraction for handling nullable enabledproductresponse types. 
type NullableEnabledProductResponse struct {
	value *EnabledProductResponse
	isSet bool
}

// Get returns the value.
func (v NullableEnabledProductResponse) Get() *EnabledProductResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableEnabledProductResponse) Set(val *EnabledProductResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableEnabledProductResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableEnabledProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableEnabledProductResponse returns a pointer to a new instance of NullableEnabledProductResponse.
func NewNullableEnabledProductResponse(val *EnabledProductResponse) *NullableEnabledProductResponse {
	return &NullableEnabledProductResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableEnabledProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableEnabledProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
