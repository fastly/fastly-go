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

// TlsBulkCertificatesResponse struct for TlsBulkCertificatesResponse
type TlsBulkCertificatesResponse struct {
	Links                *PaginationLinks                 `json:"links,omitempty"`
	Meta                 *PaginationMeta                  `json:"meta,omitempty"`
	Data                 []TlsBulkCertificateResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsBulkCertificatesResponse TlsBulkCertificatesResponse

// NewTlsBulkCertificatesResponse instantiates a new TlsBulkCertificatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsBulkCertificatesResponse() *TlsBulkCertificatesResponse {
	this := TlsBulkCertificatesResponse{}
	return &this
}

// NewTlsBulkCertificatesResponseWithDefaults instantiates a new TlsBulkCertificatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsBulkCertificatesResponseWithDefaults() *TlsBulkCertificatesResponse {
	this := TlsBulkCertificatesResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TlsBulkCertificatesResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsBulkCertificatesResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TlsBulkCertificatesResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *TlsBulkCertificatesResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TlsBulkCertificatesResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsBulkCertificatesResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TlsBulkCertificatesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *TlsBulkCertificatesResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TlsBulkCertificatesResponse) GetData() []TlsBulkCertificateResponseData {
	if o == nil || o.Data == nil {
		var ret []TlsBulkCertificateResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsBulkCertificatesResponse) GetDataOk() ([]TlsBulkCertificateResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TlsBulkCertificatesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []TlsBulkCertificateResponseData and assigns it to the Data field.
func (o *TlsBulkCertificatesResponse) SetData(v []TlsBulkCertificateResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsBulkCertificatesResponse) MarshalJSON() ([]byte, error) {
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
func (o *TlsBulkCertificatesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTlsBulkCertificatesResponse := _TlsBulkCertificatesResponse{}

	if err = json.Unmarshal(bytes, &varTlsBulkCertificatesResponse); err == nil {
		*o = TlsBulkCertificatesResponse(varTlsBulkCertificatesResponse)
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

// NullableTlsBulkCertificatesResponse is a helper abstraction for handling nullable tlsbulkcertificatesresponse types.
type NullableTlsBulkCertificatesResponse struct {
	value *TlsBulkCertificatesResponse
	isSet bool
}

// Get returns the value.
func (v NullableTlsBulkCertificatesResponse) Get() *TlsBulkCertificatesResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsBulkCertificatesResponse) Set(val *TlsBulkCertificatesResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsBulkCertificatesResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsBulkCertificatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsBulkCertificatesResponse returns a pointer to a new instance of NullableTlsBulkCertificatesResponse.
func NewNullableTlsBulkCertificatesResponse(val *TlsBulkCertificatesResponse) *NullableTlsBulkCertificatesResponse {
	return &NullableTlsBulkCertificatesResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsBulkCertificatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsBulkCertificatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
