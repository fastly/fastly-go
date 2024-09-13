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

// ConfiguredProductResponse struct for ConfiguredProductResponse
type ConfiguredProductResponse struct {
	Product *ConfiguredProductResponseProduct `json:"product,omitempty"`
	Service *EnabledProductResponseService `json:"service,omitempty"`
	Configuration *ConfiguredProductResponseConfiguration `json:"configuration,omitempty"`
	Links *ConfiguredProductResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _ConfiguredProductResponse ConfiguredProductResponse

// NewConfiguredProductResponse instantiates a new ConfiguredProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguredProductResponse() *ConfiguredProductResponse {
	this := ConfiguredProductResponse{}
	return &this
}

// NewConfiguredProductResponseWithDefaults instantiates a new ConfiguredProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguredProductResponseWithDefaults() *ConfiguredProductResponse {
	this := ConfiguredProductResponse{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ConfiguredProductResponse) GetProduct() ConfiguredProductResponseProduct {
	if o == nil || o.Product == nil {
		var ret ConfiguredProductResponseProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponse) GetProductOk() (*ConfiguredProductResponseProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ConfiguredProductResponse) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ConfiguredProductResponseProduct and assigns it to the Product field.
func (o *ConfiguredProductResponse) SetProduct(v ConfiguredProductResponseProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ConfiguredProductResponse) GetService() EnabledProductResponseService {
	if o == nil || o.Service == nil {
		var ret EnabledProductResponseService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponse) GetServiceOk() (*EnabledProductResponseService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ConfiguredProductResponse) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given EnabledProductResponseService and assigns it to the Service field.
func (o *ConfiguredProductResponse) SetService(v EnabledProductResponseService) {
	o.Service = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ConfiguredProductResponse) GetConfiguration() ConfiguredProductResponseConfiguration {
	if o == nil || o.Configuration == nil {
		var ret ConfiguredProductResponseConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponse) GetConfigurationOk() (*ConfiguredProductResponseConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ConfiguredProductResponse) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given ConfiguredProductResponseConfiguration and assigns it to the Configuration field.
func (o *ConfiguredProductResponse) SetConfiguration(v ConfiguredProductResponseConfiguration) {
	o.Configuration = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ConfiguredProductResponse) GetLinks() ConfiguredProductResponseLinks {
	if o == nil || o.Links == nil {
		var ret ConfiguredProductResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponse) GetLinksOk() (*ConfiguredProductResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ConfiguredProductResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ConfiguredProductResponseLinks and assigns it to the Links field.
func (o *ConfiguredProductResponse) SetLinks(v ConfiguredProductResponseLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ConfiguredProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
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
func (o *ConfiguredProductResponse) UnmarshalJSON(bytes []byte) (err error) {
	varConfiguredProductResponse := _ConfiguredProductResponse{}

	if err = json.Unmarshal(bytes, &varConfiguredProductResponse); err == nil {
		*o = ConfiguredProductResponse(varConfiguredProductResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "service")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableConfiguredProductResponse is a helper abstraction for handling nullable configuredproductresponse types. 
type NullableConfiguredProductResponse struct {
	value *ConfiguredProductResponse
	isSet bool
}

// Get returns the value.
func (v NullableConfiguredProductResponse) Get() *ConfiguredProductResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableConfiguredProductResponse) Set(val *ConfiguredProductResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConfiguredProductResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConfiguredProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConfiguredProductResponse returns a pointer to a new instance of NullableConfiguredProductResponse.
func NewNullableConfiguredProductResponse(val *ConfiguredProductResponse) *NullableConfiguredProductResponse {
	return &NullableConfiguredProductResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConfiguredProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableConfiguredProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
