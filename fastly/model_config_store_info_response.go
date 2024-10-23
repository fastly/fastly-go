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

// ConfigStoreInfoResponse struct for ConfigStoreInfoResponse
type ConfigStoreInfoResponse struct {
	// The number of items currently in the config store.
	ItemCount            *int32 `json:"item_count,omitempty"`
	AdditionalProperties map[string]any
}

type _ConfigStoreInfoResponse ConfigStoreInfoResponse

// NewConfigStoreInfoResponse instantiates a new ConfigStoreInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigStoreInfoResponse() *ConfigStoreInfoResponse {
	this := ConfigStoreInfoResponse{}
	return &this
}

// NewConfigStoreInfoResponseWithDefaults instantiates a new ConfigStoreInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigStoreInfoResponseWithDefaults() *ConfigStoreInfoResponse {
	this := ConfigStoreInfoResponse{}
	return &this
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *ConfigStoreInfoResponse) GetItemCount() int32 {
	if o == nil || o.ItemCount == nil {
		var ret int32
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStoreInfoResponse) GetItemCountOk() (*int32, bool) {
	if o == nil || o.ItemCount == nil {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *ConfigStoreInfoResponse) HasItemCount() bool {
	if o != nil && o.ItemCount != nil {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int32 and assigns it to the ItemCount field.
func (o *ConfigStoreInfoResponse) SetItemCount(v int32) {
	o.ItemCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ConfigStoreInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ItemCount != nil {
		toSerialize["item_count"] = o.ItemCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ConfigStoreInfoResponse) UnmarshalJSON(bytes []byte) (err error) {
	varConfigStoreInfoResponse := _ConfigStoreInfoResponse{}

	if err = json.Unmarshal(bytes, &varConfigStoreInfoResponse); err == nil {
		*o = ConfigStoreInfoResponse(varConfigStoreInfoResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableConfigStoreInfoResponse is a helper abstraction for handling nullable configstoreinforesponse types.
type NullableConfigStoreInfoResponse struct {
	value *ConfigStoreInfoResponse
	isSet bool
}

// Get returns the value.
func (v NullableConfigStoreInfoResponse) Get() *ConfigStoreInfoResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableConfigStoreInfoResponse) Set(val *ConfigStoreInfoResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConfigStoreInfoResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConfigStoreInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConfigStoreInfoResponse returns a pointer to a new instance of NullableConfigStoreInfoResponse.
func NewNullableConfigStoreInfoResponse(val *ConfigStoreInfoResponse) *NullableConfigStoreInfoResponse {
	return &NullableConfigStoreInfoResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConfigStoreInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableConfigStoreInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
