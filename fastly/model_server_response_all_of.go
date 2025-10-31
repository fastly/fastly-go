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

// ServerResponseAllOf struct for ServerResponseAllOf
type ServerResponseAllOf struct {
	ServiceId            *string `json:"service_id,omitempty"`
	Id                   *string `json:"id,omitempty"`
	PoolId               *string `json:"pool_id,omitempty"`
	AdditionalProperties map[string]any
}

type _ServerResponseAllOf ServerResponseAllOf

// NewServerResponseAllOf instantiates a new ServerResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerResponseAllOf() *ServerResponseAllOf {
	this := ServerResponseAllOf{}
	return &this
}

// NewServerResponseAllOfWithDefaults instantiates a new ServerResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerResponseAllOfWithDefaults() *ServerResponseAllOf {
	this := ServerResponseAllOf{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ServerResponseAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponseAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ServerResponseAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ServerResponseAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerResponseAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *ServerResponseAllOf) GetPoolId() string {
	if o == nil || o.PoolId == nil {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponseAllOf) GetPoolIdOk() (*string, bool) {
	if o == nil || o.PoolId == nil {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *ServerResponseAllOf) HasPoolId() bool {
	if o != nil && o.PoolId != nil {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *ServerResponseAllOf) SetPoolId(v string) {
	o.PoolId = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServerResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PoolId != nil {
		toSerialize["pool_id"] = o.PoolId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ServerResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerResponseAllOf := _ServerResponseAllOf{}

	if err = json.Unmarshal(bytes, &varServerResponseAllOf); err == nil {
		*o = ServerResponseAllOf(varServerResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "pool_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServerResponseAllOf is a helper abstraction for handling nullable serverresponseallof types.
type NullableServerResponseAllOf struct {
	value *ServerResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableServerResponseAllOf) Get() *ServerResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableServerResponseAllOf) Set(val *ServerResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServerResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServerResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServerResponseAllOf returns a pointer to a new instance of NullableServerResponseAllOf.
func NewNullableServerResponseAllOf(val *ServerResponseAllOf) *NullableServerResponseAllOf {
	return &NullableServerResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServerResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServerResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
