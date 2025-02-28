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

// AiAcceleratorResponseCustomerCustomer struct for AiAcceleratorResponseCustomerCustomer
type AiAcceleratorResponseCustomerCustomer struct {
	// Customer identifier
	ID *string `json:"id,omitempty"`
	// Name of the object
	Object               *string `json:"object,omitempty"`
	AdditionalProperties map[string]any
}

type _AiAcceleratorResponseCustomerCustomer AiAcceleratorResponseCustomerCustomer

// NewAiAcceleratorResponseCustomerCustomer instantiates a new AiAcceleratorResponseCustomerCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAcceleratorResponseCustomerCustomer() *AiAcceleratorResponseCustomerCustomer {
	this := AiAcceleratorResponseCustomerCustomer{}
	return &this
}

// NewAiAcceleratorResponseCustomerCustomerWithDefaults instantiates a new AiAcceleratorResponseCustomerCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAcceleratorResponseCustomerCustomerWithDefaults() *AiAcceleratorResponseCustomerCustomer {
	this := AiAcceleratorResponseCustomerCustomer{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *AiAcceleratorResponseCustomerCustomer) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAcceleratorResponseCustomerCustomer) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AiAcceleratorResponseCustomerCustomer) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *AiAcceleratorResponseCustomerCustomer) SetID(v string) {
	o.ID = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *AiAcceleratorResponseCustomerCustomer) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAcceleratorResponseCustomerCustomer) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *AiAcceleratorResponseCustomerCustomer) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *AiAcceleratorResponseCustomerCustomer) SetObject(v string) {
	o.Object = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AiAcceleratorResponseCustomerCustomer) MarshalJSON() ([]byte, error) {
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
func (o *AiAcceleratorResponseCustomerCustomer) UnmarshalJSON(bytes []byte) (err error) {
	varAiAcceleratorResponseCustomerCustomer := _AiAcceleratorResponseCustomerCustomer{}

	if err = json.Unmarshal(bytes, &varAiAcceleratorResponseCustomerCustomer); err == nil {
		*o = AiAcceleratorResponseCustomerCustomer(varAiAcceleratorResponseCustomerCustomer)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAiAcceleratorResponseCustomerCustomer is a helper abstraction for handling nullable aiacceleratorresponsecustomercustomer types.
type NullableAiAcceleratorResponseCustomerCustomer struct {
	value *AiAcceleratorResponseCustomerCustomer
	isSet bool
}

// Get returns the value.
func (v NullableAiAcceleratorResponseCustomerCustomer) Get() *AiAcceleratorResponseCustomerCustomer {
	return v.value
}

// Set modifies the value.
func (v *NullableAiAcceleratorResponseCustomerCustomer) Set(val *AiAcceleratorResponseCustomerCustomer) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAiAcceleratorResponseCustomerCustomer) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAiAcceleratorResponseCustomerCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiAcceleratorResponseCustomerCustomer returns a pointer to a new instance of NullableAiAcceleratorResponseCustomerCustomer.
func NewNullableAiAcceleratorResponseCustomerCustomer(val *AiAcceleratorResponseCustomerCustomer) *NullableAiAcceleratorResponseCustomerCustomer {
	return &NullableAiAcceleratorResponseCustomerCustomer{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAiAcceleratorResponseCustomerCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAiAcceleratorResponseCustomerCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
