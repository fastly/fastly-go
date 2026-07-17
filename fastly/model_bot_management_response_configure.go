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

// BotManagementResponseConfigure struct for BotManagementResponseConfigure
type BotManagementResponseConfigure struct {
	Product              *BotManagementResponseProductProduct             `json:"product,omitempty"`
	Service              *ApiDiscoveryResponseServiceService              `json:"service,omitempty"`
	Configuration        *BotManagementResponseConfigurationConfiguration `json:"configuration,omitempty"`
	Links                *BotManagementResponseLinksLinks                 `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementResponseConfigure BotManagementResponseConfigure

// NewBotManagementResponseConfigure instantiates a new BotManagementResponseConfigure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementResponseConfigure() *BotManagementResponseConfigure {
	this := BotManagementResponseConfigure{}
	return &this
}

// NewBotManagementResponseConfigureWithDefaults instantiates a new BotManagementResponseConfigure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementResponseConfigureWithDefaults() *BotManagementResponseConfigure {
	this := BotManagementResponseConfigure{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *BotManagementResponseConfigure) GetProduct() BotManagementResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret BotManagementResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseConfigure) GetProductOk() (*BotManagementResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *BotManagementResponseConfigure) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given BotManagementResponseProductProduct and assigns it to the Product field.
func (o *BotManagementResponseConfigure) SetProduct(v BotManagementResponseProductProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *BotManagementResponseConfigure) GetService() ApiDiscoveryResponseServiceService {
	if o == nil || o.Service == nil {
		var ret ApiDiscoveryResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseConfigure) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *BotManagementResponseConfigure) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given ApiDiscoveryResponseServiceService and assigns it to the Service field.
func (o *BotManagementResponseConfigure) SetService(v ApiDiscoveryResponseServiceService) {
	o.Service = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BotManagementResponseConfigure) GetConfiguration() BotManagementResponseConfigurationConfiguration {
	if o == nil || o.Configuration == nil {
		var ret BotManagementResponseConfigurationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseConfigure) GetConfigurationOk() (*BotManagementResponseConfigurationConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BotManagementResponseConfigure) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given BotManagementResponseConfigurationConfiguration and assigns it to the Configuration field.
func (o *BotManagementResponseConfigure) SetConfiguration(v BotManagementResponseConfigurationConfiguration) {
	o.Configuration = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BotManagementResponseConfigure) GetLinks() BotManagementResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret BotManagementResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseConfigure) GetLinksOk() (*BotManagementResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BotManagementResponseConfigure) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BotManagementResponseLinksLinks and assigns it to the Links field.
func (o *BotManagementResponseConfigure) SetLinks(v BotManagementResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementResponseConfigure) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementResponseConfigure) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementResponseConfigure := _BotManagementResponseConfigure{}

	if err = json.Unmarshal(bytes, &varBotManagementResponseConfigure); err == nil {
		*o = BotManagementResponseConfigure(varBotManagementResponseConfigure)
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

// NullableBotManagementResponseConfigure is a helper abstraction for handling nullable botmanagementresponseconfigure types.
type NullableBotManagementResponseConfigure struct {
	value *BotManagementResponseConfigure
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementResponseConfigure) Get() *BotManagementResponseConfigure {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementResponseConfigure) Set(val *BotManagementResponseConfigure) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementResponseConfigure) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementResponseConfigure) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementResponseConfigure returns a pointer to a new instance of NullableBotManagementResponseConfigure.
func NewNullableBotManagementResponseConfigure(val *BotManagementResponseConfigure) *NullableBotManagementResponseConfigure {
	return &NullableBotManagementResponseConfigure{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementResponseConfigure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementResponseConfigure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
