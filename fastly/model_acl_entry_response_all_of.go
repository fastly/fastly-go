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

// AclEntryResponseAllOf struct for AclEntryResponseAllOf
type AclEntryResponseAllOf struct {
	AclId                *string `json:"acl_id,omitempty"`
	Id                   *string `json:"id,omitempty"`
	ServiceId            *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _AclEntryResponseAllOf AclEntryResponseAllOf

// NewAclEntryResponseAllOf instantiates a new AclEntryResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclEntryResponseAllOf() *AclEntryResponseAllOf {
	this := AclEntryResponseAllOf{}
	return &this
}

// NewAclEntryResponseAllOfWithDefaults instantiates a new AclEntryResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclEntryResponseAllOfWithDefaults() *AclEntryResponseAllOf {
	this := AclEntryResponseAllOf{}
	return &this
}

// GetAclId returns the AclId field value if set, zero value otherwise.
func (o *AclEntryResponseAllOf) GetAclId() string {
	if o == nil || o.AclId == nil {
		var ret string
		return ret
	}
	return *o.AclId
}

// GetAclIdOk returns a tuple with the AclId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponseAllOf) GetAclIdOk() (*string, bool) {
	if o == nil || o.AclId == nil {
		return nil, false
	}
	return o.AclId, true
}

// HasAclId returns a boolean if a field has been set.
func (o *AclEntryResponseAllOf) HasAclId() bool {
	if o != nil && o.AclId != nil {
		return true
	}

	return false
}

// SetAclId gets a reference to the given string and assigns it to the AclId field.
func (o *AclEntryResponseAllOf) SetAclId(v string) {
	o.AclId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AclEntryResponseAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AclEntryResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AclEntryResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *AclEntryResponseAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponseAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *AclEntryResponseAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *AclEntryResponseAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AclEntryResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AclId != nil {
		toSerialize["acl_id"] = o.AclId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AclEntryResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAclEntryResponseAllOf := _AclEntryResponseAllOf{}

	if err = json.Unmarshal(bytes, &varAclEntryResponseAllOf); err == nil {
		*o = AclEntryResponseAllOf(varAclEntryResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "acl_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAclEntryResponseAllOf is a helper abstraction for handling nullable aclentryresponseallof types.
type NullableAclEntryResponseAllOf struct {
	value *AclEntryResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableAclEntryResponseAllOf) Get() *AclEntryResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableAclEntryResponseAllOf) Set(val *AclEntryResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAclEntryResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAclEntryResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAclEntryResponseAllOf returns a pointer to a new instance of NullableAclEntryResponseAllOf.
func NewNullableAclEntryResponseAllOf(val *AclEntryResponseAllOf) *NullableAclEntryResponseAllOf {
	return &NullableAclEntryResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAclEntryResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAclEntryResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
