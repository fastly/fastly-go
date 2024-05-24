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

// PublishRequest Contains a batch of messages to publish.
type PublishRequest struct {
	// The messages to publish.
	Items []PublishItem `json:"items"`
	AdditionalProperties map[string]any
}

type _PublishRequest PublishRequest

// NewPublishRequest instantiates a new PublishRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishRequest(items []PublishItem) *PublishRequest {
	this := PublishRequest{}
	this.Items = items
	return &this
}

// NewPublishRequestWithDefaults instantiates a new PublishRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishRequestWithDefaults() *PublishRequest {
	this := PublishRequest{}
	return &this
}

// GetItems returns the Items field value
func (o *PublishRequest) GetItems() []PublishItem {
	if o == nil {
		var ret []PublishItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PublishRequest) GetItemsOk() ([]PublishItem, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PublishRequest) SetItems(v []PublishItem) {
	o.Items = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PublishRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PublishRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPublishRequest := _PublishRequest{}

	if err = json.Unmarshal(bytes, &varPublishRequest); err == nil {
		*o = PublishRequest(varPublishRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePublishRequest is a helper abstraction for handling nullable publishrequest types. 
type NullablePublishRequest struct {
	value *PublishRequest
	isSet bool
}

// Get returns the value.
func (v NullablePublishRequest) Get() *PublishRequest {
	return v.value
}

// Set modifies the value.
func (v *NullablePublishRequest) Set(val *PublishRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePublishRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePublishRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePublishRequest returns a pointer to a new instance of NullablePublishRequest.
func NewNullablePublishRequest(val *PublishRequest) *NullablePublishRequest {
	return &NullablePublishRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePublishRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePublishRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
