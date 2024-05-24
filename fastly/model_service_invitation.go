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

// ServiceInvitation struct for ServiceInvitation
type ServiceInvitation struct {
	Data *ServiceInvitationData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceInvitation ServiceInvitation

// NewServiceInvitation instantiates a new ServiceInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInvitation() *ServiceInvitation {
	this := ServiceInvitation{}
	return &this
}

// NewServiceInvitationWithDefaults instantiates a new ServiceInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInvitationWithDefaults() *ServiceInvitation {
	this := ServiceInvitation{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceInvitation) GetData() ServiceInvitationData {
	if o == nil || o.Data == nil {
		var ret ServiceInvitationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInvitation) GetDataOk() (*ServiceInvitationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceInvitation) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ServiceInvitationData and assigns it to the Data field.
func (o *ServiceInvitation) SetData(v ServiceInvitationData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceInvitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceInvitation) UnmarshalJSON(bytes []byte) (err error) {
	varServiceInvitation := _ServiceInvitation{}

	if err = json.Unmarshal(bytes, &varServiceInvitation); err == nil {
		*o = ServiceInvitation(varServiceInvitation)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceInvitation is a helper abstraction for handling nullable serviceinvitation types. 
type NullableServiceInvitation struct {
	value *ServiceInvitation
	isSet bool
}

// Get returns the value.
func (v NullableServiceInvitation) Get() *ServiceInvitation {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceInvitation) Set(val *ServiceInvitation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceInvitation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceInvitation returns a pointer to a new instance of NullableServiceInvitation.
func NewNullableServiceInvitation(val *ServiceInvitation) *NullableServiceInvitation {
	return &NullableServiceInvitation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
