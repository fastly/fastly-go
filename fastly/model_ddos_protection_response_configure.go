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

// DdosProtectionResponseConfigure struct for DdosProtectionResponseConfigure
type DdosProtectionResponseConfigure struct {
	Product              *DdosProtectionResponseProductProduct             `json:"product,omitempty"`
	Service              *BotManagementResponseServiceService              `json:"service,omitempty"`
	Configuration        *DdosProtectionResponseConfigurationConfiguration `json:"configuration,omitempty"`
	Links                *DdosProtectionResponseLinksLinks                 `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionResponseConfigure DdosProtectionResponseConfigure

// NewDdosProtectionResponseConfigure instantiates a new DdosProtectionResponseConfigure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionResponseConfigure() *DdosProtectionResponseConfigure {
	this := DdosProtectionResponseConfigure{}
	return &this
}

// NewDdosProtectionResponseConfigureWithDefaults instantiates a new DdosProtectionResponseConfigure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionResponseConfigureWithDefaults() *DdosProtectionResponseConfigure {
	this := DdosProtectionResponseConfigure{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *DdosProtectionResponseConfigure) GetProduct() DdosProtectionResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret DdosProtectionResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseConfigure) GetProductOk() (*DdosProtectionResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *DdosProtectionResponseConfigure) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given DdosProtectionResponseProductProduct and assigns it to the Product field.
func (o *DdosProtectionResponseConfigure) SetProduct(v DdosProtectionResponseProductProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *DdosProtectionResponseConfigure) GetService() BotManagementResponseServiceService {
	if o == nil || o.Service == nil {
		var ret BotManagementResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseConfigure) GetServiceOk() (*BotManagementResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *DdosProtectionResponseConfigure) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given BotManagementResponseServiceService and assigns it to the Service field.
func (o *DdosProtectionResponseConfigure) SetService(v BotManagementResponseServiceService) {
	o.Service = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *DdosProtectionResponseConfigure) GetConfiguration() DdosProtectionResponseConfigurationConfiguration {
	if o == nil || o.Configuration == nil {
		var ret DdosProtectionResponseConfigurationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseConfigure) GetConfigurationOk() (*DdosProtectionResponseConfigurationConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *DdosProtectionResponseConfigure) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given DdosProtectionResponseConfigurationConfiguration and assigns it to the Configuration field.
func (o *DdosProtectionResponseConfigure) SetConfiguration(v DdosProtectionResponseConfigurationConfiguration) {
	o.Configuration = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DdosProtectionResponseConfigure) GetLinks() DdosProtectionResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret DdosProtectionResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseConfigure) GetLinksOk() (*DdosProtectionResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DdosProtectionResponseConfigure) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DdosProtectionResponseLinksLinks and assigns it to the Links field.
func (o *DdosProtectionResponseConfigure) SetLinks(v DdosProtectionResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionResponseConfigure) MarshalJSON() ([]byte, error) {
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
func (o *DdosProtectionResponseConfigure) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionResponseConfigure := _DdosProtectionResponseConfigure{}

	if err = json.Unmarshal(bytes, &varDdosProtectionResponseConfigure); err == nil {
		*o = DdosProtectionResponseConfigure(varDdosProtectionResponseConfigure)
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

// NullableDdosProtectionResponseConfigure is a helper abstraction for handling nullable ddosprotectionresponseconfigure types.
type NullableDdosProtectionResponseConfigure struct {
	value *DdosProtectionResponseConfigure
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionResponseConfigure) Get() *DdosProtectionResponseConfigure {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionResponseConfigure) Set(val *DdosProtectionResponseConfigure) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionResponseConfigure) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionResponseConfigure) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionResponseConfigure returns a pointer to a new instance of NullableDdosProtectionResponseConfigure.
func NewNullableDdosProtectionResponseConfigure(val *DdosProtectionResponseConfigure) *NullableDdosProtectionResponseConfigure {
	return &NullableDdosProtectionResponseConfigure{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionResponseConfigure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionResponseConfigure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
