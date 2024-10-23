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

// ACLResponseAllOf struct for ACLResponseAllOf
type ACLResponseAllOf struct {
	ServiceID *string `json:"service_id,omitempty"`
	// String representing the number identifying a version of the service.
	Version              *string `json:"version,omitempty"`
	ID                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _ACLResponseAllOf ACLResponseAllOf

// NewACLResponseAllOf instantiates a new ACLResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLResponseAllOf() *ACLResponseAllOf {
	this := ACLResponseAllOf{}
	return &this
}

// NewACLResponseAllOfWithDefaults instantiates a new ACLResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLResponseAllOfWithDefaults() *ACLResponseAllOf {
	this := ACLResponseAllOf{}
	return &this
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ACLResponseAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLResponseAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *ACLResponseAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ACLResponseAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ACLResponseAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLResponseAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ACLResponseAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ACLResponseAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ACLResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ACLResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ACLResponseAllOf) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ACLResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ACLResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varACLResponseAllOf := _ACLResponseAllOf{}

	if err = json.Unmarshal(bytes, &varACLResponseAllOf); err == nil {
		*o = ACLResponseAllOf(varACLResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableACLResponseAllOf is a helper abstraction for handling nullable aclresponseallof types.
type NullableACLResponseAllOf struct {
	value *ACLResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableACLResponseAllOf) Get() *ACLResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableACLResponseAllOf) Set(val *ACLResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableACLResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableACLResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableACLResponseAllOf returns a pointer to a new instance of NullableACLResponseAllOf.
func NewNullableACLResponseAllOf(val *ACLResponseAllOf) *NullableACLResponseAllOf {
	return &NullableACLResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableACLResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableACLResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
