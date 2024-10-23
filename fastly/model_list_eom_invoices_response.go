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

// ListEomInvoicesResponse struct for ListEomInvoicesResponse
type ListEomInvoicesResponse struct {
	Data                 []Invoice `json:"data,omitempty"`
	Meta                 *Metadata `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _ListEomInvoicesResponse ListEomInvoicesResponse

// NewListEomInvoicesResponse instantiates a new ListEomInvoicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEomInvoicesResponse() *ListEomInvoicesResponse {
	this := ListEomInvoicesResponse{}
	return &this
}

// NewListEomInvoicesResponseWithDefaults instantiates a new ListEomInvoicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEomInvoicesResponseWithDefaults() *ListEomInvoicesResponse {
	this := ListEomInvoicesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListEomInvoicesResponse) GetData() []Invoice {
	if o == nil || o.Data == nil {
		var ret []Invoice
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEomInvoicesResponse) GetDataOk() ([]Invoice, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListEomInvoicesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Invoice and assigns it to the Data field.
func (o *ListEomInvoicesResponse) SetData(v []Invoice) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListEomInvoicesResponse) GetMeta() Metadata {
	if o == nil || o.Meta == nil {
		var ret Metadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEomInvoicesResponse) GetMetaOk() (*Metadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListEomInvoicesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Metadata and assigns it to the Meta field.
func (o *ListEomInvoicesResponse) SetMeta(v Metadata) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ListEomInvoicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ListEomInvoicesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varListEomInvoicesResponse := _ListEomInvoicesResponse{}

	if err = json.Unmarshal(bytes, &varListEomInvoicesResponse); err == nil {
		*o = ListEomInvoicesResponse(varListEomInvoicesResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableListEomInvoicesResponse is a helper abstraction for handling nullable listeominvoicesresponse types.
type NullableListEomInvoicesResponse struct {
	value *ListEomInvoicesResponse
	isSet bool
}

// Get returns the value.
func (v NullableListEomInvoicesResponse) Get() *ListEomInvoicesResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableListEomInvoicesResponse) Set(val *ListEomInvoicesResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableListEomInvoicesResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableListEomInvoicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableListEomInvoicesResponse returns a pointer to a new instance of NullableListEomInvoicesResponse.
func NewNullableListEomInvoicesResponse(val *ListEomInvoicesResponse) *NullableListEomInvoicesResponse {
	return &NullableListEomInvoicesResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableListEomInvoicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableListEomInvoicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
