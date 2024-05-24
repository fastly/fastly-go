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

// ServiceInvitationResponseAllOfData struct for ServiceInvitationResponseAllOfData
type ServiceInvitationResponseAllOfData struct {
	ID *string `json:"id,omitempty"`
	Attributes *Timestamps `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceInvitationResponseAllOfData ServiceInvitationResponseAllOfData

// NewServiceInvitationResponseAllOfData instantiates a new ServiceInvitationResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInvitationResponseAllOfData() *ServiceInvitationResponseAllOfData {
	this := ServiceInvitationResponseAllOfData{}
	return &this
}

// NewServiceInvitationResponseAllOfDataWithDefaults instantiates a new ServiceInvitationResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInvitationResponseAllOfDataWithDefaults() *ServiceInvitationResponseAllOfData {
	this := ServiceInvitationResponseAllOfData{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServiceInvitationResponseAllOfData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInvitationResponseAllOfData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServiceInvitationResponseAllOfData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServiceInvitationResponseAllOfData) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServiceInvitationResponseAllOfData) GetAttributes() Timestamps {
	if o == nil || o.Attributes == nil {
		var ret Timestamps
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInvitationResponseAllOfData) GetAttributesOk() (*Timestamps, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServiceInvitationResponseAllOfData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Timestamps and assigns it to the Attributes field.
func (o *ServiceInvitationResponseAllOfData) SetAttributes(v Timestamps) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceInvitationResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceInvitationResponseAllOfData) UnmarshalJSON(bytes []byte) (err error) {
	varServiceInvitationResponseAllOfData := _ServiceInvitationResponseAllOfData{}

	if err = json.Unmarshal(bytes, &varServiceInvitationResponseAllOfData); err == nil {
		*o = ServiceInvitationResponseAllOfData(varServiceInvitationResponseAllOfData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceInvitationResponseAllOfData is a helper abstraction for handling nullable serviceinvitationresponseallofdata types. 
type NullableServiceInvitationResponseAllOfData struct {
	value *ServiceInvitationResponseAllOfData
	isSet bool
}

// Get returns the value.
func (v NullableServiceInvitationResponseAllOfData) Get() *ServiceInvitationResponseAllOfData {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceInvitationResponseAllOfData) Set(val *ServiceInvitationResponseAllOfData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceInvitationResponseAllOfData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceInvitationResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceInvitationResponseAllOfData returns a pointer to a new instance of NullableServiceInvitationResponseAllOfData.
func NewNullableServiceInvitationResponseAllOfData(val *ServiceInvitationResponseAllOfData) *NullableServiceInvitationResponseAllOfData {
	return &NullableServiceInvitationResponseAllOfData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceInvitationResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceInvitationResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
