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

// KvStoreDetails struct for KvStoreDetails
type KvStoreDetails struct {
	// ID of the store.
	Id *string `json:"id,omitempty"`
	// Name of the store.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _KvStoreDetails KvStoreDetails

// NewKvStoreDetails instantiates a new KvStoreDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvStoreDetails() *KvStoreDetails {
	this := KvStoreDetails{}
	return &this
}

// NewKvStoreDetailsWithDefaults instantiates a new KvStoreDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvStoreDetailsWithDefaults() *KvStoreDetails {
	this := KvStoreDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KvStoreDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KvStoreDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KvStoreDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KvStoreDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KvStoreDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KvStoreDetails) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o KvStoreDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *KvStoreDetails) UnmarshalJSON(bytes []byte) (err error) {
	varKvStoreDetails := _KvStoreDetails{}

	if err = json.Unmarshal(bytes, &varKvStoreDetails); err == nil {
		*o = KvStoreDetails(varKvStoreDetails)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableKvStoreDetails is a helper abstraction for handling nullable kvstoredetails types.
type NullableKvStoreDetails struct {
	value *KvStoreDetails
	isSet bool
}

// Get returns the value.
func (v NullableKvStoreDetails) Get() *KvStoreDetails {
	return v.value
}

// Set modifies the value.
func (v *NullableKvStoreDetails) Set(val *KvStoreDetails) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableKvStoreDetails) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableKvStoreDetails) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableKvStoreDetails returns a pointer to a new instance of NullableKvStoreDetails.
func NewNullableKvStoreDetails(val *KvStoreDetails) *NullableKvStoreDetails {
	return &NullableKvStoreDetails{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableKvStoreDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableKvStoreDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
