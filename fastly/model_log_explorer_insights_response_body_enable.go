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

// LogExplorerInsightsResponseBodyEnable struct for LogExplorerInsightsResponseBodyEnable
type LogExplorerInsightsResponseBodyEnable struct {
	Product              *LogExplorerInsightsResponseProductProduct `json:"product,omitempty"`
	Service              *BotManagementResponseServiceService       `json:"service,omitempty"`
	Links                *LogExplorerInsightsResponseLinksLinks     `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _LogExplorerInsightsResponseBodyEnable LogExplorerInsightsResponseBodyEnable

// NewLogExplorerInsightsResponseBodyEnable instantiates a new LogExplorerInsightsResponseBodyEnable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExplorerInsightsResponseBodyEnable() *LogExplorerInsightsResponseBodyEnable {
	this := LogExplorerInsightsResponseBodyEnable{}
	return &this
}

// NewLogExplorerInsightsResponseBodyEnableWithDefaults instantiates a new LogExplorerInsightsResponseBodyEnable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExplorerInsightsResponseBodyEnableWithDefaults() *LogExplorerInsightsResponseBodyEnable {
	this := LogExplorerInsightsResponseBodyEnable{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseBodyEnable) GetProduct() LogExplorerInsightsResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret LogExplorerInsightsResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseBodyEnable) GetProductOk() (*LogExplorerInsightsResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseBodyEnable) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given LogExplorerInsightsResponseProductProduct and assigns it to the Product field.
func (o *LogExplorerInsightsResponseBodyEnable) SetProduct(v LogExplorerInsightsResponseProductProduct) {
	o.Product = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseBodyEnable) GetService() BotManagementResponseServiceService {
	if o == nil || o.Service == nil {
		var ret BotManagementResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseBodyEnable) GetServiceOk() (*BotManagementResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseBodyEnable) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given BotManagementResponseServiceService and assigns it to the Service field.
func (o *LogExplorerInsightsResponseBodyEnable) SetService(v BotManagementResponseServiceService) {
	o.Service = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseBodyEnable) GetLinks() LogExplorerInsightsResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret LogExplorerInsightsResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseBodyEnable) GetLinksOk() (*LogExplorerInsightsResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseBodyEnable) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LogExplorerInsightsResponseLinksLinks and assigns it to the Links field.
func (o *LogExplorerInsightsResponseBodyEnable) SetLinks(v LogExplorerInsightsResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogExplorerInsightsResponseBodyEnable) MarshalJSON() ([]byte, error) {
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
func (o *LogExplorerInsightsResponseBodyEnable) UnmarshalJSON(bytes []byte) (err error) {
	varLogExplorerInsightsResponseBodyEnable := _LogExplorerInsightsResponseBodyEnable{}

	if err = json.Unmarshal(bytes, &varLogExplorerInsightsResponseBodyEnable); err == nil {
		*o = LogExplorerInsightsResponseBodyEnable(varLogExplorerInsightsResponseBodyEnable)
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

// NullableLogExplorerInsightsResponseBodyEnable is a helper abstraction for handling nullable logexplorerinsightsresponsebodyenable types.
type NullableLogExplorerInsightsResponseBodyEnable struct {
	value *LogExplorerInsightsResponseBodyEnable
	isSet bool
}

// Get returns the value.
func (v NullableLogExplorerInsightsResponseBodyEnable) Get() *LogExplorerInsightsResponseBodyEnable {
	return v.value
}

// Set modifies the value.
func (v *NullableLogExplorerInsightsResponseBodyEnable) Set(val *LogExplorerInsightsResponseBodyEnable) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogExplorerInsightsResponseBodyEnable) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogExplorerInsightsResponseBodyEnable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogExplorerInsightsResponseBodyEnable returns a pointer to a new instance of NullableLogExplorerInsightsResponseBodyEnable.
func NewNullableLogExplorerInsightsResponseBodyEnable(val *LogExplorerInsightsResponseBodyEnable) *NullableLogExplorerInsightsResponseBodyEnable {
	return &NullableLogExplorerInsightsResponseBodyEnable{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogExplorerInsightsResponseBodyEnable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogExplorerInsightsResponseBodyEnable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
