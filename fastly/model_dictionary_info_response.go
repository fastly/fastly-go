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

// DictionaryInfoResponse struct for DictionaryInfoResponse
type DictionaryInfoResponse struct {
	// Timestamp (UTC) when the dictionary was last updated or an item was added or removed.
	LastUpdated *string `json:"last_updated,omitempty"`
	// The number of items currently in the dictionary.
	ItemCount *int32 `json:"item_count,omitempty"`
	// A hash of all the dictionary content.
	Digest               *string `json:"digest,omitempty"`
	AdditionalProperties map[string]any
}

type _DictionaryInfoResponse DictionaryInfoResponse

// NewDictionaryInfoResponse instantiates a new DictionaryInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionaryInfoResponse() *DictionaryInfoResponse {
	this := DictionaryInfoResponse{}
	return &this
}

// NewDictionaryInfoResponseWithDefaults instantiates a new DictionaryInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionaryInfoResponseWithDefaults() *DictionaryInfoResponse {
	this := DictionaryInfoResponse{}
	return &this
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DictionaryInfoResponse) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryInfoResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DictionaryInfoResponse) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *DictionaryInfoResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *DictionaryInfoResponse) GetItemCount() int32 {
	if o == nil || o.ItemCount == nil {
		var ret int32
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryInfoResponse) GetItemCountOk() (*int32, bool) {
	if o == nil || o.ItemCount == nil {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *DictionaryInfoResponse) HasItemCount() bool {
	if o != nil && o.ItemCount != nil {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int32 and assigns it to the ItemCount field.
func (o *DictionaryInfoResponse) SetItemCount(v int32) {
	o.ItemCount = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *DictionaryInfoResponse) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryInfoResponse) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *DictionaryInfoResponse) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *DictionaryInfoResponse) SetDigest(v string) {
	o.Digest = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DictionaryInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.ItemCount != nil {
		toSerialize["item_count"] = o.ItemCount
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DictionaryInfoResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDictionaryInfoResponse := _DictionaryInfoResponse{}

	if err = json.Unmarshal(bytes, &varDictionaryInfoResponse); err == nil {
		*o = DictionaryInfoResponse(varDictionaryInfoResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "item_count")
		delete(additionalProperties, "digest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDictionaryInfoResponse is a helper abstraction for handling nullable dictionaryinforesponse types.
type NullableDictionaryInfoResponse struct {
	value *DictionaryInfoResponse
	isSet bool
}

// Get returns the value.
func (v NullableDictionaryInfoResponse) Get() *DictionaryInfoResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableDictionaryInfoResponse) Set(val *DictionaryInfoResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDictionaryInfoResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDictionaryInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDictionaryInfoResponse returns a pointer to a new instance of NullableDictionaryInfoResponse.
func NewNullableDictionaryInfoResponse(val *DictionaryInfoResponse) *NullableDictionaryInfoResponse {
	return &NullableDictionaryInfoResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDictionaryInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDictionaryInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
