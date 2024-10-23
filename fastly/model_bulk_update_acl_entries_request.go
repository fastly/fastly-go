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

// BulkUpdateACLEntriesRequest struct for BulkUpdateACLEntriesRequest
type BulkUpdateACLEntriesRequest struct {
	Entries              []BulkUpdateACLEntry `json:"entries,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateACLEntriesRequest BulkUpdateACLEntriesRequest

// NewBulkUpdateACLEntriesRequest instantiates a new BulkUpdateACLEntriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateACLEntriesRequest() *BulkUpdateACLEntriesRequest {
	this := BulkUpdateACLEntriesRequest{}
	return &this
}

// NewBulkUpdateACLEntriesRequestWithDefaults instantiates a new BulkUpdateACLEntriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateACLEntriesRequestWithDefaults() *BulkUpdateACLEntriesRequest {
	this := BulkUpdateACLEntriesRequest{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BulkUpdateACLEntriesRequest) GetEntries() []BulkUpdateACLEntry {
	if o == nil || o.Entries == nil {
		var ret []BulkUpdateACLEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateACLEntriesRequest) GetEntriesOk() ([]BulkUpdateACLEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BulkUpdateACLEntriesRequest) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []BulkUpdateACLEntry and assigns it to the Entries field.
func (o *BulkUpdateACLEntriesRequest) SetEntries(v []BulkUpdateACLEntry) {
	o.Entries = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateACLEntriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BulkUpdateACLEntriesRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateACLEntriesRequest := _BulkUpdateACLEntriesRequest{}

	if err = json.Unmarshal(bytes, &varBulkUpdateACLEntriesRequest); err == nil {
		*o = BulkUpdateACLEntriesRequest(varBulkUpdateACLEntriesRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateACLEntriesRequest is a helper abstraction for handling nullable bulkupdateaclentriesrequest types.
type NullableBulkUpdateACLEntriesRequest struct {
	value *BulkUpdateACLEntriesRequest
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateACLEntriesRequest) Get() *BulkUpdateACLEntriesRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateACLEntriesRequest) Set(val *BulkUpdateACLEntriesRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateACLEntriesRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateACLEntriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateACLEntriesRequest returns a pointer to a new instance of NullableBulkUpdateACLEntriesRequest.
func NewNullableBulkUpdateACLEntriesRequest(val *BulkUpdateACLEntriesRequest) *NullableBulkUpdateACLEntriesRequest {
	return &NullableBulkUpdateACLEntriesRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateACLEntriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBulkUpdateACLEntriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
