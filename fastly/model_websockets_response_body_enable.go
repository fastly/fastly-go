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

// WebsocketsResponseBodyEnable struct for WebsocketsResponseBodyEnable
type WebsocketsResponseBodyEnable struct {
	Product              *WebsocketsResponseProductProduct    `json:"product,omitempty"`
	Service              *BotManagementResponseServiceService `json:"service,omitempty"`
	Links                *WebsocketsResponseLinksLinks        `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _WebsocketsResponseBodyEnable WebsocketsResponseBodyEnable

// NewWebsocketsResponseBodyEnable instantiates a new WebsocketsResponseBodyEnable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketsResponseBodyEnable() *WebsocketsResponseBodyEnable {
	this := WebsocketsResponseBodyEnable{}
	return &this
}

// NewWebsocketsResponseBodyEnableWithDefaults instantiates a new WebsocketsResponseBodyEnable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketsResponseBodyEnableWithDefaults() *WebsocketsResponseBodyEnable {
	this := WebsocketsResponseBodyEnable{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *WebsocketsResponseBodyEnable) GetProduct() WebsocketsResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret WebsocketsResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsocketsResponseBodyEnable) GetProductOk() (*WebsocketsResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *WebsocketsResponseBodyEnable) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given WebsocketsResponseProductProduct and assigns it to the Product field.
func (o *WebsocketsResponseBodyEnable) SetProduct(v WebsocketsResponseProductProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *WebsocketsResponseBodyEnable) GetService() BotManagementResponseServiceService {
	if o == nil || o.Service == nil {
		var ret BotManagementResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsocketsResponseBodyEnable) GetServiceOk() (*BotManagementResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *WebsocketsResponseBodyEnable) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given BotManagementResponseServiceService and assigns it to the Service field.
func (o *WebsocketsResponseBodyEnable) SetService(v BotManagementResponseServiceService) {
	o.Service = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebsocketsResponseBodyEnable) GetLinks() WebsocketsResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret WebsocketsResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsocketsResponseBodyEnable) GetLinksOk() (*WebsocketsResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebsocketsResponseBodyEnable) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WebsocketsResponseLinksLinks and assigns it to the Links field.
func (o *WebsocketsResponseBodyEnable) SetLinks(v WebsocketsResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WebsocketsResponseBodyEnable) MarshalJSON() ([]byte, error) {
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
func (o *WebsocketsResponseBodyEnable) UnmarshalJSON(bytes []byte) (err error) {
	varWebsocketsResponseBodyEnable := _WebsocketsResponseBodyEnable{}

	if err = json.Unmarshal(bytes, &varWebsocketsResponseBodyEnable); err == nil {
		*o = WebsocketsResponseBodyEnable(varWebsocketsResponseBodyEnable)
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

// NullableWebsocketsResponseBodyEnable is a helper abstraction for handling nullable websocketsresponsebodyenable types.
type NullableWebsocketsResponseBodyEnable struct {
	value *WebsocketsResponseBodyEnable
	isSet bool
}

// Get returns the value.
func (v NullableWebsocketsResponseBodyEnable) Get() *WebsocketsResponseBodyEnable {
	return v.value
}

// Set modifies the value.
func (v *NullableWebsocketsResponseBodyEnable) Set(val *WebsocketsResponseBodyEnable) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWebsocketsResponseBodyEnable) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWebsocketsResponseBodyEnable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWebsocketsResponseBodyEnable returns a pointer to a new instance of NullableWebsocketsResponseBodyEnable.
func NewNullableWebsocketsResponseBodyEnable(val *WebsocketsResponseBodyEnable) *NullableWebsocketsResponseBodyEnable {
	return &NullableWebsocketsResponseBodyEnable{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWebsocketsResponseBodyEnable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWebsocketsResponseBodyEnable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
