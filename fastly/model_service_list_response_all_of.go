// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// ServiceListResponseAllOf struct for ServiceListResponseAllOf
type ServiceListResponseAllOf struct {
	ID *string `json:"id,omitempty"`
	// Current [version](/reference/api/services/version/) of the service.
	Version *int32 `json:"version,omitempty"`
	// A list of [versions](/reference/api/services/version/) associated with the service.
	Versions []SchemasVersionResponse `json:"versions,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceListResponseAllOf ServiceListResponseAllOf

// NewServiceListResponseAllOf instantiates a new ServiceListResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceListResponseAllOf() *ServiceListResponseAllOf {
	this := ServiceListResponseAllOf{}
	return &this
}

// NewServiceListResponseAllOfWithDefaults instantiates a new ServiceListResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceListResponseAllOfWithDefaults() *ServiceListResponseAllOf {
	this := ServiceListResponseAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServiceListResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceListResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServiceListResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServiceListResponseAllOf) SetID(v string) {
	o.ID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceListResponseAllOf) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceListResponseAllOf) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceListResponseAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ServiceListResponseAllOf) SetVersion(v int32) {
	o.Version = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ServiceListResponseAllOf) GetVersions() []SchemasVersionResponse {
	if o == nil || o.Versions == nil {
		var ret []SchemasVersionResponse
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceListResponseAllOf) GetVersionsOk() ([]SchemasVersionResponse, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ServiceListResponseAllOf) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []SchemasVersionResponse and assigns it to the Versions field.
func (o *ServiceListResponseAllOf) SetVersions(v []SchemasVersionResponse) {
	o.Versions = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceListResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceListResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServiceListResponseAllOf := _ServiceListResponseAllOf{}

	if err = json.Unmarshal(bytes, &varServiceListResponseAllOf); err == nil {
		*o = ServiceListResponseAllOf(varServiceListResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "versions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceListResponseAllOf is a helper abstraction for handling nullable servicelistresponseallof types. 
type NullableServiceListResponseAllOf struct {
	value *ServiceListResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableServiceListResponseAllOf) Get() *ServiceListResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceListResponseAllOf) Set(val *ServiceListResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceListResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceListResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceListResponseAllOf returns a pointer to a new instance of NullableServiceListResponseAllOf.
func NewNullableServiceListResponseAllOf(val *ServiceListResponseAllOf) *NullableServiceListResponseAllOf {
	return &NullableServiceListResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceListResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceListResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
