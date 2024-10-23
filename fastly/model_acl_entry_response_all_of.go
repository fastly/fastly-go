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

// ACLEntryResponseAllOf struct for ACLEntryResponseAllOf
type ACLEntryResponseAllOf struct {
	ACLID                *string `json:"acl_id,omitempty"`
	ID                   *string `json:"id,omitempty"`
	ServiceID            *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _ACLEntryResponseAllOf ACLEntryResponseAllOf

// NewACLEntryResponseAllOf instantiates a new ACLEntryResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLEntryResponseAllOf() *ACLEntryResponseAllOf {
	this := ACLEntryResponseAllOf{}
	return &this
}

// NewACLEntryResponseAllOfWithDefaults instantiates a new ACLEntryResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLEntryResponseAllOfWithDefaults() *ACLEntryResponseAllOf {
	this := ACLEntryResponseAllOf{}
	return &this
}

// GetACLID returns the ACLID field value if set, zero value otherwise.
func (o *ACLEntryResponseAllOf) GetACLID() string {
	if o == nil || o.ACLID == nil {
		var ret string
		return ret
	}
	return *o.ACLID
}

// GetACLIDOk returns a tuple with the ACLID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponseAllOf) GetACLIDOk() (*string, bool) {
	if o == nil || o.ACLID == nil {
		return nil, false
	}
	return o.ACLID, true
}

// HasACLID returns a boolean if a field has been set.
func (o *ACLEntryResponseAllOf) HasACLID() bool {
	if o != nil && o.ACLID != nil {
		return true
	}

	return false
}

// SetACLID gets a reference to the given string and assigns it to the ACLID field.
func (o *ACLEntryResponseAllOf) SetACLID(v string) {
	o.ACLID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ACLEntryResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ACLEntryResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ACLEntryResponseAllOf) SetID(v string) {
	o.ID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ACLEntryResponseAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponseAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *ACLEntryResponseAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ACLEntryResponseAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ACLEntryResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ACLID != nil {
		toSerialize["acl_id"] = o.ACLID
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ACLEntryResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varACLEntryResponseAllOf := _ACLEntryResponseAllOf{}

	if err = json.Unmarshal(bytes, &varACLEntryResponseAllOf); err == nil {
		*o = ACLEntryResponseAllOf(varACLEntryResponseAllOf)
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

// NullableACLEntryResponseAllOf is a helper abstraction for handling nullable aclentryresponseallof types.
type NullableACLEntryResponseAllOf struct {
	value *ACLEntryResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableACLEntryResponseAllOf) Get() *ACLEntryResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableACLEntryResponseAllOf) Set(val *ACLEntryResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableACLEntryResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableACLEntryResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableACLEntryResponseAllOf returns a pointer to a new instance of NullableACLEntryResponseAllOf.
func NewNullableACLEntryResponseAllOf(val *ACLEntryResponseAllOf) *NullableACLEntryResponseAllOf {
	return &NullableACLEntryResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableACLEntryResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableACLEntryResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
