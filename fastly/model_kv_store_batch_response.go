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

// KvStoreBatchResponse struct for KvStoreBatchResponse
type KvStoreBatchResponse struct {
	// A descriptor for the response of the entire batch
	Title *string `json:"title,omitempty"`
	// If an error is present in any of the requests, this field will describe that error
	Type *string `json:"type,omitempty"`
	// Errors which occurred while handling the items in the request
	Errors               []KvStoreBatchResponseErrors `json:"errors,omitempty"`
	AdditionalProperties map[string]any
}

type _KvStoreBatchResponse KvStoreBatchResponse

// NewKvStoreBatchResponse instantiates a new KvStoreBatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvStoreBatchResponse() *KvStoreBatchResponse {
	this := KvStoreBatchResponse{}
	return &this
}

// NewKvStoreBatchResponseWithDefaults instantiates a new KvStoreBatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvStoreBatchResponseWithDefaults() *KvStoreBatchResponse {
	this := KvStoreBatchResponse{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *KvStoreBatchResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreBatchResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *KvStoreBatchResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *KvStoreBatchResponse) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KvStoreBatchResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreBatchResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KvStoreBatchResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KvStoreBatchResponse) SetType(v string) {
	o.Type = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *KvStoreBatchResponse) GetErrors() []KvStoreBatchResponseErrors {
	if o == nil || o.Errors == nil {
		var ret []KvStoreBatchResponseErrors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreBatchResponse) GetErrorsOk() ([]KvStoreBatchResponseErrors, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *KvStoreBatchResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []KvStoreBatchResponseErrors and assigns it to the Errors field.
func (o *KvStoreBatchResponse) SetErrors(v []KvStoreBatchResponseErrors) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o KvStoreBatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *KvStoreBatchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKvStoreBatchResponse := _KvStoreBatchResponse{}

	if err = json.Unmarshal(bytes, &varKvStoreBatchResponse); err == nil {
		*o = KvStoreBatchResponse(varKvStoreBatchResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableKvStoreBatchResponse is a helper abstraction for handling nullable kvstorebatchresponse types.
type NullableKvStoreBatchResponse struct {
	value *KvStoreBatchResponse
	isSet bool
}

// Get returns the value.
func (v NullableKvStoreBatchResponse) Get() *KvStoreBatchResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableKvStoreBatchResponse) Set(val *KvStoreBatchResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableKvStoreBatchResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableKvStoreBatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableKvStoreBatchResponse returns a pointer to a new instance of NullableKvStoreBatchResponse.
func NewNullableKvStoreBatchResponse(val *KvStoreBatchResponse) *NullableKvStoreBatchResponse {
	return &NullableKvStoreBatchResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableKvStoreBatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableKvStoreBatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
