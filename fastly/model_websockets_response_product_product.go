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

// WebsocketsResponseProductProduct struct for WebsocketsResponseProductProduct
type WebsocketsResponseProductProduct struct {
	// Product identifier
	ID *string `json:"id,omitempty"`
	// Name of the object
	Object               *string `json:"object,omitempty"`
	AdditionalProperties map[string]any
}

type _WebsocketsResponseProductProduct WebsocketsResponseProductProduct

// NewWebsocketsResponseProductProduct instantiates a new WebsocketsResponseProductProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketsResponseProductProduct() *WebsocketsResponseProductProduct {
	this := WebsocketsResponseProductProduct{}
	return &this
}

// NewWebsocketsResponseProductProductWithDefaults instantiates a new WebsocketsResponseProductProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketsResponseProductProductWithDefaults() *WebsocketsResponseProductProduct {
	this := WebsocketsResponseProductProduct{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WebsocketsResponseProductProduct) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsocketsResponseProductProduct) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WebsocketsResponseProductProduct) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WebsocketsResponseProductProduct) SetID(v string) {
	o.ID = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *WebsocketsResponseProductProduct) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsocketsResponseProductProduct) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *WebsocketsResponseProductProduct) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *WebsocketsResponseProductProduct) SetObject(v string) {
	o.Object = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WebsocketsResponseProductProduct) MarshalJSON() ([]byte, error) {
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
func (o *WebsocketsResponseProductProduct) UnmarshalJSON(bytes []byte) (err error) {
	varWebsocketsResponseProductProduct := _WebsocketsResponseProductProduct{}

	if err = json.Unmarshal(bytes, &varWebsocketsResponseProductProduct); err == nil {
		*o = WebsocketsResponseProductProduct(varWebsocketsResponseProductProduct)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWebsocketsResponseProductProduct is a helper abstraction for handling nullable websocketsresponseproductproduct types.
type NullableWebsocketsResponseProductProduct struct {
	value *WebsocketsResponseProductProduct
	isSet bool
}

// Get returns the value.
func (v NullableWebsocketsResponseProductProduct) Get() *WebsocketsResponseProductProduct {
	return v.value
}

// Set modifies the value.
func (v *NullableWebsocketsResponseProductProduct) Set(val *WebsocketsResponseProductProduct) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWebsocketsResponseProductProduct) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWebsocketsResponseProductProduct) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWebsocketsResponseProductProduct returns a pointer to a new instance of NullableWebsocketsResponseProductProduct.
func NewNullableWebsocketsResponseProductProduct(val *WebsocketsResponseProductProduct) *NullableWebsocketsResponseProductProduct {
	return &NullableWebsocketsResponseProductProduct{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWebsocketsResponseProductProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWebsocketsResponseProductProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
