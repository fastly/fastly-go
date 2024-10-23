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

// BulkUpdateDictionaryListRequest struct for BulkUpdateDictionaryListRequest
type BulkUpdateDictionaryListRequest struct {
	Items                []BulkUpdateDictionaryItem `json:"items,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateDictionaryListRequest BulkUpdateDictionaryListRequest

// NewBulkUpdateDictionaryListRequest instantiates a new BulkUpdateDictionaryListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateDictionaryListRequest() *BulkUpdateDictionaryListRequest {
	this := BulkUpdateDictionaryListRequest{}
	return &this
}

// NewBulkUpdateDictionaryListRequestWithDefaults instantiates a new BulkUpdateDictionaryListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateDictionaryListRequestWithDefaults() *BulkUpdateDictionaryListRequest {
	this := BulkUpdateDictionaryListRequest{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BulkUpdateDictionaryListRequest) GetItems() []BulkUpdateDictionaryItem {
	if o == nil || o.Items == nil {
		var ret []BulkUpdateDictionaryItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateDictionaryListRequest) GetItemsOk() ([]BulkUpdateDictionaryItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BulkUpdateDictionaryListRequest) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BulkUpdateDictionaryItem and assigns it to the Items field.
func (o *BulkUpdateDictionaryListRequest) SetItems(v []BulkUpdateDictionaryItem) {
	o.Items = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateDictionaryListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BulkUpdateDictionaryListRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateDictionaryListRequest := _BulkUpdateDictionaryListRequest{}

	if err = json.Unmarshal(bytes, &varBulkUpdateDictionaryListRequest); err == nil {
		*o = BulkUpdateDictionaryListRequest(varBulkUpdateDictionaryListRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateDictionaryListRequest is a helper abstraction for handling nullable bulkupdatedictionarylistrequest types.
type NullableBulkUpdateDictionaryListRequest struct {
	value *BulkUpdateDictionaryListRequest
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateDictionaryListRequest) Get() *BulkUpdateDictionaryListRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateDictionaryListRequest) Set(val *BulkUpdateDictionaryListRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateDictionaryListRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateDictionaryListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateDictionaryListRequest returns a pointer to a new instance of NullableBulkUpdateDictionaryListRequest.
func NewNullableBulkUpdateDictionaryListRequest(val *BulkUpdateDictionaryListRequest) *NullableBulkUpdateDictionaryListRequest {
	return &NullableBulkUpdateDictionaryListRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateDictionaryListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBulkUpdateDictionaryListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
