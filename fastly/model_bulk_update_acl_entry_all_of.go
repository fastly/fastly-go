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

// BulkUpdateAclEntryAllOf struct for BulkUpdateAclEntryAllOf
type BulkUpdateAclEntryAllOf struct {
	Op                   *string `json:"op,omitempty"`
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateAclEntryAllOf BulkUpdateAclEntryAllOf

// NewBulkUpdateAclEntryAllOf instantiates a new BulkUpdateAclEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateAclEntryAllOf() *BulkUpdateAclEntryAllOf {
	this := BulkUpdateAclEntryAllOf{}
	return &this
}

// NewBulkUpdateAclEntryAllOfWithDefaults instantiates a new BulkUpdateAclEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateAclEntryAllOfWithDefaults() *BulkUpdateAclEntryAllOf {
	this := BulkUpdateAclEntryAllOf{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *BulkUpdateAclEntryAllOf) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAclEntryAllOf) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *BulkUpdateAclEntryAllOf) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *BulkUpdateAclEntryAllOf) SetOp(v string) {
	o.Op = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkUpdateAclEntryAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAclEntryAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkUpdateAclEntryAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BulkUpdateAclEntryAllOf) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateAclEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BulkUpdateAclEntryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateAclEntryAllOf := _BulkUpdateAclEntryAllOf{}

	if err = json.Unmarshal(bytes, &varBulkUpdateAclEntryAllOf); err == nil {
		*o = BulkUpdateAclEntryAllOf(varBulkUpdateAclEntryAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateAclEntryAllOf is a helper abstraction for handling nullable bulkupdateaclentryallof types.
type NullableBulkUpdateAclEntryAllOf struct {
	value *BulkUpdateAclEntryAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateAclEntryAllOf) Get() *BulkUpdateAclEntryAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateAclEntryAllOf) Set(val *BulkUpdateAclEntryAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateAclEntryAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateAclEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateAclEntryAllOf returns a pointer to a new instance of NullableBulkUpdateAclEntryAllOf.
func NewNullableBulkUpdateAclEntryAllOf(val *BulkUpdateAclEntryAllOf) *NullableBulkUpdateAclEntryAllOf {
	return &NullableBulkUpdateAclEntryAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateAclEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBulkUpdateAclEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
