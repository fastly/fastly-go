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

// TLSPrivateKeysResponse struct for TLSPrivateKeysResponse
type TLSPrivateKeysResponse struct {
	Links *PaginationLinks `json:"links,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Data []TLSPrivateKeyResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSPrivateKeysResponse TLSPrivateKeysResponse

// NewTLSPrivateKeysResponse instantiates a new TLSPrivateKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSPrivateKeysResponse() *TLSPrivateKeysResponse {
	this := TLSPrivateKeysResponse{}
	return &this
}

// NewTLSPrivateKeysResponseWithDefaults instantiates a new TLSPrivateKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSPrivateKeysResponseWithDefaults() *TLSPrivateKeysResponse {
	this := TLSPrivateKeysResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TLSPrivateKeysResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeysResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TLSPrivateKeysResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *TLSPrivateKeysResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TLSPrivateKeysResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeysResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TLSPrivateKeysResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *TLSPrivateKeysResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TLSPrivateKeysResponse) GetData() []TLSPrivateKeyResponseData {
	if o == nil || o.Data == nil {
		var ret []TLSPrivateKeyResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeysResponse) GetDataOk() ([]TLSPrivateKeyResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TLSPrivateKeysResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []TLSPrivateKeyResponseData and assigns it to the Data field.
func (o *TLSPrivateKeysResponse) SetData(v []TLSPrivateKeyResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSPrivateKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSPrivateKeysResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTLSPrivateKeysResponse := _TLSPrivateKeysResponse{}

	if err = json.Unmarshal(bytes, &varTLSPrivateKeysResponse); err == nil {
		*o = TLSPrivateKeysResponse(varTLSPrivateKeysResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSPrivateKeysResponse is a helper abstraction for handling nullable tlsprivatekeysresponse types. 
type NullableTLSPrivateKeysResponse struct {
	value *TLSPrivateKeysResponse
	isSet bool
}

// Get returns the value.
func (v NullableTLSPrivateKeysResponse) Get() *TLSPrivateKeysResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSPrivateKeysResponse) Set(val *TLSPrivateKeysResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSPrivateKeysResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSPrivateKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSPrivateKeysResponse returns a pointer to a new instance of NullableTLSPrivateKeysResponse.
func NewNullableTLSPrivateKeysResponse(val *TLSPrivateKeysResponse) *NullableTLSPrivateKeysResponse {
	return &NullableTLSPrivateKeysResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSPrivateKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSPrivateKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
