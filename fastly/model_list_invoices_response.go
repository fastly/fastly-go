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

// ListInvoicesResponse struct for ListInvoicesResponse
type ListInvoicesResponse struct {
	Data []Invoice `json:"data,omitempty"`
	Meta *Metadata `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _ListInvoicesResponse ListInvoicesResponse

// NewListInvoicesResponse instantiates a new ListInvoicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInvoicesResponse() *ListInvoicesResponse {
	this := ListInvoicesResponse{}
	return &this
}

// NewListInvoicesResponseWithDefaults instantiates a new ListInvoicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInvoicesResponseWithDefaults() *ListInvoicesResponse {
	this := ListInvoicesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListInvoicesResponse) GetData() []Invoice {
	if o == nil || o.Data == nil {
		var ret []Invoice
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInvoicesResponse) GetDataOk() ([]Invoice, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListInvoicesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Invoice and assigns it to the Data field.
func (o *ListInvoicesResponse) SetData(v []Invoice) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListInvoicesResponse) GetMeta() Metadata {
	if o == nil || o.Meta == nil {
		var ret Metadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInvoicesResponse) GetMetaOk() (*Metadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListInvoicesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Metadata and assigns it to the Meta field.
func (o *ListInvoicesResponse) SetMeta(v Metadata) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ListInvoicesResponse) MarshalJSON() ([]byte, error) {
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
func (o *ListInvoicesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varListInvoicesResponse := _ListInvoicesResponse{}

	if err = json.Unmarshal(bytes, &varListInvoicesResponse); err == nil {
		*o = ListInvoicesResponse(varListInvoicesResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableListInvoicesResponse is a helper abstraction for handling nullable listinvoicesresponse types. 
type NullableListInvoicesResponse struct {
	value *ListInvoicesResponse
	isSet bool
}

// Get returns the value.
func (v NullableListInvoicesResponse) Get() *ListInvoicesResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableListInvoicesResponse) Set(val *ListInvoicesResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableListInvoicesResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableListInvoicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableListInvoicesResponse returns a pointer to a new instance of NullableListInvoicesResponse.
func NewNullableListInvoicesResponse(val *ListInvoicesResponse) *NullableListInvoicesResponse {
	return &NullableListInvoicesResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableListInvoicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableListInvoicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
