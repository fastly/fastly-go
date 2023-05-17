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

// BulkUpdateConfigStoreItemAllOf struct for BulkUpdateConfigStoreItemAllOf
type BulkUpdateConfigStoreItemAllOf struct {
	Op *string `json:"op,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateConfigStoreItemAllOf BulkUpdateConfigStoreItemAllOf

// NewBulkUpdateConfigStoreItemAllOf instantiates a new BulkUpdateConfigStoreItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateConfigStoreItemAllOf() *BulkUpdateConfigStoreItemAllOf {
	this := BulkUpdateConfigStoreItemAllOf{}
	return &this
}

// NewBulkUpdateConfigStoreItemAllOfWithDefaults instantiates a new BulkUpdateConfigStoreItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateConfigStoreItemAllOfWithDefaults() *BulkUpdateConfigStoreItemAllOf {
	this := BulkUpdateConfigStoreItemAllOf{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *BulkUpdateConfigStoreItemAllOf) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateConfigStoreItemAllOf) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *BulkUpdateConfigStoreItemAllOf) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *BulkUpdateConfigStoreItemAllOf) SetOp(v string) {
	o.Op = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateConfigStoreItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BulkUpdateConfigStoreItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateConfigStoreItemAllOf := _BulkUpdateConfigStoreItemAllOf{}

	if err = json.Unmarshal(bytes, &varBulkUpdateConfigStoreItemAllOf); err == nil {
		*o = BulkUpdateConfigStoreItemAllOf(varBulkUpdateConfigStoreItemAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateConfigStoreItemAllOf is a helper abstraction for handling nullable bulkupdateconfigstoreitemallof types. 
type NullableBulkUpdateConfigStoreItemAllOf struct {
	value *BulkUpdateConfigStoreItemAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateConfigStoreItemAllOf) Get() *BulkUpdateConfigStoreItemAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateConfigStoreItemAllOf) Set(val *BulkUpdateConfigStoreItemAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateConfigStoreItemAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateConfigStoreItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateConfigStoreItemAllOf returns a pointer to a new instance of NullableBulkUpdateConfigStoreItemAllOf.
func NewNullableBulkUpdateConfigStoreItemAllOf(val *BulkUpdateConfigStoreItemAllOf) *NullableBulkUpdateConfigStoreItemAllOf {
	return &NullableBulkUpdateConfigStoreItemAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateConfigStoreItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBulkUpdateConfigStoreItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
