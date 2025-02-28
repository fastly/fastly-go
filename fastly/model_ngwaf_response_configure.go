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

// NgwafResponseConfigure struct for NgwafResponseConfigure
type NgwafResponseConfigure struct {
	Product              *NgwafResponseProductProduct             `json:"product,omitempty"`
	Service              *BotManagementResponseServiceService     `json:"service,omitempty"`
	Configuration        *NgwafResponseConfigurationConfiguration `json:"configuration,omitempty"`
	Links                *NgwafResponseLinksLinks                 `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _NgwafResponseConfigure NgwafResponseConfigure

// NewNgwafResponseConfigure instantiates a new NgwafResponseConfigure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgwafResponseConfigure() *NgwafResponseConfigure {
	this := NgwafResponseConfigure{}
	return &this
}

// NewNgwafResponseConfigureWithDefaults instantiates a new NgwafResponseConfigure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgwafResponseConfigureWithDefaults() *NgwafResponseConfigure {
	this := NgwafResponseConfigure{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *NgwafResponseConfigure) GetProduct() NgwafResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret NgwafResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafResponseConfigure) GetProductOk() (*NgwafResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *NgwafResponseConfigure) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NgwafResponseProductProduct and assigns it to the Product field.
func (o *NgwafResponseConfigure) SetProduct(v NgwafResponseProductProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *NgwafResponseConfigure) GetService() BotManagementResponseServiceService {
	if o == nil || o.Service == nil {
		var ret BotManagementResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafResponseConfigure) GetServiceOk() (*BotManagementResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *NgwafResponseConfigure) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given BotManagementResponseServiceService and assigns it to the Service field.
func (o *NgwafResponseConfigure) SetService(v BotManagementResponseServiceService) {
	o.Service = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *NgwafResponseConfigure) GetConfiguration() NgwafResponseConfigurationConfiguration {
	if o == nil || o.Configuration == nil {
		var ret NgwafResponseConfigurationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafResponseConfigure) GetConfigurationOk() (*NgwafResponseConfigurationConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *NgwafResponseConfigure) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NgwafResponseConfigurationConfiguration and assigns it to the Configuration field.
func (o *NgwafResponseConfigure) SetConfiguration(v NgwafResponseConfigurationConfiguration) {
	o.Configuration = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NgwafResponseConfigure) GetLinks() NgwafResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret NgwafResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafResponseConfigure) GetLinksOk() (*NgwafResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NgwafResponseConfigure) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given NgwafResponseLinksLinks and assigns it to the Links field.
func (o *NgwafResponseConfigure) SetLinks(v NgwafResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NgwafResponseConfigure) MarshalJSON() ([]byte, error) {
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
func (o *NgwafResponseConfigure) UnmarshalJSON(bytes []byte) (err error) {
	varNgwafResponseConfigure := _NgwafResponseConfigure{}

	if err = json.Unmarshal(bytes, &varNgwafResponseConfigure); err == nil {
		*o = NgwafResponseConfigure(varNgwafResponseConfigure)
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

// NullableNgwafResponseConfigure is a helper abstraction for handling nullable ngwafresponseconfigure types.
type NullableNgwafResponseConfigure struct {
	value *NgwafResponseConfigure
	isSet bool
}

// Get returns the value.
func (v NullableNgwafResponseConfigure) Get() *NgwafResponseConfigure {
	return v.value
}

// Set modifies the value.
func (v *NullableNgwafResponseConfigure) Set(val *NgwafResponseConfigure) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNgwafResponseConfigure) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNgwafResponseConfigure) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNgwafResponseConfigure returns a pointer to a new instance of NullableNgwafResponseConfigure.
func NewNullableNgwafResponseConfigure(val *NgwafResponseConfigure) *NullableNgwafResponseConfigure {
	return &NullableNgwafResponseConfigure{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNgwafResponseConfigure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNgwafResponseConfigure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
