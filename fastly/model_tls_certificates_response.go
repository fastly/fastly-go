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

// TlsCertificatesResponse struct for TlsCertificatesResponse
type TlsCertificatesResponse struct {
	Links                *PaginationLinks             `json:"links,omitempty"`
	Meta                 *PaginationMeta              `json:"meta,omitempty"`
	Data                 []TlsCertificateResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCertificatesResponse TlsCertificatesResponse

// NewTlsCertificatesResponse instantiates a new TlsCertificatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCertificatesResponse() *TlsCertificatesResponse {
	this := TlsCertificatesResponse{}
	return &this
}

// NewTlsCertificatesResponseWithDefaults instantiates a new TlsCertificatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCertificatesResponseWithDefaults() *TlsCertificatesResponse {
	this := TlsCertificatesResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TlsCertificatesResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificatesResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TlsCertificatesResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *TlsCertificatesResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TlsCertificatesResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificatesResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TlsCertificatesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *TlsCertificatesResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TlsCertificatesResponse) GetData() []TlsCertificateResponseData {
	if o == nil || o.Data == nil {
		var ret []TlsCertificateResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificatesResponse) GetDataOk() ([]TlsCertificateResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TlsCertificatesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []TlsCertificateResponseData and assigns it to the Data field.
func (o *TlsCertificatesResponse) SetData(v []TlsCertificateResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCertificatesResponse) MarshalJSON() ([]byte, error) {
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
func (o *TlsCertificatesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCertificatesResponse := _TlsCertificatesResponse{}

	if err = json.Unmarshal(bytes, &varTlsCertificatesResponse); err == nil {
		*o = TlsCertificatesResponse(varTlsCertificatesResponse)
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

// NullableTlsCertificatesResponse is a helper abstraction for handling nullable tlscertificatesresponse types.
type NullableTlsCertificatesResponse struct {
	value *TlsCertificatesResponse
	isSet bool
}

// Get returns the value.
func (v NullableTlsCertificatesResponse) Get() *TlsCertificatesResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCertificatesResponse) Set(val *TlsCertificatesResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCertificatesResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCertificatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCertificatesResponse returns a pointer to a new instance of NullableTlsCertificatesResponse.
func NewNullableTlsCertificatesResponse(val *TlsCertificatesResponse) *NullableTlsCertificatesResponse {
	return &NullableTlsCertificatesResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCertificatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCertificatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
