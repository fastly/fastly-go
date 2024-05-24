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

// BulkUpdateACLEntryAllOf struct for BulkUpdateACLEntryAllOf
type BulkUpdateACLEntryAllOf struct {
	Op *string `json:"op,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateACLEntryAllOf BulkUpdateACLEntryAllOf

// NewBulkUpdateACLEntryAllOf instantiates a new BulkUpdateACLEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateACLEntryAllOf() *BulkUpdateACLEntryAllOf {
	this := BulkUpdateACLEntryAllOf{}
	return &this
}

// NewBulkUpdateACLEntryAllOfWithDefaults instantiates a new BulkUpdateACLEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateACLEntryAllOfWithDefaults() *BulkUpdateACLEntryAllOf {
	this := BulkUpdateACLEntryAllOf{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *BulkUpdateACLEntryAllOf) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateACLEntryAllOf) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *BulkUpdateACLEntryAllOf) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *BulkUpdateACLEntryAllOf) SetOp(v string) {
	o.Op = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *BulkUpdateACLEntryAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateACLEntryAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *BulkUpdateACLEntryAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *BulkUpdateACLEntryAllOf) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateACLEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Op != nil {
		toSerialize["op"] = o.Op
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
func (o *BulkUpdateACLEntryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateACLEntryAllOf := _BulkUpdateACLEntryAllOf{}

	if err = json.Unmarshal(bytes, &varBulkUpdateACLEntryAllOf); err == nil {
		*o = BulkUpdateACLEntryAllOf(varBulkUpdateACLEntryAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateACLEntryAllOf is a helper abstraction for handling nullable bulkupdateaclentryallof types. 
type NullableBulkUpdateACLEntryAllOf struct {
	value *BulkUpdateACLEntryAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateACLEntryAllOf) Get() *BulkUpdateACLEntryAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateACLEntryAllOf) Set(val *BulkUpdateACLEntryAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateACLEntryAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateACLEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateACLEntryAllOf returns a pointer to a new instance of NullableBulkUpdateACLEntryAllOf.
func NewNullableBulkUpdateACLEntryAllOf(val *BulkUpdateACLEntryAllOf) *NullableBulkUpdateACLEntryAllOf {
	return &NullableBulkUpdateACLEntryAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateACLEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBulkUpdateACLEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
