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

// BulkUpdateDictionaryItemAllOf struct for BulkUpdateDictionaryItemAllOf
type BulkUpdateDictionaryItemAllOf struct {
	Op *string `json:"op,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateDictionaryItemAllOf BulkUpdateDictionaryItemAllOf

// NewBulkUpdateDictionaryItemAllOf instantiates a new BulkUpdateDictionaryItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateDictionaryItemAllOf() *BulkUpdateDictionaryItemAllOf {
	this := BulkUpdateDictionaryItemAllOf{}
	return &this
}

// NewBulkUpdateDictionaryItemAllOfWithDefaults instantiates a new BulkUpdateDictionaryItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateDictionaryItemAllOfWithDefaults() *BulkUpdateDictionaryItemAllOf {
	this := BulkUpdateDictionaryItemAllOf{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *BulkUpdateDictionaryItemAllOf) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateDictionaryItemAllOf) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *BulkUpdateDictionaryItemAllOf) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *BulkUpdateDictionaryItemAllOf) SetOp(v string) {
	o.Op = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateDictionaryItemAllOf) MarshalJSON() ([]byte, error) {
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
func (o *BulkUpdateDictionaryItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateDictionaryItemAllOf := _BulkUpdateDictionaryItemAllOf{}

	if err = json.Unmarshal(bytes, &varBulkUpdateDictionaryItemAllOf); err == nil {
		*o = BulkUpdateDictionaryItemAllOf(varBulkUpdateDictionaryItemAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateDictionaryItemAllOf is a helper abstraction for handling nullable bulkupdatedictionaryitemallof types. 
type NullableBulkUpdateDictionaryItemAllOf struct {
	value *BulkUpdateDictionaryItemAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateDictionaryItemAllOf) Get() *BulkUpdateDictionaryItemAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateDictionaryItemAllOf) Set(val *BulkUpdateDictionaryItemAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateDictionaryItemAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateDictionaryItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateDictionaryItemAllOf returns a pointer to a new instance of NullableBulkUpdateDictionaryItemAllOf.
func NewNullableBulkUpdateDictionaryItemAllOf(val *BulkUpdateDictionaryItemAllOf) *NullableBulkUpdateDictionaryItemAllOf {
	return &NullableBulkUpdateDictionaryItemAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateDictionaryItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBulkUpdateDictionaryItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
