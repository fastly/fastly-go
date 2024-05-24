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
	ServiceID *string `json:"service_id,omitempty"`
	ID *string `json:"id,omitempty"`
	PoolID *string `json:"pool_id,omitempty"`
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

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ServerResponseAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponseAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *ServerResponseAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ServerResponseAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServerResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServerResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServerResponseAllOf) SetID(v string) {
	o.ID = &v
}

// GetPoolID returns the PoolID field value if set, zero value otherwise.
func (o *ServerResponseAllOf) GetPoolID() string {
	if o == nil || o.PoolID == nil {
		var ret string
		return ret
	}
	return *o.PoolID
}

// GetPoolIDOk returns a tuple with the PoolID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponseAllOf) GetPoolIDOk() (*string, bool) {
	if o == nil || o.PoolID == nil {
		return nil, false
	}
	return o.PoolID, true
}

// HasPoolID returns a boolean if a field has been set.
func (o *ServerResponseAllOf) HasPoolID() bool {
	if o != nil && o.PoolID != nil {
		return true
	}

	return false
}

// SetPoolID gets a reference to the given string and assigns it to the PoolID field.
func (o *ServerResponseAllOf) SetPoolID(v string) {
	o.PoolID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServerResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.PoolID != nil {
		toSerialize["pool_id"] = o.PoolID
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
