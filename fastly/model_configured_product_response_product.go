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

// ConfiguredProductResponseProduct struct for ConfiguredProductResponseProduct
type ConfiguredProductResponseProduct struct {
	// Product identifier
	ID *string `json:"id,omitempty"`
	// Name of the object
	Object *string `json:"object,omitempty"`
	AdditionalProperties map[string]any
}

type _ConfiguredProductResponseProduct ConfiguredProductResponseProduct

// NewConfiguredProductResponseProduct instantiates a new ConfiguredProductResponseProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguredProductResponseProduct() *ConfiguredProductResponseProduct {
	this := ConfiguredProductResponseProduct{}
	return &this
}

// NewConfiguredProductResponseProductWithDefaults instantiates a new ConfiguredProductResponseProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguredProductResponseProductWithDefaults() *ConfiguredProductResponseProduct {
	this := ConfiguredProductResponseProduct{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ConfiguredProductResponseProduct) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponseProduct) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ConfiguredProductResponseProduct) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ConfiguredProductResponseProduct) SetID(v string) {
	o.ID = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ConfiguredProductResponseProduct) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponseProduct) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ConfiguredProductResponseProduct) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ConfiguredProductResponseProduct) SetObject(v string) {
	o.Object = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ConfiguredProductResponseProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *ConfiguredProductResponseProduct) UnmarshalJSON(bytes []byte) (err error) {
	varConfiguredProductResponseProduct := _ConfiguredProductResponseProduct{}

	if err = json.Unmarshal(bytes, &varConfiguredProductResponseProduct); err == nil {
		*o = ConfiguredProductResponseProduct(varConfiguredProductResponseProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableConfiguredProductResponseProduct is a helper abstraction for handling nullable configuredproductresponseproduct types. 
type NullableConfiguredProductResponseProduct struct {
	value *ConfiguredProductResponseProduct
	isSet bool
}

// Get returns the value.
func (v NullableConfiguredProductResponseProduct) Get() *ConfiguredProductResponseProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableConfiguredProductResponseProduct) Set(val *ConfiguredProductResponseProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConfiguredProductResponseProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConfiguredProductResponseProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConfiguredProductResponseProduct returns a pointer to a new instance of NullableConfiguredProductResponseProduct.
func NewNullableConfiguredProductResponseProduct(val *ConfiguredProductResponseProduct) *NullableConfiguredProductResponseProduct {
	return &NullableConfiguredProductResponseProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConfiguredProductResponseProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableConfiguredProductResponseProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
