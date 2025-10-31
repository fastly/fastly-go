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

// TlsConfigurationsResponse struct for TlsConfigurationsResponse
type TlsConfigurationsResponse struct {
	Links                *PaginationLinks               `json:"links,omitempty"`
	Meta                 *PaginationMeta                `json:"meta,omitempty"`
	Data                 []TlsConfigurationResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsConfigurationsResponse TlsConfigurationsResponse

// NewTlsConfigurationsResponse instantiates a new TlsConfigurationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsConfigurationsResponse() *TlsConfigurationsResponse {
	this := TlsConfigurationsResponse{}
	return &this
}

// NewTlsConfigurationsResponseWithDefaults instantiates a new TlsConfigurationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsConfigurationsResponseWithDefaults() *TlsConfigurationsResponse {
	this := TlsConfigurationsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TlsConfigurationsResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationsResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TlsConfigurationsResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *TlsConfigurationsResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TlsConfigurationsResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationsResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TlsConfigurationsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *TlsConfigurationsResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TlsConfigurationsResponse) GetData() []TlsConfigurationResponseData {
	if o == nil || o.Data == nil {
		var ret []TlsConfigurationResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationsResponse) GetDataOk() ([]TlsConfigurationResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TlsConfigurationsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []TlsConfigurationResponseData and assigns it to the Data field.
func (o *TlsConfigurationsResponse) SetData(v []TlsConfigurationResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsConfigurationsResponse) MarshalJSON() ([]byte, error) {
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
func (o *TlsConfigurationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTlsConfigurationsResponse := _TlsConfigurationsResponse{}

	if err = json.Unmarshal(bytes, &varTlsConfigurationsResponse); err == nil {
		*o = TlsConfigurationsResponse(varTlsConfigurationsResponse)
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

// NullableTlsConfigurationsResponse is a helper abstraction for handling nullable tlsconfigurationsresponse types.
type NullableTlsConfigurationsResponse struct {
	value *TlsConfigurationsResponse
	isSet bool
}

// Get returns the value.
func (v NullableTlsConfigurationsResponse) Get() *TlsConfigurationsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsConfigurationsResponse) Set(val *TlsConfigurationsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsConfigurationsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsConfigurationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsConfigurationsResponse returns a pointer to a new instance of NullableTlsConfigurationsResponse.
func NewNullableTlsConfigurationsResponse(val *TlsConfigurationsResponse) *NullableTlsConfigurationsResponse {
	return &NullableTlsConfigurationsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsConfigurationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
