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

// ServiceIDAndVersionString struct for ServiceIDAndVersionString
type ServiceIDAndVersionString struct {
	ServiceID *string `json:"service_id,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceIDAndVersionString ServiceIDAndVersionString

// NewServiceIDAndVersionString instantiates a new ServiceIDAndVersionString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceIDAndVersionString() *ServiceIDAndVersionString {
	this := ServiceIDAndVersionString{}
	return &this
}

// NewServiceIDAndVersionStringWithDefaults instantiates a new ServiceIDAndVersionString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceIDAndVersionStringWithDefaults() *ServiceIDAndVersionString {
	this := ServiceIDAndVersionString{}
	return &this
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ServiceIDAndVersionString) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceIDAndVersionString) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *ServiceIDAndVersionString) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ServiceIDAndVersionString) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceIDAndVersionString) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceIDAndVersionString) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceIDAndVersionString) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ServiceIDAndVersionString) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceIDAndVersionString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceIDAndVersionString) UnmarshalJSON(bytes []byte) (err error) {
	varServiceIDAndVersionString := _ServiceIDAndVersionString{}

	if err = json.Unmarshal(bytes, &varServiceIDAndVersionString); err == nil {
		*o = ServiceIDAndVersionString(varServiceIDAndVersionString)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceIDAndVersionString is a helper abstraction for handling nullable serviceidandversionstring types. 
type NullableServiceIDAndVersionString struct {
	value *ServiceIDAndVersionString
	isSet bool
}

// Get returns the value.
func (v NullableServiceIDAndVersionString) Get() *ServiceIDAndVersionString {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceIDAndVersionString) Set(val *ServiceIDAndVersionString) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceIDAndVersionString) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceIDAndVersionString) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceIDAndVersionString returns a pointer to a new instance of NullableServiceIDAndVersionString.
func NewNullableServiceIDAndVersionString(val *ServiceIDAndVersionString) *NullableServiceIDAndVersionString {
	return &NullableServiceIDAndVersionString{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceIDAndVersionString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceIDAndVersionString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
