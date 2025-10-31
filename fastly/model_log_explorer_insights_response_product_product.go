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

// LogExplorerInsightsResponseProductProduct struct for LogExplorerInsightsResponseProductProduct
type LogExplorerInsightsResponseProductProduct struct {
	// Product identifier
	Id *string `json:"id,omitempty"`
	// Name of the object
	Object               *string `json:"object,omitempty"`
	AdditionalProperties map[string]any
}

type _LogExplorerInsightsResponseProductProduct LogExplorerInsightsResponseProductProduct

// NewLogExplorerInsightsResponseProductProduct instantiates a new LogExplorerInsightsResponseProductProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExplorerInsightsResponseProductProduct() *LogExplorerInsightsResponseProductProduct {
	this := LogExplorerInsightsResponseProductProduct{}
	return &this
}

// NewLogExplorerInsightsResponseProductProductWithDefaults instantiates a new LogExplorerInsightsResponseProductProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExplorerInsightsResponseProductProductWithDefaults() *LogExplorerInsightsResponseProductProduct {
	this := LogExplorerInsightsResponseProductProduct{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseProductProduct) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseProductProduct) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseProductProduct) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogExplorerInsightsResponseProductProduct) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseProductProduct) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseProductProduct) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseProductProduct) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *LogExplorerInsightsResponseProductProduct) SetObject(v string) {
	o.Object = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogExplorerInsightsResponseProductProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogExplorerInsightsResponseProductProduct) UnmarshalJSON(bytes []byte) (err error) {
	varLogExplorerInsightsResponseProductProduct := _LogExplorerInsightsResponseProductProduct{}

	if err = json.Unmarshal(bytes, &varLogExplorerInsightsResponseProductProduct); err == nil {
		*o = LogExplorerInsightsResponseProductProduct(varLogExplorerInsightsResponseProductProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogExplorerInsightsResponseProductProduct is a helper abstraction for handling nullable logexplorerinsightsresponseproductproduct types.
type NullableLogExplorerInsightsResponseProductProduct struct {
	value *LogExplorerInsightsResponseProductProduct
	isSet bool
}

// Get returns the value.
func (v NullableLogExplorerInsightsResponseProductProduct) Get() *LogExplorerInsightsResponseProductProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableLogExplorerInsightsResponseProductProduct) Set(val *LogExplorerInsightsResponseProductProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogExplorerInsightsResponseProductProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogExplorerInsightsResponseProductProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogExplorerInsightsResponseProductProduct returns a pointer to a new instance of NullableLogExplorerInsightsResponseProductProduct.
func NewNullableLogExplorerInsightsResponseProductProduct(val *LogExplorerInsightsResponseProductProduct) *NullableLogExplorerInsightsResponseProductProduct {
	return &NullableLogExplorerInsightsResponseProductProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogExplorerInsightsResponseProductProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogExplorerInsightsResponseProductProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
