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

// BulkUpdateConfigStoreListRequest struct for BulkUpdateConfigStoreListRequest
type BulkUpdateConfigStoreListRequest struct {
	Items []BulkUpdateConfigStoreItem `json:"items,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateConfigStoreListRequest BulkUpdateConfigStoreListRequest

// NewBulkUpdateConfigStoreListRequest instantiates a new BulkUpdateConfigStoreListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateConfigStoreListRequest() *BulkUpdateConfigStoreListRequest {
	this := BulkUpdateConfigStoreListRequest{}
	return &this
}

// NewBulkUpdateConfigStoreListRequestWithDefaults instantiates a new BulkUpdateConfigStoreListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateConfigStoreListRequestWithDefaults() *BulkUpdateConfigStoreListRequest {
	this := BulkUpdateConfigStoreListRequest{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BulkUpdateConfigStoreListRequest) GetItems() []BulkUpdateConfigStoreItem {
	if o == nil || o.Items == nil {
		var ret []BulkUpdateConfigStoreItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateConfigStoreListRequest) GetItemsOk() ([]BulkUpdateConfigStoreItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BulkUpdateConfigStoreListRequest) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BulkUpdateConfigStoreItem and assigns it to the Items field.
func (o *BulkUpdateConfigStoreListRequest) SetItems(v []BulkUpdateConfigStoreItem) {
	o.Items = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateConfigStoreListRequest) MarshalJSON() ([]byte, error) {
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
func (o *BulkUpdateConfigStoreListRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateConfigStoreListRequest := _BulkUpdateConfigStoreListRequest{}

	if err = json.Unmarshal(bytes, &varBulkUpdateConfigStoreListRequest); err == nil {
		*o = BulkUpdateConfigStoreListRequest(varBulkUpdateConfigStoreListRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateConfigStoreListRequest is a helper abstraction for handling nullable bulkupdateconfigstorelistrequest types. 
type NullableBulkUpdateConfigStoreListRequest struct {
	value *BulkUpdateConfigStoreListRequest
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateConfigStoreListRequest) Get() *BulkUpdateConfigStoreListRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateConfigStoreListRequest) Set(val *BulkUpdateConfigStoreListRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateConfigStoreListRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateConfigStoreListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateConfigStoreListRequest returns a pointer to a new instance of NullableBulkUpdateConfigStoreListRequest.
func NewNullableBulkUpdateConfigStoreListRequest(val *BulkUpdateConfigStoreListRequest) *NullableBulkUpdateConfigStoreListRequest {
	return &NullableBulkUpdateConfigStoreListRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateConfigStoreListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBulkUpdateConfigStoreListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
